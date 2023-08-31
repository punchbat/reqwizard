import { configureStore } from "@reduxjs/toolkit";

// * сервисы
import { authApi } from "@app/services/auth";
import { applicationApi } from "@app/services/applicaiton";
import { ticketResponseApi } from "@app/services/ticket_response";
// * компоненты
import toastsSlice from "@widgets/toasts-renderer/store";
// * обычные
import userSlice from "./reducers/userSlice";

export function makeStore() {
    return configureStore({
        reducer: {
            // * сервисы
            [authApi.reducerPath]: authApi.reducer,
            [applicationApi.reducerPath]: applicationApi.reducer,
            [ticketResponseApi.reducerPath]: ticketResponseApi.reducer,

            // * компоненты
            toasts: toastsSlice,

            // * обычные
            user: userSlice,
        },

        middleware: getdefaultMiddleware =>
            getdefaultMiddleware().concat([
                authApi.middleware,
                applicationApi.middleware,
                ticketResponseApi.middleware,
            ]),
    });
}

const store = makeStore();

export type AppState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;

// export type AppThunk<ReturnType = void> = ThunkAction<ReturnType, AppState, unknown, Action<string>>;

export { store };
