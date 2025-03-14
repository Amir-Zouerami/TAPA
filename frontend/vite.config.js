import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite';
import viteReact from '@vitejs/plugin-react';
import { TanStackRouterVite } from '@tanstack/router-plugin/vite';

const ReactCompilerConfig = {};

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [
		tailwindcss(),
		TanStackRouterVite({ autoCodeSplitting: true }),
		viteReact({ babel: { plugins: ['babel-plugin-react-compiler', ReactCompilerConfig] } }),
	],
	test: {
		globals: true,
		environment: 'jsdom',
	},
});
