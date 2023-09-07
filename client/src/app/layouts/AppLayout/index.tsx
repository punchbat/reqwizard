import { useEffect, FC, PropsWithChildren } from "react";

import { useAppDispatch, useAppSelector } from "@hooks/index";
import { selectCsrfToken, setCsrfToken } from "@reducers/appSlice";

const getCSRFToken = async (): Promise<string | null> => {
    try {
        const response = await fetch(`${import.meta.env.VITE_API_URL}/csrf`, {
            credentials: "include",
        });
        // const data = await response.json();
        // return data.token;

        console.log(response.headers);
        return response.headers.get("X-CSRF-Token");
    } catch (error) {
        console.error("Ошибка при получении CSRF-токена:", error);
        return null;
    }
};

const AppLayout: FC<PropsWithChildren> = ({ children }) => {
    const dispatch = useAppDispatch();
    const csrfToken = useAppSelector(selectCsrfToken);

    // eslint-disable-next-line react-hooks/exhaustive-deps
    const makeFetch = async () => {
        const responseCsrfToken = await getCSRFToken();

        if (!csrfToken && responseCsrfToken && responseCsrfToken.length) {
            dispatch(setCsrfToken(responseCsrfToken));
        }
    };

    useEffect(() => {
        console.log("main");

        makeFetch();
    }, []);

    // eslint-disable-next-line react/jsx-no-useless-fragment
    return <>{children}</>;
};

export { AppLayout };
