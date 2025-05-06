import { ProjectStatus } from "../../../LLIO2025/src/lib/types/enums";
import { MockConfig } from "../types";

export const projectMocks = {
    getByIdFail: {
        url: '*/**/project/*',
        method: 'GET',
        response: {
            status: 404,
            json: {"error":"Le(La) projet n'a pas été trouvé(e)"}
        }
    },
    getByIdSuccess: {
        url: '*/**/project/*',
        method: 'GET',
        response: {
            status: 200,
            json: {"project":
                {"id": 1,
                    "name": "Projet Sherlock",
                    "description": "Développement d'application de suivi de temps",
                    "status": ProjectStatus.InProgress,
                    "billable": true,
                    "userId": 1
                }
            }
        }
    },
    addProjectSuccess: {
        url: '*/**/project',
        method: 'POST',
        response: {
            status: 201,
            json: {
                "project": {
                    "id": 1,
                    "name": "Projet Sherlock",
                    "description": "Développement d'application de suivi de temps",
                    "status": ProjectStatus.InProgress,
                    "billable": true,
                    "userId": 1
                },
                "reponse": "Le projet a bien été ajouté à la base de données."
            }
        }
    },
    addProjectSuccessNoNameNoDescription: {
        url: '*/**/project',
        method: 'POST',
        response: {
            status: 201,
            json: {
                "project": {
                    "id": 1,
                    "name": "",
                    "description": "",
                    "status": ProjectStatus.InProgress,
                    "billable": false,
                    "userId": 1
                },
                "reponse": "Le projet a bien été ajouté à la base de données."
            }
        }
    },
    addProjectFailValidation: {
        url: '*/**/project',
        method: 'POST',
        response: {
            status: 400,
            json: {
                "errors": [
                    {
                        "field": "name",
                        "message": "Le nom du projet est requis"
                    }
                ]
            }
        }
    },
    getAllProjectsSuccess: {
        url: '*/**/projects',
        method: 'GET',
        response: {
            status: 200,
            json: {"projets":[ // Changé de "projects" à "projets"
                {"id":1,"name":"Projet Sherlock","description":"Développement d'application de suivi de temps",
                 "status": ProjectStatus.InProgress,"billable":true,"userId":1},
                {"id":2,"name":"Projet Holmes","description":"Conception de l'interface utilisateur",
                 "status": ProjectStatus.InProgress,"billable":true,"userId":1},
                {"id":3,"name":"Projet Watson","description":"Tests d'intégration",
                 "status": ProjectStatus.Finished,"billable":true,"userId":1},
                {"id":4,"name":"Projet Moriarty","description":"Refactorisation du backend",
                 "status": ProjectStatus.NotStarted,"billable":false,"userId":1},
                {"id":5,"name":"Projet Lestrade","description":"Documentation technique",
                 "status": ProjectStatus.InProgress,"billable":false,"userId":1}]
            }
        }
    },
    getAllProjectsEmpty: {
        url: '*/**/projects',
        method: 'GET',
        response: {
            status: 200,
            json: {"projets":[]}
        }
    },
    updateProjectSuccess: {
        url: '*/**/project/*',
        method: 'PUT',
        response: {
            status: 200,
            json: {
                "project": {
                    "id": 1,
                    "name": "Projet Sherlock - Version 2",
                    "description": "Développement d'application de suivi de temps - Phase 2",
                    "status": ProjectStatus.Finished,
                    "billable": true,
                    "userId": 1
                },
                "reponse": "Le projet a bien été mis à jour."
            }
        }
    },
    updateProjectFailValidation: {
        url: '*/**/project/*',
        method: 'PUT',
        response: {
            status: 400,
            json: {
                "errors": [
                    {
                        "field": "name",
                        "message": "Le nom du projet est requis"
                    }
                ]
            }
        }
    },
    deleteProjectSuccess: {
        url: '*/**/project/*',
        method: 'DELETE',
        response: {
            status: 200,
            json: {
                "reponse": "Le projet a bien été supprimé de la base de données."
            }
        }
    },
    deleteProjectFail: {
        url: '*/**/project/*',
        method: 'DELETE',
        response: {
            status: 400,
            json: {
                "error": "Impossible de supprimer un projet avec des activités associées."
            }
        }
    },
    getProjectsByUserSuccess: {
        url: '*/**/projects/me',
        method: 'GET',
        response: {
            status: 200,
            json: {"projets":[
                {"id":1,"name":"Projet Sherlock","description":"Développement d'application de suivi de temps",
                 "status": ProjectStatus.InProgress,"billable":true,"userId":1},
                {"id":3,"name":"Projet Watson","description":"Tests d'intégration",
                 "status": ProjectStatus.Finished,"billable":true,"userId":1},
                {"id":5,"name":"Projet Lestrade","description":"Documentation technique",
                 "status": ProjectStatus.InProgress,"billable":false,"userId":1}]
            }
        }
    },
    getProjectsByUserEmpty: {
        url: '*/**/projects/me',
        method: 'GET',
        response: {
            status: 200,
            json: {"projets":[]}
        }
    },
    getProjectsFilterByStatusSuccess: {
        url: '*/**/projects?status=InProgress',
        method: 'GET',
        response: {
            status: 200,
            json: {"projets":[
                {"id":1,"name":"Projet Sherlock","description":"Développement d'application de suivi de temps",
                 "status": ProjectStatus.InProgress,"billable":true,"userId":1},
                {"id":2,"name":"Projet Holmes","description":"Conception de l'interface utilisateur",
                 "status": ProjectStatus.InProgress,"billable":true,"userId":1},
                {"id":5,"name":"Projet Lestrade","description":"Documentation technique",
                 "status": ProjectStatus.InProgress,"billable":false,"userId":1}]
            }
        }
    }
} satisfies Record<string, MockConfig>;