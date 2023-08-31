import { createApi } from "@reduxjs/toolkit/query/react";
import { IApplication, SuccessResponse } from "@localtypes";
import { baseQuery } from "./base";

// * RESPONSES
export type CreateResponse = undefined;

export type GetByIDResponse = SuccessResponse<IApplication>;

export type GetMyListResponse = SuccessResponse<Array<IApplication>>;

export type GetListResponse = SuccessResponse<Array<IApplication>>;

// * REQUESTS
export interface FilterInput {
    search?: string;
    statuses?: Array<string>;
    types?: Array<string>;
    subTypes?: Array<string>;
    createdAtFrom?: string;
    createdAtTo?: string;
    updatedAtFrom?: string;
    updatedAtTo?: string;
}

export const applicationApi = createApi({
    reducerPath: "applicationApi",
    baseQuery,
    endpoints: builder => ({
        create: builder.mutation<CreateResponse, FormData>({
            query: formData => ({
                url: "/api/application/v1/create",
                method: "POST",
                body: formData,
            }),
            async onQueryStarted(_, { dispatch, queryFulfilled }) {
                await queryFulfilled;

                await dispatch(
                    applicationApi.endpoints.getMyList.initiate({}, { subscribe: false, forceRefetch: true }),
                );
                await dispatch(
                    applicationApi.endpoints.getManagerList.initiate({}, { subscribe: false, forceRefetch: true }),
                );
            },
        }),
        getByID: builder.query<GetByIDResponse, string>({
            query: id => ({
                url: `/api/application/v1/${id}`,
                method: "GET",
            }),
        }),
        downloadFile: builder.query<string, string>({
            query: id => ({
                url: `/api/application/v1/download-file/${id}`,
                method: "GET",
                responseHandler: async response => {
                    const blob = await response.blob();
                    const blobUrl = URL.createObjectURL(blob);
                    return blobUrl;
                },
            }),
        }),
        getMyList: builder.query<GetMyListResponse, FilterInput>({
            query: (filters = {}) => ({
                url: "/api/application/v1/my-list",
                method: "GET",
                params: filters,
            }),
        }),
        getManagerList: builder.query<GetListResponse, FilterInput>({
            query: (filters = {}) => ({
                url: "/api/application/v1/manager-list",
                method: "GET",
                params: filters,
            }),
        }),
    }),
});

export const { useCreateMutation, useGetByIDQuery, useDownloadFileQuery, useGetMyListQuery, useGetManagerListQuery } =
    applicationApi;
