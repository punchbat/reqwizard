import { FC, PropsWithChildren } from "react";

import { Header, Footer, ToastsRenderer, Navbar } from "@widgets/index";

import { cn } from "@utils";
import "./index.scss";

const b = cn("main_layout");

const MainLayout: FC<PropsWithChildren> = ({ children }) => {
    return (
        <div className={b("")}>
            <div className={b("inner")}>
                <ToastsRenderer />
                <Header />
                <Navbar />

                <div className={b("content")}>
                    <main>{children}</main>
                </div>

                <Footer />
            </div>
        </div>
    );
};

export { MainLayout };
