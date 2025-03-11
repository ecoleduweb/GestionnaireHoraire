import { PUBLIC_BASE_URL } from "$env/static/public";
import type { Task } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

let tasks: Task[];

interface RawTask extends Omit<Task, 'startDate' | 'endDate'| 'userId' | 'projectId' | 'categoryId'> {
    start_date: string | Date;
    end_date: string | Date;
    user_id: number;
    project_id: number;
    category_id: number;
}

const transformTasksDates = (task: RawTask): Task => ({
    ...task,
    startDate: new Date(task.start_date),
    endDate: new Date(task.end_date),
    userId: task.user_id,
    projectId: task.project_id,
    categoryId: task.category_id
});

const createTask = async (task: Task): Promise<Task> => {
        
    // Préparation des données pour l'API
    const taskForApi: RawTask = {
      name: task.name,
      description: task.description || "",
      billable: task.billable || false,
      start_date: task.startDate.toISOString(),
      end_date: task.endDate.toISOString(),
      user_id: task.userId || 1,
      project_id: task.projectId || 1,
      category_id: task.categoryId || 1,
    };
    
    try {
        const response = await POST<typeof taskForApi, {task: RawTask}>("/tasks", taskForApi);
        
        // Conversion directe de l'objet task
        const taskData = response.task;
        
        return {
            id: taskData.id,
            name: taskData.name,
            description: taskData.description,
            billable: taskData.billable,
            startDate: new Date(taskData.start_date || taskData.start_date),
            endDate: new Date(taskData.end_date || taskData.end_date),
            userId: taskData.user_id || taskData.user_id,
            projectId: taskData.project_id || taskData.project_id,
            categoryId: taskData.category_id || taskData.category_id
        };
    } catch (error) {
        throw error;
    }
};

const updateTask = async (task: Task): Promise<Task> => {
    const taskForApi: RawTask = {
        ...task,
        start_date: task.startDate.toISOString(),
        end_date: task.endDate.toISOString(),
        user_id: task.userId,
        project_id: task.projectId,
        category_id: task.categoryId
    };

    const response = await PUT<RawTask, RawTask>(`/tasks/${task.id}`, taskForApi);
    return transformTasksDates(response);
};

const deleteTask = async (id: number): Promise<void> => {
    await DELETE(`/tasks/${id}`);
};

const fetchTasks = async () => {
    try {
        const response = await GET<RawTask[]>("/tasks");
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
    getTasks: fetchTasks,
    createTask,
    updateTask,
    deleteTask
}