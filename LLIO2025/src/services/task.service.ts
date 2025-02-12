import type { Task } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

let tasks: Task[];

const getTaskData = async (): Promise<Task[]> => 
    await GET<Task[]>("/task/all");


const createTask = async (task: Task): Promise<Task> => {
    return await POST<Task, Task>("/task", task);
};

const updateTask = async (task: Task): Promise<Task> => {
    return await PUT<Task, Task>("/task/${task.id}", task);
};

const deleteTask = async (id: number): Promise<void> => {
    await DELETE("/task/${id}");
};

export const cacheTasks = async () => {
    let savedData = localStorage.getItem("Task");
    let taskData: Task[];

    try {
        if (tasks?.length) {
            taskData = tasks;
        } else if (savedData) {
            taskData = JSON.parse(savedData);
        } else {
            taskData = await getTaskData();
            tasks = taskData;
            localStorage.setItem("Task", JSON.stringify(taskData));
        }
    } catch {
        taskData = await getTaskData();
        tasks = taskData;
        localStorage.setItem("Task", JSON.stringify(taskData));
    } 

    return taskData;
};

export const TaskService = {
    getTasks: () => cacheTasks(),
    createTask,
    updateTask,
    deleteTask
}