import { FC, MouseEvent } from "react";
import { useNavigate } from "react-router-dom";
import { Typography, Button } from "antd";

import { useGetMyProfileQuery } from "@app/services/auth";
import { cn, isManager } from "@utils";
import moment from "moment";
import { IApplication } from "@localtypes";
import { applicationStatusColors } from "@constants";
import { Cube } from "@components";

import "./index.scss";

const { Text } = Typography;

const b = cn("application_item");

type Props = IApplication;

const ApplicationItem: FC<Props> = function ({
    _id,
    ticketResponseId,
    status,
    title,
    description,
    type,
    subType,
    createdAt,
    updatedAt,
}) {
    const navigate = useNavigate();

    const { data, isLoading } = useGetMyProfileQuery();

    const handleClick = () => {
        navigate(`/application/${_id}`);
    };

    const handleCreateTicketResponse = (event: MouseEvent<HTMLElement>) => {
        event.stopPropagation();

        navigate(`/create-ticket-response/${_id}`);
    };

    const isTicketResponseExist = () => ticketResponseId && ticketResponseId.length;

    return (
        // eslint-disable-next-line jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions
        <div className={b()} onClick={handleClick}>
            <div className={b("inner")}>
                <div className={b("info")}>
                    <div className={b("title")}>
                        <Text>
                            {title}
                            <sup>
                                <Text
                                    style={{
                                        marginLeft: "4px",
                                        color: applicationStatusColors[status],
                                    }}
                                >
                                    {status}
                                </Text>
                            </sup>
                        </Text>
                    </div>
                    <div className={b("description")}>
                        <Text>{description}</Text>
                    </div>
                    <div className={b("time")}>
                        <Text disabled>created at: {moment(createdAt).format("MMMM Do YYYY, h:mm:ss a")}</Text>
                        <Text disabled>updated at: {moment(updatedAt).format("MMMM Do YYYY, h:mm:ss a")}</Text>
                    </div>
                </div>
                <div className={b("sub_info")}>
                    <Cube>
                        <Text>{type}</Text>
                    </Cube>
                    <Cube>
                        <Text>{subType}</Text>
                    </Cube>
                </div>
                {!isLoading && isManager(data?.payload.userRoles) && !isTicketResponseExist() && (
                    <div className={b("manager_action")}>
                        <Button onClick={handleCreateTicketResponse}>Create ticket reponse</Button>
                    </div>
                )}
            </div>
        </div>
    );
};

export { ApplicationItem };
