<script lang="ts">
  import type { Task, User, Project, Category } from '../../Models';
  import { taskTemplate } from '../../forms/task';
  import { TaskApiService } from '../../services/TaskApiService';
  import {
    getHoursFromDate,
    getMinutesFromDate,
    createDateWithTime,
    initializeTaskDates,
    validateEndTime as validateEndTimeUtil,
  } from '../../utils/date';
  import '../../style/app.css';

  type Props = {
    show: boolean;
    users: User[];
    projects: Project[];
    categories: Category[];
    taskToEdit: Task | null;
    selectedDate?: { start: Date; end: Date } | null;
    onClose: () => void;
    onDelete: (task: Task) => void;
    onSubmit: (task: Task) => void;
    onUpdate: (task: Task) => void;
  };

  let {
    show,
    users,
    projects,
    categories,
    taskToEdit,
    selectedDate = null,
    onClose,
    onDelete,
    onSubmit,
    onUpdate,
  }: Props = $props();

  const editMode = taskToEdit !== null;

  let initialTask = taskTemplate.generate();

  if (selectedDate && selectedDate.start) {
    const { startDate, endDate } = initializeTaskDates(selectedDate.start);
    initialTask.startDate = startDate;
    initialTask.endDate = endDate ? new Date(selectedDate.end) : endDate;
  }

  const task = $state<Task>(initialTask);

  const time = $state({
    startHours: getHoursFromDate(task.startDate),
    startMinutes: getMinutesFromDate(task.startDate),
    endHours: getHoursFromDate(task.endDate),
    endMinutes: getMinutesFromDate(task.endDate),
  });

  if (taskToEdit) {
    Object.assign(task, taskToEdit);
    time.startHours = getHoursFromDate(taskToEdit.startDate);
    time.startMinutes = getMinutesFromDate(taskToEdit.startDate);
    time.endHours = getHoursFromDate(taskToEdit.endDate);
    time.endMinutes = getMinutesFromDate(taskToEdit.endDate);
  }

  const {
    time: { hours, minutes },
  } = taskTemplate;

  const validateEndTime = () => {
    const result = validateEndTimeUtil(
      time.startHours,
      time.startMinutes,
      time.endHours,
      time.endMinutes
    );
    time.endHours = result.endHours;
    time.endMinutes = result.endMinutes;
  };

  const handleSubmit = async () => {
    if (task.name && task.userId && task.projectId && task.categoryId) {
      try {
        // Créer une nouvelle date de début basée sur la tâche existante et les heures/minutes sélectionnées
        const updatedStartDate = createDateWithTime(
          task.startDate,
          time.startHours,
          time.startMinutes
        );

        // Créer une date de fin basée sur la même date (même jour) que le début
        const updatedEndDate = createDateWithTime(
          task.startDate, // Utiliser la même date de base que le début
          time.endHours,
          time.endMinutes
        );

        // Mettre à jour les dates dans l'objet task
        task.startDate = updatedStartDate;
        task.endDate = updatedEndDate;

        if (editMode) {
          const updatedTask = await TaskApiService.updateTask(task);
          onUpdate(updatedTask);
        } else {
          const newTask = await TaskApiService.createTask(task);
          onSubmit(newTask);
        }
        onClose();
      } catch (error) {
        console.error('Erreur', error);
      }
    } else {
      alert('Veuillez remplir tous les champs obligatoires');
    }
  };

  const handleDelete = async () => {
    if (!task.id) return;

    if (confirm('Supprimer cette tâche ?')) {
      try {
        await TaskApiService.deleteTask(task.id);
        onDelete(task);
      } catch (error) {
        console.error('Erreur lors de la suppression', error);
      }
    }
  };

  const handleClose = () => {
    onClose();
  };
</script>

