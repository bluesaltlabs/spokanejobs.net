import { defineStore } from 'pinia'

export const useApp = defineStore('app', {
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
