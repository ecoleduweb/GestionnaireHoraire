import { Calendar } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';
import type { EventClickArg, DateSelectArg } from '@fullcalendar/core/index.js';
import type { Task } from '../Models';

export class CalendarService {
    calendar: Calendar | null = null;
    onDateSelect?: (info: DateSelectArg) => void;
    onEventClick?: (info: EventClickArg) => void;
    onViewChange?: (view: string) => void;

    initialize(element: HTMLElement) {
        this.calendar = new Calendar(element, {
            plugins: [dayGridPlugin, timeGridPlugin, interactionPlugin],
            initialView: 'dayGridMonth',
            headerToolbar: {
                left: 'prev,next today',
                center: 'title',
                right: 'dayGridMonth,timeGridWeek,timeGridDay'
            },
            editable: true,
            selectable: true,
            select: (info) => this.onDateSelect?.(info),
            eventClick: (info) => this.onEventClick?.(info),
            viewDidMount: (info) => this.onViewChange?.(info.view.type)
        });
    }

    render() {
        this.calendar?.render();
    }

    addEvent(event: any) {
        this.calendar?.addEvent(event);
    }

    updateEvent(task: Task) {
        const existingEvent = this.calendar?.getEventById(task.id.toString());
        if (existingEvent) {
            existingEvent.remove();
            this.addEvent({
                id: task.id.toString(),
                title: task.name,
                start: task.startDateTime,
                end: task.endDateTime,
                extendedProps: { ...task },
            });
        }
    }

    deleteTask(eventId: string) {
        const event = this.calendar?.getEventById(eventId);
        if (event) {
            event.remove();
        }
    }
}