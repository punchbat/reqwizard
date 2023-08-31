import { FC } from "react";
import { Tabs, TabsProps } from "antd";
import { useNavigate, useLocation } from "react-router-dom";

import { PageTitle } from "@features/index";
import { cn } from "@utils";
import qs from "qs";
import { Applications, TicketResponses } from "./ui";

import "./Page.scss";

const items: TabsProps["items"] = [
    {
        key: "applications",
        label: "All active Applications",
        children: <Applications />,
    },
    {
        key: "ticketResponses",
        label: "You responded the Ticket-Responses",
        children: <TicketResponses />,
    },
];

const b = cn("workspace");

const Workspace: FC = function () {
    const navigate = useNavigate();
    const location = useLocation();

    const { tab } = qs.parse(location.search.substring(1)) as {
        tab: string;
    };

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("title")}>
                    <PageTitle
                        title={{
                            text: "Your Workspace",
                        }}
                    />
                </div>
                <div className={b("content")}>
                    <Tabs
                        defaultActiveKey={tab || "applications"}
                        items={items}
                        onChange={(tab: string) => {
                            navigate({
                                pathname: location.pathname,
                                search: qs.stringify({
                                    tab,
                                }),
                            });
                        }}
                    />
                </div>
            </div>
        </div>
    );
};

export { Workspace };
