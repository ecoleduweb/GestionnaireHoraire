<script lang="ts">
  export let project;

  const formatHours = (hours) => {
    return hours ? `${hours}h00` : '-';
  };
</script>

<div class="border-l-10 border-b" style="border-left-color: {project.color}">
  <div class="p-4">
    <div class="flex justify-between items-center">
      <div>
        <span class="text-black">{project.id}</span>
        <span class="text-gray-500 ml-2">|</span>
        <span class="text-gray-500 ml-2">{project.name}</span>
      </div>
    </div>

    <div class="mt-2 flex items-center text-xs text-gray-500">
      <div class="mr-16">
        <div class="font-bold text-black">Total</div>
      </div>
      <div class="mr-4">
        <div class="font-bold text-black">{formatHours(project.timeSpent)}</div>
      </div>
      <div class="mr-4">
        <div class="text-gray-400">{formatHours(project.timeEstimated)}</div>
      </div>
      <div>
        {#if project.timeRemaining < 0}
          <div class="font-medium text-red-700">{formatHours(project.timeRemaining)}</div>
        {:else}
          <div class="text-gray-400">{formatHours(project.timeRemaining)}</div>
        {/if}
      </div>
    </div>
  </div>
</div>
