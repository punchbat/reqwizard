import { FC, PropsWithChildren } from "react";

import { ToastsRenderer } from "@widgets/index";
import { cn } from "@utils";

import "./index.scss";

const b = cn("domik_layout");

const DomikLayout: FC<PropsWithChildren> = ({ children }) => {
    return (
        <div className={b("")}>
            <div className={b("inner")}>
                <ToastsRenderer />

                <main>{children}</main>
            </div>
        </div>
    );
};

export { DomikLayout };
