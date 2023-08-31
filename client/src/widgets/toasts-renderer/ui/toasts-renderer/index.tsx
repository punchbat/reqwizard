/* eslint-disable react/no-array-index-key */
import React, { Fragment, FC } from "react";
import { useAppSelector } from "@hooks";
import { MessageType, MessagePositions, selectToasts, ToastProp } from "../../store";
import { Toast } from "../toast";
import { DefaultParamsProp } from "../../types";

import styles from "./index.module.scss";

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
        <div className={styles.toasts_renderer}>
            <div className={styles.toasts_renderer__inner}>
                <div className={styles.toasts_renderer__left_top}>
                    {leftToasts.map((item, index) => {
                        return (
                            <Fragment key={index}>
                                <Toast defaultParams={defaultParams} item={item} />
                            </Fragment>
                        );
                    })}
                </div>
                <div className={styles.toasts_renderer__right_top}>
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
