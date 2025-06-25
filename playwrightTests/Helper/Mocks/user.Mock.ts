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
                        "firstName": "Usager",
                        "lastName": "test",
                        "role": 2
                    },
                    {
                        "id": 2,
                        "firstName": "JérémieTest",
                        "lastName": "Lapointe",
                        "role": 1
                    },
                    {
                        "id": 3,
                        "firstName": "Charle-ÉtienneTest",
                        "lastName": "Soucy",
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
                    "firstName": "JérémieTest",
                    "lastName": "Lapointe",
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
    getAllManagersSuccess: {
        url: '/users?role=1&role=2',
        method: 'GET',
        response: {
                status: 200,
                json:
                [
                    {
                        "id": 2,
                        "firstName": "JérémieTest",
                        "lastName": "Lapointe",
                        "role": 1
                    },
                    {
                        "id": 3,
                        "firstName": "Charle-ÉtienneTest",
                        "lastName": "Soucy",
                        "role": 2
                    }
                ]
        }
    },
    getAllManagersError: {
        url: '/users/managers',
        method: 'GET',
        response: {
                status: 400,
                json:
                {
                    "error": "Erreur lors de la récupération des managers"
                }
        }
    },
    logOutSuccess: {
        url: '/logout',
        method: 'POST',
        response: {
                status: 200,
                json: {"message":"Déconnexion réussie"}
        }
    },


} satisfies Record<string, MockConfig>;