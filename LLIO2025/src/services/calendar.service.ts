// services/calendar.service.ts
import { Calendar } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';
import frLocale from '@fullcalendar/core/locales/fr';

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
      businessHours: {
        daysOfWeek: [1, 2, 3, 4, 5], // Lundi - Vendredi
        startTime: '08:00',
        endTime: '18:00',
      },
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

  updateEvent(task: any) {
    // Trouver l'événement existant
    const event = this.calendar?.getEventById(task.id.toString());
    if (event) {
      // Mettre à jour les propriétés de l'événement
      event.setProp('title', task.name);
      event.setStart(task.startDateTime);
      event.setEnd(task.endDateTime);
      event.setExtendedProp('description', task.description);
      event.setExtendedProp('billable', task.billable);
      event.setExtendedProp('userId', task.userId);
      event.setExtendedProp('projectId', task.projectId);
      event.setExtendedProp('categoryId', task.categoryId);
    }
  }

  deleteTask(taskId: string) {
    const event = this.calendar?.getEventById(taskId);
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
}
