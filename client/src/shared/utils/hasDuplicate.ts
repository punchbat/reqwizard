import { uniqWith, isEqual } from 'lodash';

const hasDuplicate = (arr: Array<string | number | boolean | Date | undefined | null>): boolean => {
    return uniqWith(arr, isEqual).length !== arr.length;
};

export { hasDuplicate };
