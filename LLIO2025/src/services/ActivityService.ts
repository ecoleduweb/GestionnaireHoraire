import type { Activity } from '../Models/index';
import { GET, POST, PUT, DELETE } from '../ts/server';

interface FormActivity extends Omit<Activity, 'startDate' | 'endDate'> {
  startDate: string;
  endDate: string;
}

const transformActivitiesDates = (activity: FormActivity): Activity => ({
  ...activity,
  startDate: new Date(activity.startDate),
  endDate: new Date(activity.endDate),
});

const createActivity = async (activity: Activity): Promise<Activity> => {
  // Préparation des données pour l'API
  const activityForApi = {
    name: activity.name,
    description: activity.description || '',
    startDate: activity.startDate.toISOString(),
    endDate: activity.endDate.toISOString(),
    userId: activity.userId || 1,
    projectId: activity.projectId || 1,
    categoryId: activity.categoryId || 1,
  };

  try {
    const response = await POST<typeof activityForApi, { activity: Activity }>(
      '/activities',
      activityForApi
    );
    return {
      ...response.activity,
      startDate: new Date(response.activity.startDate),
      endDate: new Date(response.activity.endDate),
    };
  } catch (error) {
    console.error('Erreur lors de la création de tâche:', error);
    throw error;
  }
};

const updateActivity = async (activity: Activity): Promise<Activity> => {
  const activityForApi: FormActivity = {
    ...activity,
    startDate: activity.startDate.toISOString(),
    endDate: activity.endDate.toISOString(),
  };

  const response = await PUT<FormActivity, FormActivity>(
    `/activities/${activity.id}`,
    activityForApi
  );
  return transformActivitiesDates(response);
};

const deleteActivity = async (id: number): Promise<void> => {
  await DELETE(`/activities/${id}`);
};

const getActivities = async () => {
  try {
    const response = await GET<FormActivity[]>('/activities');
    const activityData = response.map(transformActivitiesDates);
    return activityData;
  } catch (error) {
    // Gardé pour le débogage, mais peut être supprimé si nécessaire
    console.error('Erreur lors de la récupération des tâches', error);
    throw error;
  }
};

export const ActivityApiService = {
  getActivities,
  createActivity,
  updateActivity,
  deleteActivity,
};
