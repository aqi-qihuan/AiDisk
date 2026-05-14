import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
    },
  },
  server: {
    port: 8081,
    host: true,
    proxy: {
      // 代理所有 /api 请求到后端服务器
      '/api': {
        target: 'https://pan.aqi125.cn',
        changeOrigin: true,
      },
    },
  },
  preview: {
    port: 8081,
  },
  build: {
    outDir: 'dist',
    assetsDir: 'assets',
    sourcemap: false,
  },
})
