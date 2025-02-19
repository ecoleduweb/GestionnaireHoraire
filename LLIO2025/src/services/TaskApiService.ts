import type { Task } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";
import { writable } from "svelte/store";

let tasks: Task[];

interface RawTask extends Omit<Task, 'startDateTime' | 'endDateTime'> {
    startDateTime: string | Date;
    endDateTime: string | Date;
}

const transformTasksDates = (task: RawTask): Task => ({
    ...task,
    startDateTime: new Date(task.startDateTime),
    endDateTime: new Date(task.endDateTime)
});

const transformTasksArray = (tasks: Task[]): Task[] => 
    tasks.map(transformTasksDates);

const getTaskData = async (): Promise<Task[]> => {
    const response = await GET<Task[]>("/task/all");
    return transformTasksArray(response);
};

const createTask = async (task: Task): Promise<Task> => {
    const response = await POST<Task, Task>("/task", task);
    return transformTasksDates(response);
};

const updateTask = async (task: Task): Promise<Task> => {
    const response = await PUT<Task, Task>(`/task/${task.id}`, task);
    return transformTasksDates(response);
};

const deleteTask = async (id: number): Promise<void> => {
    await DELETE(`/task/${id}`);
};

export const fetchTasks = async () => {
    let savedData: string | null = null;
    let taskData: Task[];

    try {
        savedData = localStorage.getItem("Task");
    } catch (error) {
        console.error(`Erreur lors de l'accès au localStorage: ${error.message}`);
    }

    try {
        if (tasks?.length) {
            taskData = tasks;
        } else if (savedData) {
            try {
                const parsedData = JSON.parse(savedData);
                if (Array.isArray(parsedData)) {
                    taskData = transformTasksArray(parsedData);
                } else {
                    throw new Error("Données invalides");
                }
            } catch (error) {
                console.error("Erreur lors de la lecture des données sauvegardées", error);
                taskData = await getTaskData();
            }
        } else {
            taskData = await getTaskData();
            tasks = taskData;
            try {
                localStorage.setItem("Task", JSON.stringify(taskData));
            } catch (error) {
                console.error(`Erreur lors de la sauvegarde des données: ${error.message}`);
            }
        }
    } catch (error) {
        console.error("Erreur lors de la réccupération des tâches", error);
        taskData = await getTaskData();
        tasks = taskData;
        try {
            localStorage.setItem("Task", JSON.stringify(taskData));
        } catch (storageError) {
            console.error("Erreur lors de la sauvegarde des données:", storageError);
        }
    } 

    return taskData;
};

export const TaskApiService = {
    getTasks: () => fetchTasks(),
    createTask,
    updateTask,
    deleteTask
}