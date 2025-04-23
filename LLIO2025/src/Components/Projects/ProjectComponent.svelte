<script lang="ts">
  import { Plus } from 'lucide-svelte';
  interface Project {
    id: string;
    name: string;
    lead: string;
    color: string;
    timeSpent: number;
    timeEstimated: number;
    timeRemaining: number;
  }

  export let project: Project;

  const formatHours = (hours: number | null): string => {
    return hours !== null ? `${hours}h00` : "-";
  };
</script>

<style>
  .project {
    width: auto;
    background-color: white;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    margin: 15px;
  }
</style>

<div class="project">
  <!-- Contenu du dashboard -->
  <div class="project-content">
    <div>
        <!-- Bordure gauche de couleur -->
        <div class="border-l-12 rounded" style="border-left-color: {project.color}">
          <div class="p-4">
            <div class="flex justify-between items-center">
              <div>
                <span class="text-xl">{project.id}</span>
                <span class="text-xl text-gray-500 ml-2">|</span>
                <span class="text-base text-gray-500 ml-2">{project.name}</span>
              </div>
            </div>
            <div class="flex-col justify-between items-center mt-1">
              <div class="mt-1 text-xs text-gray-400">Chargé·e de projet</div>
              <div class="text-sm wrap-normal">{project.lead}</div>
              <button
                class="mt-2 inline-flex items-center px-3 py-1.5 bg-gray-100 border border-transparent rounded-4xl shadow-sm hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-grey-500 text-gray-700 text-xs font-medium"
              >
                <span class="text-xs">Réattribuer</span>
              </button>
              <hr class="mt-2 text-xs text-gray-400"/>
              <div class="mt-1 text-xs text-gray-400">Co-chargé·e de projet</div>
              {#each project.coLeads as coLead}
                <div class="text-sm wrap-normal">{coLead}</div>
              {/each}
              <button 
                class="mt-2 inline-flex items-center  bg-gray-100 border border-transparent rounded-4xl shadow-sm hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-grey-500 text-gray-700 text-xs"
              >
                <Plus class="w-3 h-3" />
              </button>
            </div>
          </div>
        </div>
    </div>
  </div>
</div>
