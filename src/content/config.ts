import { defineCollection, z } from "astro:content";

const projects = defineCollection({
	type: "content",
	schema: z.object({
		title: z.string(),
		description: z.string(),
		url: z.union([z.string().url(), z.literal("")]),
		type: z.enum(["github", "youtube", "work"]),
		references: z.array(
			z.object({
				title: z.string(),
				url: z.string().url(),
			})
		),
	}),
});

export const collections = {
	projects,
};
