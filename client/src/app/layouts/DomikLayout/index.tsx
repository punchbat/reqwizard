import React, { FC, PropsWithChildren } from "react";

import { ToastsRenderer } from "@widgets/index";

import styles from "./index.module.scss";

const DomikLayout: FC<PropsWithChildren> = ({ children }) => {
    return (
        <div className={`${styles.DomikLayout}`}>
            <div className={`${styles.DomikLayout__inner}`}>
                <ToastsRenderer />

                <main>{children}</main>
            </div>
        </div>
    );
};

export { DomikLayout };
