import type { Task } from "../Models";

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
    }),

    states: ["à faire", "en cours", "terminée"],
    time: {
        hours: Array.from({ length: 24 }, (_, i) => i.toString().padStart(2, "0")),
        minutes: ["00", "15", "30", "45"]
    }
} as const;