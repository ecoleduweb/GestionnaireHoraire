<script lang="ts">
  import { createForm } from 'felte'
  import ActivityValidationSchema from '../../Validation/Activity';
  import type { Activity, User, Project, Category } from '../../Models';
  import { activityTemplate } from '../../forms/activity';
  import { ActivityApiService } from '../../services/ActivityApiService';
  import {
    getHoursFromDate,
    getMinutesFromDate,
    createDateWithTime,
    initializeActivityDates,
    applyEndTime as applyEndTimeUtil,
  } from '../../utils/date';
  import '../../style/app.css';
  import { on } from 'svelte/events';

  type Props = {
    show: boolean;
    users: User[];
    projects: Project[];
    categories: Category[];
    activityToEdit: Activity | null;
    selectedDate?: { start: Date; end: Date } | null;
    onClose: () => void;
    onDelete: (activity: Activity) => void;
    onSubmit: (activity: Activity) => void;
    onUpdate: (activity: Activity) => void;
  };

  let {
    show,
    users,
    projects,
    categories,
    activityToEdit,
    selectedDate = null,
    onClose,
    onDelete,
    onSubmit,
    onUpdate,
  }: Props = $props();

  const editMode = activityToEdit !== null;

  let initialActivity = activityTemplate.generate();

  let isSubmitting = false;

  if (selectedDate && selectedDate.start) {
    const { startDate, endDate } = initializeActivityDates(selectedDate.start);
    initialActivity.startDate = startDate;
    initialActivity.endDate = endDate ? new Date(selectedDate.end) : endDate;
  }

  const activity = $state<Activity>(initialActivity);

  const time = $state({
    startHours: getHoursFromDate(activity.startDate),
    startMinutes: getMinutesFromDate(activity.startDate),
    endHours: getHoursFromDate(activity.endDate),
    endMinutes: getMinutesFromDate(activity.endDate),
  });

  if (activityToEdit) {
    Object.assign(activity, activityToEdit);
    time.startHours = getHoursFromDate(activityToEdit.startDate);
    time.startMinutes = getMinutesFromDate(activityToEdit.startDate);
    time.endHours = getHoursFromDate(activityToEdit.endDate);
    time.endMinutes = getMinutesFromDate(activityToEdit.endDate);
  }

  const {
    time: { hours, minutes },
  } = activityTemplate;

  const applyEndTime = () => {
    const result = applyEndTimeUtil(
      time.startHours,
      time.startMinutes,
      time.endHours,
      time.endMinutes
    );
    time.endHours = result.endHours;
    time.endMinutes = result.endMinutes;
  };

  const { form, errors } = createForm({
    validate: async (values) => {
      try {
        await ActivityValidationSchema.validate(values, { abortEarly: false });
      } catch(err) {
        
        const errors = err.inner.reduce((res, value) => ({
          [value.path]: value.message,
          ...res,
        }), {});
        
        return errors;
      }
    }
  });

  const handleSubmit = async () => {
    if (isSubmitting) return; // Empêche les soumissions multiples
    isSubmitting = true;

    try {
      if (!activity.name || !activity.userId || !activity.projectId || !activity.categoryId) {
        alert('Veuillez remplir tous les champs obligatoires.');
        return;
      }

      // Créer une nouvelle date de début avec les heures et minutes sélectionnées
      const updatedStartDate = createDateWithTime(
        activity.startDate,
        time.startHours,
        time.startMinutes
      );

      // Créer une nouvelle date de fin avec les heures et minutes sélectionnées
      const updatedEndDate = createDateWithTime(activity.endDate, time.endHours, time.endMinutes);

      // Mise à jour des dates de l'activité
      activity.startDate = updatedStartDate;
      activity.endDate = updatedEndDate;

      if (editMode) {
        const updatedActivity = await ActivityApiService.updateActivity(activity);
        onUpdate(updatedActivity);
      } else {
        const newActivity = await ActivityApiService.createActivity(activity);
        onSubmit(newActivity);
      }

      onClose();
    } catch (error) {
      console.error('Erreur lors de la soumission', error);
      alert('Une erreur est survenue. Veuillez réessayer.');
    } finally {
      isSubmitting = false;
    }
  };

  const handleDelete = async () => {
    if (!activity.id) return;

    if (confirm('Supprimer cette tâche ?')) {
      try {
        await ActivityApiService.deleteActivity(activity.id);
        onDelete(activity);
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
        {editMode ? "Modifier l'activité" : 'Nouvelle tâche'}
      </h2>
      <form use:form on:submit|preventDefault={handleSubmit}>
        <div class="mb-6">
          <label for="activity-name" class="block text-gray-600 mb-2">
            Nom
          <span class="text-red-500">*</span>
          </label>
          <input
            id="activity-name"
            name="name"
            type="text"
            bind:value={activity.name}
            placeholder="Nom de l'activité..."
            autofocus
            class="w-full py-3 px-3 border border-gray-300 rounded-md text-base focus:outline-none focus:ring-2 focus:ring-gray-500"
          />
          {#if $errors.name}
          <span class="text-red-500 text-sm">{$errors.name}</span>
          {/if}
        </div>

        <div class="mb-6">
          <label for="activity-description" class="block text-gray-600 mb-2">
            Description
            <span class="text-gray-500 text-sm">(optionnel)</span>
          </label>
          <textarea
            id="activity-description"
            name="description"
            placeholder="Description de l'activité..."
            bind:value={activity.description}
            rows="3"
            class="w-full py-3 px-3 border border-gray-300 rounded-md text-base focus:outline-none focus:ring-2 focus:ring-gray-500"
          ></textarea>
          {#if $errors.description}
          <span class="text-red-500 text-sm">{$errors.description}</span>
          {/if}
        </div>
        <div class="grid grid-cols-2 gap-4 mb-6">
          <div class="flex flex-col gap-2">
            <label class="text-gray-600">
              Heure de début
              <span class="text-red-500">*</span>
            </label>
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
            <label class="text-gray-600">
              Heure de fin
              <span class="text-red-500">*</span>
            </label>
            <div class="flex gap-2">
              <select
                bind:value={time.endHours}
                on:change={applyEndTime}
                class="p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
              >
                {#each hours as hour}
                  <option value={hour}>{hour}h</option>
                {/each}
              </select>
              <select
                bind:value={time.endMinutes}
                on:change={applyEndTime}
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
            <label for="activity-user" class="block text-gray-600 mb-2">
              Utilisateur
              <span class="text-red-500">*</span>
            </label>
            <select
              id="activity-user"
              name="userId"
              bind:value={activity.userId}
              required
              class="w-full p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
            >
              <option value="">Sélectionner...</option>
              {#each users as user}
                <option value={user.id}>{user.name}</option>
              {/each}
            </select>
            {#if $errors.userId}
            <span class="text-red-500 text-sm">{$errors.userId}</span>
            {/if}
          </div>

          <div>
            <label for="activity-project" class="block text-gray-600 mb-2">
              Projet
              <span class="text-red-500">*</span>
            </label>
            <select
              id="activity-project"
              name="projectId"
              bind:value={activity.projectId}
              required
              class="w-full p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
            >
              <option value="">Sélectionner...</option>
              {#each projects as project}
                <option value={project.id}>{project.name}</option>
              {/each}
            </select>
            {#if $errors.projectId}
            <span class="text-red-500 text-sm">{$errors.projectId}</span>
            {/if}
          </div>

          <div>
            <label for="activity-category" class="block text-gray-600 mb-2">
              Catégorie
              <span class="text-red-500">*</span>
            </label>
            <select
              id="activity-category"
              name="categoryId"
              bind:value={activity.categoryId}
              required
              class="w-full p-2 border border-gray-300 rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-gray-500"
            >
              <option value="">Sélectionner...</option>
              {#each categories as category}
                <option value={category.id}>{category.name}</option>
              {/each}
            </select>
            {#if $errors.categoryId}
            <span class="text-red-500 text-sm">{$errors.categoryId}</span>
            {/if}
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
              class="py-3 px-6 bg-gray-900 text-white rounded-md hover:bg-gray-700 {isSubmitting
                ? 'opacity-50 cursor-not-allowed'
                : ''}"
              disabled={isSubmitting}
            >
              {isSubmitting ? 'En cours...' : 'Créer'}
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
