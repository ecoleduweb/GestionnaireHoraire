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
    projectName: string;
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

export interface Project {
    id: number;
    name: string;
    description: string;
    color: string;
    lead: string;
    coLeads: string[];
    employees: Employee[];
    totalTimeEstimated: number;
    totalTimeRemaining: number;
    totalTimeSpent: number;
    isArchived?: boolean;
}

export interface Employee {
    name: string;
    categories: Category[];
}

export interface Category {
    id: number;
    name: string;
    description?: string;
    billable?: boolean;
    timeSpent: number;
    timeEstimated: number;
}

export interface ActivityUpdateResponse {
    updated_activity: Activity;
}

export interface CreateProject {
    name: string;
    description: string;
    status: ProjectStatus;
    billable: boolean;
}

export interface UpdateProject {
    name: string;
    status: ProjectStatus;
    billable: boolean;
}