import React, { FC } from "react";

import { Typography } from "antd";
import { GithubLogo, LinkedinLogo, TelegramLogo } from "../logos";

import styles from "./index.module.scss";

const { Text } = Typography;

const Footer: FC = () => {
    return (
        <footer className={`${styles.footer}`}>
            <div className={styles.footer__inner}>
                <div className={styles.footer__summary}>
                    <Text disabled>© 2023 By Abat</Text>
                </div>
                <div className={styles.footer__links}>
                    <GithubLogo />
                    <LinkedinLogo />
                    <TelegramLogo />
                </div>
            </div>
        </footer>
    );
};

export { Footer };
