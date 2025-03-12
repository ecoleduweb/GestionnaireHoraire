import { PUBLIC_BASE_URL } from "$env/static/public";
import type { Task } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

let tasks: Task[];

interface FormTask extends Omit<Task, 'startDate' | 'endDate'> {
    startDate: string;
    endDate: string;
}

const transformTasksDates = (task: FormTask): Task => ({
    ...task,
    startDate: new Date(task.startDate),
    endDate: new Date(task.endDate)
});

const createTask = async (task: Task): Promise<Task> => {
        
    // Préparation des données pour l'API
    const taskForApi = {
        name: task.name,
        description: task.description || "",
        billable: task.billable || false,
        startDate: task.startDate.toISOString(),
        endDate: task.endDate.toISOString(),
        userId: task.userId || 1,
        projectId: task.projectId || 1,
        categoryId: task.categoryId || 1
    };
    
    try {
        const response = await POST<typeof taskForApi, {task: Task}>("/tasks", taskForApi);
        return {
            ...response.task,
            startDate: new Date(response.task.startDate),
            endDate: new Date(response.task.endDate)
        };
    } catch (error) {
        console.error("Erreur lors de la création de tâche:", error);
        throw error;
    }
};

const updateTask = async (task: Task): Promise<Task> => {
    const taskForApi: FormTask = {
        ...task,
        startDate: task.startDate.toISOString(),
        endDate: task.endDate.toISOString()
    };

    const response = await PUT<FormTask, FormTask>(`/tasks/${task.id}`, taskForApi);
    return transformTasksDates(response);
};

const deleteTask = async (id: number): Promise<void> => {
    await DELETE(`/tasks/${id}`);
};

const getTasks = async () => {
    try {
        const response = await GET<FormTask[]>("/tasks");
        const taskData = response.map(transformTasksDates);
        tasks = taskData;
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