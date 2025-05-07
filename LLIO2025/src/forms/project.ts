import { ProjectStatus } from "$lib/types/enums";
import type { ProjectBase } from "../Models";

export const projectTemplate = {
    generate: (): ProjectBase => ({
        name: "",
        description: "",
        status: ProjectStatus.InProgress,
        billable: false,
        manager_id: 0
    })
};