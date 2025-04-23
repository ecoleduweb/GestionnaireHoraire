<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { ProjectStatus } from "../../lib/types/enums";
  import type { CreateProject } from '../../Models/index';

  const dispatch = createEventDispatcher();
  
  export let isOpen = false;
  
  let projectToAdd: CreateProject = {
    name: '',
    description: '',
    status: ProjectStatus.NotStart,
    billable: false
  };
  
  function handleClose() {
    dispatch('close');
  }
  
  function handleSubmit() {
    dispatch('submit', projectToAdd); // Envoyer les données au composant parent
    
    handleClose();
    
    // Réinitialiser le formulaire
    projectToAdd = {
      name: '',
      description: '',
      status: ProjectStatus.NotStart,
      billable: false
    };
  }
</script>

{#if isOpen }
<div class="modal-overlay">
<div class="modal">
  <div class="modal-header">
    <h2 class="modal-title">Créer un nouveau projet</h2>
    <button class="close-button" on:click={handleClose}>×</button>
  </div>
  
  <div class="modal-content">
    <form on:submit|preventDefault={handleSubmit}>
      <div class="form-group">
        <label for="projectName">Identifiant unique du projet*</label>
        <input 
          type="text" 
          id="projectName" 
          maxlength="50"
          bind:value={projectToAdd.name} 
          required
        />
      </div>
      
      <div class="form-group">
        <label for="projectDescription">Description</label>
        <textarea 
          id="projectDescription" 
          bind:value={projectToAdd.description}
          rows="3"
        ></textarea>
      </div>
              
      <div class="form-group">
        <label>
          <input type="checkbox" bind:checked={projectToAdd.billable}>
          Facturable
        </label>
      </div>
      
      <div class="modal-footer">
        <button type="button" class="btn-cancel" on:click={handleClose}>Annuler</button>
        <button type="submit" class="btn-create">Créer</button>
      </div>
    </form>
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

.close-button {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
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

input, select, textarea {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

textarea {
  resize: vertical;
}

input[type="checkbox"] {
  width: auto;
  margin-right: 8px;
}

.time-inputs {
  display: flex;
  gap: 8px;
}

.time-inputs select {
  flex: 1;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 24px;
}

.btn-create {
  background-color: #014446;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-weight: 500;
  transition: background-color 0.2s;
}

.btn-create:hover {
  background-color: #015658;
}

.btn-cancel {
  background-color: transparent;
  border: none;
  color: #666;
  padding: 8px 16px;
  cursor: pointer;
}
</style>