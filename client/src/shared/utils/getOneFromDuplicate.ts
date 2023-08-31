import { allEqual } from './allEqual';
import { hasDuplicate } from './hasDuplicate';

const getOneFromDuplicate = (arr: Array<string | number | boolean | Date | undefined | null>): string | undefined => {
    if (arr.length === 1) {
        return arr[0] ? String(arr[0]) : '-';
    }

    if (hasDuplicate(arr) && allEqual(arr)) {
        return arr[0] ? String(arr[0]) : '-';
    }

    return undefined;
};

export { getOneFromDuplicate };
