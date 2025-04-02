// Extraire les heures et minutes d'un objet Date
export function getHoursFromDate(date: Date | null | undefined): string {
  if (!date || !(date instanceof Date) || isNaN(date.getTime())) {
    return '12';
  }
  return date.getHours().toString().padStart(2, '0');
};

export function getMinutesFromDate(date: Date | null | undefined): string {
  if (!date || !(date instanceof Date) || isNaN(date.getTime())) {
    return '00';
  }
  return date.getMinutes().toString().padStart(2, '0');
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
export const initializeActivityDates = (targetDate?: Date): { startDate: Date, endDate: Date } => {
  const date = targetDate ? new Date(targetDate) : new Date();
  // Réinitialiser les secondes et millisecondes
  date.setSeconds(0);
  date.setMilliseconds(0);
  
  // Créer des copies pour éviter les références partagées
  const startDate = new Date(date);
  // Par défaut, l'activité dure 1 heure
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

export function getDateOrDefault(dateToValidate: Date | null | undefined, defaultDate: Date): Date {
  if (!dateToValidate || !(dateToValidate instanceof Date) || isNaN(dateToValidate.getTime())) {
    return new Date(defaultDate);
  }
  return dateToValidate;
}