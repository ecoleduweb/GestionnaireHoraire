export const getMinutesFromDate = (date: Date): string => {
    return date.getMinutes().toString().padStart(2, "0");
};

export const getHoursFromDate = (date: Date): string => {
    return date.getHours().toString().padStart(2, "0");
};