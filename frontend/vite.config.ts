import {defineConfig} from 'vite'
import {sveltekit} from '@sveltejs/kit/vite'
import {vitePreprocess} from '@sveltejs/vite-plugin-svelte'
import adapter from '@sveltejs/adapter-static'
import tailwindcss from "@tailwindcss/vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    tailwindcss(),
    sveltekit({
      preprocess: vitePreprocess(),
      // SPA mode: no server inside Wails, so emit a single fallback page.
      // Output goes to dist/ because main.go embeds frontend/dist.
      adapter: adapter({
        pages: 'dist',
        assets: 'dist',
        fallback: 'index.html',
      }),
    }),
  ],
})