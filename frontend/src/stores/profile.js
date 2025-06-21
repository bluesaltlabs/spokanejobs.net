import { defineStore } from 'pinia'
import { openDB } from 'idb'

const DB_NAME = 'userProfileDB'
const STORE_NAME = 'profile'

async function getDb() {
  return openDB(DB_NAME, 1, {
    upgrade(db) {
      if (!db.objectStoreNames.contains(STORE_NAME)) {
        db.createObjectStore(STORE_NAME)
      }
    },
  })
}

export const useProfileStore = defineStore('profile', {
  state: () => ({
    first_name: '',
    last_name: '',
    email: '',
    avatar: '',
    dark_mode: false, // add dark_mode to state
    resumeEntries: [], // add resumeEntries to state
  }),
  actions: {
    async loadProfile() {
      const db = await getDb()
      const data = await db.get(STORE_NAME, 'user')
      if (data) {
        this.$patch(data)
      }
    },
    async saveProfile() {
      const db = await getDb()
      // De-proxy and deeply clone the data before saving
      const dataToSave = JSON.parse(JSON.stringify({
        first_name: this.first_name,
        last_name: this.last_name,
        email: this.email,
        avatar: this.avatar,
        dark_mode: this.dark_mode,
        resumeEntries: this.resumeEntries,
      }))
      await db.put(STORE_NAME, dataToSave, 'user')
    },
    updateField(field, value) {
      this[field] = value
    },
    setDarkMode(value) {
      this.dark_mode = value
      this.saveProfile()
    },
    // Resume Entries Actions
    addResumeEntry(entry) {
      // entry: { id, jobTitle, company, startDate, endDate, description, ... }
      this.resumeEntries.push(entry)
      this.saveProfile()
    },
    editResumeEntry(id, updatedEntry) {
      const idx = this.resumeEntries.findIndex(e => e.id === id)
      if (idx !== -1) {
        this.resumeEntries[idx] = { ...this.resumeEntries[idx], ...updatedEntry }
        this.saveProfile()
      }
    },
    removeResumeEntry(id) {
      this.resumeEntries = this.resumeEntries.filter(e => e.id !== id)
      this.saveProfile()
    },
  },
})
