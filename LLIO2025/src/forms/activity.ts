import type { Activity } from '../Models';
import { initializeActivityDates } from '../utils/date';

export const activityTemplate = {
  generate: (): Activity => {
    const { startDate, endDate } = initializeActivityDates();
    return {
      name: '',
      description: '',
      billable: false,
      userId: 0,
      projectId: 0,
      categoryId: 0,
      startDate,
      endDate,
    };
  },

  time: {
    hours: Array.from({ length: 24 }, (_, i) => i.toString().padStart(2, '0')),
    minutes: ['00', '15', '30', '45'],
  },
} as const;
