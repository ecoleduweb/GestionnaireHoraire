<script lang="ts">
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

<style>
  /* Animation pour le panneau latéral - ne peut pas être fait en Tailwind standard */
  @keyframes slideIn {
    from {
      transform: translateX(100%);
    }
    to {
      transform: translateX(0);
    }
  }

  .animate-slideIn {
    animation: slideIn 0.3s forwards;
  }

  /* Classe pour les flèches de sélection */
  .select-arrow {
    background-image: url('data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="%23606060"%3E%3Cpath stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"%3E%3C/path%3E%3C/svg%3E');
    background-position: right 0.75rem center;
    background-size: 1em;
    background-repeat: no-repeat;
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    padding-right: 2.5rem; /* Espace pour la flèche */
  }

  /* Supprimer la flèche par défaut sur IE */
  .select-arrow::-ms-expand {
    display: none;
  }
</style>

{#if show}
  <!-- Structure principale avec Tailwind -->
  <div class="fixed inset-0 z-40 flex justify-end">
    <!-- Overlay semi-transparent avec opacité à 40% comme dans l'original -->
    <div
      class="absolute inset-0 bg-gray bg-opacity-40 transition-opacity"
      on:click={handleClose}
    ></div>

    <!-- Panneau latéral avec bordure et ombre à gauche pour délimiter -->
    <div
      class="w-full max-w-[480px] bg-white h-full overflow-y-auto relative flex flex-col z-50 animate-slideIn border-l border-gray-300 shadow-xl"
    >
      <!-- En-tête avec titre et bouton de fermeture -->
      <div class="flex items-center justify-between bg-[#015e61] text-white px-6 py-4">
        <h2 class="text-xl font-medium">
          {editMode ? "Modifier l'activité" : 'Nouvelle activité'}
        </h2>
        <button type="button" class="text-white hover:text-gray-200" on:click={handleClose}>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            ></path>
          </svg>
        </button>
      </div>

      <!-- Contenu du formulaire - espace vertical ajusté -->
      <div class="p-6 flex-grow">
        <form on:submit|preventDefault={handleSubmit} class="flex flex-col h-full">
          <!-- Champs de formulaire avec espacement vertical uniforme -->
          <div class="space-y-6">
            <!-- Champ Nom -->
            <div>
              <label for="activity-name" class="block text-gray-700 font-medium mb-2">Nom*</label>
              <input
                id="activity-name"
                type="text"
                bind:value={activity.name}
                placeholder="Nom de l'activité..."
                required
                autofocus
                class="w-full py-3 px-4 border border-gray-300 rounded-lg focus:outline-none focus:border-[#015e61] focus:ring-2 focus:ring-[#015e61] focus:ring-opacity-20 transition"
              />
            </div>

            <!-- Champ Description -->
            <div>
              <label for="activity-description" class="block text-gray-700 font-medium mb-2"
                >Description*</label
              >
              <textarea
                id="activity-description"
                bind:value={activity.description}
                required
                rows="3"
                class="w-full py-3 px-4 border border-gray-300 rounded-lg focus:outline-none focus:border-[#015e61] focus:ring-2 focus:ring-[#015e61] focus:ring-opacity-20 transition"
              ></textarea>
            </div>

            <!-- Sélecteurs d'heure côte à côte -->
            <div>
              <label class="block text-gray-700 font-medium mb-2">Heure de début*</label>
              <div class="flex gap-3">
                <select
                  bind:value={time.startHours}
                  class="flex-1 appearance-none bg-white border border-gray-300 rounded-lg py-3 px-4 text-gray-700 focus:outline-none focus:border-[#015e61] focus:ring-2 focus:ring-[#015e61] focus:ring-opacity-20 pr-10 bg-no-repeat transition select-arrow"
                >
                  {#each hours as hour}
                    <option value={hour}>{hour} h</option>
                  {/each}
                </select>
                <select
                  bind:value={time.startMinutes}
                  class="flex-1 appearance-none bg-white border border-gray-300 rounded-lg py-3 px-4 text-gray-700 focus:outline-none focus:border-[#015e61] focus:ring-2 focus:ring-[#015e61] focus:ring-opacity-20 pr-10 bg-no-repeat transition select-arrow"
                >
                  {#each minutes as minute}
                    <option value={minute}>{minute} min</option>
                  {/each}
                </select>
              </div>
            </div>

            <!-- Sélecteurs d'heure de fin -->
            <div>
              <label class="block text-gray-700 font-medium mb-2">Heure de fin*</label>
              <div class="flex gap-3">
                <select
                  bind:value={time.endHours}
                  on:change={applyEndTime}
                  class="flex-1 appearance-none bg-white border border-gray-300 rounded-lg py-3 px-4 text-gray-700 focus:outline-none focus:border-[#015e61] focus:ring-2 focus:ring-[#015e61] focus:ring-opacity-20 pr-10 bg-no-repeat transition select-arrow"
                >
                  {#each hours as hour}
                    <option value={hour}>{hour} h</option>
                  {/each}
                </select>
                <select
                  bind:value={time.endMinutes}
                  on:change={applyEndTime}
                  class="flex-1 appearance-none bg-white border border-gray-300 rounded-lg py-3 px-4 text-gray-700 focus:outline-none focus:border-[#015e61] focus:ring-2 focus:ring-[#015e61] focus:ring-opacity-20 pr-10 bg-no-repeat transition select-arrow"
                >
                  {#each minutes as minute}
                    <option value={minute}>{minute} min</option>
                  {/each}
                </select>
              </div>
            </div>

            <!-- Utilisateur -->
            <div>
              <label for="activity-user" class="block text-gray-700 font-medium mb-2"
                >Utilisateur*</label
              >
              <div class="relative">
                <select
                  id="activity-user"
                  bind:value={activity.userId}
                  required
                  class="w-full appearance-none bg-white border border-gray-300 rounded-lg py-3 px-4 text-gray-700 focus:outline-none focus:border-[#015e61] focus:ring-2 focus:ring-[#015e61] focus:ring-opacity-20 pr-10 bg-no-repeat transition select-arrow"
                >
                  <option value="">Sélectionner un utilisateur...</option>
                  {#each users as user}
                    <option value={user.id}>{user.name}</option>
                  {/each}
                </select>
              </div>
            </div>

            <!-- Projet -->
            <div>
              <label for="activity-project" class="block text-gray-700 font-medium mb-2"
                >Projet*</label
              >
              <div class="relative">
                <select
                  id="activity-project"
                  bind:value={activity.projectId}
                  required
                  class="w-full appearance-none bg-white border border-gray-300 rounded-lg py-3 px-4 text-gray-700 focus:outline-none focus:border-[#015e61] focus:ring-2 focus:ring-[#015e61] focus:ring-opacity-20 pr-10 bg-no-repeat transition select-arrow"
                >
                  <option value="">Sélectionner un projet...</option>
                  {#each projects as project}
                    <option value={project.id}>{project.name}</option>
                  {/each}
                </select>
              </div>
            </div>

            <!-- Catégorie -->
            <div>
              <label for="activity-category" class="block text-gray-700 font-medium mb-2"
                >Catégorie*</label
              >
              <div class="relative">
                <select
                  id="activity-category"
                  bind:value={activity.categoryId}
                  required
                  class="w-full appearance-none bg-white border border-gray-300 rounded-lg py-3 px-4 text-gray-700 focus:outline-none focus:border-[#015e61] focus:ring-2 focus:ring-[#015e61] focus:ring-opacity-20 pr-10 bg-no-repeat transition select-arrow"
                >
                  <option value="">Sélectionner une catégorie...</option>
                  {#each categories as category}
                    <option value={category.id}>{category.name}</option>
                  {/each}
                </select>
              </div>
            </div>
          </div>

          <!-- Séparateur et boutons d'action -->
          <div class="mt-auto">
            <!-- Ligne de séparation -->
            <div class="border-t border-gray-200 my-6"></div>

            <!-- Actions en bas du formulaire -->
            <div class="flex justify-end gap-3">
              {#if editMode}
                <button
                  type="button"
                  class="py-3 px-6 bg-red-500 text-white rounded-lg font-medium hover:bg-red-600 hover:-translate-y-0.5 hover:shadow-md active:translate-y-0 transition"
                  on:click={handleDelete}
                >
                  Supprimer
                </button>
                <button
                  type="submit"
                  class="py-3 px-6 bg-[#015e61] text-white rounded-lg font-medium hover:bg-[#014446] hover:-translate-y-0.5 hover:shadow-md active:translate-y-0 transition disabled:opacity-50"
                  disabled={isSubmitting}
                >
                  {isSubmitting ? 'En cours...' : 'Modifier'}
                </button>
              {:else}
                <button
                  type="button"
                  class="py-3 px-6 bg-gray-100 text-gray-700 rounded-lg font-medium hover:bg-gray-200 hover:-translate-y-0.5 hover:shadow-sm active:translate-y-0 transition border border-gray-200"
                  on:click={handleClose}
                >
                  Annuler
                </button>
                <button
                  type="submit"
                  class="py-3 px-6 bg-[#015e61] text-white rounded-lg font-medium hover:bg-[#014446] hover:-translate-y-0.5 hover:shadow-md active:translate-y-0 transition disabled:opacity-50"
                  disabled={isSubmitting}
                >
                  {isSubmitting ? 'En cours...' : 'Créer'}
                </button>
              {/if}
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
{/if}
