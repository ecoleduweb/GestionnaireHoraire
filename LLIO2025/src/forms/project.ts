import type { CreateProject } from "../Models";

export const projectTemplate = {
    generate: (): CreateProject => {
        return {
            name: '',
            description: '',
            status: 0,
            billable: false,
        };
    },
} as const