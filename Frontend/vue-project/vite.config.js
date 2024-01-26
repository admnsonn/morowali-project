import { fileURLToPath, URL } from 'node:url'

import { defineConfig  } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig ({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src/assets/*', import.meta.url))
    }
  },

  build: {
    rollupOptions: {
      external: ['src/assets/img/*.svg'],
    },
    outDir: 'dist',
    assetsDir: '',
    sourcemap: false,
    minify: true,
  }
})


