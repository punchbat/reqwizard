import { FC, PropsWithChildren } from "react";
import { cn } from "@utils";

import "./index.scss";

const b = cn("cube");

type Props = PropsWithChildren;

export const Cube: FC<Props> = function ({ children }) {
    return (
        <div className={b("")}>
            <div className={b("inner")}>{children}</div>
        </div>
    );
};
