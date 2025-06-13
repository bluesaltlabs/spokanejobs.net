import { defineStore } from 'pinia'
import { ref, onMounted } from 'vue'
import { supabase } from '@/lib/supabaseClient'

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
      return state.companies?.sort((a, b) => a.name.localeCompare(b.name))
    }
  },
  actions: {
    async fetchCompanies() {
      const { data, error } = await supabase.from('companies').select('*') // todo: this probably needs to use filters

      if (error) {
        console.error(error)
        throw error
      }
      this.companies = data
    },
    updateFilters(newFilters) {
      const { search, ...rest } = newFilters
      this.filters = { search: search ?? '', ...rest }
      // fetchCompanies()
    },
    // addCompany(attributes) {
    //   this.companies.push({
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
