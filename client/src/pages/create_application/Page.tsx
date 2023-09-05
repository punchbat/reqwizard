/* eslint-disable @typescript-eslint/no-unsafe-argument */
import { useState, FC } from "react";
import { useNavigate } from "react-router-dom";
import { Spin, Button, Form, Input, Select, Upload } from "antd";
import { InboxOutlined } from "@ant-design/icons";
import { useAppDispatch } from "@hooks/index";

import { cn, getBase64 } from "@utils";
import { useCreateMutation } from "@app/services/applicaiton";
import { ToastStore } from "@widgets/index";
import { PageTitle } from "@features/index";
import { RcFile, UploadChangeParam, UploadFile } from "antd/es/upload";
import { FailResponse } from "../../shared/types/api";

import "./Page.scss";

const { Option } = Select;
const { Dragger } = Upload;

const b = cn("application");

interface FormValues {
    type: string;
    subType: string;
    title: string;
    description: string;
    file: UploadChangeParam<UploadFile>;
}

const CreateApplication: FC = function () {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const [createApplication, { isLoading }] = useCreateMutation();

    const [subTypes, setSubTypes] = useState<string[]>([]);
    const [form] = Form.useForm();

    const handleTypeChange = (value: string) => {
        if (value === "general") {
            setSubTypes(["information", "account_help"]);
            form.setFieldsValue({ subType: undefined });
        } else if (value === "financial") {
            setSubTypes(["refunds", "payment"]);
            form.setFieldsValue({ subType: undefined });
        }
    };

    const onFinish = async (values: FormValues) => {
        try {
            const formData = new FormData();
            formData.append("type", values.type);
            formData.append("subType", values.subType);
            formData.append("title", values.title);
            formData.append("description", values.description);

            if (values.file && values.file.file) {
                const content = await getBase64(values.file.file as RcFile);

                if (content) {
                    formData.append(
                        "file",
                        new File([content], values.file.file.name, {
                            type: values.file.file.type,
                        }),
                    );
                }
            }

            await createApplication(formData).unwrap();

            navigate("/");
        } catch (err) {
            const { data } = err as { data: FailResponse };
            dispatch(
                ToastStore.notify({
                    // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access, @typescript-eslint/no-unsafe-assignment
                    message: data?.message,
                    options: {
                        type: ToastStore.MessageType.ERROR,
                        duration: 3000,
                        position: ToastStore.MessagePositions["RIGHT-TOP"],
                    },
                }),
            );
        }
    };

    if (isLoading) {
        return <Spin />;
    }

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("title")}>
                    <PageTitle title={{ text: "Create Application" }} />
                </div>
                <div className={b("content")}>
                    <Form onFinish={onFinish} form={form}>
                        <Form.Item name="type" rules={[{ required: true, message: "Please select a type" }]}>
                            <Select placeholder="Select type" onChange={handleTypeChange}>
                                <Option value="financial">Financial</Option>
                                <Option value="general">General</Option>
                            </Select>
                        </Form.Item>
                        {subTypes.length ? (
                            <Form.Item name="subType" rules={[{ required: true, message: "Please select a sub-type" }]}>
                                <Select placeholder="Select sub-type">
                                    {subTypes.map(subType => (
                                        <Option key={subType} value={subType}>
                                            {subType}
                                        </Option>
                                    ))}
                                </Select>
                            </Form.Item>
                        ) : null}
                        <Form.Item name="title" rules={[{ required: true, message: "Please enter a title" }]}>
                            <Input placeholder="Title" />
                        </Form.Item>
                        <Form.Item
                            name="description"
                            rules={[{ required: true, message: "Please enter a description" }]}
                        >
                            <Input.TextArea rows={4} placeholder="Description" />
                        </Form.Item>
                        <Form.Item
                            name="file"
                            rules={[
                                {
                                    // eslint-disable-next-line @typescript-eslint/require-await
                                    validator: async (_, value) => {
                                        if (value && value.file) {
                                            const allowedTypes = ["text/plain", "application/json"];
                                            // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
                                            if (!allowedTypes.includes(value.file.type)) {
                                                throw new Error("You can only upload .txt and .json files!");
                                            }
                                        }
                                    },
                                },
                            ]}
                        >
                            <Dragger
                                maxCount={1}
                                beforeUpload={() => {
                                    return false;
                                }}
                            >
                                <p className="ant-upload-drag-icon">
                                    <InboxOutlined />
                                </p>
                                <p className="ant-upload-text">Click or drag file to this area to upload (optional)</p>
                            </Dragger>
                        </Form.Item>
                        <Form.Item>
                            <Button type="primary" htmlType="submit" className="ant-btn-block">
                                Create Application
                            </Button>
                        </Form.Item>
                    </Form>
                </div>
            </div>
        </div>
    );
};

export { CreateApplication };