{#if show}
  <div class="fixed inset-0 flex items-center justify-center z-10" on:click={handleClose}>
    <div
      class="bg-white p-8 rounded-lg w-11/12 max-w-lg border-2 border-gray-300 shadow-xl"
      on:click|stopPropagation
    >
      <h2 class="text-2xl text-gray-800 font-medium mb-6">
        {editMode ? 'Modifier la tâche' : 'Nouvelle tâche'}
      </h2>
      <form on:submit|preventDefault={handleSubmit}>
        <div class="mb-6">
          <label for="task-name" class="block text-gray-600 mb-2">Nom*</label>
          <input
            id="task-name"
            type="text"
            bind:value={task.name}
            placeholder="Nom de la tâche..."
            required
            autofocus
            class="w-full py-3 px-3 border border-gray-300 rounded-md text-base focus:outline-none focus:ring-2 focus:ring-gray-500"
          />
        </div>

        <div class="mb-6">
          <label for="task-description" class="block text-gray-600 mb-2">Description</label>
          <textarea
            id="task-description"
            bind:value={task.description}
            rows="3"
            class="w-full py-3 px-3 border border-gray-300 rounded-md text-base focus:outline-none focus:ring-2 focus:ring-gray-500"
          ></textarea>
        </div>

        <div class="flex gap-4 mb-6">
          <div class="flex items-center">
            <label class="flex items-center gap-2 text-gray-600">
              <input type="checkbox" bind:checked={task.billable} class="h-5 w-5" />
              Facturable
            </label>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-4 mb-6">
          <div class="flex flex-col gap-2">
            <label class="text-gray-600">Heure de début*</label>
            <div class="flex gap-2">
              <select
                bind:value={time.startHours}
                class="p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
              >
                {#each hours as hour}
                  <option value={hour}>{hour}h</option>
                {/each}
              </select>
              <select
                bind:value={time.startMinutes}
                class="p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
              >
                {#each minutes as minute}
                  <option value={minute}>{minute}</option>
                {/each}
              </select>
            </div>
          </div>
          <div class="flex flex-col gap-2">
            <label class="text-gray-600">Heure de fin*</label>
            <div class="flex gap-2">
              <select
                bind:value={time.endHours}
                on:change={validateEndTime}
                class="p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
              >
                {#each hours as hour}
                  <option value={hour}>{hour}h</option>
                {/each}
              </select>
              <select
                bind:value={time.endMinutes}
                on:change={validateEndTime}
                class="p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
              >
                {#each minutes as minute}
                  <option value={minute}>{minute}</option>
                {/each}
              </select>
            </div>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
          <div>
            <label for="task-user" class="block text-gray-600 mb-2">Utilisateur*</label>
            <select
              id="task-user"
              bind:value={task.userId}
              required
              class="w-full p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
            >
              <option value="">Sélectionner...</option>
              {#each users as user}
                <option value={user.id}>{user.name}</option>
              {/each}
            </select>
          </div>

          <div>
            <label for="task-project" class="block text-gray-600 mb-2">Projet*</label>
            <select
              id="task-project"
              bind:value={task.projectId}
              required
              class="w-full p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
            >
              <option value="">Sélectionner...</option>
              {#each projects as project}
                <option value={project.id}>{project.name}</option>
              {/each}
            </select>
          </div>

          <div>
            <label for="task-category" class="block text-gray-600 mb-2">Catégorie*</label>
            <select
              id="task-category"
              bind:value={task.categoryId}
              required
              class="w-full p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
            >
              <option value="">Sélectionner...</option>
              {#each categories as category}
                <option value={category.id}>{category.name}</option>
              {/each}
            </select>
          </div>
        </div>

        <div class="flex justify-end gap-4">
          {#if editMode}
            <button
              type="button"
              class="py-3 px-6 bg-red-500 text-white rounded-md hover:bg-red-600"
              on:click={handleDelete}
            >
              Supprimer
            </button>
            <button
              type="submit"
              class="py-3 px-6 bg-gray-900 text-white rounded-md hover:bg-gray-700"
            >
              Modifier
            </button>
          {:else}
            <button
              type="submit"
              class="py-3 px-6 bg-gray-900 text-white rounded-md hover:bg-gray-700"
            >
              Créer
            </button>
          {/if}
          <button
            type="button"
            class="py-3 px-6 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
            on:click={handleClose}
          >
            Annuler
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}
