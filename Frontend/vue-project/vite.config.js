import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': '/absolute/path/to/src'
    }
  },
  build: {
    rollupOptions: {
      external: ['path/to/external/file.svg'],
    },
    outDir: 'dist',
    assetsDir: '',
    sourcemap: false,
    minify: true,
  }
})
