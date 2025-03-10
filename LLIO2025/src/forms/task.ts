import { initializeTaskDates } from "../utils/date";
import type { Task } from "../Models";

export const taskTemplate = {
  generate: (): Task => {
    const { startDate, endDate } = initializeTaskDates();
    
    return {
      name: "",
      description: "",
      billable: false,
      userId: 0,
      projectId: 0,
      categoryId: 0,
      startDate,
      endDate
    };
  },

  time: {
    hours: Array.from({ length: 24 }, (_, i) => i.toString().padStart(2, "0")),
    minutes: ["00", "15", "30", "45"]
  }
} as const;