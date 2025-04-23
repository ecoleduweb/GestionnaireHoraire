import type { Project } from "../Models/index";
import { GET } from "../ts/server";

interface ProjectsResponse {
  projets: Project[]; 
}

// Récupérer tous les projets
const getAllProjects = async (): Promise<Project[]> => {
  try {
    const response = await GET<ProjectsResponse>("/projects");
    return response.projets; 
  } catch (error) {
    console.error("Erreur lors de la récupération des projets:", error);
    return [];
  }
};

export const ProjectApiService = {
  getAllProjects
};