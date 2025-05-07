import { bool } from "yup";
import { ProjectStatus, UserRole } from "../lib/types/enums";

export interface Activity {
    id?: number;
    name: string;
    description: string;
    startDate: Date;
    endDate: Date;
    userId: number;
    projectId: number;
    categoryId: number;
}

export interface User {
    id: number;
    name: string;
}

export interface UserInfo {
    firstName: string;
    lastName: string;
    role : UserRole;
}

export interface User {
    id: number;
    first_name: string;
    last_name: string;
    email: string;
    role: UserRole;
}

export interface Employee {
    name: string;
    categories: Category[];
}

export interface Category {
    name: string;
    timeSpent: number;
    timeEstimated: number;
}

export interface ActivityUpdateResponse {
    updated_activity: Activity;
}

/*************** Project ***************/
// Base interface avec les propriétés communes
export interface ProjectBase {
    name: string;
    description: string;
    manager_id: number;
    billable: boolean;
    status?: ProjectStatus;
}

// Interface pour la création
export interface UpdateProject extends ProjectBase {
    id: number;
}

// Interface complète pour un projet existant
export interface Project extends ProjectBase {
    id: number;
    color: string;
    lead: string;
    coLeads: string[];
    employees: Employee[];
    totalTimeEstimated: number;
    totalTimeRemaining: number;
    totalTimeSpent: number;
    isArchived?: boolean;
    status: ProjectStatus; // Devenu obligatoire
}