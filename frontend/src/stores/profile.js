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
      await db.put(STORE_NAME, {
        first_name: this.first_name,
        last_name: this.last_name,
        email: this.email,
        avatar: this.avatar,
      }, 'user')
    },
    updateField(field, value) {
      this[field] = value
    },
  },
})
