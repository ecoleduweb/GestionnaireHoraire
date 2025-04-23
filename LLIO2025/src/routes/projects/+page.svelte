<script>
  import { onMount } from "svelte";
  import { ClientTelemetry } from "$lib/tracer"
  import { env } from "$env/dynamic/public"
  import { goto } from "$app/navigation"
  import "../../style/app.css"
  import ProjectsLeftPane from "../../Components/Calendar/ProjectsLeftPane.svelte";
  import ProjectComponent from "../../Components/Projects/ProjectComponent.svelte";
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

<div class="h-screen bg-gray-100">
  <ProjectsLeftPane />

  <!-- Main Content -->
  <div class="flex flex-col gap-6" style="padding-left: 300px;">
      
    <!-- Project Details -->
    <div class="p-4 border-b">
      <h2 class="text-lg font-medium text-gray-800">Vos projets en cours</h2>
    </div>
    {#each projects as project}
      <ProjectComponent {project} />
    {/each}
  </div>
</div>
