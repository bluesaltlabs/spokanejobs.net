import { defineStore } from 'pinia'
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
      return state.companies?.slice().sort((a, b) => a.name.localeCompare(b.name))
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
    async createCompany(company) {
      const { error } = await supabase.from('companies').insert([company])

      if (error) throw error
      await this.fetchCompanies()
    },
    async updateCompany(company) {
      const { error } = await supabase.from('companies').update(company).eq('id', company.id)

      if (error) throw error
      await this.fetchCompanies()
    },
    async deleteCompany(id) {
      const { error } = await supabase.from('companies').delete().eq('id', id)

      if (error) throw error
      await this.fetchCompanies()
    },
    updateFilters(newFilters) {
      const { search, ...rest } = newFilters
      this.filters = { search: search ?? '', ...rest }
      // fetchCompanies()
    },
  },
})
