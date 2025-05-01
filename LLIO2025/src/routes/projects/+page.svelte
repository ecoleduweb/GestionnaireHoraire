<script lang="ts">
  import { onMount } from "svelte";
  import "../../style/app.css"
  import ProjectsLeftPane from "../../Components/Projects/ProjectsLeftPane.svelte";
  import ProjectComponent from "../../Components/Projects/ProjectComponent.svelte";
  import type { Project } from '../../Models/index.ts';
  import { ProjectApiService } from "../../services/ProjectApiService";

  // État des projets
  let projects = $state<Project[]>([]);
  let isLoading = $state(true);
  let error = $state<string | null>(null);

  async function loadProjects() {
    try {
      isLoading = true;
      error = null;
      const response = await ProjectApiService.getProjects();
      projects = response;
    } catch (err) {
      console.error('Erreur lors de la récupération des projets:', err);
      error = "Impossible de charger les projets. Veuillez réessayer plus tard.";
      projects = [];
    } finally {
      isLoading = false;
    }
  }

  function mockProjects() {
    return [
      {
      id: 1,
      name: "AT-123",
      description: "Nommer le projet",
      color: "blue",
      lead: "Katell Arnault de la Ménardière",
      isArchived: false,
      coLeads: ["Jean-François Jasmin"],
      employees: [
        {
        name: "Katell Arnault de la Ménardière",
        categories: [
          { name: "Développement", timeSpent: 30, timeEstimated: 25 },
          { name: "Graphisme", timeSpent: 15, timeEstimated: 30 }
        ]
        },
        {
        name: "Jean-François Jasmin",
        categories: [
          { name: "Développement", timeSpent: 20, timeEstimated: 50 }
        ]
        },
        {
        name: "Laure Desjardins",
        categories: []
        }
      ],
      totalTimeSpent: 65
      },
      {
      id: 2,
      name: "AT-456",
      description: "Le projet de Marie Amélie",
      color: "pink",
      lead: "Katell Arnault de la Ménardière",
      isArchived: false,
      coLeads: ["Marie Amélie Dubé", "Jimmy Paquet-Cormier"],
      employees: [
        {
        name: "Katell Arnault de la Ménardière",
        categories: [
          { name: "Développement", timeSpent: 12, timeEstimated: 15 }
        ]
        },
        {
        name: "Marie Amélie Dubé",
        categories: [
          { name: "Graphisme", timeSpent: 18, timeEstimated: 20 }
        ]
        },
        {
        name: "Ariane Dionne-Santerre",
        categories: [
          { name: "Rédaction", timeSpent: 8, timeEstimated: 10 }
        ]
        },
        {
        name: "Jimmy Paquet-Cormier",
        categories: [
          { name: "Développement", timeSpent: 14, timeEstimated: 15 }
        ]
        }
      ],
      totalTimeSpent: 52
      },
      {
      id: 3,
      name: "FO-115",
      description: "Graphisme 101",
      color: "yellow",
      lead: "Katell Arnault de la Ménardière",
      isArchived: false,
      coLeads: [],
      employees: [
        {
        name: "Katell Arnault de la Ménardière",
        categories: [
          { name: "Gestion", timeSpent: 5, timeEstimated: 8 }
        ]
        },
        {
        name: "Annie Côté",
        categories: [
          { name: "Graphisme", timeSpent: 22, timeEstimated: 25 }
        ]
        }
      ],
      totalTimeSpent: 27
      },
      {
      id: 4,
      name: "RA-224",
      description: "Noisette Steve",
      color: "red",
      lead: "Steve Joncoux",
      isArchived: false,
      coLeads: ["Katell Arnault de la Ménardière", "Marjolaine Poirier", "Jimmy Paquet-Cormier"],
      employees: [
        {
        name: "Steve Joncoux",
        categories: [
          { name: "Développement", timeSpent: 16, timeEstimated: 20 }
        ]
        },
        {
        name: "Katell Arnault de la Ménardière",
        categories: [
          { name: "Gestion", timeSpent: 7, timeEstimated: 10 }
        ]
        },
        {
        name: "Marjolaine Poirier",
        categories: [
          { name: "Rédaction", timeSpent: 12, timeEstimated: 15 }
        ]
        },
        {
        name: "Jimmy Paquet-Cormier",
        categories: [
          { name: "Développement", timeSpent: 9, timeEstimated: 12 }
        ]
        },
        {
        name: "Marie Amélie Dubé",
        categories: [
          { name: "Graphisme", timeSpent: 11, timeEstimated: 10 }
        ]
        }
      ],
      totalTimeSpent: 55
      },
      {
      id: 5,
      name: "AT-789",
      description: "Nommer le projet",
      color: "blue",
      lead: "Katell Arnault de la Ménardière",
      isArchived: true,
      coLeads: ["Jean-François Jasmin"],
      employees: [
        {
        name: "Katell Arnault de la Ménardière",
        categories: [
          { name: "Développement", timeSpent: 30, timeEstimated: 25 },
          { name: "Graphisme", timeSpent: 15, timeEstimated: 30 }
        ]
        },
        {
        name: "Jean-François Jasmin",
        categories: [
          { name: "Développement", timeSpent: 20, timeEstimated: 50 }
        ]
        },
        {
        name: "Laure Desjardins",
        categories: []
        }
      ],
      totalTimeSpent: 65
      },
      {
      id: 6,
      name: "AT-987",
      description: "Le projet de Marie Amélie",
      color: "pink",
      lead: "Katell Arnault de la Ménardière",
      isArchived: true,
      coLeads: ["Marie Amélie Dubé", "Jimmy Paquet-Cormier"],
      employees: [
        {
        name: "Katell Arnault de la Ménardière",
        categories: [
          { name: "Développement", timeSpent: 12, timeEstimated: 15 }
        ]
        },
        {
        name: "Marie Amélie Dubé",
        categories: [
          { name: "Graphisme", timeSpent: 18, timeEstimated: 20 }
        ]
        },
        {
        name: "Ariane Dionne-Santerre",
        categories: [
          { name: "Rédaction", timeSpent: 8, timeEstimated: 10 }
        ]
        },
        {
        name: "Jimmy Paquet-Cormier",
        categories: [
          { name: "Développement", timeSpent: 14, timeEstimated: 15 }
        ]
        }
      ],
      totalTimeSpent: 52
      },
      {
      id: 7,
      name: "FO-789",
      description: "Graphisme 101",
      color: "yellow",
      lead: "Katell Arnault de la Ménardière",
      isArchived: true,
      coLeads: [],
      employees: [
        {
        name: "Katell Arnault de la Ménardière",
        categories: [
          { name: "Gestion", timeSpent: 5, timeEstimated: 8 }
        ]
        },
        {
        name: "Annie Côté",
        categories: [
          { name: "Graphisme", timeSpent: 22, timeEstimated: 25 }
        ]
        }
      ],
      totalTimeSpent: 27
      },
      {
      id: 8,
      name: "RA-987",
      description: "Noisette Steve",
      color: "red",
      lead: "Steve Joncoux",
      isArchived: true,
      coLeads: ["Katell Arnault de la Ménardière", "Marjolaine Poirier", "Jimmy Paquet-Cormier"],
      employees: [
        {
        name: "Steve Joncoux",
        categories: [
          { name: "Développement", timeSpent: 16, timeEstimated: 20 }
        ]
        },
        {
        name: "Katell Arnault de la Ménardière",
        categories: [
          { name: "Gestion", timeSpent: 7, timeEstimated: 10 }
        ]
        },
        {
        name: "Marjolaine Poirier",
        categories: [
          { name: "Rédaction", timeSpent: 12, timeEstimated: 15 }
        ]
        },
        {
        name: "Jimmy Paquet-Cormier",
        categories: [
          { name: "Développement", timeSpent: 9, timeEstimated: 12 }
        ]
        },
        {
        name: "Marie Amélie Dubé",
        categories: [
          { name: "Graphisme", timeSpent: 11, timeEstimated: 10 }
        ]
        }
      ],
      totalTimeSpent: 55
      },
    ];
  }

  onMount(() => {
    loadProjects();
  });
</script>

<div class="bg-gray-100">
  <ProjectsLeftPane {projects} />
  <div class="flex flex-col" style="padding-left: 300px;">
    <!-- Project Details -->
    <div class="p-4">
      <h1 class="text-2xl font-medium text-gray-800">Vos projets en cours</h1>
    </div>
    {#if isLoading}
    <div class="flex justify-center items-center h-screen">
      <p class="text-gray-500">Chargement des projets...</p>
    </div>
  {:else if error}
    <div class="flex justify-center items-center h-screen">
      <p class="text-red-500">{error}</p>
    </div>
  {:else}
    {#each projects as project}
      <ProjectComponent {project} />
    {/each}
  {/if}
  </div>
</div>
