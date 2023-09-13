import { MouseEvent, FC } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Button, Spin, Typography } from "antd";

import { useGetMyProfileQuery } from "@app/services/auth";
import { useGetByIDQuery, useDownloadFileQuery } from "@app/services/applicaiton";
import { Avatar, PageTitle } from "@features/index";
import { cn, isManager } from "@utils";
import moment from "moment";
import { Cube } from "@components";
import { applicationStatusColors } from "@constants";

import "./Page.scss";

const { Title, Text, Link } = Typography;

const b = cn("application");

const Application: FC = function () {
    const navigate = useNavigate();
    const { id } = useParams();

    const { data: dataProfile, isLoading: isLoadingProfile } = useGetMyProfileQuery();

    const { data, error, isLoading } = useGetByIDQuery(id || "");
    const {
        data: dataFileUrl,
        error: errorFile,
        isLoading: isFileLoading,
    } = useDownloadFileQuery(data?.payload?.fileName || "", {
        skip: !data?.payload?.fileName,
    });

    const handleCreateTicketResponse = (event: MouseEvent<HTMLElement>) => {
        event.stopPropagation();

        navigate(`/create-ticket-response/${id}`);
    };

    if (error || errorFile) {
        const errorMessage = error
            ? "An error occurred while loading the application details"
            : "An error occurred while downloading the file";
        return (
            <div className={b()}>
                <div className={b("inner")}>
                    <div className={b("title")}>
                        <PageTitle title={{ text: "Error" }} />
                    </div>
                    <div className={b("content")}>
                        <div className={b("error-message")}>
                            <Text strong>{errorMessage}</Text>
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    if (isLoading || isLoadingProfile || isFileLoading) {
        return <Spin />;
    }

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("title")}>
                    <PageTitle
                        title={{
                            text:
                                (
                                    <>
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
                                    </>
                                ) || "Application Title",
                        }}
                        avatar={data?.payload.user && <Avatar {...data?.payload.user} />}
                    />
                </div>
                <div className={b("content")}>
                    <div className={b("type")}>
                        <Cube>
                            <Text>{data?.payload?.type}</Text>
                        </Cube>
                        <Cube>
                            <Text>{data?.payload?.subType}</Text>
                        </Cube>
                    </div>
                    <div className={b("description")}>
                        <Title level={3}>{data?.payload?.description}</Title>
                    </div>
                    {dataFileUrl && (
                        <div className={b("file")}>
                            <Link download href={dataFileUrl} target="_blank" rel="noopener noreferrer">
                                Download File
                            </Link>
                        </div>
                    )}

                    {data?.payload.ticketResponseId && (
                        <div className={b("ticketresponse")}>
                            {data?.payload.manager && <Avatar {...data?.payload.manager} />}

                            <Link href={`/ticket-response/${data?.payload.ticketResponseId}`} rel="noopener noreferrer">
                                Go to Ticket-Response
                            </Link>
                        </div>
                    )}

                    <div className={b("time")}>
                        <Text disabled>
                            created at: {moment(data?.payload?.createdAt).format("MMMM Do YYYY, h:mm:ss a")}
                        </Text>
                        <Text disabled>
                            updated at: {moment(data?.payload?.updatedAt).format("MMMM Do YYYY, h:mm:ss a")}
                        </Text>
                    </div>

                    {!data?.payload.ticketResponseId && isManager(dataProfile?.payload.userRoles) && (
                        <div className={b("action")}>
                            <Button onClick={handleCreateTicketResponse}>Create Ticket-Response</Button>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export { Application };
