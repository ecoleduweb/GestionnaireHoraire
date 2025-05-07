<!-- HoursWorkedDashboard.svelte -->
<script>
    // Importation des props avec la syntaxe Svelte 5
    const { activities = [], dateStart, dateEnd, textHoursWorked } = $props();
    
    const formatDate = (date) => {
    let dateObj;
    
    if (date instanceof Date) {
        dateObj = date;
    } else if (typeof date === 'string') {
        // Pour une chaîne comme "2025-05-05", extraire année, mois, jour
        const parts = date.split('-');
        dateObj = new Date(parseInt(parts[0]), parseInt(parts[1]) - 1, parseInt(parts[2]), 12, 0, 0);
    }
    
    return new Intl.DateTimeFormat('fr-FR', {
        day: 'numeric',
        month: 'long'
    }).format(dateObj);
    };
    
    // Vérification si les dates sont identiques
    const areDatesEqual = () => {
        if (!dateStart || !dateEnd) return false;
        return dateStart == dateEnd;
    };
    
    const totalHours = $derived(calculateTotalHours());
    
    // Fonction pour calculer le total
    function calculateTotalHours() {
      let total = 0;
      if (activities && activities.length > 0) {
        activities.forEach(activity => {
          const start = activity.startDate instanceof Date ? activity.startDate : new Date(activity.startDate);
          const end = activity.endDate instanceof Date ? activity.endDate : new Date(activity.endDate);
          
          const diffMilliseconds = end.getTime() - start.getTime();
          const hours = diffMilliseconds / (1000 * 60 * 60);
          
          total += hours;
        });
      }
      return total;
    }
  </script>
  
  <div class="bilan-container">
    <div class="header">
      <h2>
        {#if areDatesEqual()}
          Bilan du {formatDate(dateStart)}
        {:else}
          Bilan du {formatDate(dateStart)} au {formatDate(dateEnd)}
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