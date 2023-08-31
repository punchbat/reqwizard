import { FC, PropsWithChildren } from "react";

import { ConfigProvider, theme } from "antd";
import { PALETTE } from "@constants";

const ThemeLayout: FC<PropsWithChildren> = ({ children }) => (
    <ConfigProvider
        theme={{
            algorithm: theme.compactAlgorithm,
            token: {
                borderRadius: PALETTE.borderRadius,
                colorPrimary: PALETTE.colorPrimary,
            },
        }}
    >
        {children}
    </ConfigProvider>
);

export { ThemeLayout };
