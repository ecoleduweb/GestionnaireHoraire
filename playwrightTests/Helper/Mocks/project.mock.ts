import { MockConfig } from "../types";

export const projectMocks = {
  getDetailedProjectsSuccess: {
    url: "/projects/detailed",
    method: "GET",
    response: {
      status: 200,
      json: {
        projects: [
          {
            id: 1,
            name: "AT-123",
            description: "Nommer le projet",
            color: "blue",
            lead: "Katell Arnault de la Ménardière",
            isArchived: false,
            coLeads: ["Jean-François Jasmin"],
            employees: [
              {
                name: "Katell Arnault de la Ménardière",
                categories: [
                  { name: "Développement", timeSpent: 30, timeEstimated: 25 },
                  { name: "Graphisme", timeSpent: 15, timeEstimated: 30 },
                ],
              },
              {
                name: "Jean-François Jasmin",
                categories: [
                  { name: "Développement", timeSpent: 20, timeEstimated: 50 },
                ],
              },
              {
                name: "Laure Desjardins",
                categories: [],
              },
            ],
            totalTimeEstimated: 20,
            totalTimeRemaining: -45,
            totalTimeSpent: 65,
          },
          {
            id: 2,
            name: "AT-456",
            description: "Le projet de Marie Amélie",
            color: "pink",
            lead: "Katell Arnault de la Ménardière",
            isArchived: false,
            coLeads: ["Marie Amélie Dubé", "Jimmy Paquet-Cormier"],
            employees: [
              {
                name: "Katell Arnault de la Ménardière",
                categories: [
                  { name: "Développement", timeSpent: 12, timeEstimated: 15 },
                ],
              },
              {
                name: "Marie Amélie Dubé",
                categories: [
                  { name: "Graphisme", timeSpent: 18, timeEstimated: 20 },
                ],
              },
              {
                name: "Ariane Dionne-Santerre",
                categories: [
                  { name: "Rédaction", timeSpent: 8, timeEstimated: 10 },
                ],
              },
              {
                name: "Jimmy Paquet-Cormier",
                categories: [
                  { name: "Développement", timeSpent: 14, timeEstimated: 15 },
                ],
              },
            ],
            totalTimeEstimated: 20,
            totalTimeRemaining: -32,
            totalTimeSpent: 52,
          },
          {
            id: 3,
            name: "FO-115",
            description: "Graphisme 101",
            color: "yellow",
            lead: "Katell Arnault de la Ménardière",
            isArchived: false,
            coLeads: [],
            employees: [
              {
                name: "Katell Arnault de la Ménardière",
                categories: [
                  { name: "Gestion", timeSpent: 5, timeEstimated: 8 },
                ],
              },
              {
                name: "Annie Côté",
                categories: [
                  { name: "Graphisme", timeSpent: 22, timeEstimated: 25 },
                ],
              },
            ],
            totalTimeEstimated: 20,
            totalTimeRemaining: -7,
            totalTimeSpent: 27,
          },
          {
            id: 4,
            name: "RA-224",
            description: "Noisette Steve",
            color: "red",
            lead: "Steve Joncoux",
            isArchived: false,
            coLeads: [
              "Katell Arnault de la Ménardière",
              "Marjolaine Poirier",
              "Jimmy Paquet-Cormier",
            ],
            employees: [
              {
                name: "Steve Joncoux",
                categories: [
                  { name: "Développement", timeSpent: 16, timeEstimated: 20 },
                ],
              },
              {
                name: "Katell Arnault de la Ménardière",
                categories: [
                  { name: "Gestion", timeSpent: 7, timeEstimated: 10 },
                ],
              },
              {
                name: "Marjolaine Poirier",
                categories: [
                  { name: "Rédaction", timeSpent: 12, timeEstimated: 15 },
                ],
              },
              {
                name: "Jimmy Paquet-Cormier",
                categories: [
                  { name: "Développement", timeSpent: 9, timeEstimated: 12 },
                ],
              },
              {
                name: "Marie Amélie Dubé",
                categories: [
                  { name: "Graphisme", timeSpent: 11, timeEstimated: 10 },
                ],
              },
            ],
            totalTimeEstimated: 20,
            totalTimeRemaining: -35,
            totalTimeSpent: 55,
          },
          {
            id: 5,
            name: "AT-789",
            description: "Nommer le projet",
            color: "blue",
            lead: "Katell Arnault de la Ménardière",
            isArchived: true,
            coLeads: ["Jean-François Jasmin"],
            employees: [
              {
                name: "Katell Arnault de la Ménardière",
                categories: [
                  { name: "Développement", timeSpent: 30, timeEstimated: 25 },
                  { name: "Graphisme", timeSpent: 15, timeEstimated: 30 },
                ],
              },
              {
                name: "Jean-François Jasmin",
                categories: [
                  { name: "Développement", timeSpent: 20, timeEstimated: 50 },
                ],
              },
              {
                name: "Laure Desjardins",
                categories: [],
              },
            ],
            totalTimeEstimated: 20,
            totalTimeRemaining: -45,
            totalTimeSpent: 65,
          },
          {
            id: 6,
            name: "AT-987",
            description: "Le projet de Marie Amélie",
            color: "pink",
            lead: "Katell Arnault de la Ménardière",
            isArchived: true,
            coLeads: ["Marie Amélie Dubé", "Jimmy Paquet-Cormier"],
            employees: [
              {
                name: "Katell Arnault de la Ménardière",
                categories: [
                  { name: "Développement", timeSpent: 12, timeEstimated: 15 },
                ],
              },
              {
                name: "Marie Amélie Dubé",
                categories: [
                  { name: "Graphisme", timeSpent: 18, timeEstimated: 20 },
                ],
              },
              {
                name: "Ariane Dionne-Santerre",
                categories: [
                  { name: "Rédaction", timeSpent: 8, timeEstimated: 10 },
                ],
              },
              {
                name: "Jimmy Paquet-Cormier",
                categories: [
                  { name: "Développement", timeSpent: 14, timeEstimated: 15 },
                ],
              },
            ],
            totalTimeEstimated: 20,
            totalTimeRemaining: -32,
            totalTimeSpent: 52,
          },
          {
            id: 7,
            name: "FO-789",
            description: "Graphisme 101",
            color: "yellow",
            lead: "Katell Arnault de la Ménardière",
            isArchived: true,
            coLeads: [],
            employees: [
              {
                name: "Katell Arnault de la Ménardière",
                categories: [
                  { name: "Gestion", timeSpent: 5, timeEstimated: 8 },
                ],
              },
              {
                name: "Annie Côté",
                categories: [
                  { name: "Graphisme", timeSpent: 22, timeEstimated: 25 },
                ],
              },
            ],
            totalTimeEstimated: 20,
            totalTimeRemaining: -7,
            totalTimeSpent: 27,
          },
          {
            id: 8,
            name: "RA-987",
            description: "Noisette Steve",
            color: "red",
            lead: "Steve Joncoux",
            isArchived: true,
            coLeads: [
              "Katell Arnault de la Ménardière",
              "Marjolaine Poirier",
              "Jimmy Paquet-Cormier",
            ],
            employees: [
              {
                name: "Steve Joncoux",
                categories: [
                  { name: "Développement", timeSpent: 16, timeEstimated: 20 },
                ],
              },
              {
                name: "Katell Arnault de la Ménardière",
                categories: [
                  { name: "Gestion", timeSpent: 7, timeEstimated: 10 },
                ],
              },
              {
                name: "Marjolaine Poirier",
                categories: [
                  { name: "Rédaction", timeSpent: 12, timeEstimated: 15 },
                ],
              },
              {
                name: "Jimmy Paquet-Cormier",
                categories: [
                  { name: "Développement", timeSpent: 9, timeEstimated: 12 },
                ],
              },
              {
                name: "Marie Amélie Dubé",
                categories: [
                  { name: "Graphisme", timeSpent: 11, timeEstimated: 10 },
                ],
              },
            ],
            totalTimeEstimated: 20,
            totalTimeRemaining: -35,
            totalTimeSpent: 55,
          },
        ],
      },
    },
  },
  getProjectsListSuccess: {
    url: "/projects",
    method: "GET",
    response: {
      status: 200,
      json: {
        projects: [
          {
            id: 1,
            managerId: 1,
            name: "Projet sous-sol",
            description: "Nommer le projet",
            status: 1,
            billable: false,
            activities: [],
            categories: [],
            createdAt: "2025-03-22T08:00:00",
            updatedAt: "2025-03-22T08:00:00",
            end_at: null,
          },
        ],
      },
    },
  },
  addProjectSuccess: {
    url: "/project",
    method: "POST",
    response: {
      status: 201,
      json: {
          "project": {
              "id": 6,
              "managerId": 3,
              "name": "Jérémie Lapointe",
              "description": "das",
              "status": 0,
              "billable": true,
              "activities": [],
              "categories": [],
              "createdAt": "2025-03-23T15:07:14.991-04:00",
              "updatedAt": "2025-03-23T15:07:14.991-04:00",
              "endAt": "0001-01-01T00:00:00Z"
          },
          "response": "Le projet a bien été ajouté à la base de données"
      },
    },
  },
  getDetailedProjectsByUserSuccess:{
    url: "/projects/me/detailed",
    method: "GET",
    response: {
      status: 200,
      json: {
        "projects": [
          {
            "billable": false,
            "coLeads": [],
            "createdAt": "2025-06-12T14:00:37-04:00",
            "employees": [
              {
                "categories": [
                  {
                    "name": "Par défaut",
                    "timeEstimated": 0,
                    "timeSpent": 76
                  }
                ],
                "name": " "
              }
            ],
            "id": 2,
            "isArchived": false,
            "lead": "Usager test",
            "managerId": 1,
            "name": "un admin 2 est charge",
            "totalTimeEstimated": 0,
            "totalTimeRemaining": -76,
            "totalTimeSpent": 76,
            "uniqueId": "skiidi",
            "updatedAt": "2025-06-12T14:00:37-04:00"
          },
          {
            "billable": false,
            "coLeads": [],
            "createdAt": "2025-06-12T14:03:19-04:00",
            "employees": [
              {
                "categories": [
                  {
                    "name": "Par défaut",
                    "timeEstimated": 0,
                    "timeSpent": 13
                  }
                ],
                "name": " "
              }
            ],
            "id": 3,
            "isArchived": false,
            "lead": "Usager test",
            "managerId": 1,
            "name": "adfasfd",
            "totalTimeEstimated": 0,
            "totalTimeRemaining": -13,
            "totalTimeSpent": 13,
            "uniqueId": "User vincent.antoineccmmm",
            "updatedAt": "2025-06-12T14:03:19-04:00"
          },
          {
            "billable": true,
            "coLeads": [],
            "createdAt": "2025-06-23T16:27:08-04:00",
            "employees": [
              {
                "categories": [
                  {
                    "name": "Par défaut",
                    "timeEstimated": 0,
                    "timeSpent": 13.5
                  }
                ],
                "name": " "
              }
            ],
            "id": 8,
            "isArchived": false,
            "lead": "Usager test",
            "managerId": 1,
            "name": "projet ! apre smigration ",
            "totalTimeEstimated": 0,
            "totalTimeRemaining": -13.5,
            "totalTimeSpent": 13.5,
            "uniqueId": "migr-1324",
            "updatedAt": "2025-06-23T16:27:08-04:00"
          },
          {
            "billable": false,
            "coLeads": [],
            "createdAt": "2025-06-23T16:43:12-04:00",
            "employees": [
              {
                "categories": [
                  {
                    "name": "Par défaut",
                    "timeEstimated": 0,
                    "timeSpent": 24
                  }
                ],
                "name": " "
              }
            ],
            "id": 11,
            "isArchived": false,
            "lead": "Toumani Camara",
            "managerId": 4,
            "name": "new commut",
            "totalTimeEstimated": 0,
            "totalTimeRemaining": -24,
            "totalTimeSpent": 24,
            "uniqueId": "new-1",
            "updatedAt": "2025-06-23T16:43:12-04:00"
          },
          {
            "billable": false,
            "coLeads": [],
            "createdAt": "2025-06-12T15:06:22-04:00",
            "employees": [
              {
                "categories": [
                  {
                    "name": "Par défaut",
                    "timeEstimated": 0,
                    "timeSpent": 1
                  }
                ],
                "name": " "
              }
            ],
            "id": 4,
            "isArchived": false,
            "lead": "Usager test",
            "managerId": 1,
            "name": "kkkkkk",
            "totalTimeEstimated": 0,
            "totalTimeRemaining": -1,
            "totalTimeSpent": 1,
            "uniqueId": "ad",
            "updatedAt": "2025-06-12T15:06:22-04:00"
          }
        ]
      },
    },
  }
} satisfies Record<string, MockConfig>;
