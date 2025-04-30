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
}

export interface Project {
    id: string;
    name: string;
    color: string;
    lead: string;
    coLeads: string[];
    employees: Employee[];
    totalTimeSpent: number;
    isArchived?: boolean;
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

export interface ActivityUpdateResponse 
{
    updated_activity: Activity;
}