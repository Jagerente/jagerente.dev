// https://astro.build/config
import { defineConfig } from 'astro/config';
import mdx from '@astrojs/mdx';
import sitemap from '@astrojs/sitemap';

import cloudflare from "@astrojs/cloudflare";

export default defineConfig({
	site: 'https://jagerente.xyz',
	integrations: [mdx(), sitemap()],
	output: 'server',
	adapter: cloudflare(),
});