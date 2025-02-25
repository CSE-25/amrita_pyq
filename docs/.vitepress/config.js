import { defineConfig } from "vitepress";

export default defineConfig({
    base: "/amrita_pyq/", // Not to be changed as it's the /<repository-name>/ for deployment.
    title: "Amrita PYQ",
    titleTemplate: ":title | Amrita PYQ",
    description: "A CLI tool for accessing PYQs.",
    cleanUrls: true,

    themeConfig: {
        nav: [
            { text: "Home", link: "/" },
            { text: "Docs", link: "/getting-started" },
        ],
        search: {
            provider: "local",
            placeholder: "Search",
        },
        sidebar: [
            {
                text: "Docs",
                items: [
                    { text: "Getting Started", link: "/getting-started" },
                    { text: "Testing", link: "/testing" },
                ],
            },
        ],
        socialLinks: [
            {
                icon: "github",
                link: "https://github.com/CSE-25/amrita_pyq",
            },
        ],
    },
});
