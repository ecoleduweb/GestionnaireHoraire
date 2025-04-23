<script>
  import { onMount } from "svelte";
  import { ClientTelemetry } from "$lib/tracer"
  import { env } from "$env/dynamic/public"
  import "../../style/app.css"
  import ProjectsLeftPane from "../../Components/Projects/ProjectsLeftPane.svelte";
  import ProjectComponent from "../../Components/Projects/ProjectComponent.svelte";
  const ENABLED_TELEMETRY = env.PUBLIC_ENABLED_TELEMETRY === "true";

  if (ENABLED_TELEMETRY) {
    const telemetry = ClientTelemetry.getInstance();
    telemetry.start();
  }

  // État des projets
  let projects = [];
  let loading = true;

  function mockProjects() {
    return [
      {
        id: "AT-123",
        name: "Nommer le projet",
        color: "blue",
        lead: "Katell Arnault de la Ménardière",
        coLeads: ["Jean-François Jasmin"],
        timeSpent: 205,
        timeEstimated: 300,
        timeRemaining: 95
      },
      {
        id: "AT-456",
        name: "Le projet de Marie Amélie",
        color: "pink",
        lead: "Katell Arnault de la Ménardière",
        coLeads: ["Marie Amélie Dubé", "Jimmy Paquet-Cormier"],
        timeSpent: 85,
        timeEstimated: 450,
        timeRemaining: 365
      },
      {
        id: "FO-115",
        name: "Graphisme 101",
        color: "yellow",
        lead: "Katell Arnault de la Ménardière",
        coLeads: [],
        timeSpent: 40,
        timeEstimated: 0,
        timeRemaining: 0
      },
      {
        id: "RA-224",
        name: "Noisette Steve",
        color: "red",
        lead: "Steve Joncoux",
        coLeads: ["Katell Arnault de la Ménardière", "Marjolaine Poirier", "Jimmy Paquet-Cormier"],
        timeSpent: 450,
        timeEstimated: 400,
        timeRemaining: -50
      }
    ];
  }

  onMount(() => {
    projects = mockProjects();
    loading = false;
  });
</script>

<div class="h-screen bg-gray-100">
  <ProjectsLeftPane />
  <div class="flex flex-col" style="padding-left: 300px;">
    <!-- Project Details -->
    <div class="p-4">
      <h1 class="text-2xl font-medium text-gray-800">Vos projets en cours</h1>
    </div>
    {#each projects as project}
      <ProjectComponent {project} />
    {/each}
  </div>
</div>
