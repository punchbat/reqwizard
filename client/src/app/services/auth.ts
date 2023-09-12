import { createApi } from "@reduxjs/toolkit/query/react";
import { IUser, SuccessResponse } from "@localtypes";
import { baseQuery } from "./base";

// * RESPONSES
export type SignResponse = undefined;

export type TokenResponse = SuccessResponse<string>;

export type GetProfileResponse = SuccessResponse<IUser>;

export type GetRolesResponse = SuccessResponse<
    Array<{
        _id: string;
        name: string;
        createdAt: string;
        updatedAt: string;
    }>
>;

// * REQUESTS
export interface SignInput {
    email: string;
    password: string;
}

export interface CheckVerifyCodeInput {
    email: string;
    password: string;
    verifyCode: string;
}

export const authApi = createApi({
    reducerPath: "authApi",
    baseQuery,
    endpoints: builder => ({
        signUp: builder.mutation<void, FormData>({
            query: body => ({
                url: "/auth/v1/sign-up",
                method: "POST",
                body,
            }),
        }),
        sendVerifyCode: builder.mutation<void, SignInput>({
            query: body => ({
                url: "/auth/v1/send-verify-code",
                method: "POST",
                body,
            }),
        }),
        checkVerifyCode: builder.mutation<TokenResponse, CheckVerifyCodeInput>({
            query: body => ({
                url: "/auth/v1/check-verify-code",
                method: "POST",
                body,
            }),
            async onQueryStarted(_, { dispatch, queryFulfilled }) {
                await queryFulfilled;

                await dispatch(authApi.endpoints.getMyProfile.initiate(undefined, { forceRefetch: true }));
            },
        }),
        signIn: builder.mutation<void, SignInput>({
            query: body => ({
                url: "/auth/v1/sign-in",
                method: "POST",
                body,
            }),
        }),
        getMyProfile: builder.query<GetProfileResponse, void>({
            query: () => ({
                url: "/auth/v1/get-my-profile",
                method: "GET",
            }),
        }),
        getProfile: builder.query<GetProfileResponse, string>({
            query: id => ({
                url: `/auth/v1/get-profile/${id}`,
                method: "GET",
            }),
        }),
        updateProfile: builder.mutation<void, FormData>({
            query: body => ({
                url: "/auth/v1/update-profile",
                method: "PUT",
                body,
            }),
            async onQueryStarted(_, { dispatch, queryFulfilled }) {
                await queryFulfilled;

                await dispatch(authApi.endpoints.getMyProfile.initiate(undefined, { forceRefetch: true }));
            },
        }),
        getRoles: builder.query<GetRolesResponse, void>({
            query: () => ({
                url: "/api/role/v1/list",
                method: "GET",
            }),
        }),
        logout: builder.mutation<void, void>({
            query: () => ({
                url: "/auth/v1/logout",
                method: "POST",
            }),
        }),
    }),
});

export const {
    useSignUpMutation,
    useSendVerifyCodeMutation,
    useCheckVerifyCodeMutation,
    useSignInMutation,
    useGetMyProfileQuery,
    useGetProfileQuery,
    useUpdateProfileMutation,
    useGetRolesQuery,
    useLogoutMutation,
} = authApi;
