import React, { FC } from "react";
import { Link } from "react-router-dom";

interface Props {
    to?: string;
}

const AppLogo: FC<Props> = ({ to = "/" }) => {
    return (
        <Link to={to}>
            <svg
                width="60px"
                height="60px"
                viewBox="0 0 1024 1024"
                className="icon"
                version="1.1"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path
                    d="M211.5 595.5l434.2-434.2c18.4-18.4 48.3-18.4 66.8 0L946.2 395c18.4 18.4 18.4 48.3 0 66.8L545.4 862.6c-18.4 18.4-48.3 18.4-66.8 0L211.5 595.5z"
                    fill="#FFEA00"
                />
                <path
                    d="M478.6 862.6L77.8 461.9c-18.4-18.4-18.4-48.3 0-66.8l233.7-233.7c18.4-18.4 48.3-18.4 66.8 0l400.8 400.8c18.4 18.4 18.4 48.3 0 66.8L545.4 862.6c-18.5 18.5-48.3 18.5-66.8 0z"
                    fill="#536DFE"
                />
                <path
                    d="M345 729.1L77.8 461.9c-18.4-18.4-18.4-48.3 0-66.8l233.7-233.7c18.4-18.4 48.3-18.4 66.8 0l400.8 400.8L345 729.1z"
                    fill="#3D5AFE"
                />
                <path
                    d="M294.8 344.9c-9.2-9.2-9.2-24.2 0-33.4L411.3 195c55.5-55.5 145.4-55.5 200.9 0L462.3 344.9c-46.3 46.2-121.3 46.2-167.5 0z"
                    fill="#FFEA00"
                />
            </svg>
        </Link>
    );
};

export { AppLogo };
