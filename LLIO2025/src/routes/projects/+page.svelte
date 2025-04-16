<script>
  import { ClientTelemetry } from "$lib/tracer"
  import { env } from "$env/dynamic/public"
  import { goto } from "$app/navigation";

  const apiBaseUrl = import.meta.env.VITE_API_BASE_URL;

  const ENABLED_TELEMETRY = env.PUBLIC_ENABLED_TELEMETRY === "true";

  if (ENABLED_TELEMETRY) {
      const telemetry = ClientTelemetry.getInstance()
      telemetry.start()
  }

</script>

<div class="w-full h-full bg-white px-4 py-6">
  <p>Salut</p>
  <div class="inline-flex rounded-md shadow-xs" role="group">
    <button 
      onclick={() => (goto('/calendar'))}
      type="button" 
      class="py-2 px-4 text-sm transition-colors font-semibold
              {index === 0 ? 'rounded-l-lg rounded-r-none' : index === timeRanges.length - 1 ? 'rounded-r-lg rounded-l-none' : 'rounded-none border-x border-white/20'}
              {activeTimeRange.label === range.label
                ? 'bg-[#015e61] text-white'
                : 'bg-gray-100 text-gray-700 hover:bg-gray-200'}
              "
      >
      Calendrier
    </button>
    <button 
      onclick={() => (goto('/projects'))}
      type="button" 
      class="px-4 py-2 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-e-lg hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:text-white dark:hover:bg-gray-700 dark:focus:ring-blue-500 dark:focus:text-white"
      >
      Projets
    </button>
  </div>
</div>
