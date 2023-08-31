export const parseGateDesc = (description: string) : {contractor?: string; consumer?: string;} =>
    description.split(';').reduce((obj, item) => {
        const [key, value] = item.split('=');
        obj[key] = value;
        return {
            ...obj,
        };
    }, {});
