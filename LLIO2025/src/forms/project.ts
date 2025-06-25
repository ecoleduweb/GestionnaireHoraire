import { ProjectStatus } from "$lib/types/enums";
import type { ProjectBase } from "../Models";

export const projectTemplate = {
    generate: (): ProjectBase => ({
        uniqueId: "",
        name: "",
        status: ProjectStatus.InProgress,
        billable: false,
        managerId: 0,
        estimatedHours:0
    })
};