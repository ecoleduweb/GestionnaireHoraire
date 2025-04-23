<script>
  import { onMount } from "svelte";
  import { ClientTelemetry } from "$lib/tracer"
  import { env } from "$env/dynamic/public"
  import { goto } from "$app/navigation"
  import "../../style/app.css"
  import ProjectsLeftPane from "../../Components/Calendar/ProjectsLeftPane.svelte";
  const apiBaseUrl = import.meta.env.VITE_API_BASE_URL;
  const ENABLED_TELEMETRY = env.PUBLIC_ENABLED_TELEMETRY === "true";

  if (ENABLED_TELEMETRY) {
    const telemetry = ClientTelemetry.getInstance();
    telemetry.start();
  }

  // État utilisateur
  let userName = "Katell Arnault de la Ménardière";
  let userRole = "Chargé·e de projet";

  // État des projets
  let projects = [];
  let selectedProject = null;
  let loading = true;
  let error = null;

  // Fonction pour formater les heures
  const formatHours = (hours) => {
    return hours ? `${hours}h00` : "-";
  };

  // Fonction pour chercher les projets de l'API
  async function fetchProjects() {
    try {
      loading = true;
      const response = await fetch(`${apiBaseUrl}/projects`);
      
      if (!response.ok) {
        throw new Error(`Erreur API: ${response.status}`);
      }
      
      const data = await response.json();
      projects = data.map(project => ({
        ...project,
        expanded: false
      }));
      
      // Sélectionner le premier projet par défaut
      if (projects.length > 0) {
        selectProject(projects[0]);
      }
    } catch (err) {
      console.error("Erreur lors du chargement des projets:", err);
      error = err.message;
    } finally {
      loading = false;
    }
  }

  // Fonction pour sélectionner un projet
  function selectProject(project) {
    selectedProject = project;
    projects = projects.map(p => ({
      ...p,
      expanded: p.id === project.id ? !p.expanded : false
    }));
  }

  // Fonction pour se déconnecter
  function logout() {
    goto('/login');
  }

  // Simuler des données pour le développement
  function mockProjects() {
    return [
      {
        id: "AT-123",
        name: "Nommer le projet",
        color: "blue",
        lead: "Katell Arnault de la Ménardière",
        collaborators: [
          { name: "Katell Arnault de la Ménardière", role: "Chargé·e de projet", timeSpent: 80, timeEstimated: 100, timeRemaining: 20 },
          { name: "Laure Desjardins", role: "Co-chargé·e de projet", timeSpent: 20, timeEstimated: 50, timeRemaining: 30 },
          { name: "Jean-François Jasmin", role: "", timeSpent: 105, timeEstimated: 150, timeRemaining: 45 }
        ],
        timeSpent: 205,
        timeEstimated: 300,
        timeRemaining: 95
      },
      {
        id: "AT-456",
        name: "Le projet de Marie Amélie",
        color: "pink",
        lead: "Katell Arnault de la Ménardière",
        collaborators: [
          { name: "Katell Arnault de la Ménardière", role: "Chargé·e de projet", timeSpent: 20, timeEstimated: 35, timeRemaining: 15 },
          { name: "Marie-Amélie Dubé", role: "Co-chargé·e de projet", timeSpent: 50, timeEstimated: 350, timeRemaining: 300 },
          { name: "Ariane Dionne-Santerre", role: "", timeSpent: 10, timeEstimated: 15, timeRemaining: 5 },
          { name: "Jimmy Paquet-Cormier", role: "", timeSpent: 5, timeEstimated: 50, timeRemaining: 45 }
        ],
        timeSpent: 85,
        timeEstimated: 450,
        timeRemaining: 365
      },
      {
        id: "FO-115",
        name: "Graphisme 101",
        color: "yellow",
        lead: "Katell Arnault de la Ménardière",
        collaborators: [
          { name: "Katell Arnault de la Ménardière", role: "Chargé·e de projet", timeSpent: 40, timeEstimated: 0, timeRemaining: 0 }
        ],
        timeSpent: 40,
        timeEstimated: 0,
        timeRemaining: 0
      },
      {
        id: "RA-224",
        name: "Noisette Steve",
        color: "red",
        lead: "Katell Arnault de la Ménardière",
        collaborators: [
          { name: "Katell Arnault de la Ménardière", role: "Co-chargé·e de projet", timeSpent: 450, timeEstimated: 400, timeRemaining: 50 }
        ],
        timeSpent: 450,
        timeEstimated: 400,
        timeRemaining: 50
      }
    ];
  }

  onMount(() => {
    projects = mockProjects();
    if (projects.length > 0) {
      selectProject(projects[0]);
    }
    loading = false;
  });
