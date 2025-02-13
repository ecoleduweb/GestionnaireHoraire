import type { Task } from "../../Models";

export const taskTemplate = {
    generate: (): Task => ({
        name: "",
        description: "",
        state: "à faire",
        billable: false,
        userId: null,
        projectId: null,
        categoryId: null,
        startDateTime: new Date(),
        endDateTime: new Date(),
    })
};