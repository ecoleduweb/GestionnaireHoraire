<script lang="ts">
  import { validateProjectForm } from '../../Validation/Project';
  import { projectTemplate } from '../../forms/project';
  import { ProjectStatus } from '../../lib/types/enums';
  import { ProjectApiService } from '../../services/ProjectApiService';
  import type { CreateProject } from '../../Models/index';
  import '../../style/app.css';
  import { X } from 'lucide-svelte';
  type Props = {
    show: boolean;
    projectToEdit: CreateProject | null;
    onClose: () => void;
    onSubmit: (project: CreateProject) => void;
    onUpdate: (project: CreateProject) => void;
  };

  let { show, projectToEdit, onClose, onSubmit, onUpdate }: Props = $props();

  const editMode = projectToEdit !== null;

  let initialProject = projectTemplate.generate();

  let isSubmitting = false;

  const project = $state<CreateProject>(initialProject);

  if (projectToEdit) {
    Object.assign(project, projectToEdit);
  }
  
  const handleClose = () => {
    onClose();
  };

  const handleSubmit = async () => {
    if (isSubmitting) return; // Empêche les soumissions multiples
    isSubmitting = true;

    try {
      if (editMode) {
        // On appelle le update de l'api
      } else {
        const newProject = await ProjectApiService.createProject(project);
        onSubmit(newProject);
      }

      onClose();
    } catch (error) {
      console.error('Erreur lors de la soumission du projet', error);
      alert('Une erreur est survenue. Veuillez réessayer.');
    } finally {
      isSubmitting = false;
    }
  };

  const { form, errors } = validateProjectForm(handleSubmit, project);
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
  textarea {
    width: 100%;
    padding: 8px 12px;
    border: 1px solid #ccc;
    border-radius: 4px;
    box-sizing: border-box;
  }

  textarea {
    resize: vertical;
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
</style>

{#if show}
  <div class="modal-overlay">
    <div class="modal">
      <div class="modal-header">
        <h2 class="modal-title">Créer un nouveau projet</h2>
        <button type="button" class="text-white hover:text-gray-200" onclick={handleClose}>
          <X />
        </button>
      </div>

      <div class="modal-content">
        <form
          class="flex flex-col h-full"
          use:form
          onsubmit={(e) => {
            e.preventDefault();
          }}
        >
          <div class="form-group">
            <label for="project-name">Identifiant unique du projet*</label>
            <input id="project-name" name="name" type="text" bind:value={project.name} required />
          </div>

          <div class="form-group">
            <label for="project-description">Description</label>
            <textarea id="project-description" bind:value={project.description} rows="3"></textarea>
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
                {isSubmitting ? 'En cours...' : 'Créer'}
              </button>
            {/if}
          </div>
        </form>
      </div>
    </div>
  </div>
{/if}
