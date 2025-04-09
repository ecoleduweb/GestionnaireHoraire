export interface MockResponse {
    status: number;
    json: any;
    headers?: Record<string, string>;
}

export interface MockConfig {
    url: string;
    method?: string;
    response: MockResponse;
}