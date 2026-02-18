// https://astro.build/config
import { defineConfig } from 'astro/config';
import mdx from '@astrojs/mdx';
import sitemap from '@astrojs/sitemap';

import cloudflare from "@astrojs/cloudflare";

export default defineConfig({
	site: 'https://jagerente.dev',
	integrations: [
		mdx(),
		sitemap({
			filter: (page) => !page.includes("/projects"),
		}),
	],
	output: 'server',
	adapter: cloudflare(),
});
