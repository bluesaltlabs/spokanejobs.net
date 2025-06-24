import { defineStore } from 'pinia'
import { githubClient } from '@/lib/githubClient'

export const useCompanies = defineStore('companies', {
  // state
  state: () => ({
    companies: [],
    filters: { search: '' },
    loading: false,
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
      this.loading = true
      try {
        const { data, error } = await githubClient.fetchCompanies()

        if (error) {
          console.error(error)
          throw error
        }
        this.companies = data
      } finally {
        this.loading = false
      }
    },

    updateFilters(newFilters) {
      const { search, ...rest } = newFilters
      this.filters = { search: search ?? '', ...rest }
      // fetchCompanies()
    },
  },
})
