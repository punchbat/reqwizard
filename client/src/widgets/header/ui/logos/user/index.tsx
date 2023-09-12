import { FC } from "react";
import { useNavigate } from "react-router-dom";
import { cn } from "@utils";
import { Spin } from "antd";
import { useGetMyProfileQuery } from "@app/services/auth";

import "./index.scss";

const b = cn("user");

const UserLogo: FC = () => {
    const navigate = useNavigate();

    const { data, error, isLoading } = useGetMyProfileQuery();

    const handleClick = () => {
        // eslint-disable-next-line @typescript-eslint/restrict-template-expressions, no-underscore-dangle
        navigate(`/profile/${data?.payload._id}`);
    };

    if (error) {
        return <div>Something went wrong!</div>;
    }

    if (isLoading) {
        return <Spin />;
    }

    return (
        // eslint-disable-next-line jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions
        <div className={b()} onClick={handleClick}>
            <div className={b("inner")}>
                <div className={b("avatar")}>
                    <img src={`${import.meta.env.VITE_API_URL}/uploads/avatars/${data?.payload.avatar}`} alt="" />
                </div>
            </div>
        </div>
    );
};

export { UserLogo };
