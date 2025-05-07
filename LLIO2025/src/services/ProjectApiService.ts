// ProjectApiService.ts
import { ProjectStatus } from "$lib/types/enums";
import type { CreateProject, Project, UpdateProject } from "../Models/index";
import { GET, POST, PUT } from "../ts/server";

interface ProjectResponse {
  project: CreateProject;
  projets: Project[];
}

// Récupérer tous les projets
const getAllProjects = async (): Promise<Project[]> => {
  try {
    const response = await GET<ProjectResponse>("/projects");
    return response.projets; 
  } catch (error) {
    console.error("Erreur lors de la récupération des projets:", error);
    return [];
  }
};

const createProject = async (project: CreateProject): Promise<CreateProject> => {
  if (project.status === undefined) {
    project.status = ProjectStatus.InProgress;
  }
  try {
    const response = await POST<CreateProject, ProjectResponse>("/project", project);
    return response.project;
  } catch (error) {
    console.error("Erreur lors de la création du projet:", error);
    throw new Error("Échec de la création du projet. Veuillez réessayer.");
  }
};

const updateProject = async (project: UpdateProject): Promise<CreateProject> => {
  if (project.status === undefined) {
    project.status = ProjectStatus.InProgress;
  }
  try {
    const response = await PUT<UpdateProject, ProjectResponse>("/project", project);
    return response.project;
  } catch (error) {
    console.error("Erreur lors de la création du projet:", error);
    throw new Error("Échec de la création du projet. Veuillez réessayer.");
  }
};

const getProjects = async(): Promise<Project[]> => {
  try {
    const response = await GET<{projects: Project[]}>("/projects");
    return response.projects;
  } catch (error) {
    console.error("Erreur lors de la récupération des projets:", error);
    throw new Error("Échec de la récupération des projets. Veuillez réessayer.");
  }
}

export const ProjectApiService = {
  createProject,
  updateProject,
  getProjects,
};
