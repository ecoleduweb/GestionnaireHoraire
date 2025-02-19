import type { Task } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

let tasks: Task[];

const transformTasksDates = (task: any): Task => ({
    ...task,
    startDate: new Date(task.startDate),
    endDate: new Date(task.endDate)
});

const transformTasksArray = (tasks: any[]): Task[] => 
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
    const response = await PUT<Task, Task>("/task/${task.id}", task);
    return transformTasksDates(response);
};

const deleteTask = async (id: number): Promise<void> => {
    await DELETE("/task/${id}");
};

export const fetchTasks = async () => {
    let savedData = localStorage.getItem("Task");
    let taskData: Task[];

    try {
        if (tasks?.length) {
            taskData = tasks;
        } else if (savedData) {
            taskData = transformTasksArray(JSON.parse(savedData));
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

export const TaskApiService = {
    getTasks: () => fetchTasks(),
    createTask,
    updateTask,
    deleteTask
}