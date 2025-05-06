import { MockConfig } from "../types";

export const categoryMocks = {
    getCategoriesByProjectSuccess: {
        url: '*/**/project/*/categories',
        method: 'GET',
        response: {
            status: 200,
            json: {
                "categories": []
            }
        }
    }
} satisfies Record<string, MockConfig>;