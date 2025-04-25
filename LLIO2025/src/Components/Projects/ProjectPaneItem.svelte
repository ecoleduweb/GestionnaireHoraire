<script lang="ts">
  import { formatHours } from "../../utils/date";

  let { project } = $props();
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
        <div class="text-gray-400">{formatHours(project.timeEstimated)}</div>
      </div>
      <div>
        <span>Temps restant</span>
        <hr class="my-1" />
        {#if project.timeRemaining < 0}
          <div class="font-medium text-red-700">{formatHours(project.timeRemaining)}</div>
        {:else}
          <div class="text-gray-400">{formatHours(project.timeRemaining)}</div>
        {/if}
      </div>
    </div>
  </div>
</div>
