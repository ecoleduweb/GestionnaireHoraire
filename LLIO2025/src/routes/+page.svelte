<script>
  import '../style/app.css';
  import { writable } from 'svelte/store';

  const apiBaseUrl = import.meta.env.VITE_API_BASE_URL;
  const isLoading = writable(false);

  function handleAzureadLogin() {
    $isLoading = true;
    window.location.href = `${apiBaseUrl}/auth/azureadv2`;
  }
</script>

<link
  rel="stylesheet"
  href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css"
  integrity="sha512-1ycn6IcaQQ40/MKBW2W4Rhis/DbILU74C1vSrLJxCq57o941Ym01SwNsOMqvEBFlcgUa6xLiPY/NS5R+E6ztJQ=="
  crossorigin="anonymous"
  referrerpolicy="no-referrer"
/>

<div
  class="w-full h-screen flex items-center justify-center bg-[url('/grungy-painted-wall-textured-background.jpg')] bg-cover bg-center relative"
>
  <div
    class="w-[90%] max-w-sm md:max-w-md lg:max-w-md p-5 bg-gray-50 flex-col flex items-center gap-3 rounded-xl shadow-slate-500 shadow-lg"
  >
    <img
      fetchpriority="high"
      decoding="async"
      width="1024"
      height="383"
      src="llio.gif"
      class="attachment-large size-large wp-image-930"
      alt="Animation du logo du LLio pour expliquer que LLio signifie Le Laboratoire en innovation ouverte"
    />
    <p class="text-center text-lg">Bienvenue, connectez-vous à votre compte.</p>
    <button
      class="w-full bg-[#005f61] hover:bg-[#00968f] text-white font-bold py-2 px-4 rounded"
      on:click={handleAzureadLogin}
    >
      <img src="microsoft.svg" alt="Logo de Microsoft" class="w-6 h-6 inline" />
      Se connecter avec Microsoft
    </button>
  </div>
  <img
    src="cegep+llio.png"
    alt="Logo du Cégep LLio"
    class="absolute bottom-4 right-4 w-50 h-auto"
  />

  {#if $isLoading}
    <div class="w-full h-full fixed top-0 left-0 bg-white opacity-75 z-50">
      <div class="flex justify-center items-center mt-[50vh]">
        <div class="fas fa-circle-notch fa-spin fa-5x text-[#005f61]"></div>
      </div>
    </div>
  {/if}
</div>
