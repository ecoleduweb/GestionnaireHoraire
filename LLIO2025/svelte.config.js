import adapter from '@sveltejs/adapter-node';
import { loadEnv } from 'vite';

// Charger les variables d'environnement
const env = loadEnv(process.env.NODE_ENV || 'development', process.cwd(), '');

// Utiliser la valeur de BASE_PATH depuis .env ou une valeur par défaut
const basePath = env.BASE_PATH || '';
console.log("🚀 ~ basePath:", basePath)

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter(),
    paths: {
      base: basePath
    }
  }
};

export default config;