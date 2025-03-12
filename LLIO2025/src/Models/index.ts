export interface Task {
    id?: number;
    name: string;
    description: string;
    billable: boolean;
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

export interface Project {
    id: number;
    name: string;
}

export interface Category {
    id: number;
    name: string;
}