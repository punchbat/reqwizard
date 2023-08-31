import { FC } from "react";
import { cn } from "@utils";

import "./Page.scss";

const b = cn("error");

const Error: FC = function () {
    return (
        <div className={b()}>
            <div className={b("inner")}>Error page :)</div>
        </div>
    );
};

export { Error };
