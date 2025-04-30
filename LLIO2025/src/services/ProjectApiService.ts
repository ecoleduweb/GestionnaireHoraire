// ProjectApiService.ts
import { ProjectStatus } from "$lib/types/enums";
import type { CreateProject } from "../Models/index";
import { POST } from "../ts/server";

interface ProjectResponse {
  project: CreateProject;
}

const createProject = async (project: CreateProject): Promise<CreateProject> => {
  if (project.status === undefined) {
    project.status = ProjectStatus.InProgress;
  }
  try {
    const response = await POST<CreateProject, ProjectResponse>("/project", project);
    return response.project;
  } catch (error) {
    console.error("Erreur lors de la création du projet:", error);
    throw error;
  }
};

export const ProjectApiService = {
  createProject
};