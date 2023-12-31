import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  base: "/mobile",
  build: {
    chunkSizeWarningLimit: 2048,
    outDir: 'mobile', //指定输出路径
    minify: 'terser' // 混淆器，terser构建后文件体积更小
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
