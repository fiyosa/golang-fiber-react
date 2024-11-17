import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],

  envDir: '../..',

  server: {
    port: 3000,
    host: 'localhost',
  },

  build: {
    outDir: '../../public',
    copyPublicDir: true,
    emptyOutDir: false,
    manifest: true,
    rollupOptions: {
      input: './src/main.tsx',
      output: {
        inlineDynamicImports: true,
        entryFileNames: 'main.js',
        assetFileNames: (assetInfo) => {
          if (assetInfo.name === 'main.css') return 'main.css'
          return 'assets/[name]-[hash].[ext]'
        },
      },
    },
  },
})
