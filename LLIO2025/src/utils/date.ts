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
export const initializeActivityDates = (targetDate?: Date): { startDate: Date; endDate: Date } => {
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
): { endHours: string; endMinutes: string } => {
  if (
    parseInt(endHours) < parseInt(startHours) ||
    (parseInt(endHours) === parseInt(startHours) && parseInt(endMinutes) < parseInt(startMinutes))
  ) {
    return {
      endHours: startHours,
      endMinutes: startMinutes,
    };
  }
  return { endHours, endMinutes };
};
// Formatage du titre de la vue calendrier (jour, semaine, mois)
export const formatViewTitle = (viewType: string, date: Date): string => {
  if (viewType === 'dayGridMonth') {
    // Format pour la vue mois: "Février 2025"
    return date.toLocaleDateString('fr-FR', {
      month: 'long',
      year: 'numeric',
    });
  } else if (viewType === 'timeGridWeek') {
    // Format pour la vue semaine: "10 mars – 16 mars 2025"
    // Trouver le lundi de la semaine
    const startDate = new Date(date);
    const day = startDate.getDay();
    const diff = startDate.getDate() - day + (day === 0 ? -6 : 1); // Ajuster lorsque jour = dimanche
    startDate.setDate(diff);

    // Le dimanche est le 6ème jour après le lundi
    const endDate = new Date(startDate);
    endDate.setDate(startDate.getDate() + 6);

    const startFormatted = startDate.getDate();
    const startMonth = startDate.toLocaleDateString('fr-FR', { month: 'short' });

    const endFormatted = endDate.getDate();
    const endMonth = endDate.toLocaleDateString('fr-FR', { month: 'long' });
    const year = endDate.getFullYear();

    return `${startFormatted} ${startMonth}. – ${endFormatted} ${endMonth} ${year}`;
  } else if (viewType === 'timeGridDay') {
    // Format pour la vue jour: "Mercredi 26 février 2025"
    return date.toLocaleDateString('fr-FR', {
      weekday: 'long',
      day: 'numeric',
      month: 'long',
      year: 'numeric',
    });
  }

  // Format par défaut si le type de vue n'est pas reconnu
  return date.toLocaleDateString('fr-FR', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
  });
};