</script>

<div class="flex flex-col h-screen bg-gray-100">
  <ProjectsLeftPane />

  <!-- Main Content -->
  <main class="flex-1 container mx-auto p-4 overflow-y-auto">
    <div class="flex flex-col md:flex-row gap-6">
      <!-- Projet List -->
      <div class="w-full md:w-1/3 bg-white rounded-lg shadow-sm">
        <div class="p-4 border-b">
          <h2 class="text-lg font-medium text-gray-800">Tableau de bord</h2>
        </div>
        
        {#if loading}
          <div class="p-4 text-center">Chargement des projets...</div>
        {:else if error}
          <div class="p-4 text-center text-red-500">{error}</div>
        {:else}
          <div class="overflow-y-auto max-h-[calc(100vh-250px)]">
            {#each projects as project}
              <div 
                class="border-l-4 hover:bg-gray-50 cursor-pointer border-b" 
                style="border-left-color: {project.color === 'blue' ? '#93c5fd' : project.color === 'pink' ? '#f9a8d4' : project.color === 'yellow' ? '#fde68a' : project.color === 'red' ? '#fca5a5' : '#d1d5db'}"
                on:click={() => selectProject(project)}
              >
                <div class="p-4">
                  <div class="flex justify-between items-center">
                    <div>
                      <span class="text-sm font-medium text-gray-600">{project.id}</span>
                      <span class="text-xs text-gray-500 ml-2">|</span>
                      <span class="text-xs text-gray-500 ml-2">{project.lead.includes('Chargé') ? 'Chargé·e de projet' : 'Co-chargé·e de projet'}</span>
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
                      <div class="font-medium text-gray-700">{formatHours(project.timeEstimated)}</div>
                    </div>
                    <div>
                      <span>Temps restant</span>
                      <div class="font-medium text-gray-700">{formatHours(project.timeRemaining)}</div>
                    </div>
                  </div>
                </div>
              </div>
            {/each}
          </div>
          
          <div class="p-4 border-t flex items-center">
            <button class="text-sm text-gray-500 flex items-center">
              Projets archivés
              <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
              </svg>
            </button>
          </div>
        {/if}
      </div>
      
      <!-- Projet Details -->
      <div class="w-full md:w-2/3">
        {#if selectedProject}
          <div class="bg-white rounded-lg shadow-sm mb-6">
            <div class="p-4 border-b flex items-center justify-between">
              <div>
                <div class="flex items-center">
                  <div 
                    class="w-6 h-full rounded-l-lg mr-2" 
                    style="background-color: {selectedProject.color === 'blue' ? '#93c5fd' : selectedProject.color === 'pink' ? '#f9a8d4' : selectedProject.color === 'yellow' ? '#fde68a' : selectedProject.color === 'red' ? '#fca5a5' : '#d1d5db'}"
                  ></div>
                  <h2 class="text-lg font-medium text-gray-800">{selectedProject.id}</h2>
                </div>
                <p class="text-sm text-gray-500 mt-1">{selectedProject.name}</p>
              </div>
            </div>
            
            <div class="p-4">
              <div class="mb-4">
                <h3 class="text-sm font-medium text-gray-700 mb-2">Chargé·e de projet</h3>
                <div class="text-sm">{selectedProject.lead}</div>
              </div>
              
              <div class="mb-4">
                <h3 class="text-sm font-medium text-gray-700 mb-2">Co-chargé·e de projet</h3>
                {#each selectedProject.collaborators.filter(c => c.role === 'Co-chargé·e de projet') as collaborator}
                  <div class="flex justify-between items-center mb-2">
                    <div class="flex items-center">
                      <div class="text-sm">{collaborator.name}</div>
                      <button class="ml-2 text-gray-400">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                        </svg>
                      </button>
                    </div>
                    <div class="flex items-center space-x-4">
                      <div class="text-right">
                        <div class="text-xs text-gray-500">Temps passé</div>
                        <div class="font-medium text-sm">{formatHours(collaborator.timeSpent)}</div>
                      </div>
                      <div class="text-right">
                        <div class="text-xs text-gray-500">Temps estimé</div>
                        <div class="font-medium text-sm">{formatHours(collaborator.timeEstimated)}</div>
                      </div>
                      <div class="text-right">
                        <div class="text-xs text-gray-500">Temps restant</div>
                        <div class="font-medium text-sm">{formatHours(collaborator.timeRemaining)}</div>
                      </div>
                    </div>
                  </div>
                {/each}
                <button class="text-sm text-gray-500 mt-2 flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                  Ajouter
                </button>
              </div>
              
              <div>
                <h3 class="text-sm font-medium text-gray-700 mb-2">Collaborateurs</h3>
                {#each selectedProject.collaborators.filter(c => c.role !== 'Chargé·e de projet' && c.role !== 'Co-chargé·e de projet') as collaborator}
                  <div class="flex justify-between items-center mb-2">
                    <div class="flex items-center">
                      <div class="text-sm">{collaborator.name}</div>
                      <button class="ml-2 text-gray-400">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                        </svg>
                      </button>
                    </div>
                    <div class="flex items-center space-x-4">
                      <div class="text-right">
                        <div class="text-xs text-gray-500">Temps passé</div>
                        <div class="font-medium text-sm">{formatHours(collaborator.timeSpent)}</div>
                      </div>
                      <div class="text-right">
                        <div class="text-xs text-gray-500">Temps estimé</div>
                        <div class="font-medium text-sm">{formatHours(collaborator.timeEstimated)}</div>
                      </div>
                      <div class="text-right">
                        <div class="text-xs text-gray-500">Temps restant</div>
                        <div class="font-medium text-sm">{formatHours(collaborator.timeRemaining)}</div>
                      </div>
                    </div>
                  </div>
                {/each}
                <button class="text-sm text-gray-500 mt-2 flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                  Ajouter
                </button>
              </div>
              
              <div class="mt-6 pt-4 border-t">
                <div class="flex justify-between items-center">
                  <h3 class="text-base font-medium">Total</h3>
                  <div class="flex items-center space-x-4">
                    <div class="text-right">
                      <div class="text-xs text-gray-500">Temps passé</div>
                      <div class="font-medium text-sm">{formatHours(selectedProject.timeSpent)}</div>
                    </div>
                    <div class="text-right">
                      <div class="text-xs text-gray-500">Temps estimé</div>
                      <div class="font-medium text-sm">{formatHours(selectedProject.timeEstimated)}</div>
                    </div>
                    <div class="text-right">
                      <div class="text-xs text-gray-500">Temps restant</div>
                      <div class="font-medium text-sm">{formatHours(selectedProject.timeRemaining)}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        {:else}
          <div class="bg-white rounded-lg shadow-sm p-6 text-center">
            <p class="text-gray-500">Sélectionnez un projet pour voir les détails</p>
          </div>
        {/if}
      </div>
    </div>
  </main>
</div>
