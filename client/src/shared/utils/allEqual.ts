type Item = string | number | boolean | Date | undefined | null;

const allEqual = (arr: Array<Item>) => arr.every(v => v === arr[0]);

export { allEqual };
