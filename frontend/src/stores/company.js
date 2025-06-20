import { defineStore } from 'pinia'
import { supabase } from '@/lib/supabaseClient'

export const useCompany = defineStore('company', {
  // state
  state: () => ({
    company: null,
  }),
  actions: {
    async fetchCompany(slug) {
      const { data, error } = await supabase.from('companies').select('*').eq('slug', slug).single()
      if (error) {
        console.error(error)
        throw error
      }
      this.company = data
    },
  },
})
