<script lang="ts">
  import { formatHours } from '../../utils/date';

  let { project } = $props();
</script>

<div class="border-l-10 border-b" style="border-left-color: {project.color}">
  <div class="p-4">
    <div class="flex justify-between items-center">
      <div>
        <span class="text-black">{project.name}</span>
        <span class="text-gray-500 ml-2">|</span>
        <span class="text-gray-500 ml-2">{project.description}</span>
      </div>
    </div>

    <div class="mt-2 flex items-center text-xs">
      <div class="mr-16">
        <div class="font-bold">Total</div>
      </div>
      <div class="flex">
        <div class="w-14">
          <div class="font-bold">{formatHours(project.totalTimeSpent)}</div>
        </div>
        <div class="w-14">
          <div class="text-gray-400">
            {formatHours(project.totalTimeEstimated)}
          </div>
        </div>
        <div>
          {#if project.totalTimeRemaining < 0}
            <div class="font-medium text-red-700">
              {formatHours(project.totalTimeRemaining)}
            </div>
          {:else}
            <div class="text-gray-400">
              {formatHours(project.totalTimeRemaining)}
            </div>
          {/if}
        </div>
      </div>
    </div>
  </div>
</div>
