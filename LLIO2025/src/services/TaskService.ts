import { TaskApiService } from './TaskApiService';
import type { Task } from '../Models';

export class TaskService {
    static addTask = async (task: Task) => {
        try {
            const taskToSend = {
                ...task,
                startDateTime: new Date(task.startDateTime),
                endDateTime: new Date(task.endDateTime)
            };
            return await TaskApiService.createTask(taskToSend);
        } catch (error) {
            console.error("Erreur lors de la création de la tâche", error);
            throw error;
        }
    }

    static updateTask = async (task: Task) => {
        try {
            const taskToUpdate = {
                ...task,
                startDateTime: new Date(task.startDateTime),
                endDateTime: new Date(task.endDateTime)
            };
            return await TaskApiService.updateTask(taskToUpdate);
        } catch (error) {
            console.error("Erreur lors de la mise à jour de la tâche", error);
            throw error;
        }
    }
}