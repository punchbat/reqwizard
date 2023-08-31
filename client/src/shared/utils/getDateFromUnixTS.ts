const LOCALES = 'en-GB';
const MULTI_NUMBER = 1e3;

export const getDateFromUnixTS = (unix_ts: number) => {
    return new Date(unix_ts * MULTI_NUMBER).toLocaleDateString(LOCALES, {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
    });
};
