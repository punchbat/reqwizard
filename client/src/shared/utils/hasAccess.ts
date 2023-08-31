import { State } from '@cshared/types/blackbox';
import { AccessEnums } from 'src/server/types/controllers';

enum AccessType {
    READ = 'READ',
    WRITE = 'WRITE',
}

const hasAccess = (userData: Partial<State>, key: AccessEnums | undefined, accessType: AccessType): boolean => {
    const arr =
        (accessType === AccessType.READ ? userData?.userInfo?.roles?.read : userData?.userInfo?.roles?.write) ?? [];
    return Boolean(arr.find(item => item === key));
};

export { hasAccess, AccessType };
