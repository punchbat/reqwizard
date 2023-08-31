import { createApi } from "@reduxjs/toolkit/query/react";
import { IUser, SuccessResponse } from "@localtypes";
import { USER_TOKEN } from "@constants";
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
export interface SignUpInput {
    email: string;
    password: string;
    passwordConfirm: string;
    role: string;
}

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
        signUp: builder.mutation<void, SignUpInput>({
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
                const { data } = await queryFulfilled;

                const token = data?.payload;
                if (token) {
                    localStorage.setItem(USER_TOKEN, token);

                    await dispatch(authApi.endpoints.getProfile.initiate(undefined, { forceRefetch: true }));
                }
            },
        }),
        signIn: builder.mutation<void, SignInput>({
            query: body => ({
                url: "/auth/v1/sign-in",
                method: "POST",
                body,
            }),
        }),
        getProfile: builder.query<GetProfileResponse, void>({
            query: () => ({
                url: "/auth/v1/get-profile",
                method: "GET",
            }),
        }),
        getRoles: builder.query<GetRolesResponse, void>({
            query: () => ({
                url: "/api/role/v1/list",
                method: "GET",
            }),
        }),
    }),
});

export const {
    useSignUpMutation,
    useSendVerifyCodeMutation,
    useCheckVerifyCodeMutation,
    useSignInMutation,
    useGetProfileQuery,
    useGetRolesQuery,
} = authApi;
