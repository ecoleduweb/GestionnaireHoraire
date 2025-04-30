<script>
  import { ClientTelemetry } from "$lib/tracer";
  import { env } from "$env/dynamic/public"

  let { children } = $props();

  const ENABLED_TELEMETRY = env.PUBLIC_ENABLED_TELEMETRY === "true";

  if (ENABLED_TELEMETRY) {
      const telemetry = ClientTelemetry.getInstance()
      telemetry.start()
  }
</script>

{@render children()}