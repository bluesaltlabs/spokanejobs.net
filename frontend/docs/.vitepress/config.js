import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Bedrock",
  description: "Bedrock Job Search Tools",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
    ],

    sidebar: [
      {
        text: 'Overview',
        items: [
          { text: 'Overview', link: '/overview' },
        ]
      },
      {
        text: 'Profile',
        items: [
          { text: 'Overview', link: '/profile/overview' },
        ]
      },
      {
        text: 'Misc',
        items: [
          { text: 'Markdown Examples', link: '/misc/markdown-examples' },
          { text: 'Runtime API Examples', link: '/misc/api-examples' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
    ]
  }
})
