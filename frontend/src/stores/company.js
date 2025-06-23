import { defineStore } from 'pinia'
import { githubClient } from '@/lib/githubClient'

export const useCompany = defineStore('company', {
  // state
  state: () => ({
    company: null,
    loading: false,
  }),
  actions: {
    async fetchCompany(slug) {
      this.loading = true
      try {
        const { data, error } = await githubClient.fetchCompanies()
        
        if (error) {
          console.error(error)
          throw error
        }
        
        // Find the company by slug from the fetched data
        const company = data.find(company => company.slug === slug)
        
        if (!company) {
          throw new Error(`Company with slug '${slug}' not found`)
        }
        
        this.company = company
      } finally {
        this.loading = false
      }
    },
  },
})
