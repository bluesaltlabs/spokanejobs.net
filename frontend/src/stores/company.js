import { defineStore } from 'pinia'
import { supabase } from '@/lib/supabaseClient'

export const useCompany = defineStore('company', {
  // state
  state: () => ({
    company: [],
    // company: {
    //   id: 0,
    //   slug: '',
    //   //
    // },
  }),
  getters: {
    //
  },
  actions: {
    async fetchCompany(slug) {
      const { data, error } = await supabase.from('companies').select('id, slug').eq('slug', slug)
      //const { data, error } = await supabase.from('companies').select('where id = ') // todo: this probably needs to use filters

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
