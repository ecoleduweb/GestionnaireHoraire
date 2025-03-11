// Extraire les heures et minutes d'un objet Date
export const getHoursFromDate = (date: Date): string => {
  const hours = date.getHours();
  return hours.toString().padStart(2, '0');
};

export const getMinutesFromDate = (date: Date): string => {
  const minutes = date.getMinutes();
  return minutes.toString().padStart(2, '0');
};

// Créer une date avec année/mois/jour spécifiques et heures/minutes données
export const createDateWithTime = (
  baseDate: Date,
  hours: string | number,
  minutes: string | number
): Date => {
  const result = new Date(baseDate);
  result.setHours(Number(hours));
  result.setMinutes(Number(minutes));
  result.setSeconds(0);
  result.setMilliseconds(0);
  return result;
};

// Initialiser les dates de début et fin pour une nouvelle tâche
export const initializeTaskDates = (targetDate?: Date): { startDate: Date, endDate: Date } => {
  const date = targetDate ? new Date(targetDate) : new Date();
  // Réinitialiser les secondes et millisecondes
  date.setSeconds(0);
  date.setMilliseconds(0);
  
  // Créer des copies pour éviter les références partagées
  const startDate = new Date(date);
  // Par défaut, la tâche dure 1 heure
  const endDate = new Date(date);
  endDate.setHours(endDate.getHours() + 1);
  
  return { startDate, endDate };
};

// Valider que l'heure de fin est après l'heure de début
export const applyEndTime = (
  startHours: string, 
  startMinutes: string, 
  endHours: string, 
  endMinutes: string
): { endHours: string, endMinutes: string } => {
  if (
    parseInt(endHours) < parseInt(startHours) ||
    (parseInt(endHours) === parseInt(startHours) &&
     parseInt(endMinutes) < parseInt(startMinutes))
  ) {
    return { 
      endHours: startHours, 
      endMinutes: startMinutes 
    };
  }
  return { endHours, endMinutes };
};