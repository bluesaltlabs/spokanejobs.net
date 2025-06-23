<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useProfileStore } from '@/stores/profile'
import BaseButton from '@/components/BaseButton.vue'

const router = useRouter()
const profile = useProfileStore()
const saved = ref(false)

// Auto-save functionality
const autoSave = async () => {
  await profile.saveProfile()
  saved.value = true
  setTimeout(() => (saved.value = false), 1500)
}

// Watch for changes in profile fields and auto-save
watch(
  () => [profile.first_name, profile.last_name, profile.email, profile.avatar],
  () => {
    autoSave()
  },
  { deep: true }
)

onMounted(() => {
  profile.loadProfile()
})

// Resume Entry Management
const newEntry = ref({
  id: '',
  jobTitle: '',
  company: '',
  startDate: '',
  endDate: '',
  description: ''
})
const editingId = ref(null)
const showEntryForm = ref(false)

// Education Entry Management
const newEducationEntry = ref({
  id: '',
  degree: '',
  institution: '',
  startDate: '',
  endDate: '',
  description: ''
})
const editingEducationId = ref(null)
const showEducationForm = ref(false)

function resetNewEntry() {
  newEntry.value = {
    id: '',
    jobTitle: '',
    company: '',
    startDate: '',
    endDate: '',
    description: ''
  }
}

function resetNewEducationEntry() {
  newEducationEntry.value = {
    id: '',
    degree: '',
    institution: '',
    startDate: '',
    endDate: '',
    description: ''
  }
}

function startAddEntry() {
  editingId.value = null
  resetNewEntry()
  showEntryForm.value = true
}

function startAddEducationEntry() {
  editingEducationId.value = null
  resetNewEducationEntry()
  showEducationForm.value = true
}

function startEditEntry(entry) {
  editingId.value = entry.id
  newEntry.value = { ...entry }
  showEntryForm.value = true
}

function startEditEducationEntry(entry) {
  editingEducationId.value = entry.id
  newEducationEntry.value = { ...entry }
  showEducationForm.value = true
}

async function saveEntry() {
  if (!newEntry.value.jobTitle || !newEntry.value.company) return
  if (editingId.value) {
    profile.editResumeEntry(editingId.value, { ...newEntry.value })
  } else {
   profile.addResumeEntry({ ...newEntry.value, id: Date.now().toString() })
  }
  await profile.loadProfile()
  editingId.value = null
  resetNewEntry()
  showEntryForm.value = false
}

async function saveEducationEntry() {
  if (!newEducationEntry.value.degree || !newEducationEntry.value.institution) return
  if (editingEducationId.value) {
    profile.editEducationEntry(editingEducationId.value, { ...newEducationEntry.value })
  } else {
   profile.addEducationEntry({ ...newEducationEntry.value, id: Date.now().toString() })
  }
  await profile.loadProfile()
  editingEducationId.value = null
  resetNewEducationEntry()
  showEducationForm.value = false
}

function cancelEntry() {
  resetNewEntry()
  editingId.value = null
  showEntryForm.value = false
}

function cancelEducationEntry() {
  resetNewEducationEntry()
  editingEducationId.value = null
  showEducationForm.value = false
}

async function removeEntry(id) {
  profile.removeResumeEntry(id)
  await profile.loadProfile()
}

async function removeEducationEntry(id) {
  profile.removeEducationEntry(id)
  await profile.loadProfile()
}

function goToView() {
  router.push('/profile')
}
</script>

<style scoped>
.profile-edit-view {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
  background: var(--background-color);
  border-radius: var(--border-radius-medium);
  box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}

.profile-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 2rem;
}

.profile-header h1 {
  margin: 0;
}

.profile-edit-content {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

/* Desktop 2-column layout */
@media (min-width: 768px) {
  .profile-edit-content {
    flex-direction: row;
    align-items: flex-start;
    gap: 3rem;
  }

  .user-info-section {
    flex: 0 0 400px;
    margin-bottom: 0;
  }

  .entries-section {
    flex: 1;
    margin-top: 0;
    padding-top: 0;
    border-top: none;
  }

  .resume-section {
    margin-top: 0;
    padding-top: 0;
    border-top: none;
  }

  .education-section {
    margin-top: 2rem;
    padding-top: 2rem;
    border-top: 1px solid var(--color-border);
  }
}

.user-info-section {
  margin-bottom: 2rem;
}

.user-info-section h2 {
  margin-bottom: 1.5rem;
  color: var(--color-heading);
  font-size: 1.5rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: var(--color-text);
}

.themed-input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #b3c6e0;
  border-radius: var(--border-radius-small);
  font-size: 1rem;
  background: #fafdff;
  color: #222;
  transition: border 0.2s, box-shadow 0.2s;
}

.themed-input:focus {
  border: 1.5px solid #4f8cff;
  outline: none;
  box-shadow: 0 0 0 2px #4f8cff22;
}

