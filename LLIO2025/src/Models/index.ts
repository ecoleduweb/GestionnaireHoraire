import { bool } from "yup";
import { ProjectStatus } from "../lib/types/enums";

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
}

export interface Project {
    id: number;
    name: string;
}

export interface Category {
    id: number;
    name: string;
}

export interface ActivityUpdateResponse {
    updated_activity: Activity;
}

export interface Project {
    name: string;
    description: string;
    status: ProjectStatus;
    billable: boolean;
}