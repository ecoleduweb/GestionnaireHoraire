<script lang="ts">
  import { goto } from "$app/navigation"
  import { quintOut } from "svelte/easing";
  import { slide } from "svelte/transition";
  import DashboardProjectItem from "../Projects/DashboardPaneProjectItem.svelte";
  import HoursWorkedDashboard from "./HoursWorkedDashboard.svelte";
  import { $props } from "svelte";
  const { activities }: { activities: any[] } = $props();

  let { projects } = $props();
  let isArchivedVisible = $state(false);

</script>

<div class="dashboard-container">
  <!-- En-tête du dashboard -->
  <div class="dashboard-header">Tableau de bord</div>

  <!-- Contenu du dashboard -->
  <div class="dashboard-content">
    <!-- Contenu à venir -->
    <div class="dashboard-item">
      <div class="inline-flex rounded-md shadow-xs" role="group">
        <button
          type="button"
          class="px-4 py-2 text-sm transition-colors font-semibold bg-[#014446] text-white rounded-l-lg"
        >
          Calendrier
        </button>
        <button 
          onclick={() => goto('./projects')}
          type="button" 
          class="py-2 px-4 text-sm transition-colors font-semibold bg-gray-200 text-gray-900 rounded-r-lg hover:bg-[#014446] hover:text-white cursor-pointer"
        >
          Projets
        </button>
      </div>
    </div>

    <!-- Projets en cours -->
    <div class="overflow-y-auto max-h-[calc(100vh-150px)]">
      {#each projects.filter(x => !x.isArchived) as project}
        <DashboardProjectItem project={project} />
      {/each}

      <!-- Projets archivés -->
      {#if projects.some(x => x.isArchived)}
      <div>
        <button
          class="w-full p-4 flex items-center justify-between text-gray-600 hover:bg-gray-50 cursor-pointer"
          onclick={() => isArchivedVisible = !isArchivedVisible}
        >
          <span class="font-medium">Projets archivés ({projects.filter(x => x.isArchived).length})</span>
          <svg
            class="w-5 h-5 transform transition-transform {isArchivedVisible ? 'rotate-180' : ''}"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </button>

        {#if isArchivedVisible}
          <div transition:slide={{ duration: 300, easing: quintOut }}>
            {#each projects.filter(x => x.isArchived) as project}
              <DashboardProjectItem project={project} />
            {/each}
          </div>
        {/if}
      </div>
      {/if}
    </div>
  </div>
  <HoursWorkedDashboard />
</div>

<style>
  .dashboard-container {
    width: 300px;
    height: 100vh;
    background-color: white;
    border-right: 1px solid #e5e7eb;
    box-shadow: 2px 0 5px rgba(0, 0, 0, 0.05);
    position: fixed;
    left: 0;
    top: 0;
    bottom: 0;
    z-index: 30;
    display: flex;
    flex-direction: column;
  }

  .dashboard-header {
    background-color: #005e61;
    color: white;
    padding: 16px;
    font-weight: 600;
    font-size: 1.25rem;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  }

  .dashboard-content {
    flex: 1;
    overflow-y: auto;
  }

  .dashboard-item {
    padding: 16px;
    border-bottom: 1px solid #f0f0f0;
    transition: background-color 0.2s;
  }

  .dashboard-item:hover {
    background-color: #f5f5f5;
  }
</style>