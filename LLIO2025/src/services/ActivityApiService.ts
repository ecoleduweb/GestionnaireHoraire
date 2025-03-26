import type { Activity } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

let activites: Activity[];

interface RawActivity extends Omit<Activity, 'startDate' | 'endDate'> {
    startDate: string;
    endDate: string;
    updated_activity: Activity;
}

const transformActivitesDates = (activity: RawActivity): Activity => ({
    ...activity,
    startDate: new Date(activity.startDate),
    endDate: new Date(activity.endDate)
});

const createActivity = async (activity: Activity): Promise<Activity> => {
        
    // Préparation des données pour l'API
    const activityForApi = {
        name: activity.name,
        description: activity.description || "",
        billable: activity.billable || false,
        startDate: activity.startDate.toISOString(),
        endDate: activity.endDate.toISOString(),
        userId: activity.userId || 1,
        projectId: activity.projectId || 1,
        categoryId: activity.categoryId || 1
    };
    
    try {
        const response = await POST<typeof activityForApi, {activity: Activity}>("/activity", activityForApi);
        return {
            ...response.activity,
            startDate: new Date(response.activity.startDate),
            endDate: new Date(response.activity.endDate)
        };
    } catch (error) {
        console.error("Erreur lors de la création de tâche:", error);
        throw error;
    }
};

const updateActivity = async (activity: Activity): Promise<Activity> => {
    if (!activity) {
        console.error("La tâche est undefined");
        throw new Error("La tâche ou son ID est manquant");
    }

    if (!activity.id) {
        console.error("L'ID de la tâche est manquant", activity);
        throw new Error("L'ID de la tâche est manquant");
    }

    // Préparation des données pour l'API - ATTENTION: UTILISER L'URL RACINE /activitys
    const activityForApi = {
        id: activity.id,
        name: activity.name,
        description: activity.description || "",
        billable: activity.billable || false,
        startDate: activity.startDate.toISOString(),
        endDate: activity.endDate.toISOString(),
        userId: activity.userId || 1,
        projectId: activity.projectId || 1,
        categoryId: activity.categoryId || 1
    };

    try {
        const response = await PUT<typeof activityForApi, {updated_activity: Activity}>("/activity", activityForApi);
        
        if (response && response.updated_activity) {
            // Extraire la tâche mise à jour de l'objet de réponse
            const updatedactivityData = response.updated_activity;
            
            // Convertir les dates
            return {
                ...updatedactivityData,
                startDate: new Date(updatedactivityData.startDate),
                endDate: new Date(updatedactivityData.endDate)
            };
        } else {
            return activity;
        }
    } catch (error) {
        console.error("Erreur lors de la mise à jour de la tâche:", error);
        throw error;
    }
};

const deleteActivity = async (id: number): Promise<void> => {
    await DELETE(`/activity/${id}`);
};

const getActivites = async () => {
    try {
        const response = await GET<RawActivity[]>("/activities");
        const activityData = response.map(transformActivitesDates);
        activites = activityData;
        return activityData;
    } catch (error) {
        // Gardé pour le débogage, mais peut être supprimé si nécessaire
        console.error("Erreur lors de la récupération des tâches", error);
        throw error;
    }
};

export const ActivityApiService = {
    getActivites,
    createActivity,
    updateActivity,
    deleteActivity
}