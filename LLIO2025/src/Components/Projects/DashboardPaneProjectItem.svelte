<script lang="ts">
  import { formatHours } from '../../utils/date';
  import type { Project } from '../../Models/index.ts';

  let { project } = $props();

  function getProjectHoursColorBasedOnEstimatedHours(project : Project){
    if (project.totalTimeEstimated <= 0) {
      return 'text-gray-400';
    } else if (project.totalTimeSpent > project.totalTimeEstimated) {
      return 'text-red-700';
    } else {
      return 'text-gray-700';
    }
  }
</script>

<div class="border-l-10 border-b" style="border-left-color: {project.color}">
  <div class="p-4">
    <div class="flex justify-between items-center">
      <div>
        <span class="text-black">{project.uniqueId}</span>
        <span class="text-gray-500 ml-2">|</span>
        <span class="text-gray-500 ml-2">{project.name}</span>
      </div>
    </div>

    <div class="mt-2 flex items-center text-xs">
      <div class="mr-10">
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
              <div class="{getProjectHoursColorBasedOnEstimatedHours(project)}">
                {formatHours(project.totalTimeRemaining)}
              </div>
        </div>
      </div>
    </div>
  </div>
</div>
