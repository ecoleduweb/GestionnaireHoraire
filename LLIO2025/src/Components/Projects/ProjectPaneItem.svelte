<script lang="ts">
  import { formatHours } from '../../utils/date';
  import { Pencil } from 'lucide-svelte';
  import type { Project } from '../../Models';
  
  type Props = {
    project: Project;
    onEdit: (project: Project) => void;
  };
  
  let { project, onEdit } = $props();
  
  function handleEdit(event: MouseEvent) {
    event.stopPropagation(); // Empêche la propagation du clic aux éléments parents
    onEdit(project);
  }
</script>

<div class="border-l-10 border-b" style="border-left-color: {project.color}">
  <div class="p-4">
    <div class="flex justify-between items-center">
      <div>
        <span class="text-black">{project.id}</span>
        <span class="text-gray-500 ml-2">|</span>
        <span class="text-black ml-2">{project.lead}</span>
        <span class="text-gray-500 ml-2">|</span>
        <span class="text-gray-500 ml-2">{project.name}</span>
      </div>
      <button 
        class="p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-full transition-colors"
        on:click={handleEdit}
        aria-label="Modifier le projet"
      >
        <Pencil size={16} />
      </button>
    </div>

    <div class="mt-2 flex items-center text-xs text-gray-500">
      <div class="mr-3">
        <span>Temps passé</span>
        <hr class="my-1" />
        <div class="font-bold text-black mr-4">{formatHours(project.totalTimeSpent)}</div>
      </div>
      <div class="mr-3">
        <span>Temps estimé</span>
        <hr class="my-1" />
        <div class="text-gray-400">
          {formatHours(
            project.employees.reduce(
              (sum, emp) =>
                sum + emp.categories.reduce((catSum, cat) => catSum + cat.timeEstimated, 0),
              0
            )
          )}
        </div>
      </div>
      <div>
        <span>Temps restant</span>
        <hr class="my-1" />
        {#if project.employees.reduce((sum, emp) => sum + emp.categories.reduce((catSum, cat) => catSum + cat.timeEstimated, 0), 0) - project.employees.reduce((sum, emp) => sum + emp.categories.reduce((catSum, cat) => catSum + cat.timeSpent, 0), 0) < 0}
          <div class="font-medium text-red-700">
            {formatHours(
              project.employees.reduce(
                (sum, emp) =>
                  sum + emp.categories.reduce((catSum, cat) => catSum + cat.timeEstimated, 0),
                0
              ) -
                project.employees.reduce(
                  (sum, emp) =>
                    sum + emp.categories.reduce((catSum, cat) => catSum + cat.timeSpent, 0),
                  0
                )
            )}
          </div>
        {:else}
          <div class="text-gray-400">
            {formatHours(
              project.employees.reduce(
                (sum, emp) =>
                  sum + emp.categories.reduce((catSum, cat) => catSum + cat.timeEstimated, 0),
                0
              ) -
                project.employees.reduce(
                  (sum, emp) =>
                    sum + emp.categories.reduce((catSum, cat) => catSum + cat.timeSpent, 0),
                  0
                )
            )}
          </div>
        {/if}
      </div>
    </div>
  </div>
</div>
