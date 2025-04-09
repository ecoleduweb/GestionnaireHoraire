<script lang="ts">
  import type { CalendarService } from '../../services/calendar.service';
  import { CalendarService as CS } from '../../services/calendar.service';
  import { onMount } from 'svelte';
  import ActivityModal from '../../Components/Calendar/ActivityModal.svelte';
  import { ActivityApiService } from '../../services/ActivityApiService';
  import type { Activity } from '../../Models/index.ts';
  // Importez le fichier CSS
  import '../../style/modern-calendar.css';
  import { getDateOrDefault } from '../../utils/date';
  // Importer FullCalendar en français
  import frLocale from '@fullcalendar/core/locales/fr';
  import { formatViewTitle } from '../../utils/date';
  import { ClientTelemetry } from "$lib/tracer"
  import { env } from "$env/dynamic/public"
  import { Plus, Calendar, ChevronLeft, ChevronRight } from 'lucide-svelte';

  const ENABLED_TELEMETRY = env.PUBLIC_ENABLED_TELEMETRY === "true";

  if (ENABLED_TELEMETRY) {
      const telemetry = ClientTelemetry.getInstance()
      telemetry.start()
  }
  
  let calendarEl = $state<HTMLElement | null>(null);
  let calendarService = $state<CalendarService | null>(null);
  let showModal = $state(false);
  let selectedDate: { start: Date; end: Date } | null = null;
  let editMode = $state(false);
  let editActivity = $state(null);
  let activeView = $state('timeGridWeek');
  let currentViewTitle = $state('');

  const timeRanges = [
  { label: 'Heures de bureau (8h-17h)', start: '08:00:00', end: '17:00:00', default: true },
  { label: 'Toute la journée (24h)', start: '00:00:00', end: '24:00:00' }
  ];

  let activeTimeRange = $state(timeRanges.find(range => range.default));

  const users = [{ id: 1, name: 'Test ManuDev' }];
  const projects = [{ id: 1, name: 'Projet manudev' }];
  const categories = [{ id: 1, name: 'Categorie Test ManuDev' }];

  // Fonction pour attribuer une couleur à chaque événement
  function getEventClassName(eventInfo: any) {
    const eventTypes = [
      'event-blue',
      'event-green',
      'event-teal',
      'event-dark-teal',
      'event-light-teal',
    ];

    if (eventInfo.event.extendedProps.categoryId) {
      const categoryId = parseInt(eventInfo.event.extendedProps.categoryId);
      return eventTypes[categoryId % eventTypes.length];
    }

    const hash = eventInfo.event.title.split('').reduce((acc, char) => acc + char.charCodeAt(0), 0);
    return eventTypes[hash % eventTypes.length];
  }

  // Fonction pour mettre à jour le titre de la période courante
  function updateViewTitle() {
    if (calendarService?.calendar) {
      const dateAPI = calendarService.calendar.getDate();
      const viewType = calendarService.calendar.view.type;
      currentViewTitle = formatViewTitle(viewType, dateAPI);
    }
  }

  onMount(() => {
    if (calendarEl) {
      calendarService = new CS();

      // Configuration personnalisée pour FullCalendar
      const calendarOptions = {
        initialView: activeView,
        locale: frLocale, // Utiliser la locale française
        firstDay: 1, // 1 = lundi (standard français)
        buttonText: {
          today: "Aujourd'hui",
          month: 'Mois',
          week: 'Semaine',
          day: 'Jour',
        },
        slotDuration: '00:30:00', // Durée de chaque intervalle de temps
        allDaySlot: false,
        slotMinTime: activeTimeRange.start,
        slotMaxTime: activeTimeRange.end,
        nowIndicator: true,

        // Gestion du drag
        editable: true,
        eventDrop: handleEventDropOrResize,
        eventResize: handleEventDropOrResize,

        height: 'auto',
        contentHeight: 'auto', // Hauteur automatique
        slotHeight: 25, // Hauteur réduite des slots (plus compact)
        expandRows: true,
        dayHeaderFormat: { weekday: 'short', day: '2-digit', month: '2-digit' },
        eventClassNames: getEventClassName,
        eventTimeFormat: {
          hour: '2-digit',
          minute: '2-digit',
          meridiem: false,
          hour12: false,
        },
        slotLabelFormat: {
          hour: 'numeric',
          minute: '2-digit',
          hour12: false,
        },
        datesSet: () => {
          // Appelé à chaque changement de dates ou de vue
          updateViewTitle();
        },
      };

      calendarService.onDateSelect = (info) => {
        editMode = false;
        editActivity = null;
        selectedDate = info;
        showModal = true;
      };

      calendarService.onEventClick = (info) => {
        editMode = true;
        editActivity = {
          id: info.event.extendedProps.id,
          name: info.event.title,
          description: info.event.extendedProps.description,
          userId: info.event.extendedProps.userId,
          projectId: info.event.extendedProps.projectId,
          categoryId: info.event.extendedProps.categoryId,
          startDate: info.event.start,
          endDate: info.event.end,
        };
        selectedDate = {
          start: info.event.start,
          end: info.event.end,
        };
        showModal = true;
      };

      // Initialiser avec les options personnalisées
      calendarService.initialize(calendarEl, calendarOptions);
      calendarService.render();

      // Mettre à jour le titre initial
      updateViewTitle();
    }
  });

  function setView(viewName: string) {
    if (calendarService) {
      calendarService.setView(viewName);
      activeView = viewName;
      updateViewTitle();
    }
  }

  async function handleActivitySubmit(activityData: Activity) {
    calendarService.addEvent({
      id: activityData.id.toString(),
      title: activityData.name,
      start: activityData.startDate,
      end: activityData.endDate,
      extendedProps: { ...activityData },
    });
  }

  async function handleActivityUpdate(activity: Activity) {
    if (!calendarService?.calendar) return;

    try {
      // Utiliser la fonction pour valider les dates
      const now = new Date();
      activity.startDate = getDateOrDefault(activity.startDate, now);

      // si la date est invalide, définir par défaut 30 minutes après le début
      const defaultEndDate = new Date(activity.startDate.getTime() + 30 * 60000);
      activity.endDate = getDateOrDefault(activity.endDate, defaultEndDate);

      const updatedActivity = await ActivityApiService.updateActivity(activity);
      calendarService.updateEvent(updatedActivity);
    } catch (error) {
      console.error('Erreur lors de la mise à jour de la tâche', error);

      alert("Une erreur est survenue lors de la mise à jour de l'activité.");

      throw error;
    }
  }

  // Fonction pour gérer le déplacement et le redimmensionnement d'une tâche
  async function handleEventDropOrResize(info) {
    try {
      const activity = calendarService.eventToActivity(info);

      const updatedActivity = await ActivityApiService.updateActivity(activity);

      calendarService.updateEvent(updatedActivity);
    } catch (error) {
      console.error("Erreur lors de la mise à jour de l'activité", error);
      alert("Une erreur est survenue lors de la mise à jour de l'activité.");
      info.revert();
    }
  }

  async function handleActivityDelete(activity: Activity) {
    if (!calendarService?.calendar || !activity.id) return;
    try {
      await ActivityApiService.deleteActivity(activity.id);
      calendarService.deleteActivity(activity.id.toString());
    } catch (error) {
      console.error("Erreur lors de la suppression de l'activité", error);
      throw error;
    }
  }

  function handleNewActivity() {
    editMode = false;
    editActivity = null;
    selectedDate = {
      start: new Date(),
      end: new Date(new Date().getTime() + 30 * 60000), // 30 minutes par défaut
    };
    showModal = true;
  }

  function prevPeriod() {
    calendarService?.prev();
    updateViewTitle();
  }

  function nextPeriod() {
    calendarService?.next();
    updateViewTitle();
  }

  function goToday() {
    calendarService?.today();
    updateViewTitle();
  }
  
  function setTimeRange(range) {
    activeTimeRange = range;
    
    if (calendarService?.calendar) {
      calendarService.calendar.setOption('slotMinTime', range.start);
      calendarService.calendar.setOption('slotMaxTime', range.end);
      calendarService.calendar.render();
      
      // Modifier directement les styles
      const tableElement = calendarEl.querySelector('.fc-timegrid-slots table') as HTMLTableElement;
      if (tableElement) {
        tableElement.style.height = range.start === '00:00:00' ? '1200px' : '540px';
      }
      
      // Ajuster la hauteur des cellules
      const cells = calendarEl.querySelectorAll('.fc-timegrid-slot');
      const cellHeight = range.start === '00:00:00' ? '20px' : '30px';
      cells.forEach(cell => {
        (cell as HTMLElement).style.height = cellHeight;
      });
    }
  }

  const currentDate = new Date();
  const formattedDate = currentDate.toLocaleDateString('fr-FR', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
  });
