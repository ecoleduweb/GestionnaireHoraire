<script lang="ts">
  import { goto } from "$app/navigation"
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
          onclick={() => goto('/projects')}
          type="button"
          class="py-2 px-4 text-sm transition-colors font-semibold bg-gray-200 text-gray-900 rounded-r-lg hover:bg-[#014446] hover:text-white cursor-pointer"
        >
          Projets
        </button>
      </div>
    </div>
    <div class="overflow-y-auto max-h-[calc(100vh-250px)]">
      {#each projects as project}
        <div 
          class="border-l-10 hover:bg-gray-50 cursor-pointer border-b" 
          style="border-left-color: {project.color}"
        >
          <div class="p-4">
            <div class="flex justify-between items-center">
              <div>
                <span class="text-black">{project.id}</span>
                <span class="text-gray-500 ml-2">|</span>
                <span class="text-gray-500 ml-2">{project.name}</span>
              </div>
            </div>
            
            <div class="mt-2 flex items-center text-xs text-gray-500">
              <div class="mr-16">
                <div class="font-bold text-black">Total</div>
              </div>
              <div class="mr-4">
                <div class="font-bold text-black">{formatHours(project.timeSpent)}</div>
              </div>
              <div class="mr-4">
                <div class="text-gray-400">{formatHours(project.timeEstimated)}</div>
              </div>
              <div>
                {#if project.timeRemaining < 0}
                <div class="font-medium text-red-700">{formatHours(project.timeRemaining)}</div>
                {:else}
                <div class="text-gray-400">{formatHours(project.timeRemaining)}</div>
                {/if}
              </div>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>
