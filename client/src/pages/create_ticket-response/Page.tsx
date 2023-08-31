import { useEffect } from "react";
import { Button, Form, Input, Spin } from "antd";
import { useNavigate, useParams } from "react-router-dom";

import { useCreateTicketResponseMutation } from "@app/services/ticket_response";
import { ToastStore } from "@widgets/index";
import { cn } from "@utils";
import { useAppDispatch } from "@hooks/index";
import { PageTitle } from "@features/index";
import { UID_LENGTH } from "@constants";

import "./Page.scss";

const b = cn("ticket_response");

const CreateTicketResponse = () => {
    const { id } = useParams();
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const [createTicketResponse, { isLoading }] = useCreateTicketResponseMutation();
    const [form] = Form.useForm();

    const onFinish = async (values: { text: string }) => {
        try {
            await createTicketResponse({
                applicationId: id || "",
                text: values.text,
            }).unwrap();

            navigate("/");
        } catch (error) {
            dispatch(
                ToastStore.notify({
                    message: "Failed to create ticket response",
                    options: {
                        type: ToastStore.MessageType.ERROR,
                        duration: 3000,
                    },
                }),
            );
        }
    };

    useEffect(() => {
        if (!id || id.length !== UID_LENGTH) {
            navigate("/");
        }
    }, [id, navigate]);

    if (isLoading) {
        return <Spin />;
    }

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("title")}>
                    <PageTitle title={{ text: `Ticket Response for ${id}`, level: 2 }} />
                </div>
                <div className={b("content")}>
                    <Form onFinish={onFinish} form={form}>
                        <Form.Item name="text" rules={[{ required: true, message: "Please enter text" }]}>
                            <Input.TextArea rows={7} placeholder="Enter your response here" />
                        </Form.Item>
                        <Form.Item>
                            <Button type="primary" htmlType="submit" className="ant-btn-block">
                                Send response
                            </Button>
                        </Form.Item>
                    </Form>
                </div>
            </div>
        </div>
    );
};

export { CreateTicketResponse };
