export interface Activity {
    id?: number;
    name: string;
    description: string;
    state: "à faire" | "en cours" | "terminé";
    startDateTime: Date;
    endDateTime: Date;
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