/* eslint-disable react/no-array-index-key */
import { Fragment, FC } from "react";
import { useAppSelector } from "@hooks/index";
import { cn } from "@utils";
import { MessageType, MessagePositions, selectToasts, ToastProp } from "../../store";
import { Toast } from "../toast";
import { DefaultParamsProp } from "../../types";

import "./index.scss";

const b = cn("toasts_renderer");

const defaultParams: DefaultParamsProp = {
    message: "not message",
    options: {
        type: MessageType.DEFAULT,
        duration: 3000,
        position: MessagePositions["LEFT-TOP"],
    },
};

const ToastsRenderer: FC = () => {
    const toasts: Array<ToastProp> = useAppSelector(selectToasts);

    const leftToasts: Array<ToastProp> = toasts.filter(
        item => !item.options.position || item.options.position === MessagePositions["LEFT-TOP"],
    );
    const rightToasts: Array<ToastProp> = toasts.filter(
        item => item.options.position === MessagePositions["RIGHT-TOP"],
    );

    return (
        <div className={b("")}>
            <div className={b("inner")}>
                <div className={b("left_top")}>
                    {leftToasts.map((item, index) => {
                        return (
                            <Fragment key={index}>
                                <Toast defaultParams={defaultParams} item={item} />
                            </Fragment>
                        );
                    })}
                </div>
                <div className={b("right_top")}>
                    {rightToasts.map((item, index) => {
                        return (
                            <Fragment key={index}>
                                <Toast defaultParams={defaultParams} item={item} />
                            </Fragment>
                        );
                    })}
                </div>
            </div>
        </div>
    );
};

export type { DefaultParamsProp };
export { ToastsRenderer };
