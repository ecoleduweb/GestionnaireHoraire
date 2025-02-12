export interface Task {
    id?: number;
    name: string;
    description: string;
    state: "à faire" | "en cours" | "terminé";
    billable: boolean;
    startDate: Date;
    endDate: Date;
    userId: number;
    projectId: number;
    categoryId: number;
    startTime?: string;
    endTime?: string;
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