import { fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { USER_TOKEN } from "@constants";

export const baseQuery = fetchBaseQuery({
    baseUrl: import.meta.env.VITE_API_URL,
    prepareHeaders: headers => {
        const token = localStorage.getItem(USER_TOKEN);
        if (token) {
            headers.set("authorization", `${token}`);
        }
        return headers;
    },
});
