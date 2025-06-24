import { defineStore } from 'pinia'
import { githubClient } from '@/lib/githubClient'

export const useJobs = defineStore('jobs', {
  // state
  state: () => ({
    jobs: [],
    filters: { search: '' },
    loading: false,
  }),
  getters: {
    countJobs(state) {
      return state.jobs?.length ?? 0
    },
    sortedJobs(state) {
      // return the jobs sorted by name
      return state.jobs?.sort((a, b) => a.name.localeCompare(b.name))
    }
  },
  actions: {
    async fetchJobs() {
      this.loading = true
      try {
        const { data, error } = await githubClient.fetchJobs()

        if (error) {
          console.error(error)
          throw error
        }
        this.jobs = data
      } finally {
        this.loading = false
      }
    },

    updateFilters(newFilters) {
      const { search, ...rest } = newFilters
      this.filters = { search: search ?? '', ...rest }
      // fetchJobs()
    },
  },
})
