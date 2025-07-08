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
          { text: 'Personal Information', link: '/profile/personal-information' },
          { text: 'Work Experience', link: '/profile/work-experience' },
          { text: 'Education Experience', link: '/profile/education-experience' },
          { text: 'Licenses & Certifications', link: '/profile/licenses-certifications' },
          { text: 'Memberships', link: '/profile/memberships' },
          { text: 'Skills', link: '/profile/skills' },
          { text: 'Interests', link: '/profile/interests' },
          { text: 'Projects', link: '/profile/projects' },
          { text: 'References', link: '/profile/references' },
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
