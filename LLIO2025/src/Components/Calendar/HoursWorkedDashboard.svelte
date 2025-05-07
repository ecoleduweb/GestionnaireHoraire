<script>
  import { formatDateHoursWorked, areDatesEqual } from '../../utils/date.ts';
  const { hoursTotal, dateStart, dateEnd, textHoursWorked } = $props();
  areDatesEqual(dateStart, dateEnd);
</script>

<style>
  .bilan-container {
    padding: 1rem;
    border: 1px solid #ddd;
    border-radius: 5px;
  }
  .header {
    margin-bottom: 1rem;
  }
  .activites {
    margin-top: 1rem;
  }
</style>

<div class="bilan-container">
  <div class="header">
    <h2>
      {#if areDatesEqual(dateStart, dateEnd)}
        Bilan du {formatDateHoursWorked(dateStart)}
      {:else}
        Bilan du {formatDateHoursWorked(dateStart)} au {formatDateHoursWorked(dateEnd)}
      {/if}
    </h2>
  </div>
  <span>Vous avez travaillé <strong>{hoursTotal}</strong> heures {textHoursWorked}.</span>
</div>
