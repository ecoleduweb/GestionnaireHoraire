import { PUBLIC_BASE_URL } from "$env/static/public";
import type { Task } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

let tasks: Task[];

interface RawTask extends Omit<Task, 'startDate' | 'endDate'> {
    startDate: string | Date;
    endDate: string | Date;
}

const transformTasksDates = (task: RawTask): Task => ({
    ...task,
    startDate: new Date(task.startDate),
    endDate: new Date(task.endDate)
});

const transformTasksArray = (tasks: RawTask[]): Task[] => 
    tasks.map(transformTasksDates);

const getTaskData = async (): Promise<Task[]> => {
    const response = await GET<RawTask[]>("/tasks");
    return transformTasksArray(response);
};

const createTask = async (task: Task): Promise<Task> => {
    // Vérifier que les dates sont valides
    const startDate = task.startDate instanceof Date && !isNaN(task.startDate.getTime())
      ? task.startDate
      : new Date();
      
    const endDate = task.endDate instanceof Date && !isNaN(task.endDate.getTime())
      ? task.endDate
      : new Date(startDate.getTime() + 60 * 60 * 1000);
    
    // Préparation des données pour l'API
    const taskForApi = {
      name: task.name,
      description: task.description || "",
      billable: task.billable || false,
      start_date: startDate.toISOString(),
      end_date: endDate.toISOString(),
      user_id: task.userId || 1,
      project_id: task.projectId || 1,
      category_id: task.categoryId || 1,
    };
    
    try {
        const response = await POST<typeof taskForApi, { reponse: string, task: any }>("/tasks", taskForApi);
        
        // Extraire l'objet task de la réponse
        let taskData: any;
        
        if (response && response.task) {
            taskData = response.task;
        } else if (response && typeof response === 'object') {
            // Si la réponse est un objet, utiliser directement
            taskData = response;
        } else {
            // Remplacé par console.error qui peut être utile pour le débogage
            // mais peut aussi être supprimé si vous voulez éliminer tous les messages
            console.error('Format de réponse inattendu:', response);
            // Utiliser les données envoyées comme fallback
            taskData = {
                ...taskForApi,
                id: Date.now() // ID temporaire
            };
        }
        
        // Conversion du format snake_case vers camelCase si nécessaire
        const newTask: Task = {
            id: taskData.id || undefined,
            name: taskData.name,
            description: taskData.description || "",
            billable: taskData.billable || false,
            startDate: new Date(taskData.start_date || taskData.startDate),
            endDate: new Date(taskData.end_date || taskData.endDate),
            userId: task.userId || 1,
            projectId: task.projectId || 1,
            categoryId: task.categoryId || 1
        };
        
        return newTask;
    } catch (error) {
        throw error;
    }
};

const updateTask = async (task: Task): Promise<Task> => {
    const taskForApi = {
        ...task,
        startDate: task.startDate.toISOString(),
        endDate: task.endDate.toISOString()
    };

    const response = await PUT<typeof taskForApi, RawTask>(`/tasks/${task.id}`, taskForApi);
    return transformTasksDates(response);
};

const deleteTask = async (id: number): Promise<void> => {
    await DELETE(`/tasks/${id}`);
};

export const fetchTasks = async () => {
    try {
        const taskData = await getTaskData();
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