interface IUser {
    _id: string;
    email: string;
    verified: boolean;
    userRoles: Array<IUserRole>;
    applicationCreatedAt: string;
    createdAt: string;
    updatedAt: string;
}

interface IUserRole {
    name: EUserRoleName;
    status: EUserRoleStatus;
    createdAt: string;
    updatedAt: string;
}

enum EUserRoleName {
    USER = "user",
    MANAGER = "manager",
}

enum EUserRoleStatus {
    CANCELED = "canceled",
    PENDING = "pending",
    APPROVED = "approved",
}

interface IApplication {
    _id: string;
    ticketResponseId: string;
    userId: string;
    managerId: string;
    status: EApplicationStatus;
    title: string;
    description: string;
    type: string;
    subType: string;
    fileName: string;
    createdAt: string;
    updatedAt: string;
}

enum EApplicationStatus {
    CANCELED = "canceled",
    WAITING = "waiting",
    WORKING = "working",
    DONE = "done",
}

interface ITicketResponse {
    _id: string;
    applicationId: string;
    userId: string;
    managerId: string;
    text: string;
    createdAt: string;
    updatedAt: string;
}

export type { IUser, IUserRole, IApplication, ITicketResponse };
export { EApplicationStatus, EUserRoleName, EUserRoleStatus };
