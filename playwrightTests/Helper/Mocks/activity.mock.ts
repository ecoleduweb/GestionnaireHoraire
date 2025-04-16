import { MockConfig } from "../types";

export const activityMocks = {
    getByIdFail: {
        url: '*/**/activity/*',
        method: 'GET',
        response: {
            status: 404,
            json: {"error":"Le(La) activité n'a pas été trouvé(e)"}
        }
    },
    getByIdSuccess: {
        url: '*/**/activity/*',
        method: 'GET',
        response: {
            status: 200,
            json: {"activity":
                {"id": 1,
                    "name": "Fête de Jean-Félix et Sherlock",
                    "description": "Fête de Jean-Félix et Sherlock",
                    "startDate": "2025-03-22T18:04:00Z",
                    "endDate": "2025-03-22T20:34:00Z",
                    "userId": 1,
                    "projectId": 1,
                    "categoryId": 1
                }
            }
        }
    },
    addActivitySuccess: {
        url: '*/**/activity',
        method: 'POST',
        response: {
            status: 201,
            json: {
                "activity": {
                    "id": 1,
                    "name": "Fête de Jean-Félix et Sherlock",
                    "description": "Fête de Jean-Félix et Sherlock",
                    "startDate": "2025-03-22T12:04:00Z",
                    "endDate": "2025-03-22T14:34:00Z",
                    "userId": 1,
                    "projectId": 1,
                    "categoryId": 1
                },
                "reponse": "L'activité a bien été ajoutée à la base de données."
            }
        }
    },
    addActivityFailEndDateBeforeStartDate: {
        url: '*/**/activity',
        method: 'POST',
        response: {
            status: 400,
            json: {
                "errors": [
                    {
                        "field": "startDate",
                        "message": "La date de début doit être avant la date de fin"
                    }
                ]
            }
        }
    },
    getAllActivitySuccess: {
        url: '*/**/activities',
        method: 'GET',
        response: {
            status: 200,
            json: {"activities":[
                {"id":1,"name":"A","description":"aaa","billable":false,"startDate":"2025-02-05T05:25:47-05:00",
                 "endDate":"2025-02-05T07:25:47-05:00","userId":1,"projectId":1,"categoryId":1},
                {"id":2,"name":"asd","description":"asd","billable":false,"startDate":"2025-04-02T08:00:00-04:00",
                 "endDate":"2025-04-02T15:00:00-04:00","userId":1,"projectId":1,"categoryId":1},{"id":3,"name":"new!","description":"new!!","billable":false,"startDate":"2025-04-03T08:00:00-04:00","endDate":"2025-04-03T11:00:00-04:00","userId":1,"projectId":1,"categoryId":1},{"id":4,"name":"asd","description":"asd","billable":false,"startDate":"2025-04-03T08:00:00-04:00","endDate":"2025-04-03T11:00:00-04:00","userId":1,"projectId":1,"categoryId":1},
                {"id":5,"name":"new!","description":"new!","billable":false,"startDate":"2025-04-03T08:00:00-04:00",
                 "endDate":"2025-04-03T11:00:00-04:00","userId":1,"projectId":1,"categoryId":1}]
            }
        }
    },
    getAllActivityEmpty: {
        url: '*/**/activities',
        method: 'GET',
        response: {
            status: 200,
            json: {"activities":[]}
        }
    },
} satisfies Record<string, MockConfig>;