</script>

<!-- CSS supplémentaire pour rendre le calendrier plus compact -->
<style>
  :global(.fc .fc-timegrid-slot) {
    height: 25px !important;
    min-height: 25px !important;
    max-height: 25px !important;
  }

  :global(.fc-timegrid-event) {
    min-height: 20px !important;
    max-height: none !important;
  }

  :global(.fc-timegrid-slot-label) {
    vertical-align: top !important;
    padding-top: 2px !important;
  }

  :global(.fc-theme-standard td),
  :global(.fc-theme-standard th) {
    padding: 1px !important;
  }

  :global(.fc .fc-daygrid-day-frame) {
    padding: 2px !important;
  }

  /* Couleur personnalisée pour l'indicateur de l'heure actuelle */
  :global(.fc .fc-timegrid-now-indicator-line) {
    border-color: #015e61 !important;
  }

  :global(.fc .fc-timegrid-now-indicator-arrow) {
    border-color: #015e61 !important;
    background-color: #015e61 !important;
  }
</style>

<div class="w-full h-full bg-white px-4 py-6">
  <div class="max-w-7xl mx-auto">
    <!-- En-tête du calendrier -->
    <div class="flex flex-col md:flex-row justify-between items-center mb-6 gap-4">
      <!-- Titre avec icône -->
      <div class="flex items-center">
        <Calendar class="mr-2" />
        <h1 class="text-xl font-semibold text-gray-800">Aujourd'hui, {formattedDate}</h1>
      </div>

      <!-- Boutons de vue alignés au centre -->
      <div class="flex items-center bg-gray-100 p-1 rounded-lg">
        <button
          class="px-5 py-2 rounded-lg transition-colors {activeView === 'timeGridDay'
            ? 'bg-white text-[#015e61] font-medium'
            : 'text-gray-500 hover:bg-white hover:text-[#015e61]'}"
          on:click={() => setView('timeGridDay')}
        >
          Jour
        </button>
        <button
          class="px-5 py-2 rounded-lg transition-colors {activeView === 'timeGridWeek'
            ? 'bg-white text-[#015e61] font-medium'
            : 'text-gray-500 hover:bg-white hover:text-[#015e61]'}"
          on:click={() => setView('timeGridWeek')}
        >
          Semaine
        </button>
        <button
          class="px-5 py-2 rounded-lg transition-colors {activeView === 'dayGridMonth'
            ? 'bg-white text-[#015e61] font-medium'
            : 'text-gray-500 hover:bg-white hover:text-[#015e61]'}"
          on:click={() => setView('dayGridMonth')}
        >
          Mois
        </button>
      </div>
      <div class="flex items-center ml-4">
        {#each timeRanges as range}
          <button
            class="px-4 py-2 mx-1 rounded-lg text-sm transition-colors {activeTimeRange.label === range.label
              ? 'bg-[#015e61] text-white'
              : 'bg-gray-100 text-gray-700 hover:bg-gray-200'}"
            on:click={() => setTimeRange(range)}
          >
            {range.label}
          </button>
        {/each}
      </div>

      <!-- Bouton New Activity -->
      <button
        on:click={handleNewActivity}
        class="bg-[#015e61] hover:bg-[#014446] text-white py-2 px-6 rounded-xl flex items-center gap-2 font-semibold transition-colors"
      >
        <Plus class="h-5 w-5" />
        Nouvelle activité
      </button>
    </div>

    <!-- Contrôles de navigation -->
    <div class="flex items-center justify-between mb-4">
      <div class="flex items-center space-x-3">
        <button
          on:click={prevPeriod}
          class="p-2 rounded-lg bg-gray-100 hover:bg-gray-200 transition-colors"
        >
          <ChevronLeft class="w-6 h-6 text-gray-600" />
        </button>
        <button
          on:click={nextPeriod}
          class="p-2 rounded-lg bg-gray-100 hover:bg-gray-200 transition-colors"
        >
          <ChevronRight class="w-6 h-6 text-gray-600" />
        </button>
        <button
          on:click={goToday}
          class="px-4 py-2 rounded-lg bg-gray-100 hover:bg-gray-200 text-gray-700 transition-colors"
        >
          Aujourd'hui
        </button>
      </div>
      <div class="text-lg font-medium text-gray-700">
        <!-- Titre dynamique de la période courante -->
        <span>{currentViewTitle}</span>
      </div>
      <div class="w-28">
        <!-- Élément vide pour maintenir l'alignement -->
      </div>
    </div>

    <!-- Calendrier -->
    <div class="border border-gray-200 rounded-lg overflow-hidden">
      <div bind:this={calendarEl} class="w-full"></div>
    </div>
  </div>
</div>

{#if showModal}
  <ActivityModal
    show={showModal}
    {users}
    {projects}
    {categories}
    activityToEdit={editActivity}
    {selectedDate}
    onDelete={handleActivityDelete}
    onSubmit={handleActivitySubmit}
    onUpdate={handleActivityUpdate}
    onClose={() => {
      showModal = false;
    }}
  />
{/if}
