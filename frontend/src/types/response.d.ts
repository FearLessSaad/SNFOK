
export type Meta = {
    total_count: number;
    current_page: number;
    next_page: number;
    request_id: string;
    code: number;
}

export type Response<T> = {
    status: string;
    message: string;
    data: T;
    meta: Meta;
}