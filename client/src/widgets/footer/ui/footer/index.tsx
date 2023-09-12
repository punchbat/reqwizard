import { FC } from "react";

import { Typography } from "antd";
import { cn } from "@utils";
import { GithubLogo, LinkedinLogo, TelegramLogo } from "../logos";

import "./index.scss";

const { Text } = Typography;

const b = cn("footer");

const Footer: FC = () => {
    return (
        <footer className={b("")}>
            <div className={b("inner")}>
                <div className={b("summary")}>
                    <Text disabled>© 2023 By Abat</Text>
                </div>
                <div className={b("links")}>
                    <GithubLogo />
                    <LinkedinLogo />
                    <TelegramLogo />
                </div>
            </div>
        </footer>
    );
};

export { Footer };
