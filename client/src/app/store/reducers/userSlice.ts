import { createSelector, createSlice, PayloadAction } from "@reduxjs/toolkit";

import type { AppState } from "@app/store";
import { IUser } from "@localtypes";

interface InitialState {
    loading: boolean;

    user: IUser;
}

const initialState: Partial<InitialState> = {
    loading: false,

    user: {
        _id: "",
        email: "",
        verified: false,
        userRoles: [],
        applicationCreatedAt: "",
        createdAt: "",
        updatedAt: "",
    },
};

export const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {
        setLoading: (state, action: PayloadAction<boolean>) => {
            state.loading = action.payload;
        },
        setUser: (state, action: PayloadAction<IUser>) => {
            state.user = action.payload;
        },
    },
});

export const { setLoading, setUser } = userSlice.actions;

const selectUserFunc = (state: AppState): InitialState => state.user as InitialState;
export const selectUser = createSelector(selectUserFunc, user => user);

export default userSlice.reducer;
