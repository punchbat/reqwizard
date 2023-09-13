import { FC } from "react";
import { Spin } from "antd";

import { useGetMyProfileQuery } from "@app/services/auth";
import { Avatar } from "@features/index";

const UserLogo: FC = () => {
    const { data, isLoading } = useGetMyProfileQuery();

    // eslint-disable-next-line react/jsx-props-no-spreading
    return <div>{!isLoading && data?.payload ? <Avatar {...data?.payload} /> : <Spin />}</div>;
};

export { UserLogo };
