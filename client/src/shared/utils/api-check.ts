import { IUserRole, EUserRoleName, EUserRoleStatus, IUser } from "@localtypes";

const isUser = (arr: Array<IUserRole> | undefined) =>
    arr ? arr?.findIndex(i => i.status === EUserRoleStatus.APPROVED && i.name === EUserRoleName.USER) !== -1 : false;

const isManager = (arr: Array<IUserRole> | undefined) =>
    arr ? arr?.findIndex(i => i.status === EUserRoleStatus.APPROVED && i.name === EUserRoleName.MANAGER) !== -1 : false;

const isAuthenticated = (user: IUser) => {
    return user.email && user.verified;
};

const isAuthenticatedAndHasRole = (user: IUser | undefined) => {
    return user ? isAuthenticated(user) && (isUser(user.userRoles) || isManager(user.userRoles)) : false;
};

export { isUser, isManager, isAuthenticatedAndHasRole };
