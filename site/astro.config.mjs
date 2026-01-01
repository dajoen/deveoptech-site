import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

// https://astro.build/config
export default defineConfig({
  integrations: [
    starlight({
      title: "DevOpsTech",
      social: {
        github: "https://github.com/jverhoeven/devopstech-site",
      },
      sidebar: [
        {
          label: "Linux",
          autogenerate: { directory: "linux" },
        },
        {
          label: "CLI Toolchain",
          autogenerate: { directory: "cli" },
        },
        {
          label: "DevOps",
          autogenerate: { directory: "devops" },
        },
      ],
      editLink: {
        baseUrl:
          "https://github.com/jverhoeven/devopstech-site/edit/main/docs/",
      },
    }),
  ],
});
