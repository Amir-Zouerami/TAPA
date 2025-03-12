import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
// @ts-ignore
import tailwindcss from '@tailwindcss/vite';

const ReactCompilerConfig = {};

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [tailwindcss(), react({ babel: { plugins: ['babel-plugin-react-compiler', ReactCompilerConfig] } })],
});
