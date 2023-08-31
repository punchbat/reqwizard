/* eslint-disable jsx-a11y/anchor-is-valid */
import { FC } from "react";
import { Typography, Spin } from "antd";

import { cn } from "@utils";
import { useGetTicketResponsesByUserIDQuery } from "@app/services/ticket_response";
import { TicketResponseItem } from "@features/index";

import "./Page.scss";

const { Title, Text } = Typography;

const b = cn("home");

const MyTicketResponses: FC = function () {
    const { data, isLoading } = useGetTicketResponsesByUserIDQuery();

    if (isLoading) {
        return <Spin />;
    }

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("title")}>
                    <Title level={1}>Your Ticket-Responses</Title>
                </div>
                <div className={b("content")}>
                    {data?.payload?.length ? (
                        <div className={b("items")}>
                            {data?.payload?.map((i, index) => (
                                // eslint-disable-next-line react/no-array-index-key, react/jsx-props-no-spreading
                                <TicketResponseItem key={index} {...i} />
                            ))}
                        </div>
                    ) : (
                        <div>
                            <Text>You don`t have Ticket-Responses</Text>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export { MyTicketResponses };
