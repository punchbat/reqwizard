export function format2Date(date: string) {
    if (date?.[date.length - 1] === 'Z') {
        return date.slice(0, -1);
    }
    return date;
}

export function date2Format(date: string) {
    if (date?.[date.length - 1] !== 'Z') {
        return date + 'Z';
    }
    return date;
}
