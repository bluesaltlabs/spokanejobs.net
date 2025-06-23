import { defineStore } from 'pinia'
import { githubClient } from '@/lib/githubClient'

export const useCompany = defineStore('company', {
  // state
  state: () => ({
    company: null,
  }),
  actions: {
    async fetchCompany(slug) {
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
    },
  },
})
