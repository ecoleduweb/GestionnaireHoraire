import { TaskApiService } from './TaskApiService';
import type { Task } from '../Models';

export class TaskService {
    static addTask = async (task: Task): Promise<Task> => {
        const taskToSend = {
            ...task,
            startDateTime: new Date(task.startDateTime),
            endDateTime: new Date(task.endDateTime)
        };
        return await TaskApiService.createTask(taskToSend);
    }

    static updateTask = async (task: Task): Promise<Task> => {
        const taskToUpdate = {
            ...task,
            startDateTime: new Date(task.startDateTime),
            endDateTime: new Date(task.endDateTime)
        };
        return await TaskApiService.updateTask(taskToUpdate);
    }
}