import { ProjectStatus } from "$lib/types/enums";
import type { CreateProject } from "../Models";

export const projectTemplate = {
    generate: (): CreateProject => ({
        name: "",
        description: "",
        status: ProjectStatus.InProgress,
        billable: false,
        manager_id: 0
    })
};