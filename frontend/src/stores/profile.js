import { defineStore } from 'pinia'
import { openDB } from 'idb'
import {
  PersonalInformation, WorkExperience, EducationExperience
} from '@/models'

// import {
//   PersonalInformation, WorkExperience, EducationExperience, LicenseCertification,
//   Membership, Skill, Interest, Project, Reference
// } from '@/models'

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
    personalInformation: new PersonalInformation(),
    dark_mode: false, // add dark_mode to state
    resumeEntries: [], // array of WorkExperience
    educationEntries: [], // array of EducationExperience
  }),
  actions: {
    async loadProfile() {
      const db = await getDb()
      const data = await db.get(STORE_NAME, 'user')
      if (data) {
        this.dark_mode = data.dark_mode || false
        this.personalInformation = new PersonalInformation(data.personalInformation)
        this.resumeEntries = Array.isArray(data.resumeEntries)
          ? data.resumeEntries.map(e => new WorkExperience(e))
          : []
        this.educationEntries = Array.isArray(data.educationEntries)
          ? data.educationEntries.map(e => new EducationExperience(e))
          : []
      }
    },
    async saveProfile() {
      const db = await getDb()
      // De-proxy and deeply clone the data before saving
      const dataToSave = JSON.parse(JSON.stringify({
        dark_mode: this.dark_mode,
        personalInformation: this.personalInformation,
        resumeEntries: this.resumeEntries,
        educationEntries: this.educationEntries,
      }))
      await db.put(STORE_NAME, dataToSave, 'user')
    },
    updatePersonalInformationField(field, value) {
      if (field in this.personalInformation) {
        this.personalInformation[field] = value
      }
    },
    setDarkMode(value) {
      this.dark_mode = value
      this.saveProfile()
    },
    // Resume Entries Actions
    addResumeEntry(entry) {
      // entry: WorkExperience or plain object
      this.resumeEntries.push(entry instanceof WorkExperience ? entry : new WorkExperience(entry))
      this.saveProfile()
    },
    editResumeEntry(id, updatedEntry) {
      const idx = this.resumeEntries.findIndex(e => e.id === id)
      if (idx !== -1) {
        this.resumeEntries[idx] = new WorkExperience({ ...this.resumeEntries[idx], ...updatedEntry })
        this.saveProfile()
      }
    },
    removeResumeEntry(id) {
      this.resumeEntries = this.resumeEntries.filter(e => e.id !== id)
      this.saveProfile()
    },
    // Education Entries Actions
    addEducationEntry(entry) {
      // entry: EducationExperience or plain object
      this.educationEntries.push(entry instanceof EducationExperience ? entry : new EducationExperience(entry))
      this.saveProfile()
    },
    editEducationEntry(id, updatedEntry) {
      const idx = this.educationEntries.findIndex(e => e.id === id)
      if (idx !== -1) {
        this.educationEntries[idx] = new EducationExperience({ ...this.educationEntries[idx], ...updatedEntry })
        this.saveProfile()
      }
    },
    removeEducationEntry(id) {
      this.educationEntries = this.educationEntries.filter(e => e.id !== id)
      this.saveProfile()
    },
  },
})
