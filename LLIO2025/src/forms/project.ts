import type { CreateProject } from "../Models";

export const projectTemplate = {
    generate: (): CreateProject => {
        return {
            name: '',
            description: '',
import type { CreateProject } from "../Models";
import { ProjectStatus } from "../lib/types/enums";

export const projectTemplate = {
    generate: (): CreateProject => {
        return {
            name: '',
            description: '',
            status: ProjectStatus.InProgress,
            billable: false,
        };
    },
} as const
            billable: false,
        };
    },
} as const