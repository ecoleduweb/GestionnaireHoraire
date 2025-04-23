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
      // Validation et soumission du formulaire
      // Appelle vers l'API
      // Ecoute et attend le retour
      handleClose();
      
      projectToAdd = {
        name: '',
        description: '',
        status: ProjectStatus.NotStart,
        billable: false
      };
    }
  </script>
  
  {#if isOpen}
  <div class="modal-overlay">
    <div class="modal">
      <div class="modal-header">
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
                    
          <!-- <div class="form-group">
            <label>Temps estimé</label>
            <div class="time-inputs">
              <select>
                <option value="00">00h</option>
                {#each Array(24) as _, i}
                  <option value={i < 10 ? `0${i}` : `${i}`}>{i < 10 ? `0${i}` : `${i}`}h</option>
                {/each}
              </select>
              
              <select>
                <option value="00">00</option>
                {#each [0, 15, 30, 45] as min}
                  <option value={min < 10 ? `0${min}` : `${min}`}>{min < 10 ? `0${min}` : `${min}`}</option>
                {/each}
              </select>
            </div>
          </div> -->
          
          <div class="form-group">
            <label>
              <input type="checkbox" bind:checked={projectToAdd.billable}>
              Facturable
            </label>
          </div>
          
          <div class="modal-footer">
            <button type="submit" class="btn-create">Créer</button>
            <button type="button" class="btn-cancel" on:click={handleClose}>Annuler</button>
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
      padding: 12px;
      text-align: right;
    }
    
    .close-button {
      background: none;
      border: none;
      font-size: 24px;
      cursor: pointer;
    }
    
    .modal-content {
      padding: 0 24px 24px;
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
    
    input, select {
      width: 100%;
      padding: 8px 12px;
      border: 1px solid #ccc;
      border-radius: 4px;
      box-sizing: border-box;
    }
    
    input[type="checkbox"] {
      width: auto;
      margin-right: 8px;
    }
    
    .add-secondary {
      background: none;
      border: none;
      color: #666;
      padding: 8px 0;
      cursor: pointer;
      font-size: 14px;
      text-decoration: underline;
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
      background-color: #1a202c;
      color: white;
      border: none;
      padding: 8px 16px;
      border-radius: 4px;
      cursor: pointer;
    }
    
    .btn-cancel {
      background-color: transparent;
      border: none;
      color: #666;
      padding: 8px 16px;
      cursor: pointer;
    }
  </style>