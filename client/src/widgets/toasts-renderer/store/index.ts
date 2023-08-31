import { createSelector, createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { AppState } from "@app/store";

import { getHashString } from "@utils";

interface MessageCallback {
    id: string;
    onClose: () => void;
}

export enum MessageType {
    "DEFAULT" = "DEFAULT",
    "SUCCESS" = "SUCCESS",
    "ERROR" = "ERROR",
}

export enum MessagePositions {
    "LEFT-TOP" = "LEFT-TOP",
    "RIGHT-TOP" = "RIGHT-TOP",
}

export interface MessageOptions {
    id: string;
    duration: number | null;
    type: MessageType;
    onRequestRemove: () => void;
    onRequestClose: () => void;
    showing: boolean;
    position: MessagePositions;
}

interface MessageCallback {
    id: string;
    onClose: () => void;
}

export type MessageProp = React.ReactNode | ((callback: MessageCallback) => React.ReactNode) | string;

export interface ToastProp {
    message: MessageProp;
    options: Partial<MessageOptions>;
}

interface InitialState {
    loading: boolean;
    items: Array<ToastProp>;
}

const initialState: InitialState = {
    loading: false,
    items: [],
};

export const toastssSlice = createSlice({
    name: "toasts",
    initialState,
    reducers: {
        setToastsLoading: (state, action: PayloadAction<boolean>) => {
            state.loading = action.payload;
        },
        notify: (state, action: PayloadAction<ToastProp>) => {
            const arr = state.items;
            if (!action.payload.options.id) {
                action.payload.options.id = getHashString();
            }
            arr.push(action.payload);
            state.items = arr;
        },
        close: (state, action: PayloadAction<string | undefined>) => {
            const arr = state.items;
            if (!action.payload) {
                arr.pop();
            } else {
                arr.splice(
                    arr.findIndex(item => item.options.id === action.payload),
                    1,
                );
            }
            state.items = arr;
        },
        closeAll: state => {
            state.items = [];
        },
    },
});

export const { setToastsLoading, notify, close, closeAll } = toastssSlice.actions;

const selectToastsFunc = (state: AppState) => state.toasts.items;

export const selectToasts = createSelector(selectToastsFunc, items => items);

export default toastssSlice.reducer;
