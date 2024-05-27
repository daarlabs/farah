import { defineConfig } from 'vite'
import tsconfigPaths from "vite-tsconfig-paths";

export default defineConfig({
	build: {
		manifest: "manifest.json",
		rollupOptions: {
			input: "./src/main.ts",
		},
		outDir: "../web/public/dist",
		assetsDir: "",
		emptyOutDir: true,
	},
	plugins: [tsconfigPaths()],
})
