import { defineStore } from 'pinia'
import { supabase } from '@/lib/supabaseClient'

export const useJobs = defineStore('jobs', {
  // state
  state: () => ({
    jobs: [],
    filters: { search: '' },
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
      const { data, error } = await supabase.from('jobs').select('*') // todo: this probably needs to use filters

      if (error) {
        console.error(error)
        throw error
      }
      this.jobs = data
    },
    updateFilters(newFilters) {
      const { search, ...rest } = newFilters
      this.filters = { search: search ?? '', ...rest }
      // fetchJobs()
    },
    // addJob(attributes) {
    //   this.jobs.push({
    //     name: attributes?.name,
    //     slug: attributes?.slug,
    //     description: attributes?.description,
    //     website: attributes?.website,
    //     job_board_url: attributes?.job_board_url,
    //     logo_url: attributes?.logo_url,
    //     industry: attributes?.industry,
    //   })
    // },
  },
})
