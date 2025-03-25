import type { Task } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

let tasks: Task[];

interface FormTask extends Omit<Task, 'startDate' | 'endDate'> {
    startDate: string;
    endDate: string;
    updated_task?: Task;
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
    if (!task) {
        console.error("La tâche est undefined");
        throw new Error("La tâche ou son ID est manquant");
    }

    if (!task.id) {
        console.error("L'ID de la tâche est manquant", task);
        throw new Error("L'ID de la tâche est manquant");
    }

    // Préparation des données pour l'API - ATTENTION: UTILISER L'URL RACINE /tasks
    const taskForApi = {
        id: task.id,
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
        const response = await PUT<typeof taskForApi, {updated_task: Task}>("/tasks", taskForApi);
        
        if (response && response.updated_task) {
            // Extraire la tâche mise à jour de l'objet de réponse
            const updatedTaskData = response.updated_task;
            
            // Convertir les dates
            return {
                ...updatedTaskData,
                startDate: new Date(updatedTaskData.startDate),
                endDate: new Date(updatedTaskData.endDate)
            };
        } else {
            return task;
        }
    } catch (error) {
        console.error("Erreur lors de la mise à jour de la tâche:", error);
        throw error;
    }
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