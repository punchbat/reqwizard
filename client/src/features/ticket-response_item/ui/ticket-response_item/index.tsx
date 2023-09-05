/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable jsx-a11y/click-events-have-key-events */
import { FC, MouseEvent } from "react";
import { useNavigate } from "react-router-dom";
import { Typography, Spin } from "antd";

import { useGetByIDQuery } from "@app/services/applicaiton";
import { cn } from "@utils";
import moment from "moment";
import { ITicketResponse } from "@localtypes";
import { applicationStatusColors } from "@constants";

import "./index.scss";

const { Text } = Typography;

const b = cn("ticket_response_item");

type Props = ITicketResponse;

const TicketResponseItem: FC<Props> = function ({ _id, applicationId, text, createdAt, updatedAt }) {
    const navigate = useNavigate();

    const { data, isLoading } = useGetByIDQuery(applicationId || "");

    const handleClick = () => {
        navigate(`/ticket-response/${_id}`);
    };

    const handleGoToApplication = (e: MouseEvent) => {
        e.stopPropagation();

        navigate(`/application/${applicationId}`);
    };

    return (
        <div className={b()} onClick={handleClick}>
            <div className={b("inner")}>
                <div className={b("info")}>
                    <div className={b("title")} onClick={handleGoToApplication}>
                        {isLoading ? (
                            <Spin />
                        ) : (
                            <Text>
                                {data?.payload?.title}

                                {data?.payload?.status && (
                                    <sup>
                                        <Text
                                            style={{
                                                marginLeft: "4px",
                                                color: applicationStatusColors[data?.payload?.status],
                                            }}
                                        >
                                            {data?.payload?.status}
                                        </Text>
                                    </sup>
                                )}
                            </Text>
                        )}
                    </div>
                    <div className={b("text")}>
                        <Text>{text}</Text>
                    </div>
                    <div className={b("time")}>
                        <Text disabled>created at: {moment(createdAt).format("MMMM Do YYYY, h:mm:ss a")}</Text>
                        <Text disabled>updated at: {moment(updatedAt).format("MMMM Do YYYY, h:mm:ss a")}</Text>
                    </div>
                </div>
            </div>
        </div>
    );
};

export { TicketResponseItem };
