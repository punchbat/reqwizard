interface IUser {
    _id: string;
    email: string;
    verified: boolean;
    userRoles: Array<IUserRole>;
    name: string;
    surname: string;
    gender: EUserGender;
    birthday: string;
    avatar: string;
    applicationCreatedAt: string;
    createdAt: string;
    updatedAt: string;
}

enum EUserGender {
    MALE = "male",
    FEMALE = "female",
    OTHER = "other",
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
export { EApplicationStatus, EUserRoleName, EUserRoleStatus, EUserGender };
