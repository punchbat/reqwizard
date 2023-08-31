import { AxiosResponse } from "axios";

interface FailResponse {
    status: number;
    message: string;
}

interface SuccessResponse<T> {
    status: number;
    payload: T;
}

type ResponseInterface<T> = SuccessResponse<T> | FailResponse;

const isSuccessResponse = <T>(response: ResponseInterface<T>): response is SuccessResponse<T> => {
    return !(response as FailResponse).message;
};

type FetchInterface<T> = AxiosResponse<ResponseInterface<T>>;

export type { FetchInterface, ResponseInterface, SuccessResponse, FailResponse };
export { isSuccessResponse };
