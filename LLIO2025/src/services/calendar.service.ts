import { Calendar } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';

export class CalendarService {
    /** @type {Calendar | null} */
    calendar;

    constructor() {
        this.calendar = null;
    }

    /**
     * @param {HTMLElement} element
     */
    initialize(element) {
        this.calendar = new Calendar(element, {
            plugins: [dayGridPlugin, timeGridPlugin, interactionPlugin],
            initialView: 'dayGridMonth',
            headerToolbar: {
                left: 'prev,next today',
                center: 'title',
                right: 'dayGridMonth,timeGridWeek,timeGridDay'
            },
            events: [ 
                {
                    title: 'Test Event',
                    start: new Date(),
                    end: new Date(),
                },
                {
                    title: 'Test Event 2',
                    start: new Date('2025-01-21 12:00:00'),
                    end: new Date('2025-01-21 13:00:00'),
                }
            ],
            editable: true,
            selectable: true,
            select: this.handleDateSelect.bind(this),
            eventClick: this.handleEventClick.bind(this)
        });
    }

    render() {
        this.calendar?.render();
    }

    /**
     * @param {import("@fullcalendar/core").DateSelectArg} info
     */
    handleDateSelect(info) {
        console.log('Date sélectionnée:', info.startStr);
    }

    /**
     * @param {import("@fullcalendar/core").EventClickArg} info
     */
    handleEventClick(info) {
        console.log('Événement cliqué:', info.event.title);
    }
}