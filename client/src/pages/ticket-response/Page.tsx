/* eslint-disable no-underscore-dangle */
import { FC } from "react";
import { useParams } from "react-router-dom";
import { Spin, Typography } from "antd";

import { useGetTicketResponseByIDQuery } from "@app/services/ticket_response";
import { Avatar, PageTitle } from "@features/index";
import { cn } from "@utils";
import moment from "moment";

import "./Page.scss";

const { Title, Text } = Typography;

const b = cn("ticket_response");

const TicketResponse: FC = function () {
    const { id } = useParams();

    const { data, error, isLoading } = useGetTicketResponseByIDQuery(id || "");

    if (error) {
        return (
            <div className={b()}>
                <div className={b("inner")}>
                    <div className={b("title")}>
                        <PageTitle title={{ text: "Error" }} />
                    </div>
                    <div className={b("content")}>
                        <div className={b("error-message")}>
                            <Text strong>An error occurred while loading the ticket-response details</Text>
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    if (isLoading) {
        return <Spin />;
    }

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("title")}>
                    <PageTitle
                        title={{ text: `Ticket-Response: ${data?.payload?._id}` || "Ticket-Response Title" }}
                        avatar={data?.payload.manager && <Avatar {...data?.payload.manager} />}
                    />
                </div>
                <div className={b("content")}>
                    <div className={b("description")}>
                        <Title level={3}>{data?.payload?.text}</Title>
                    </div>
                    <div className={b("time")}>
                        <Text disabled>
                            created at: {moment(data?.payload?.createdAt).format("MMMM Do YYYY, h:mm:ss a")}
                        </Text>
                        <Text disabled>
                            updated at: {moment(data?.payload?.updatedAt).format("MMMM Do YYYY, h:mm:ss a")}
                        </Text>
                    </div>
                </div>
            </div>
        </div>
    );
};

export { TicketResponse };
