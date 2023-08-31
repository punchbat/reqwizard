import React, { FC, PropsWithChildren } from "react";

import { Header, Footer, ToastsRenderer, Navbar } from "@widgets/index";

import styles from "./index.module.scss";

const MainLayout: FC<PropsWithChildren> = ({ children }) => {
    return (
        <div className={`${styles.MainLayout}`}>
            <div className={`${styles.MainLayout__inner}`}>
                <ToastsRenderer />
                <Header />
                <Navbar />

                <div className={styles.MainLayout__content}>
                    <main>{children}</main>
                </div>

                <Footer />
            </div>
        </div>
    );
};

export { MainLayout };
