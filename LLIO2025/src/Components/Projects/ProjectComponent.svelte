<script>
  // Vous pourriez importer vos données ici ou les recevoir comme props
  export let projectNumber = 'AT-123';
  export let projectManager = 'Kateri Arnault de la Ménardière';
  export let coProjectManager = 'Jean-François Jasmin';

  export let timeEntries = [
    {
      name: 'Kateri Arnault de la Ménardière',
      timeSpent: '80h00',
      timeRemaining: '150h00',
      timeAllocated: '20h00',
    },
    {
      name: 'Laurie Desjardins',
      timeSpent: '20h00',
      timeRemaining: '50h00',
      timeAllocated: '30h00',
    },
    {
      name: 'Jean-François Jasmin',
      timeSpent: '105h00',
      timeRemaining: '150h00',
      timeAllocated: '45h00',
    },
  ];

  // Calculer les totaux
  $: totalTimeSpent = '205h00'; // Ces valeurs pourraient être calculées dynamiquement
  $: totalTimeRemaining = '350h00';
  $: totalTimeAllocated = '95h00';
</script>

<div class="max-w-4xl mx-auto my-8">
  <div class="bg-white rounded-lg shadow-md overflow-hidden border border-gray-200">
    <div class="flex">
      <!-- Sidebar -->
      <div class="w-32 bg-indigo-100 p-4 border-r border-indigo-200">
        <div class="mb-6">
          <div class="font-bold text-gray-700">{projectNumber}</div>
          <div class="text-xs text-gray-600">Numéro de projet</div>
        </div>

        <div class="mb-6">
          <div class="text-xs text-gray-500">Chargé(e) du projet</div>
          <div class="text-sm text-gray-700">{projectManager}</div>
        </div>

        <div class="mb-6">
          <div class="text-xs text-gray-500">Co-chargé(e) du projet</div>
          <div class="text-sm text-gray-700">{coProjectManager}</div>
        </div>
      </div>

      <!-- Main content -->
      <div class="flex-1">
        <table class="w-full">
          <thead>
            <tr class="text-left text-gray-600">
              <th class="px-4 py-2 w-2/5"></th>
              <th class="px-4 py-2 text-right">Temps passé</th>
              <th class="px-4 py-2 text-right">Temps restant</th>
              <th class="px-4 py-2 text-right">Temps imputé</th>
            </tr>
          </thead>
          <tbody>
            {#each timeEntries as entry}
              <tr class="border-t border-gray-200">
                <td class="px-4 py-3">
                  <div class="flex items-center">
                    <span class="text-gray-800">{entry.name}</span>
                    <svg class="w-4 h-4 ml-1 text-gray-500" viewBox="0 0 20 20" fill="currentColor">
                      <path
                        fill-rule="evenodd"
                        d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                        clip-rule="evenodd"
                      ></path>
                    </svg>
                  </div>
                </td>
                <td class="px-4 py-3 text-right text-gray-700">{entry.timeSpent}</td>
                <td class="px-4 py-3 text-right text-gray-700">{entry.timeRemaining}</td>
                <td class="px-4 py-3 text-right text-gray-700">{entry.timeAllocated}</td>
              </tr>
            {/each}
            <tr class="border-t border-gray-200 font-semibold">
              <td class="px-4 py-3 text-gray-800">Total</td>
              <td class="px-4 py-3 text-right text-gray-800">{totalTimeSpent}</td>
              <td class="px-4 py-3 text-right text-gray-800">{totalTimeRemaining}</td>
              <td class="px-4 py-3 text-right text-gray-800">{totalTimeAllocated}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>