.avatar-preview {
  text-align: center;
  margin-top: 1rem;
}

.avatar {
  width: 100px;
  height: 100px;
  background: var(--color-surface-hover);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-subtle);
  border-radius: var(--border-radius-medium);
  margin: 0 auto;

  img {
    width: 100px;
    height: 100px;
    border-radius: var(--border-radius-small);
    object-fit: cover;
  }
}

.save-indicator {
  margin-top: 1rem;
  padding: 0.5rem;
  border-radius: var(--border-radius-small);
  text-align: center;
  font-weight: 500;
}

.save-indicator.saved {
  background: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.resume-section {
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid var(--color-border);
}

.resume-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.resume-header h2 {
  margin: 0;
  color: var(--color-heading);
  font-size: 1.5rem;
}

.resume-form {
  margin-bottom: 1.5rem;
  padding: 1rem;
  background: #f5f7fa;
  border-radius: var(--border-radius-small);
  border: 1px solid var(--color-border);
}

.resume-form input, .resume-form textarea {
  width: 100%;
  margin-bottom: 0.5rem;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: var(--border-radius-small);
}

.resume-entry {
  border: 1.5px solid var(--color-border);
  border-radius: var(--border-radius-medium);
  padding: 1.25rem 1rem 1rem 1rem;
  margin-bottom: 1.25rem;
  background: var(--color-surface);
  box-shadow: 0 1px 4px var(--color-shadow);
  transition: background 0.2s, border 0.2s;
}

[data-theme="dark"] .resume-entry {
  background: var(--color-surface-dark);
  border-color: var(--color-border-dark);
  box-shadow: 0 1px 6px var(--color-shadow-elevated);
}

.resume-entry strong {
  color: var(--color-primary-600);
  font-size: 1.1rem;
  font-weight: 700;
}

.resume-entry em {
  color: var(--color-primary-500);
  font-style: normal;
  font-weight: 600;
}

.resume-entry span {
  color: var(--color-text-muted);
  font-size: 0.95rem;
  display: block;
  margin-bottom: 0.5rem;
}

.resume-entry p {
  color: var(--color-text);
  margin: 0.5rem 0 0 0;
  font-size: 1rem;
}

.resume-entry-actions {
  margin-top: 0.75rem;
  display: flex;
  gap: 0.5rem;
}

.empty-resume {
  text-align: center;
  padding: 2rem;
  color: var(--color-text-subtle);
  background: var(--color-surface-hover);
  border-radius: 8px;
  border: 1px dashed var(--color-border);
}

.empty-resume p {
  margin: 0;
  font-size: 1rem;
}

.education-section {
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid var(--color-border);
}

.education-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.education-header h2 {
  margin: 0;
  color: var(--color-heading);
  font-size: 1.5rem;
}

.education-form {
  margin-bottom: 1.5rem;
  padding: 1rem;
  background: #f5f7fa;
  border-radius: var(--border-radius-small);
  border: 1px solid var(--color-border);
}

.education-form input, .education-form textarea {
  width: 100%;
  margin-bottom: 0.5rem;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: var(--border-radius-small);
}

.education-entry {
  border: 1.5px solid var(--color-border);
  border-radius: var(--border-radius-medium);
  padding: 1.25rem 1rem 1rem 1rem;
  margin-bottom: 1.25rem;
  background: var(--color-surface);
  box-shadow: 0 1px 4px var(--color-shadow);
  transition: background 0.2s, border 0.2s;
}

[data-theme="dark"] .education-entry {
  background: var(--color-surface-dark);
  border-color: var(--color-border-dark);
  box-shadow: 0 1px 6px var(--color-shadow-elevated);
}

.education-entry strong {
  color: var(--color-primary-600);
  font-size: 1.1rem;
  font-weight: 700;
}

.education-entry em {
  color: var(--color-primary-500);
  font-style: normal;
  font-weight: 600;
}

.education-entry span {
  color: var(--color-text-muted);
  font-size: 0.95rem;
  display: block;
  margin-bottom: 0.5rem;
}

.education-entry p {
  color: var(--color-text);
  margin: 0.5rem 0 0 0;
  font-size: 1rem;
}

.education-entry-actions {
  margin-top: 0.75rem;
  display: flex;
  gap: 0.5rem;
}

.empty-education {
  text-align: center;
  padding: 2rem;
  color: var(--color-text-subtle);
  background: var(--color-surface-hover);
  border-radius: 8px;
  border: 1px dashed var(--color-border);
}

.empty-education p {
  margin: 0;
  font-size: 1rem;
}
</style>

<template>
  <div class="profile-edit-view">
    <div class="profile-header">
      <h1>Edit Profile</h1>
      <BaseButton @click="goToView" variant="secondary">View Profile</BaseButton>
    </div>

    <div class="profile-edit-content">
      <!-- Left Column - User Info -->
      <div class="user-info-section">
        <h2>Personal Information</h2>
        <div class="avatar-preview">
          <div class="avatar">
            <img v-if="profile.avatar" :src="profile.avatar" alt="Avatar preview" />
            <span v-else>{{ profile.first_name?.charAt(0) || 'U' }}</span>
          </div>
        </div>
        <form>
          <div class="form-group">
            <label for="firstName">First Name</label>
            <input
              id="firstName"
              v-model="profile.first_name"
              type="text"
              class="themed-input"
              placeholder="Enter your first name"
            />
          </div>

          <div class="form-group">
            <label for="lastName">Last Name</label>
            <input
              id="lastName"
              v-model="profile.last_name"
              type="text"
              class="themed-input"
              placeholder="Enter your last name"
            />
          </div>

          <!-- todo: update this so it's the first option -->
          <!--       and loads from gravatar when a valid email address is entered -->
          <div class="form-group">
            <label for="email">Email</label>
            <input
              id="email"
              v-model="profile.email"
              type="email"
              class="themed-input"
              placeholder="Enter your email address"
            />
          </div>

          <div class="form-group">
            <label for="avatar">Avatar URL</label>
            <input
              id="avatar"
              v-model="profile.avatar"
              type="text"
              class="themed-input"
              placeholder="Enter avatar image URL"
            />
          </div>

          <div v-if="saved" class="save-indicator saved">
            ✓ Changes saved automatically
          </div>
        </form>
      </div>

      <!-- Right Column - Resume and Education Entries -->
      <div class="entries-section">
        <!-- Resume Entries -->
        <div class="resume-section">
          <div class="resume-header">
            <h2>Resume Entries</h2>
            <BaseButton @click="startAddEntry" variant="primary">Add Entry</BaseButton>
          </div>

          <div v-if="showEntryForm" class="resume-form">
            <input v-model="newEntry.jobTitle" placeholder="Job Title" class="themed-input" />
            <input v-model="newEntry.company" placeholder="Company" class="themed-input" />
            <input v-model="newEntry.startDate" placeholder="Start Date" type="date" class="themed-input" />
            <input v-model="newEntry.endDate" placeholder="End Date" type="date" class="themed-input" />
            <textarea v-model="newEntry.description" placeholder="Description" class="themed-input"></textarea>
            <BaseButton @click="saveEntry" variant="primary">{{ editingId ? 'Update' : 'Add' }} Entry</BaseButton>
            <BaseButton @click="cancelEntry" variant="secondary">Cancel</BaseButton>
          </div>

          <div v-for="entry in profile.resumeEntries" :key="entry.id" class="resume-entry">
            <strong>{{ entry.jobTitle }}</strong> at <em>{{ entry.company }}</em>
            <span>{{ entry.startDate }} - {{ entry.endDate }}</span>
            <p>{{ entry.description }}</p>
            <div class="resume-entry-actions">
              <BaseButton @click="startEditEntry(entry)" variant="primary" size="small">Edit</BaseButton>
              <BaseButton @click="removeEntry(entry.id)" variant="danger" size="small">Delete</BaseButton>
            </div>
          </div>

          <div v-if="profile.resumeEntries.length === 0" class="empty-resume">
            <p>No resume entries yet. Click "Add Entry" to get started.</p>
          </div>
        </div>

        <!-- Education History Section -->
        <div class="education-section">
          <div class="education-header">
            <h2>Education History</h2>
            <BaseButton @click="startAddEducationEntry" variant="primary">Add Education</BaseButton>
          </div>

          <div v-if="showEducationForm" class="education-form">
            <input v-model="newEducationEntry.degree" placeholder="Degree" class="themed-input" />
            <input v-model="newEducationEntry.institution" placeholder="Institution" class="themed-input" />
            <input v-model="newEducationEntry.startDate" placeholder="Start Date" type="date" class="themed-input" />
            <input v-model="newEducationEntry.endDate" placeholder="End Date" type="date" class="themed-input" />
            <textarea v-model="newEducationEntry.description" placeholder="Description" class="themed-input"></textarea>
            <BaseButton @click="saveEducationEntry" variant="primary">{{ editingEducationId ? 'Update' : 'Add' }} Education</BaseButton>
            <BaseButton @click="cancelEducationEntry" variant="secondary">Cancel</BaseButton>
          </div>

          <div v-for="entry in profile.educationEntries" :key="entry.id" class="education-entry">
            <strong>{{ entry.degree }}</strong> from <em>{{ entry.institution }}</em>
            <span>{{ entry.startDate }} - {{ entry.endDate }}</span>
            <p>{{ entry.description }}</p>
            <div class="education-entry-actions">
              <BaseButton @click="startEditEducationEntry(entry)" variant="primary" size="small">Edit</BaseButton>
              <BaseButton @click="removeEducationEntry(entry.id)" variant="danger" size="small">Delete</BaseButton>
            </div>
          </div>

          <div v-if="profile.educationEntries.length === 0" class="empty-education">
            <p>No education entries yet. Click "Add Education" to get started.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
