<script>
  import { ClientTelemetry } from "$lib/tracer"
  import { env } from "$env/dynamic/public"
  import { goto } from "$app/navigation"
  import "../../style/app.css"
  import ProjectsLeftPane from "../../Components/Calendar/ProjectsLeftPane.svelte";
  const apiBaseUrl = import.meta.env.VITE_API_BASE_URL;

  const ENABLED_TELEMETRY = env.PUBLIC_ENABLED_TELEMETRY === "true";

  if (ENABLED_TELEMETRY) {
      const telemetry = ClientTelemetry.getInstance()
      telemetry.start()
  }

</script>

<div class="w-full h-full bg-white px-4 py-6">
  <ProjectsLeftPane />
</div>
