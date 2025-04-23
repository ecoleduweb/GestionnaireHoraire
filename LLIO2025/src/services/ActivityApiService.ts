import type { Activity, ActivityUpdateResponse } from '../Models/index';
import { GET, POST, PUT, DELETE } from '../ts/server';

interface RawActivity extends Omit<Activity, 'startDate' | 'endDate'> {
  startDate: string;
  endDate: string;
}

const transformActivityStringToDates = (activity: RawActivity): Activity => ({
  ...activity,
  startDate: new Date(activity.startDate),
  endDate: new Date(activity.endDate),
});

const toStringDatesToString = (activity: Activity) => ({
  ...activity,
  startDate: activity.startDate.toISOString(),
  endDate: activity.endDate.toISOString(),
});

const createActivity = async (activity: Activity): Promise<Activity> => {
  // Préparation des données pour l'API
  const activityForApi = toStringDatesToString(activity);

  try {
    const response = await POST<typeof activityForApi, { activity: RawActivity }>(
      '/activity',
      activityForApi
    );
    return transformActivityStringToDates(response.activity);
  } catch (error) {
    console.error('Erreur lors de la création de tâche:', error);
    throw error;
  }
};

const updateActivity = async (activity: Activity): Promise<Activity> => {
  if (!activity) {
    console.error('La tâche est undefined');
    throw new Error('La tâche ou son ID est manquant');
  }

  if (!activity.id) {
    console.error("L'ID de la tâche est manquant", activity);
    throw new Error("L'ID de la tâche est manquant");
  }

  const activityForApi = toStringDatesToString(activity);

  try {
    const response = await PUT<typeof activityForApi, ActivityUpdateResponse>(
      '/activity',
      activityForApi
    );

    // Extraire la tâche mise à jour de l'objet de réponse
    const updatedactivityData = response.updated_activity;

    // Convertir les dates
    return {
      ...updatedactivityData,
      startDate: new Date(updatedactivityData.startDate),
      endDate: new Date(updatedactivityData.endDate),
    };
  } catch (error) {
    console.error('Erreur lors de la mise à jour de la tâche:', error);
    throw error;
  }
};

const deleteActivity = async (id: number): Promise<void> => {
  await DELETE(`/activity/${id}`);
};

const getAllActivites = async () => {
  try {
    // La réponse est un objet avec une propriété 'activities'
    const response = await GET<{ activities: RawActivity[] }>('/activities');

    // Vérifier si la réponse contient la propriété 'activities' et si c'est un tableau
    if (response && response.activities && Array.isArray(response.activities)) {
      return response.activities.map(transformActivityStringToDates);
    } else {
      console.error('Format de réponse inattendu :', response);
      alert('Erreur: Format de réponse inattendu lors de la récupération des tâches.');
      return []; // Retourner un tableau vide en cas de problème
    }
  } catch (error) {
    console.error('Erreur lors de la récupération des tâches', error);
    alert('Erreur lors de la récupération des tâches. Veuillez réessayer plus tard.');
    throw error;
  }
};

const getAllActivitesFromRange = async (startDate: string, endDate: string) => {
  try {
    const response = await GET<{ activities: RawActivity[] }>(
      '/activities/me/' + startDate + '/' + endDate
    );
    console.log('response', response);

    // Vérifier si la réponse contient la propriété 'activities' et si c'est un tableau
    if (response && response.activities && Array.isArray(response.activities)) {
      return response.activities.map(transformActivityStringToDates);
    } else {
      console.error('Format de réponse inattendu :', response);
      alert('Erreur: Format de réponse inattendu lors de la récupération des tâches.');
      return []; // Retourner un tableau vide en cas de problème
    }
  } catch (error) {
    console.error('Erreur lors de la récupération des tâches', error);
    alert('Erreur lors de la récupération des tâches. Veuillez réessayer plus tard.');
    throw error;
  }
};

export const ActivityApiService = {
  getAllActivites,
  createActivity,
  updateActivity,
  deleteActivity,
  getAllActivitesFromRange,
};
