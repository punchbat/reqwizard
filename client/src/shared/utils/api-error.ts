interface IAPIError {
    code: string;
    status: number;
    statusText: string;
    message?: string;
}

export class BaseError extends Error {
    public readonly code: string;

    public constructor(code: string, message?: string) {
        super(message);
        this.code = code;

        Object.setPrototypeOf(this, new.target.prototype);
    }
}

export class APIError extends BaseError implements IAPIError {
    public readonly status: number;
    public readonly statusText: string;
    public readonly message: string;

    public constructor(error: IAPIError) {
        const { code, statusText, status, message } = error;

        super(code, message);

        this.status = status;
        this.statusText = statusText;
        this.message = message || '';
    }
}
