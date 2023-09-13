import React, { FC, ReactNode } from "react";
import { useNavigate } from "react-router-dom";

import { Typography, Button } from "antd";
import { LeftOutlined } from "@ant-design/icons";
import { cn } from "@utils";

import "./index.scss";

const { Title } = Typography;

const b = cn("page_title");
interface Props {
    title: {
        text: string | ReactNode;
        level?: 5 | 1 | 2 | 3 | 4 | undefined;
    };
    back?: {
        text?: string;
        handleBack?: () => {};
    };
    avatar?: ReactNode;
}

const PageTitle: FC<Props> = ({ title, back, avatar }) => {
    const navigate = useNavigate();

    const handleBackOrBackToHome = () => {
        if (back?.handleBack) {
            back.handleBack();

            return;
        }

        navigate(-1);
    };

    return (
        <div className={b()}>
            <div className={b("inner", { widthAvatar: !!avatar })}>
                <Button className={b("btn")} type="text" icon={<LeftOutlined />} onClick={handleBackOrBackToHome}>
                    {back?.text || "Go to back"}
                </Button>
                <div className={b("avatar")}>{avatar}</div>
                <Title className={b("title")} level={title.level || 1}>
                    {title.text}
                </Title>
            </div>
        </div>
    );
};

export { PageTitle };
