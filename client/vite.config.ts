import path from 'path';
import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';

export default defineConfig({
  plugins: [solidPlugin()],
  server: {
    host: true,
    port: 15010,
    proxy: {
      '/api': {
        target: process.env.VITE_PROXY_API_TARGET,
        changeOrigin: true
      }
    }
  },
  root: path.join(__dirname, 'src'),
  publicDir: path.join(__dirname, 'static'),
  build: {
    target: 'esnext',
  },
});
