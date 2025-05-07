<script>
    import {formatDateHoursWorked, areDatesEqual, getHoursFromRange} from '../../utils/date.js';
    const { activities = [], dateStart, dateEnd, textHoursWorked } = $props();
    const totalHours = $derived(calculateTotalHours());

    areDatesEqual(dateStart, dateEnd);
    
    function calculateTotalHours() {
      let total = 0;
      if (activities && activities.length > 0) {
        activities.forEach(activity => {
          total += getHoursFromRange(activity);
        });
      }
      return total;
    }
  </script>
  
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
    <span>Vous avez travaillé <strong>{totalHours.toFixed(2)}</strong> heures {textHoursWorked}.</span>
  </div>
  
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