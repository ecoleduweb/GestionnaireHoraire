<script lang="ts">
  import { onMount } from "svelte";
  import "../../style/app.css"
  import ProjectsLeftPane from "../../Components/Projects/ProjectsLeftPane.svelte";
  import ProjectComponent from "../../Components/Projects/ProjectComponent.svelte";
  import type { Project, UserInfo } from '../../Models/index.ts';
  import { ProjectApiService } from "../../services/ProjectApiService";
  import { UserApiService } from "../../services/UserApiService";

  // État des projets
  let projects = $state<Project[]>([]);
  let isLoading = $state(true);
  let error = $state<string | null>(null);

  let currentUser = $state<UserInfo | null>(null);

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

  onMount(async () => {
    try {
        currentUser = await UserApiService.getUserInfo();
        console.log("Current user")
        console.log(currentUser)
      } catch (error) {
        console.error('Erreur lors du chargement des informations utilisateur:', error);
      }
    loadProjects();
  });
</script>

<div class="bg-gray-100">
  {#if currentUser}
  <ProjectsLeftPane {projects} {currentUser} />
  {/if}
  
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
