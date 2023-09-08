/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable react/jsx-no-comment-textnodes */
import { FC, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { cn } from "@utils";
import { Modal, Spin, Typography, Button } from "antd";
import { useGetProfileQuery } from "@app/services/auth";
import { PALETTE } from "@constants";

import "./index.scss";

const { Text } = Typography;

const b = cn("user");

const UserLogo: FC = () => {
    const navigate = useNavigate();

    const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
    const { data, error, isLoading } = useGetProfileQuery();

    const handleOpen = () => {
        setIsModalOpen(!isModalOpen);
    };

    const handleExit = () => {
        navigate("/sign-in");
    };

    useEffect(() => {
        fetch("http://localhost:8080/uploads/avatars/4d7ff447-81c6-4951-b8d2-f7d9edb9d16a.jpg", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        }).then(response => {
            console.log(response);
        });
    }, []);

    if (error) {
        return <div>Something went wrong!</div>;
    }

    return (
        <>
            <div className={b()} onClick={handleOpen}>
                <div className={b("inner")}>
                    <img src="http://localhost:8080/uploads/avatars/4d7ff447-81c6-4951-b8d2-f7d9edb9d16a.jpg" alt="" />

                    <svg width="40px" height="40px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <g id="SVGRepo_bgCarrier" strokeWidth="0" />

                        <g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round" />

                        <g id="SVGRepo_iconCarrier">
                            {" "}
                            <path
                                d="M13.7517 4.69329C13.7517 3.85839 13.1416 3.14883 12.3161 3.0238C12.1066 2.99207 11.8935 2.99207 11.6839 3.0238C10.8584 3.14883 10.2482 3.85839 10.2482 4.6933V5.87397C9.77662 6.00858 9.32734 6.19618 8.90726 6.42992L8.07205 5.59471C7.48168 5.00435 6.54849 4.93407 5.87637 5.42937C5.70578 5.55509 5.55509 5.70578 5.42937 5.87637C4.93407 6.54849 5.00434 7.4817 5.59471 8.07207L6.42992 8.90728C6.19617 9.32735 6.00857 9.77663 5.87397 10.2483H4.6933C3.85839 10.2483 3.14883 10.8584 3.0238 11.6839C2.99207 11.8935 2.99207 12.1066 3.0238 12.3161C3.14883 13.1416 3.85839 13.7518 4.69329 13.7518H5.87396C6.00857 14.2234 6.19617 14.6727 6.42992 15.0927L5.59471 15.9279C5.00434 16.5183 4.93407 17.4515 5.42937 18.1236C5.55509 18.2942 5.70578 18.4449 5.87638 18.5706C6.5485 19.0659 7.48169 18.9957 8.07205 18.4053L8.90726 17.5701C9.32734 17.8038 9.77662 17.9914 10.2482 18.126V19.3067C10.2482 20.1416 10.8584 20.8512 11.6839 20.9762C11.8934 21.0079 12.1066 21.0079 12.3161 20.9762C13.1416 20.8512 13.7517 20.1416 13.7517 19.3067V18.1261C14.2234 17.9914 14.6727 17.8038 15.0927 17.5701L15.9279 18.4053C16.5183 18.9957 17.4515 19.0659 18.1236 18.5706C18.2942 18.4449 18.4449 18.2942 18.5706 18.1236C19.0659 17.4515 18.9957 16.5183 18.4053 15.928L17.5701 15.0928C17.8038 14.6727 17.9914 14.2234 18.1261 13.7518H19.3067C20.1416 13.7518 20.8512 13.1416 20.9762 12.3161C21.0079 12.1066 21.0079 11.8935 20.9762 11.6839C20.8512 10.8584 20.1416 10.2483 19.3067 10.2483H18.126C17.9914 9.77662 17.8038 9.32734 17.5701 8.90726L18.4053 8.07205C18.9957 7.48168 19.0659 6.54849 18.5706 5.87637C18.4449 5.70578 18.2942 5.55509 18.1236 5.42937C17.4515 4.93407 16.5183 5.00434 15.9279 5.59471L15.0927 6.42992C14.6727 6.19617 14.2234 6.00857 13.7517 5.87396V4.69329Z"
                                stroke="#FFEA00"
                                strokeWidth="1.5"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                            />{" "}
                            <path
                                d="M18.1236 18.5706C17.4515 19.0659 16.5183 18.9957 15.9279 18.4053L15.0927 17.5701C14.6727 17.8038 14.2234 17.9914 13.7517 18.1261V19.3067C13.7517 20.1416 13.1416 20.8512 12.3161 20.9762C12.1066 21.0079 11.8934 21.0079 11.6839 20.9762C10.8584 20.8512 10.2482 20.1416 10.2482 19.3067V18.126C9.77662 17.9914 9.32734 17.8038 8.90726 17.5701L8.07205 18.4053C7.48169 18.9957 6.5485 19.0659 5.87638 18.5706C5.70578 18.4449 5.55509 18.2942 5.42937 18.1236C4.93407 17.4515 5.00434 16.5183 5.59471 15.9279L6.42992 15.0927C6.19617 14.6727 6.00857 14.2234 5.87396 13.7518H4.69329C3.85839 13.7518 3.14883 13.1416 3.0238 12.3161C2.99207 12.1066 2.99207 11.8935 3.0238 11.6839C3.14883 10.8584 3.85839 10.2483 4.6933 10.2483H5.87397C6.00857 9.77663 6.19617 9.32735 6.42992 8.90728L5.59471 8.07207C5.00434 7.4817 4.93407 6.54849 5.42937 5.87637C5.55509 5.70578 5.70578 5.55509 5.87637 5.42937C6.54849 4.93407 7.48168 5.00435 8.07205 5.59471L8.90726 6.42992C9.32734 6.19618 9.77662 6.00858 10.2482 5.87397V4.6933C10.2482 3.85839 10.8584 3.14883 11.6839 3.0238C11.8935 2.99207 12.1066 2.99207 12.3161 3.0238C13.1416 3.14883 13.7517 3.85839 13.7517 4.69329V5.87396C14.2234 6.00857 14.6727 6.19617 15.0927 6.42992L15.9279 5.59471C16.5183 5.00434 17.4515 4.93407 18.1236 5.42937"
                                stroke={PALETTE.colorPrimary}
                                strokeWidth="1.5"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                            />{" "}
                            <path
                                d="M9.5 12C9.5 10.6193 10.6193 9.5 12 9.5C13.3807 9.5 14.5 10.6193 14.5 12C14.5 13.3807 13.3807 14.5 12 14.5C10.6193 14.5 9.5 13.3807 9.5 12Z"
                                stroke="#FFEA00"
                                strokeWidth="1.5"
                            />{" "}
                        </g>
                    </svg>
                </div>
            </div>
            <Modal open={isModalOpen} onCancel={handleOpen} footer={[]}>
                {isLoading ? (
                    <Spin />
                ) : (
                    <div className={b("modal")}>
                        <div className={b("picture")}>
                            <img src="/assets/user.png" alt="" />
                        </div>
                        <div className={b("content")}>
                            <div className={b("email")}>
                                <Text strong>{data?.payload?.email}</Text>
                            </div>
                            <div className={b("roles")}>
                                {data?.payload?.userRoles?.map((i, index) => (
                                    // eslint-disable-next-line react/no-array-index-key
                                    <div className={b("role")} key={index}>
                                        <Text disabled>
                                            Role {i.name} is {i.status}
                                        </Text>
                                    </div>
                                ))}
                            </div>
                        </div>
                        <div className={b("additional_info")}>
                            <div className={b("last_application_time")}>
                                <Text style={{ fontSize: "10px" }}>Last Application is created at: </Text>
                                <Text style={{ fontSize: "10px" }} strong>
                                    {data?.payload?.applicationCreatedAt
                                        ? new Date(data?.payload?.applicationCreatedAt).toLocaleDateString("en-US", {
                                              year: "numeric",
                                              month: "long",
                                              day: "numeric",
                                              hour: "2-digit",
                                              minute: "2-digit",
                                              second: "2-digit",
                                              timeZoneName: "short",
                                          })
                                        : "Not found"}
                                </Text>
                            </div>
                            <div className={b("created_at")}>
                                <Text style={{ fontSize: "10px" }}>Account is created at: </Text>
                                <Text style={{ fontSize: "10px" }} strong>
                                    {data?.payload?.createdAt
                                        ? new Date(data?.payload?.createdAt).toLocaleDateString("en-US", {
                                              year: "numeric",
                                              month: "long",
                                              day: "numeric",
                                              hour: "2-digit",
                                              minute: "2-digit",
                                              second: "2-digit",
                                              timeZoneName: "short",
                                          })
                                        : "Not found"}
                                </Text>
                            </div>
                        </div>
                        <div className={b("actions")}>
                            <Button type="link" htmlType="button" onClick={handleExit}>
                                Exit
                            </Button>
                        </div>
                    </div>
                )}
            </Modal>
        </>
    );
};

export { UserLogo };
