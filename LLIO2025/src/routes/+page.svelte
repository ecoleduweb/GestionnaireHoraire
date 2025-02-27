<script>
  import '../style/app.css';
</script>

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
      src="https://llio.quebec/wp-content/uploads/2022/01/signification-llio.gif"
      class="attachment-large size-large wp-image-930"
      alt="Animation du logo du LLio pour expliquer que LLio signifie Le Laboratoire en innovation ouverte"
    />
    <p class="text-center text-lg">Bienvenue, connectez-vous à votre compte.</p>
    <button class="w-full bg-[#005f61] hover:bg-[#00968f] text-white font-bold py-2 px-4 rounded">
      <img src="microsoft.svg" alt="Logo de Microsoft" class="w-6 h-6 inline" />
      Se connecter avec Microsoft
    </button>
  </div>
  <img
    src="cegep+llio.png"
    alt="Logo du Cégep LLio"
    class="absolute bottom-4 right-4 w-50 h-auto"
  />
</div>
