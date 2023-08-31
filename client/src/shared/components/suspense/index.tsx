import { Suspense as ReactSuspense, FC, PropsWithChildren } from "react";
import { cn } from "@utils";

import { LoadingOutlined } from "@ant-design/icons";
import { Spin } from "antd";

import "./index.scss";

const b = cn("spin");

const antIcon = <LoadingOutlined style={{ fontSize: 74 }} spin />;

type Props = PropsWithChildren;

export const Suspense: FC<Props> = function ({ children }) {
    return (
        <ReactSuspense
            fallback={
                <div className={b()}>
                    <Spin indicator={antIcon} />
                </div>
            }
        >
            {children}
        </ReactSuspense>
    );
};
