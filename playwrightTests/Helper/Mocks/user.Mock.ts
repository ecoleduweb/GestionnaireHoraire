import { MockConfig } from "../types";

export const UserMocks = {
    userMeSuccess: {
        url: '*/**/user/me',
        method: 'GET',
        response: {
                status: 200,
                json:
                {
                    "firstName": "Usager",
                    "lastName": "test"
                }
        }
    },

} satisfies Record<string, MockConfig>;