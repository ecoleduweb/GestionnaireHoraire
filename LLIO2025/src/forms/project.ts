import { ProjectStatus } from "$lib/types/enums";
import type { CreateProject } from "../Models";

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