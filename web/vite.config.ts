import { createLogger as createLoggerRaw, defineConfig, splitVendorChunkPlugin } from 'vite'
import tsconfigPaths from 'vite-tsconfig-paths'
import react from '@vitejs/plugin-react'

const logger = {
  base: createLoggerRaw('info', { prefix: '[vite-proxy]', allowClearScreen: false }),
  info: (msg: string) => {
    logger.base.info(msg, { timestamp: true, clear: false })
  },
  warn: (msg: string) => {
    logger.base.warn(msg, { timestamp: true, clear: false })
  },
  error: (msg: string) => {
    logger.base.error(msg, { timestamp: true, clear: false })
  },
}

export default defineConfig({
  build: {
    manifest: true,
    sourcemap: false,
    rollupOptions: {
      output: {
        manualChunks: {
          'react-libs': ['react', 'react-dom', 'react-router-dom'],
          'redux-libs': ['react-redux', 'redux-logger', 'redux-persist'],
        },
      },
    },
    outDir: 'build',
  },
  server: {
    host: 'localhost',
    port: 8888,
    proxy: {
      '^/((auth|api)/.*)|meta': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        xfwd: true,
        secure: false,
        configure: (proxy) => {
          proxy.on('error', (err) => {
            logger.error(`error when starting dev server:\n${err.stack}`)
          })
          proxy.on('proxyReq', (proxyReq, req) => {
            logger.info(`request - method: ${req.method} url: ${req.url}`)
          })
          proxy.on('proxyRes', (proxyRes, req) => {
            logger.info(`response - method: ${req.method} url: ${req.url} status_code: ${proxyRes.statusCode}`)
          })
        },
      },
    },
  },
  plugins: [react(), splitVendorChunkPlugin(), tsconfigPaths()],
})
