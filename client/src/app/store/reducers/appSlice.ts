import { createSelector, createSlice, PayloadAction } from "@reduxjs/toolkit";

import type { AppState } from "@app/store";

interface InitialState {
    csrfToken: string;
}

const initialState: Partial<InitialState> = {
    csrfToken: "",
};

export const appSlice = createSlice({
    name: "app",
    initialState,
    reducers: {
        setCsrfToken: (state, action: PayloadAction<string>) => {
            state.csrfToken = action.payload;
        },
    },
});

export const { setCsrfToken } = appSlice.actions;

const selectCsrfTokenFunc = (state: AppState): InitialState => state.app as InitialState;
export const selectCsrfToken = createSelector(selectCsrfTokenFunc, app => app.csrfToken);

export default appSlice.reducer;
