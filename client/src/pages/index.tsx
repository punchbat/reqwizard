import { FC, lazy, ReactElement } from "react";
import { createBrowserRouter, Navigate, RouterProvider } from "react-router-dom";

import { Suspense } from "@components";
import { useGetProfileQuery } from "@app/services/auth";
import { DomikLayout, MainLayout } from "@app/layouts";
import { isManager, isAuthenticatedAndHasRole } from "@utils";
import { Error } from "./error";

const SignUp = lazy(() => import("./domik/sign-up"));
const SignIn = lazy(() => import("./domik/sign-in"));

const Home = lazy(() => import("./home"));
const CreateApplication = lazy(() => import("./create_application"));
const Application = lazy(() => import("./application"));

const MyTicketResponses = lazy(() => import("./my_ticket-responses"));
const TicketResponse = lazy(() => import("./ticket-response"));

// * manager
const CreateTicketResponse = lazy(() => import("./create_ticket-response"));
const Workspace = lazy(() => import("./workspace"));

const Whois = lazy(() => import("./whois"));

const AuthGuard: FC<{ children: ReactElement }> = ({ children }) => {
    const { data, error, isLoading } = useGetProfileQuery();

    if (isLoading) {
        return null;
    }

    if (!error && isAuthenticatedAndHasRole(data?.payload)) {
        return children;
    }

    return <Navigate to="/sign-in" />;
};

const ManagerGuard: FC<{ children: ReactElement }> = ({ children }) => {
    const { data, error, isLoading } = useGetProfileQuery();

    if (isLoading) {
        return null;
    }

    if (!error && isManager(data?.payload?.userRoles)) {
        return children;
    }

    return <Navigate to="/sign-in" />;
};

const router = createBrowserRouter([
    {
        path: "/sign-up",
        element: (
            <DomikLayout>
                <Suspense>
                    <SignUp />
                </Suspense>
            </DomikLayout>
        ),
        errorElement: <Error />,
    },
    {
        path: "/sign-in",
        element: (
            <DomikLayout>
                <Suspense>
                    <SignIn />
                </Suspense>
            </DomikLayout>
        ),
        errorElement: <Error />,
    },
    {
        path: "/",
        element: (
            <AuthGuard>
                <MainLayout>
                    <Suspense>
                        <Home />
                    </Suspense>
                </MainLayout>
            </AuthGuard>
        ),
        errorElement: <Error />,
    },
    {
        path: "/create-application",
        element: (
            <AuthGuard>
                <MainLayout>
                    <Suspense>
                        <CreateApplication />
                    </Suspense>
                </MainLayout>
            </AuthGuard>
        ),
        errorElement: <Error />,
    },
    {
        path: "/application/:id",
        element: (
            <AuthGuard>
                <MainLayout>
                    <Suspense>
                        <Application />
                    </Suspense>
                </MainLayout>
            </AuthGuard>
        ),
        errorElement: <Error />,
    },
    {
        path: "/ticket-response/:id",
        element: (
            <AuthGuard>
                <MainLayout>
                    <Suspense>
                        <TicketResponse />
                    </Suspense>
                </MainLayout>
            </AuthGuard>
        ),
        errorElement: <Error />,
    },
    {
        path: "/my-ticket-responses",
        element: (
            <AuthGuard>
                <MainLayout>
                    <Suspense>
                        <MyTicketResponses />
                    </Suspense>
                </MainLayout>
            </AuthGuard>
        ),
        errorElement: <Error />,
    },

    // * manager
    {
        path: "/create-ticket-response/:id",
        element: (
            <AuthGuard>
                <ManagerGuard>
                    <MainLayout>
                        <Suspense>
                            <CreateTicketResponse />
                        </Suspense>
                    </MainLayout>
                </ManagerGuard>
            </AuthGuard>
        ),
        errorElement: <Error />,
    },
    {
        path: "/workspace",
        element: (
            <AuthGuard>
                <ManagerGuard>
                    <MainLayout>
                        <Suspense>
                            <Workspace />
                        </Suspense>
                    </MainLayout>
                </ManagerGuard>
            </AuthGuard>
        ),
        errorElement: <Error />,
    },

    {
        path: "/whois",
        element: (
            <Suspense>
                <Whois />
            </Suspense>
        ),
        errorElement: <Error />,
    },
]);

export const Routing: FC = function () {
    return <RouterProvider router={router} />;
};
