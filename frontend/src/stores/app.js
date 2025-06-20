import { defineStore } from 'pinia'
import { supabase } from '@/lib/supabaseClient'

export const useCompanies = defineStore('app', {
  // state
  state: () => ({
    expandSidebar: true
  }),
  getters: {
    
  },
  actions: {
    expandSidebar() {
      this.expandSidebar = !this.expandSidebar
    }
  }
})
