import { defineStore } from 'pinia'
import { githubClient } from '@/lib/githubClient'

export const useCompanies = defineStore('companies', {
  // state
  state: () => ({
    companies: [],
    filters: { search: '' },
  }),
  getters: {
    countCompanies(state) {
      return state.companies?.length ?? 0
    },
    sortedCompanies(state) {
      // return the companies sorted by name
      return state.companies?.slice().sort((a, b) => a.name.localeCompare(b.name))
    }
  },
  actions: {
    async fetchCompanies() {
      const { data, error } = await githubClient.fetchCompanies()

      if (error) {
        console.error(error)
        throw error
      }
      this.companies = data
    },
    // Note: Create, update, and delete operations are removed since we're reading from a static JSON file
    // These operations would need to be handled differently if write access is required
    updateFilters(newFilters) {
      const { search, ...rest } = newFilters
      this.filters = { search: search ?? '', ...rest }
      // fetchCompanies()
    },
  },
})
