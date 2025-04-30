<script lang="ts">
  import { onMount } from "svelte";
  import "../../style/app.css"
  import ProjectsLeftPane from "../../Components/Projects/ProjectsLeftPane.svelte";
  import ProjectComponent from "../../Components/Projects/ProjectComponent.svelte";
  import type { Project } from '../../Models/index.ts';

  // État des projets
  let projects = $state<Project[]>([]);
  let isLoading = $state(false);

  function mockProjects() {
    return [
      {
      id: "AT-123",
      name: "Nommer le projet",
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
      id: "AT-456",
      name: "Le projet de Marie Amélie",
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
      id: "FO-115",
      name: "Graphisme 101",
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
      id: "RA-224",
      name: "Noisette Steve",
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
      id: "AT-789",
      name: "Nommer le projet",
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
      id: "AT-987",
      name: "Le projet de Marie Amélie",
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
      id: "FO-789",
      name: "Graphisme 101",
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
      id: "RA-987",
      name: "Noisette Steve",
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
    projects = mockProjects();
    isLoading = false;
  });
</script>

<div class="bg-gray-100">
  <ProjectsLeftPane {projects} />
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
