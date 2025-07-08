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
    dark_mode: false, // add dark_mode to state
    personal_information: new PersonalInformation(),
    work_experiences: [],
    education_experiences: [],
  }),
  actions: {
    async loadProfile() {
      const db = await getDb()
      const data = await db.get(STORE_NAME, 'user')
      if (data) {
        this.dark_mode = data?.dark_mode || false
        // Personal Information
        this.personal_information = new PersonalInformation(data?.personal_information ?? [])
        // Work Experience
        this.work_experiences = (data?.work_experiences ?? []).map(e => new WorkExperience(e))
        // Education Experience
        this.education_experiences = (data?.education_experiences ?? []).map(e => new EducationExperience(e))
        // todo:
      }
    },
    async saveProfile() {
      const db = await getDb()

      // Prep data for saving.
      const dataToSave = JSON.parse(JSON.stringify({
        dark_mode: this.dark_mode,
        personal_information: this.personal_information,
        work_experiences: this.work_experiences,
        education_experiences: this.education_experiences,
      }))
      await db.put(STORE_NAME, dataToSave, 'user')
    },

    setDarkMode(value) {
      this.dark_mode = value
      this.saveProfile()
    },

    // Work Experience Actions
    addWorkExperience(entry) {
      this.work_experiences.push(entry instanceof WorkExperience ? entry : new WorkExperience(entry))
      this.saveProfile()
    },
    editWorkExperience(id, updatedEntry) {
      const idx = this.work_experiences.findIndex(e => e.id === id)
      if (idx !== -1) {
        this.work_experiences[idx] = new WorkExperience({ ...this.work_experiences[idx], ...updatedEntry })
        this.saveProfile()
      }
    },
    removeWorkExperience(id) {
      this.work_experiences = this.work_experiences.filter(e => e.id !== id)
      this.saveProfile()
    },

    // Education Experience Actions
    addEducationExperience(entry) {
      this.education_experiences.push(entry instanceof EducationExperience ? entry : new EducationExperience(entry))
      this.saveProfile()
    },
    editEducationExperience(id, updatedEntry) {
      const idx = this.education_experiences.findIndex(e => e.id === id)
      if (idx !== -1) {
        this.education_experiences[idx] = new EducationExperience({ ...this.education_experiences[idx], ...updatedEntry })
        this.saveProfile()
      }
    },
    removeEducationExperience(id) {
      this.education_experiences = this.education_experiences.filter(e => e.id !== id)
      this.saveProfile()
    },
  },
})
