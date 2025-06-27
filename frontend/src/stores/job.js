import { defineStore } from 'pinia'
import { githubClient } from '@/lib/githubClient'

export const useJob = defineStore('job', {
  // state
  state: () => ({
    job: null,
    loading: false,
  }),
  actions: {
    async fetchJob(id) {
      this.loading = true
      try {
        const { data, error } = await githubClient.fetchJobs()
        
        if (error) {
          console.error(error)
          throw error
        }
        
        // Find the job by id from the fetched data
        const job = data.find(job => job.id === id || job.job_id === id)
        
        if (!job) {
          throw new Error(`Job with id '${id}' not found`)
        }
        
        this.job = job
      } finally {
        this.loading = false
      }
    },
  },
})
