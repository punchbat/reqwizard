import React, { FC } from "react";

import { Typography } from "antd";
import { AppLogo, UserLogo } from "../logos";

import styles from "./index.module.scss";

const { Text } = Typography;

const Header: FC = () => {
    return (
        <header className={styles.header}>
            <div className={`${styles.header__inner}`}>
                <div className={`${styles.header__app}`}>
                    <AppLogo to="/sign-in" />

                    <Text strong>Communication between users and reqwizard</Text>
                </div>

                <div className={`${styles.header__user}`}>
                    <UserLogo />
                </div>
            </div>
        </header>
    );
};

export { Header };
