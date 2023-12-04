import { fileURLToPath, URL } from 'node:url'

import {  } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default ({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },

  build: {
    rollupOptions: {
        output:{
            manualChunks(id) {
                if (id.includes('node_modules')) {
                    return id.toString().split('node_modules/')[1].split('/')[0].toString();
                }
            }
        }
    }
}
})


