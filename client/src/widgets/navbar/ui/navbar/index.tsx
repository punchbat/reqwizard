import { FC } from "react";
import { useLocation, useNavigate } from "react-router-dom";

import { Typography, Spin } from "antd";
import { cn, isManager } from "@utils";
import { useGetProfileQuery } from "@app/services/auth";

import "./index.scss";

const { Text } = Typography;

const b = cn("navbar");

const Navbar: FC = () => {
    const location = useLocation();
    const navigate = useNavigate();

    const { data, isLoading } = useGetProfileQuery();

    const handleClick = (path: string) => {
        navigate(path);
    };

    const items = [
        {
            isActive: location.pathname === "/",
            href: "/",
            text: "Home",
        },
        {
            isActive: location.pathname === "/my-ticket-responses",
            href: "/my-ticket-responses",
            text: "My Ticket-Responses",
        },
    ];

    if (isManager(data?.payload?.userRoles)) {
        items.push({
            isActive: location.pathname === "/workspace",
            href: "/workspace",
            text: "Workspace",
        });
    }

    if (isLoading) {
        return <Spin />;
    }

    return (
        <nav className={b("")}>
            <div className={b("inner")}>
                <div className={b("items")}>
                    {items.map((item, index) => (
                        // eslint-disable-next-line jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions
                        <div
                            className={b("item")}
                            // eslint-disable-next-line react/no-array-index-key
                            key={index}
                            onClick={() => handleClick(item.href)}
                        >
                            <Text strong={item.isActive}>{item.text}</Text>
                            {/* {item.subItems?.length && (
                                    <div className={}>
                                        {item.subItems?.map((item, index) => (
                                            <div className={item.className} key={index}>
                                                <Link href={item.href}>{item.text}</Link>
                                            </div>
                                        ))}
                                    </div>
                                )} */}
                        </div>
                    ))}
                </div>
            </div>
        </nav>
    );
};

export { Navbar };
