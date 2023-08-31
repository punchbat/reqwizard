/* eslint-disable jsx-a11y/anchor-is-valid */
import { FC } from "react";
import { Typography, Spin, Button } from "antd";
import { useNavigate } from "react-router-dom";

import { cn } from "@utils";
import { useGetMyListQuery } from "@app/services/applicaiton";
import { ApplicationItem } from "@features/index";

import "./Page.scss";

const { Title, Text, Link } = Typography;

const b = cn("home");

const Home: FC = function () {
    const navigate = useNavigate();
    const { data: dataApplications, isLoading: isLoadingApplications } = useGetMyListQuery();

    const handleGoToApplication = () => {
        navigate("/create-application");
    };

    if (isLoadingApplications) {
        return <Spin />;
    }

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("title")}>
                    <Title level={1}>Your Applications</Title>
                    <Button onClick={handleGoToApplication}>Create new Application</Button>
                </div>
                <div className={b("content")}>
                    {dataApplications?.payload?.length ? (
                        <div className={b("items")}>
                            {dataApplications?.payload?.map((i, index) => (
                                // eslint-disable-next-line react/no-array-index-key, react/jsx-props-no-spreading
                                <ApplicationItem key={index} {...i} />
                            ))}
                            {/* <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} />
                            <ApplicationItem {...dataApplications?.payload[0]} /> */}
                        </div>
                    ) : (
                        <div>
                            <Text>
                                You have no active applications,{" "}
                                <Link onClick={handleGoToApplication}>you can create a new one</Link>
                            </Text>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export { Home };
