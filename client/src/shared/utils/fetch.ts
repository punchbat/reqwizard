import axios, { AxiosRequestConfig, AxiosResponse, RawAxiosRequestHeaders } from "axios";

import { APIError, getSecretkey } from "@utils";

const DEFAULT_HEADERS: RawAxiosRequestHeaders = {
    "Content-Type": "application/json",
    Accept: "application/json",
};

async function handleErrorResponse(response: AxiosResponse) {
    const { status, statusText } = response;

    let data;
    let code;
    let message;

    try {
        data = await response.data;
    } finally {
        code = (data && data.internalCode) || "500_ISE";
        message = data && data.message;
    }

    throw new APIError({ code, status, statusText, message });
}

async function fetch<T>(path: string, options: AxiosRequestConfig = {}): Promise<T> {
    const { method = "GET", headers = {}, ...restOptions } = options;
    if (method !== "GET") {
        headers["x-csrf-token"] = getSecretkey();
    }

    const response = await axios(path, {
        headers: {
            ...DEFAULT_HEADERS,
            ...headers,
        },
        method,
        ...restOptions,
    });

    if (!response.data && response.status !== 204) {
        await handleErrorResponse(response);
    }

    return response.data;
}

export { fetch };
