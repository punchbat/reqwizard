import { useCallback, FC, ReactNode } from "react";
import { useAppDispatch, useTimeout } from "@hooks/index";
import { Typography } from "antd";
import { ExclamationCircleOutlined, CheckOutlined } from "@ant-design/icons";
import { close, MessageType, ToastProp } from "../../store";

import { DefaultParamsProp } from "../../types";

import styles from "./index.module.scss";

const { Text } = Typography;

interface Props {
    defaultParams: DefaultParamsProp;
    item: ToastProp;
}

const Toast: FC<Props> = ({ defaultParams, item }) => {
    const dispatch = useAppDispatch();

    const message = item.message || defaultParams?.message;
    const type = item?.options?.type || defaultParams?.options?.type;
    const duration = item?.options?.duration || defaultParams?.options?.duration;

    const closeToast = useCallback(() => {
        dispatch(close(item?.options?.id));
    }, [dispatch, item?.options?.id]);

    useTimeout(closeToast, duration);

    const getToast = (): ReactNode => {
        switch (type) {
            case MessageType.SUCCESS:
                return (
                    <div className={styles.toast__content}>
                        <div className={styles.toast__icon}>
                            <CheckOutlined />
                        </div>
                        <div className={styles.toast__message}>
                            <Text type="success">{message as ReactNode}</Text>
                        </div>
                    </div>
                );
            case MessageType.ERROR:
                return (
                    <div className={styles.toast__content}>
                        <div className={styles.toast__icon}>
                            <ExclamationCircleOutlined />
                        </div>
                        <div className={styles.toast__message}>
                            <Text type="danger">{message as ReactNode}</Text>
                        </div>
                    </div>
                );
            default:
                return <div>{message as ReactNode}</div>;
        }
    };

    return (
        // eslint-disable-next-line jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions
        <div className={styles.toast} onClick={closeToast}>
            <div className={styles.toast__inner}>
                <div className={styles[`toast__${type}`]}>{getToast()}</div>
            </div>
        </div>
    );
};

export { Toast };
