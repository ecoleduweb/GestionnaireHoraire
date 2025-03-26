import type { Activity, ActivityUpdateResponse } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

interface RawActivity extends Omit<Activity, 'startDate' | 'endDate'> {
    startDate: string;
    endDate: string;
}

const transformActivityDates = (activity: RawActivity): Activity => ({
    ...activity,
    startDate: new Date(activity.startDate),
    endDate: new Date(activity.endDate)
});

const toStringDates = (activity: Activity) => ({
    ...activity,
    startDate: activity.startDate.toISOString(),
    endDate: activity.endDate.toISOString()
});

const createActivity = async (activity: Activity): Promise<Activity> => {
        
    // Préparation des données pour l'API
    const activityForApi = toStringDates(activity);
    
    try {
        const response = await POST<typeof activityForApi, {activity: RawActivity}>("/activity", activityForApi);
        return transformActivityDates(response.activity);
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

    const activityForApi = toStringDates(activity);

    try {
        const response = await PUT<typeof activityForApi, ActivityUpdateResponse>("/activity", activityForApi);
        
        // Extraire la tâche mise à jour de l'objet de réponse
        const updatedactivityData = response.updated_activity;
        
        // Convertir les dates
        return {
            ...updatedactivityData,
            startDate: new Date(updatedactivityData.startDate),
            endDate: new Date(updatedactivityData.endDate)
        };
       
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
        return response.map(transformActivityDates);
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