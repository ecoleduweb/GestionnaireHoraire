export const TASK_CONFIG = {
    states: ["à faire", "en cours", "terminée"],
    time: {
        hours: Array.from({ length: 24 }, (_, i) => i.toString().padStart(2, "0")),
        minutes: ["00", "15", "30", "45"]
    },
    defaultTask: {
        name: "",
        description: "",
        state: "à faire",
        billable: false,
        userId: undefined,
        projectId: undefined,
        categoryId: undefined,
    }
} as const;