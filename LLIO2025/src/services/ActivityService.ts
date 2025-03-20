import type { Activity } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

interface FormActivity extends Omit<Activity, 'startDate' | 'endDate'> {
    startDate: string;
    endDate: string;
}

const transformTasksDates = (activity: FormTask): Activity => ({
    ...activity,
    startDate: new Date(activity.startDate),
    endDate: new Date(activity.endDate)
});

const createTask = async (activity: Activity): Promise<Activity> => {
        
    // Préparation des données pour l'API
    const taskForApi = {
        name: activity.name,
        description: activity.description || "",
        startDate: activity.startDate.toISOString(),
        endDate: activity.endDate.toISOString(),
        userId: activity.userId || 1,
        projectId: activity.projectId || 1,
        categoryId: activity.categoryId || 1
    };
    
    try {
        const response = await POST<typeof taskForApi, {activity: Activity}>("/activities", taskForApi);
        return {
            ...response.activity,
            startDate: new Date(response.activity.startDate),
            endDate: new Date(response.activity.endDate)
        };
    } catch (error) {
        console.error("Erreur lors de la création de tâche:", error);
        throw error;
    }
};

const updateTask = async (activity: Activity): Promise<Activity> => {
    const taskForApi: FormTask = {
        ...activity,
        startDate: activity.startDate.toISOString(),
        endDate: activity.endDate.toISOString()
    };

    const response = await PUT<FormTask, FormTask>(`/activities/${activity.id}`, taskForApi);
    return transformTasksDates(response);
};

const deleteTask = async (id: number): Promise<void> => {
    await DELETE(`/activities/${id}`);
};

const getTasks = async () => {
    try {
        const response = await GET<FormTask[]>("/activities");
        const taskData = response.map(transformTasksDates);
        return taskData;
    } catch (error) {
        // Gardé pour le débogage, mais peut être supprimé si nécessaire
        console.error("Erreur lors de la récupération des tâches", error);
        throw error;
    }
};

export const TaskApiService = {
    getTasks,
    createTask,
    updateTask,
    deleteTask
}