<script lang="ts">
  import { projectTemplate } from '../../forms/project';
  import { ProjectApiService } from '../../services/ProjectApiService';
  import { UserApiService } from '../../services/UserApiService';
  import type { ProjectBase, User, Project } from '../../Models/index';
  import '../../style/app.css';
  import { X } from 'lucide-svelte';
  import { onMount } from 'svelte';

  type Props = {
    show: boolean;
    projectIdToEdit: number | null;
    onClose: () => void;
    onSubmit: (project: ProjectBase) => void;
    onUpdate: (project: Project) => void;
  };

  let { show, projectIdToEdit, onClose, onSubmit, onUpdate }: Props = $props();

  const editMode = $derived(projectIdToEdit !== null);

  const getInitialProject = () => ({ ...projectTemplate.generate() });

  let isSubmitting = $state(false);
  let isLoadingManagers = $state(true);
  let isLoadingProject = $state(false);
  let managers = $state<User[]>([]);
  let error = $state<string | null>(null);
  let project = $state<ProjectBase>(getInitialProject());
  let fullProject = $state<Project | null>(null);

  async function loadProjectIfNeeded() {
    if (show && editMode && projectIdToEdit) {
      try {
        isLoadingProject = true;
        fullProject = await ProjectApiService.getProject(projectIdToEdit);

        project = {
          name: fullProject.name,
          description: fullProject.description,
          manager_id: fullProject.manager_id,
          billable: fullProject.billable,
          status: fullProject.status,
        };

        isLoadingProject = false;
      } catch (err) {
        console.error(err);
        error = 'Impossible de charger le projet';
        isLoadingProject = false;
      }
    }
  }

  $effect(() => {
    if (!show) {
      project = getInitialProject();
      fullProject = null;
      error = null;
      isLoadingProject = false;
    }
  });

  $effect(() => {
    if (show) {
      if (editMode && projectIdToEdit) {
        loadProjectIfNeeded();
      } else {
        project = getInitialProject();
      }
    }
  });

  onMount(async () => {
    try {
      isLoadingManagers = true;
      error = null;
      managers = await UserApiService.getManagerUsers();
    } catch (err) {
      console.error('Failed to load managers:', err);
      error = 'Impossible de charger la liste des managers';
    } finally {
      isLoadingManagers = false;
    }
  });

  const handleClose = () => {
    console.log('handleClose called');
    onClose();
  };

  const handleSubmit = async () => {
    if (isSubmitting) return;
    isSubmitting = true;
    try {
      if (editMode && fullProject) {
        const updatedProject = await ProjectApiService.updateProject({
          ...fullProject,
          ...project,
        });
        onUpdate(updatedProject);
      } else {
        const newProject = await ProjectApiService.createProject(project);
        onSubmit(newProject);
      }
      onClose();
    } catch (err) {
      console.error('Erreur lors de la soumission du projet', err);
      error = 'Une erreur est survenue. Veuillez réessayer.';
    } finally {
      isSubmitting = false;
    }
  };
</script>

<style>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }

  .modal {
    background-color: white;
    border-radius: 4px;
    width: 400px;
    max-width: 90%;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }

  .modal-header {
    padding: 12px 24px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #eee;
  }

  .modal-title {
    font-size: 18px;
    margin: 0;
    color: #333;
  }

  .modal-content {
    padding: 24px;
  }

  .form-group {
    margin-bottom: 16px;
  }

  label {
    display: block;
    margin-bottom: 8px;
    color: #666;
    font-size: 14px;
  }

  input,
  textarea,
  select {
    width: 100%;
    padding: 8px 12px;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-sizing: border-box;
  }

  input[type='checkbox'] {
    width: auto;
    margin-right: 8px;
  }

  .modal-footer {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
    margin-top: 24px;
  }

  .error-text {
    color: #e53e3e;
    font-size: 14px;
    margin-top: 4px;
  }
</style>

{#if show}
  <div class="modal-overlay">
    <div class="modal">
      <div class="modal-header">
        <h2 class="modal-title">{editMode ? 'Modifier le projet' : 'Créer un nouveau projet'}</h2>
        <button type="button" class="text-black hover:text-gray-600" onclick={handleClose}>
          <X />
        </button>
      </div>

      <div class="modal-content">
        {#if isLoadingProject}
          <div class="py-2 px-4 bg-gray-100 rounded">Chargement du projet...</div>
        {:else}
          <form class="flex flex-col h-full" onsubmit={handleSubmit}>
            <div class="form-group">
              <label for="project-name">Identifiant unique du projet*</label>
              <input id="project-name" name="name" type="text" bind:value={project.name} required />
            </div>

            <div class="form-group">
              <label for="project-manager">Chargé de projet*</label>
              {#if isLoadingManagers}
                <div class="py-2 px-4 bg-gray-100 rounded">Chargement des managers...</div>
              {:else if error}
                <div class="error-text">{error}</div>
              {:else}
                <select
                  id="project-manager"
                  name="manager_id"
                  bind:value={project.manager_id}
                  required
                >
                  <option value="">-- Sélectionner un manager --</option>
                  {#each managers as manager}
                    <option value={manager.id}>
                      {manager.first_name}
                      {manager.last_name}
                    </option>
                  {/each}
                </select>
              {/if}
            </div>

            <div class="form-group">
              <label for="project-description">Description</label>
              <textarea
                id="project-description"
                name="description"
                bind:value={project.description}
                rows="3"
              >
              </textarea>
            </div>

            <div class="form-group">
              <label>
                <input type="checkbox" bind:checked={project.billable} />
                Facturable
              </label>
            </div>

            <div class="modal-footer">
              {#if editMode}
                <button
                  type="button"
                  class="py-3 px-6 bg-gray-100 text-gray-700 rounded-lg font-medium hover:bg-gray-200 hover:-translate-y-0.5 hover:shadow-sm active:translate-y-0 transition border border-gray-200"
                  onclick={handleClose}
                >
                  Retour
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
                  onclick={handleClose}
                >
                  Annuler
                </button>
                <button
                  type="submit"
                  class="py-3 px-6 bg-[#015e61] text-white rounded-lg font-medium hover:bg-[#014446] hover:-translate-y-0.5 hover:shadow-md active:translate-y-0 transition disabled:opacity-50"
                  disabled={isSubmitting}
                >
                  {isSubmitting ? 'En cours...' : 'Soumettre'}
                </button>
              {/if}
            </div>
          </form>
        {/if}
      </div>
    </div>
  </div>
{/if}
