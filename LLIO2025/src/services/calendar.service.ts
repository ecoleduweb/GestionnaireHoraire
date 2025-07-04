// services/calendar.service.ts
import { Calendar } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';
import frLocale from '@fullcalendar/core/locales/fr';
import type { Activity } from '../Models/index';

export class CalendarService {
  calendar: Calendar | null = null;
  onDateSelect: ((info: any) => void) | null = null;
  onEventClick: ((info: any) => void) | null = null;

  initialize(element: HTMLElement, customOptions = {}) {
    // Options par défaut
    const defaultOptions = {
      plugins: [dayGridPlugin, timeGridPlugin, interactionPlugin],
      initialView: 'timeGridWeek',
      selectable: true,
      editable: true,
      locale: frLocale,
      headerToolbar: false, // Désactivé car nous avons un en-tête personnalisé
      firstDay: 1, // Lundi comme premier jour
      slotDuration: '00:30:00',
      nowIndicator: true,
      allDaySlot: false,
      height: 'auto',
      expandRows: true,
      dayHeaderFormat: { weekday: 'short', month: 'numeric', day: 'numeric', omitCommas: true },
      eventTimeFormat: {
        hour: '2-digit',
        minute: '2-digit',
        meridiem: false,
        hour12: false,
      },
      slotLabelFormat: {
        hour: 'numeric',
        minute: '2-digit',
        hour12: true,
      },
      eventDidMount: (info: any) => {
        let tooltipContent = info.event.title;
        info.el.setAttribute('title', tooltipContent);
        info.el.style.cursor = 'pointer';
      },

      // Gestionnaires d'événements
      select: (info: any) => {
        if (this.onDateSelect) {
          this.onDateSelect(info);
        }
      },
      eventClick: (info: any) => {
        if (this.onEventClick) {
          this.onEventClick(info);
        }
      },
    };

    // Fusionner les options par défaut avec les options personnalisées
    const options = { ...defaultOptions, ...customOptions };

    // Initialiser le calendrier avec les options
    this.calendar = new Calendar(element, options as any);
  }

  render() {
    this.calendar?.render();
  }

  addEvent(eventData: any) {
    this.calendar?.addEvent(eventData);
  }

  clearCalendar() {
    this.calendar?.removeAllEvents();
  }
  

  updateEvent(activity: any) {
    // Trouver l'événement existant
    const event = this.calendar?.getEventById(activity.id.toString());
    if (event) {
      // Mettre à jour les propriétés de l'événement
      event.setProp('title', activity.projectName);
      event.setStart(activity.startDate);
      event.setEnd(activity.endDate);
      event.setExtendedProp('name', activity.name);
      event.setExtendedProp('description', activity.description);
      event.setExtendedProp('userId', activity.userId);
      event.setExtendedProp('projectId', activity.projectId);
      event.setExtendedProp('categoryId', activity.categoryId);
    }
  }

  deleteActivity(activityId: string) {
    const event = this.calendar?.getEventById(activityId);
    if (event) {
      event.remove();
    }
  }

  setView(viewName: string) {
    this.calendar?.changeView(viewName);
  }

  next() {
    this.calendar?.next();
  }

  prev() {
    this.calendar?.prev();
  }

  today() {
    this.calendar?.today();
  }

  eventToActivity(eventInfo: any): Activity {
    return {
      id: parseInt(eventInfo.event.id),
      name: eventInfo.event.extendedProps.name,
      description: eventInfo.event.extendedProps.description,
      startDate: eventInfo.event.start,
      endDate: eventInfo.event.end,
      userId: eventInfo.event.extendedProps.userId,
      projectId: eventInfo.event.extendedProps.projectId,
      projectName: eventInfo.event.title,
      categoryId: eventInfo.event.extendedProps.categoryId,
    };
  }

  loadActivities(activities: Activity[]) {
    // Ajouter chaque activité au calendrier
    this.clearCalendar(); 
    activities.forEach((activity) => {
      this.addEvent({
        id: activity.id.toString(),
        title: activity.projectName,
        start: activity.startDate,
        end: activity.endDate,
        extendedProps: { ...activity },
      });
    });
  }

  getTotalHours(){
    const events = this.calendar?.getEvents() || [];
    let totalHours = 0;
    events.forEach((event) => {
      const start = event.start as Date;
      const end = event.end as Date;
      if (start && end) {
        const duration = (end.getTime() - start.getTime()) / (1000 * 60 * 60); // Convertir en heures
        totalHours += duration;
      }
    });
    return totalHours;

  }
}
