<script lang="ts">
  import { CalendarService } from '../services/calendar.service';
  import { onMount } from 'svelte';
  import TaskModal from '../Components/Calendar/TaskModal.svelte';
  import '../style/calendar.css';
  import { TaskApiService } from '../services/TaskApiService';
  import type { Task } from '../Models';

  let calendarEl = $state<HTMLElement | null>(null);
  let calendarService = $state<CalendarService | null>(null);
  let showModal = $state(false);
  let selectedDate: { start: Date; end: Date } | null = null;
  let editMode = $state(false);
  let editTask = $state(null);

  const users = [];
  const projects = [];
  const categories = [];

  // async function loadTasks() {
  //   try {
  //     const tasks = await TaskApiService.getTasks();
  //     tasks.forEach((task) => {
  //       calendarService?.addEvent({
  //         title: task.name,
  //         start: task.startDateTime,
  //         end: task.endDateTime,
  //         extendedProps: { ...task },
  //       });
  //     });
  //   } catch (error) {
  //     console.error('Error loading tasks', error);
  //   }
  // }

  onMount(() => {
    if (calendarEl) {
      calendarService = new CalendarService();

      calendarService.onDateSelect = (info) => {
        editMode = false;
        editTask = null;
        selectedDate = info;
        showModal = true;
      };

      calendarService.onEventClick = (info) => {
        editMode = true;
        editTask = {
          id: info.event.extendedProps.id,
          name: info.event.title,
          description: info.event.extendedProps.description,
          state: info.event.extendedProps.state,
          billable: info.event.extendedProps.billable,
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

      calendarService.initialize(calendarEl);
      calendarService.render();
      // loadTasks();
    }
  });

  async function handleTaskSubmit(taskData: Task) {
    if (!calendarService) return;
    try {
      // 1. Créer la tâche via l'API
      const newTask = await TaskApiService.createTask(taskData);

      // 2. Mettre à jour le calendrier
      calendarService.addEvent({
        id: newTask.id.toString(),
        title: newTask.name,
        start: newTask.startDateTime,
        end: newTask.endDateTime,
        extendedProps: { ...newTask },
      });
    } catch (error) {
      console.error("Erreur lors d'ajout de tâche", error);
      throw error;
    }
  }

  async function handleTaskUpdate(task: Task) {
    if (!calendarService?.calendar) return;
    try {
      // 1. Mettre à jour via l'API
      const updatedTask = await TaskApiService.updateTask(task);

      // 2. Mettre à jour le calendrier
      calendarService.updateEvent(updatedTask);
    } catch (error) {
      console.error('Erreur lors de la mise à jour de la tâche', error);
      throw error;
    }
  }

  async function handleTaskDelete(task: Task) {
    if (!calendarService?.calendar || !task.id) return;
    try {
      // 1. Supprimer via l'API
      await TaskApiService.deleteTask(task.id);

      // 2. Mettre à jour le calendrier
      calendarService.deleteTask(task.id.toString());
    } catch (error) {
      console.error('Erreur lors de la suppression de la tâche', error);
      throw error;
    }
  }
</script>

<style>
</style>

<div class="calendar-container">
  <div bind:this={calendarEl}></div>
</div>

{#if showModal}
  <TaskModal
    show={showModal}
    {users}
    {projects}
    {categories}
    taskToEdit={editTask}
    onDelete={handleTaskDelete}
    onSubmit={handleTaskSubmit}
    onUpdate={handleTaskUpdate}
    onClose={() => {
      showModal = false;
    }}
  />
{/if}
