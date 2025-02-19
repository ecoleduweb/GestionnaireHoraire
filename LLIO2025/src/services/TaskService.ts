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
            const errorMessage = `Erreur lors de l'ajout d'une tâche: ${error.message}`;
            console.error(errorMessage)
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
            const errorMessage = `Erreur lors de la modification d'une tâche: ${error.message}`;
            console.error(errorMessage)
            throw error;
        }
    }
}
