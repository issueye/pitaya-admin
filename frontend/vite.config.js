import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'
import path from 'path'
import Components from 'unplugin-vue-components/vite'
import { VxeResolver } from '@vxecli/import-unplugin-vue-components'
import Icons from 'unplugin-icons/vite'

// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src/'),
    }
  },

  plugins: [
    vue(),
    Icons({}),
    createSvgIconsPlugin({
      // 指定需要缓存的图标文件夹
      iconDirs: [path.resolve(process.cwd(), 'src/icons')],
      // 指定symbolId格式
      symbolId: 'icon-[name]',
    }),

    Components({
      resolvers: [
        VxeResolver({ libraryName: 'vxe-table' }),
      ],
    }),
  ],

  server: {
    host: '0.0.0.0',
    port: 8089,
    proxy: {
      '/resources': {
        target: 'http://127.0.0.1:10065',
        rewrite: (path) => path.replace(/^\/resources/, "/resources")
      },
      '/page/api': {
        target: 'http://127.0.0.1:10065',
        rewrite: (path) => path.replace(/^\/api/, "/api")
      },
      '/api': {
        target: 'http://127.0.0.1:10065',
        rewrite: (path) => path.replace(/^\/api/, "/api")
      },
    },
  }
})
