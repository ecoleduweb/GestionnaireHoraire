import type { Activity } from "../Models/index";
import { GET, POST, PUT, DELETE } from "../ts/server";

let activites: Activity[];

interface RawActivity extends Omit<Activity, 'startDateTime' | 'endDateTime'> {
    startDateTime: string | Date;
    endDateTime: string | Date;
}

const transformActivitysDates = (activity: RawActivity): Activity => ({
    ...activity,
    startDateTime: new Date(activity.startDateTime),
    endDateTime: new Date(activity.endDateTime)
});

const transformActivitysArray = (activites: Activity[]): Activity[] => 
    activites.map(transformActivitysDates);

// const getActivityData = async (): Promise<Activity[]> => {
//     const response = await GET<Activity[]>("/activity/all");
//     return transformActivitysArray(response);
// };

const createActivity = async (activities: Activity): Promise<Activity> => {
    const response = await POST<Activity, Activity>("/activity", activities);
    return transformActivitysDates(response);
};

const updateActivity = async (activity: Activity): Promise<Activity> => {
    const response = await PUT<Activity, Activity>(`/activity/${activity.id}`, activity);
    return transformActivitysDates(response);
};

const deleteActivity = async (id: number): Promise<void> => {
    await DELETE(`/activity/${id}`);
};

// export const fetchActivities = async () => {
//     let savedData: string | null = null;
//     let activitiesData: Activity[];

//     try {
//         savedData = localStorage.getItem("Activity");
//     } catch (error) {
//         console.error(`Erreur lors de l'accès au localStorage: ${error.message}`);
//     }

//     try {
//         if (activities?.length) {
//             activitiesData = activities;
//         } else if (savedData) {
//             try {
//                 const parsedData = JSON.parse(savedData);
//                 if (Array.isArray(parsedData)) {
//                     activitiesData = transformActivitysArray(parsedData);
//                 } else {
//                     throw new Error("Données invalides");
//                 }
//             } catch (error) {
//                 console.error("Erreur lors de la lecture des données sauvegardées", error);
//                 activitiesData = await getActivityData();
//             }
//         } else {
//             activitiesData = await getActivityData();
//             activities = activitiesData;
//             try {
//                 localStorage.setItem("Activity", JSON.stringify(activitiesData));
//             } catch (error) {
//                 console.error(`Erreur lors de la sauvegarde des données: ${error.message}`);
//             }
//         }
//     } catch (error) {
//         console.error("Erreur lors de la réccupération des tâches", error);
//         activitiesData = await getActivityData();
//         activities = activitiesData;
//         try {
//             localStorage.setItem("Activity", JSON.stringify(activitiesData));
//         } catch (storageError) {
//             console.error("Erreur lors de la sauvegarde des données:", storageError);
//         }
//     } 

//     return activitiesData;
// };

export const ActivityApiService = {
    // fetchActivities(),
    createActivity,
    updateActivity,
    deleteActivity
}