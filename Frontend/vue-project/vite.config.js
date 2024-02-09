import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { splitVendorChunkPlugin } from 'vite'


export default defineConfig({
  plugins: [vue(),splitVendorChunkPlugin()],
  resolve: {

  },

})
