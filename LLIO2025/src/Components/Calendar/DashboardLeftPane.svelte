<script lang="ts">
</script>

<style>
  .dashboard-panel {
    width: 300px;
    height: 100vh;
    background-color: white;
    border-right: 1px solid #e5e7eb;
    box-shadow: 2px 0 5px rgba(0, 0, 0, 0.05);
    overflow-y: auto;
    position: fixed;
    left: 0;
    top: 0;
    bottom: 0;
    z-index: 30;
  }

  .dashboard-header {
    background-color: #005e61;
    color: white;
    padding: 16px;
    font-weight: 600;
    font-size: 1.25rem;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  }

  .dashboard-content {
    padding: 0px;
  }

  .dashboard-item {
    padding: 16px;
    border-bottom: 1px solid #f0f0f0;
    transition: background-color 0.2s;
  }

  .dashboard-item:hover {
    background-color: #f5f5f5;
  }
</style>

<div class="dashboard-panel">
  <!-- En-tête du dashboard -->
  <div class="dashboard-header">Dashboard</div>

  <!-- Contenu du dashboard -->
  <div class="dashboard-content">
    <!-- Éléments du dashboard -->
    <div class="dashboard-item">
      <div class="text-gray-600">Contenu du dashboard à venir...</div>
    </div>

    <div class="dashboard-item">
      <div class="font-medium text-gray-700 mb-2">Statistiques</div>
      <div class="text-gray-500 text-sm">Aucune statistique disponible pour le moment.</div>
    </div>

    <div class="dashboard-item">
      <div class="font-medium text-gray-700 mb-2">Activités récentes</div>
      <div class="text-gray-500 text-sm">Aucune activité récente.</div>
    </div>

    <div class="dashboard-item">
      <div class="font-medium text-gray-700 mb-2">Projets</div>
      <div class="text-gray-500 text-sm">Aucun projet disponible.</div>
    </div>
  </div>
</div>
