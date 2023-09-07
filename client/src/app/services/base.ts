import { fetchBaseQuery, type BaseQueryFn, FetchBaseQueryError } from "@reduxjs/toolkit/query";

let csrfToken = "12345";

const query = fetchBaseQuery({
    baseUrl: import.meta.env.VITE_API_URL,
    credentials: "include",
});

const baseQuery: BaseQueryFn<any, unknown, FetchBaseQueryError> = async (args, api, extraOptions) => {
    const headers = {
        "X-Csrf-Token": csrfToken,
    };

    let updatedArgs = {
        url: args,
        headers: {
            ...headers,
        },
    };

    if (typeof args !== "string") {
        updatedArgs = {
            ...args,
            headers: {
                ...args.headers,
                ...headers,
            },
        };
    }

    const result = await query(updatedArgs, api, extraOptions);

    const responseCsrfToken = result.meta?.response?.headers.get("x-csrf-token");

    if (responseCsrfToken && responseCsrfToken.length) {
        csrfToken = responseCsrfToken;
    }
    return result;
};
export { baseQuery };
