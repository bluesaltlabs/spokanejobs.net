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
      // filter by search
      let filtered = state.jobs
      const search = state.filters.search?.toLowerCase().trim()
      if (search) {
        filtered = filtered.filter(j =>
          j.title?.toLowerCase().includes(search) ||
          j.company?.toLowerCase().includes(search)
        )
      }
      // Defensive: sort by title if present, fallback to job_id
      return [...(filtered ?? [])].sort((a, b) => {
        if (a.title && b.title) {
          return a.title.localeCompare(b.title)
        }
        if (a.title) return -1
        if (b.title) return 1
        // fallback to job_id string compare
        if (a.job_id && b.job_id) {
          return a.job_id.localeCompare(b.job_id)
        }
        return 0
      })
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
