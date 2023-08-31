import { createApi } from "@reduxjs/toolkit/query/react";
import { ITicketResponse, SuccessResponse } from "@localtypes";
import { baseQuery } from "./base";

export type CreateTicketResponseInput = {
    applicationId: string;
    text: string;
};

export type TicketResponseReponse = SuccessResponse<ITicketResponse>;

export type TicketResponsesReponse = SuccessResponse<Array<ITicketResponse>>;

// * REQUESTS
export interface FilterInput {
    search?: string;
    createdAtFrom?: string;
    createdAtTo?: string;
    updatedAtFrom?: string;
    updatedAtTo?: string;
}

export const ticketResponseApi = createApi({
    reducerPath: "ticketResponseApi",
    baseQuery,
    endpoints: builder => ({
        // * only for manager
        createTicketResponse: builder.mutation<void, CreateTicketResponseInput>({
            query: body => ({
                url: "/api/ticket-response/v1/create",
                method: "POST",
                body,
            }),
        }),
        getTicketResponsesByManagerID: builder.query<TicketResponsesReponse, FilterInput>({
            query: (filters = {}) => ({
                url: "/api/ticket-response/v1/manager-list",
                method: "GET",
                params: filters,
            }),
        }),
        // * general
        getTicketResponseByID: builder.query<TicketResponseReponse, string>({
            query: id => `/api/ticket-response/v1/${id}`,
        }),
        getTicketResponsesByUserID: builder.query<TicketResponsesReponse, FilterInput>({
            query: (filters = {}) => ({
                url: "/api/ticket-response/v1/my-list",
                method: "GET",
                params: filters,
            }),
        }),
    }),
});

export const {
    // * only for manager
    useCreateTicketResponseMutation,
    useGetTicketResponsesByManagerIDQuery,
    // * general
    useGetTicketResponseByIDQuery,
    useGetTicketResponsesByUserIDQuery,
} = ticketResponseApi;
