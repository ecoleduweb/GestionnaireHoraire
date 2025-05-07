<script lang="ts">
  import { projectTemplate } from '../../forms/project';
  import { ProjectApiService } from '../../services/ProjectApiService';
  import { UserApiService } from '../../services/UserApiService';
  import type { ProjectBase, User, Project } from '../../Models/index';
  import '../../style/app.css';
  import { X } from 'lucide-svelte';
  import { onMount } from 'svelte';
  import * as yup from 'yup';

  type Props = {
    show: boolean;
    projectToEdit: Project | null;
    onClose: () => void;
    onSubmit: (project: ProjectBase) => void;
    onUpdate: (project: Project) => void;
  };

  const schema = yup.object({
    name: yup.string().required('Le nom du projet est requis').max(50),
    manager_id: yup
      .number()
      .required('Un chargé de projet est requis')
      .positive('Un chargé de projet est requis'),
    description: yup.string().max(20000).nullable(),
    billable: yup.boolean().required(),
    status: yup.number(),
  });

  let { show, projectToEdit: projectToEdit, onClose, onSubmit, onUpdate }: Props = $props();

  const editMode = $derived(projectToEdit !== null);

  const getInitialProject = () => ({ ...projectTemplate.generate() });

  let errors = $state<Record<string, string>>({});

  const validateField = async (field: string) => {
    try {
      await (yup.reach(schema, field) as yup.Schema<any>).validate(project[field]);
      errors = { ...errors, [field]: '' };
    } catch (err) {
      if (hasSubmitted || touchedFields[field]) {
        errors = { ...errors, [field]: err.message };
      }
    }
  };

  let isSubmitting = $state(false);
  let isLoadingManagers = $state(true);
  let isLoadingProject = $state(false);
  let managers = $state<User[]>([]);
  let error = $state<string | null>(null);
  let project = $state<ProjectBase>(getInitialProject());
  let fullProject = $state<Project | null>(null);
  let hasSubmitted = $state(false);
  let touchedFields = $state<Record<string, boolean>>({});

  $effect(() => {
    (async () => {
      await validateField('name');
    })();
  });
  $effect(() => {
    (async () => {
      await validateField('manager_id');
    })();
  });
  $effect(() => {
    (async () => {
      await validateField('description');
    })();
  });
  $effect(() => {
    (async () => {
      await validateField('billable');
    })();
  });
  $effect(() => {
    (async () => {
      await validateField('status');
    })();
  });

  const loadProjectIfNeeded = async () => {
    if (show && editMode && projectToEdit) {
      try {
        isLoadingProject = true;
        fullProject = projectToEdit;

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
  };

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
      if (editMode && projectToEdit) {
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

    // Marquer tous les champs comme touchés et valider
    touchedFields = Object.keys(project).reduce(
      (acc, key) => {
        acc[key] = true;
        return acc;
      },
      {} as Record<string, boolean>
    );

    hasSubmitted = true;

    await Promise.all(Object.keys(project).map((field) => validateField(field)));

    const hasErrors = Object.values(errors).some((error) => error !== '');
    if (hasErrors) {
      error = 'Veuillez corriger les erreurs dans le formulaire';
      return;
    }

    isSubmitting = true;
    error = null;

    try {
      await schema.validate(project, { abortEarly: false });

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

      if (err instanceof yup.ValidationError) {
        error = err.errors.join(', ');
      } else {
        error = 'Une erreur est survenue. Veuillez réessayer.';
      }
    } finally {
      isSubmitting = false;
    }
  };
</script>

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
        {#if error}
          <div class="error-text">{error}</div>
        {/if}
        {#if isLoadingProject}
          <div class="py-2 px-4 bg-gray-100 rounded">Chargement du projet...</div>
        {:else}
          <form class="flex flex-col h-full" onsubmit={handleSubmit}>
            <div class="form-group">
              <label for="project-name">Identifiant unique du projet*</label>
              {#if editMode}
                <label>
                  <input
                    id="project-name"
                    name="name"
                    type="text"
                    bind:value={project.name}
                    readonly
                  />
                </label>
              {:else}
                <label>
                  <input
                    id="project-name"
                    name="name"
                    type="text"
                    bind:value={project.name}
                    onblur={() => {
                      touchedFields = { ...touchedFields, name: true };
                      validateField('name');
                    }}
                  />
                </label>
              {/if}
              {#if errors.name}
                <div class="error-text">{errors.name}</div>
              {/if}
            </div>

            <div class="form-group">
              <label for="project-manager">Chargé de projet*</label>
              {#if isLoadingManagers}
                <div class="py-2 px-4 bg-gray-100 rounded">Chargement des managers...</div>
              {:else}
                <select
                  id="project-manager"
                  name="manager_id"
                  bind:value={project.manager_id}
                  onblur={() => {
                    touchedFields = { ...touchedFields, manager_id: true };
                    validateField('manager_id');
                  }}
                >
                  <option value="">-- Sélectionner un manager --</option>
                  {#each managers as manager}
                    <option value={manager.id}>
                      {manager.first_name}
                      {manager.last_name}
                    </option>
                  {/each}
                </select>
                {#if errors.manager_id}
                  <div class="error-text">{errors.manager_id}</div>
                {/if}
              {/if}
            </div>

            <div class="form-group">
              <label for="project-description">Description</label>
              <textarea
                id="project-description"
                name="description"
                bind:value={project.description}
                rows="3"
                onblur={() => {
                  touchedFields = { ...touchedFields, description: true };
                  validateField('description');
                }}
              >
              </textarea>
              {#if errors.description}
                <div class="error-text">{errors.description}</div>
              {/if}
            </div>

            <div class="form-group">
              {#if editMode}
                <label>
                  <input type="checkbox" bind:checked={project.billable} disabled />
                  Facturable
                </label>
              {:else}
                <label>
                  <input type="checkbox" bind:checked={project.billable} />
                  Facturable
                </label>
              {/if}
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
