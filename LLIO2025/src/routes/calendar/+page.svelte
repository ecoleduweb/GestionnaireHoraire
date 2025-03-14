<script lang="ts">
  import type { CalendarService } from '../../services/calendar.service';
  import { CalendarService as CS } from '../../services/calendar.service';
  import { onMount } from 'svelte';
  import ActivityModal from '../../Components/Calendar/ActivityModal.svelte';
  import { ActivityApiService } from '../../services/ActivityApiService';
  import type { Activity } from '../../Models/index.ts';
  // Importez le fichier CSS
  import '../../style/modern-calendar.css';

  let calendarEl = $state<HTMLElement | null>(null);
  let calendarService = $state<CalendarService | null>(null);
  let showModal = $state(false);
  let selectedDate: { start: Date; end: Date } | null = null;
  let editMode = $state(false);
  let editActivity = $state(null);
  let activeView = $state('timeGridWeek');
  let currentViewTitle = $state('');

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
      const view = calendarService.calendar.view;

      if (view.type === 'dayGridMonth') {
        // Format pour la vue mois: "Février 2025"
        currentViewTitle = dateAPI.toLocaleDateString('fr-FR', {
          month: 'long',
          year: 'numeric',
        });
      } else if (view.type === 'timeGridWeek') {
        // Format pour la vue semaine: "24 févr. – 2 mars 2025"
        const endDate = new Date(dateAPI);
        endDate.setDate(dateAPI.getDate() + 6);

        const startFormatted = dateAPI.getDate();
        const startMonth = dateAPI.toLocaleDateString('fr-FR', { month: 'short' });

        const endFormatted = endDate.getDate();
        const endMonth = endDate.toLocaleDateString('fr-FR', { month: 'long' });
        const year = endDate.getFullYear();

        currentViewTitle = `${startFormatted} ${startMonth}. – ${endFormatted} ${endMonth} ${year}`;
      } else if (view.type === 'timeGridDay') {
        // Format pour la vue jour: "Mercredi 26 février 2025"
        currentViewTitle = dateAPI.toLocaleDateString('fr-FR', {
          weekday: 'long',
          day: 'numeric',
          month: 'long',
          year: 'numeric',
        });
      }
    }
  }

  onMount(() => {
    if (calendarEl) {
      calendarService = new CS();

      // Configuration personnalisée pour FullCalendar
      const calendarOptions = {
        initialView: activeView,
        buttonText: {
          today: 'today',
          month: 'Month',
          week: 'Week',
          day: 'Day',
        },
        slotDuration: '00:30:00', // Durée de chaque intervalle de temps
        allDaySlot: false,
        slotMinTime: '06:00:00',
        slotMaxTime: '20:00:00',
        nowIndicator: true,
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
          hour12: true,
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
          state: info.event.extendedProps.state,
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
      const updatedActivity = await ActivityApiService.updateActivity(activity);
      calendarService.updateEvent(updatedActivity);
    } catch (error) {
      console.error('Erreur lors de la mise à jour de la tâche', error);
      throw error;
    }
  }

  async function handleActivityDelete(activity: Activity) {
    if (!calendarService?.calendar || !activity.id) return;
    try {
      await ActivityApiService.deleteActivity(activity.id);
      calendarService.deleteActivity(activity.id.toString());
    } catch (error) {
      console.error('Erreur lors de la suppression de la tâche', error);
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
    height: 25px !important; /* Hauteur réduite des slots */
    min-height: 25px !important;
  }

  :global(.fc-timegrid-event) {
    min-height: 20px !important;
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
        <svg
          class="w-6 h-6 mr-2 text-gray-700"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="1.5"
            d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
          ></path>
        </svg>
        <h1 class="text-xl font-semibold text-gray-800">Today, {formattedDate}</h1>
      </div>

      <!-- Boutons de vue alignés au centre -->
      <div class="flex items-center bg-gray-100 p-1 rounded-lg">
        <button
          class="px-5 py-2 rounded-lg transition-colors {activeView === 'timeGridDay'
            ? 'bg-white text-[#015e61] font-medium'
            : 'text-gray-500 hover:bg-white hover:text-[#015e61]'}"
          on:click={() => setView('timeGridDay')}
        >
          Day
        </button>
        <button
          class="px-5 py-2 rounded-lg transition-colors {activeView === 'timeGridWeek'
            ? 'bg-white text-[#015e61] font-medium'
            : 'text-gray-500 hover:bg-white hover:text-[#015e61]'}"
          on:click={() => setView('timeGridWeek')}
        >
          Week
        </button>
        <button
          class="px-5 py-2 rounded-lg transition-colors {activeView === 'dayGridMonth'
            ? 'bg-white text-[#015e61] font-medium'
            : 'text-gray-500 hover:bg-white hover:text-[#015e61]'}"
          on:click={() => setView('dayGridMonth')}
        >
          Month
        </button>
      </div>

      <!-- Bouton New Activity -->
      <button
        on:click={handleNewActivity}
        class="bg-[#015e61] hover:bg-[#014446] text-white py-2 px-6 rounded-xl flex items-center gap-2 font-semibold transition-colors"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-5 w-5"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fill-rule="evenodd"
            d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z"
            clip-rule="evenodd"
          ></path>
        </svg>
        New Activity
      </button>
    </div>

    <!-- Contrôles de navigation -->
    <div class="flex items-center justify-between mb-4">
      <div class="flex items-center space-x-3">
        <button
          on:click={prevPeriod}
          class="p-2 rounded-lg bg-gray-100 hover:bg-gray-200 transition-colors"
        >
          <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 19l-7-7 7-7"
            ></path>
          </svg>
        </button>
        <button
          on:click={nextPeriod}
          class="p-2 rounded-lg bg-gray-100 hover:bg-gray-200 transition-colors"
        >
          <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"
            ></path>
          </svg>
        </button>
        <button
          on:click={goToday}
          class="px-4 py-2 rounded-lg bg-gray-100 hover:bg-gray-200 text-gray-700 transition-colors"
        >
          today
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
