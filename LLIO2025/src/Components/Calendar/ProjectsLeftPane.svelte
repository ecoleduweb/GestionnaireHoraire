<script lang="ts">
  import { goto } from "$app/navigation"
  import { LogOut } from "lucide-svelte";
  import { onMount } from "svelte";

  let projects = [];

  const formatHours = (hours) => {
    return hours ? `${hours}h00` : "-";
  };

  function mockProjects() {
    return [
      {
        id: "AT-123",
        name: "Nommer le projet",
        color: "blue",
        lead: "Katell Arnault de la Ménardière",
        timeSpent: 205,
        timeEstimated: 300,
        timeRemaining: 95
      },
      {
        id: "AT-456",
        name: "Le projet de Marie Amélie",
        color: "pink",
        lead: "Katell Arnault de la Ménardière",
        timeSpent: 85,
        timeEstimated: 450,
        timeRemaining: 365
      },
      {
        id: "FO-115",
        name: "Graphisme 101",
        color: "yellow",
        lead: "Katell Arnault de la Ménardière",
        timeSpent: 40,
        timeEstimated: 0,
        timeRemaining: 0
      },
      {
        id: "RA-224",
        name: "Noisette Steve",
        color: "red",
        lead: "Katell Arnault de la Ménardière",
        timeSpent: 450,
        timeEstimated: 400,
        timeRemaining: -50
      }
    ];
  }

  onMount(() => {
    projects = mockProjects();
  });
</script>

<style>
  .dashboard-panel {
    width: 300px;
    height: 100vh;
    background-color: white;
    border-right: 1px solid #e5e7eb;
    box-shadow: 2px 0 5px rgba(0, 0, 0, 0.05);
    overflow-y: auto;
    position: fixed;
    left: 0;
    top: 0;
    bottom: 0;
    z-index: 30;
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
    padding: 0px;
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

<div class="dashboard-panel">
  <!-- En-tête du dashboard -->
  <div class="dashboard-header">Dashboard
    <button 
          class="ml-2 mt-1 p-1.5 rounded-full hover:bg-gray-100 text-gray-600 text-[#015e61] transition-colors" 
          title="Se déconnecter">
          <LogOut class="w-5 h-5" />
        </button>
  </div>

  <!-- Contenu du dashboard -->
  <div class="dashboard-content">
    <!-- Éléments du dashboard -->
    <div class="dashboard-item">
      <div class="inline-flex rounded-md shadow-xs" role="group">
        <button 
          onclick={() => goto('/calendar')}
          type="button" 
          class="py-2 px-4 text-sm transition-colors font-semibold bg-gray-200 text-gray-900 rounded-l-lg hover:bg-[#014446] hover:text-white cursor-pointer"
          >
          Calendrier
        </button>
        <button 
          type="button" 
          class="px-4 py-2 text-sm transition-colors font-semibold bg-[#014446] text-white rounded-r-lg"
          >
          Projets
        </button>
      </div>
      <button 
          type="button" 
          class="ml-6 px-3 py-2 text-sm transition-colors font-semibold bg-[#014446] text-white rounded-lg"
          >
          Créer
        </button>
    </div>

    <div class="overflow-y-auto max-h-[calc(100vh-250px)]">
      {#each projects as project}
        <div 
          class="border-l-4 hover:bg-gray-50 cursor-pointer border-b" 
          style="border-left-color: {project.color}"
        >
          <div class="p-4">
            <div class="flex justify-between items-center">
              <div>
                <span class="text-sm font-medium text-gray-600">{project.id}</span>
                <span class="text-xs text-gray-500 ml-2">|</span>
                <span class="text-xs text-gray-500 ml-2">{project.lead}</span>
              </div>
            </div>
            <div class="mt-1 text-sm font-medium">{project.name}</div>
            
            <div class="mt-2 flex items-center text-xs text-gray-500">
              <div class="mr-4">
                <span>Temps passé</span>
                <div class="font-medium text-gray-700">{formatHours(project.timeSpent)}</div>
              </div>
              <div class="mr-4">
                <span>Temps estimé</span>
                <div class="font-medium text-black-700">{formatHours(project.timeEstimated)}</div>
              </div>
              <div>
                <span>Temps restant</span>
                {#if project.timeRemaining < 0}
                <div class="font-medium text-red-700">{formatHours(project.timeRemaining)}</div>
                {:else}
                <div class="font-medium text-black-700">{formatHours(project.timeRemaining)}</div>
                {/if}
              </div>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>
