<script lang="ts"> 
  import { CalendarService } from './services/calendar.service';
  import { onMount } from 'svelte';

  let calendarEl = $state<HTMLElement | null>(null);
  let calendarService = $state<CalendarService | null>(null);

  onMount(() => {
      if (calendarEl) {  // Vérification de null
          calendarService = new CalendarService();
          calendarService.initialize(calendarEl);
          calendarService.render();
      }
  });
</script>

<style>
  :global(.calendar-container) {
      margin: 20px;
      height: 800px;
  }
</style>

<div class="calendar-container" bind:this={calendarEl}></div>