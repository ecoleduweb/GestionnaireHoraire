import { MockConfig } from "../types";

export const userMocks = {
    userMeSuccess: {
        url: '/user/me',
        method: 'GET',
        response: {
                status: 200,
                json:
                {
                    "firstName": "Usager",
                    "lastName": "test",
                    "role": 2,
                }
        }
    },
    usersSuccess: {
        url: '/users',
        method: 'GET',
        response: {
                status: 200,
                json:
                [
                    {
                        "id": 1,
                        "first_name": "Usager",
                        "last_name": "test",
                        "role": 2
                    },
                    {
                        "id": 2,
                        "first_name": "JérémieTest",
                        "last_name": "Lapointe",
                        "role": 1
                    },
                    {
                        "id": 3,
                        "first_name": "Charle-ÉtienneTest",
                        "last_name": "Soucy",
                        "role": 0
                    }
                ]
        }
    },
    updateUserRoleSuccess: {
        url: '/user/*/role',
        method: 'PATCH',
        response: {
                status: 200,
                json:
                {
                    "id": 2,
                    "first_name": "JérémieTest",
                    "last_name": "Lapointe",
                    "email": "jeremietest.lapointe@llio.com",
                    "role": 2
                }
        }
    },
    updateUserRoleError: {
        url: '/user/*/role',
        method: 'PATCH',
        response: {
                status: 400,
                json:
                {
                    "error": "Erreur lors de la mise à jour du rôle"
                }
        }
    },


} satisfies Record<string, MockConfig>;