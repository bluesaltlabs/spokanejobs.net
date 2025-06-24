<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useProfileStore } from '@/stores/profile'
import { UiButton, UiTextInput, UiTextareaInput, UiDateInput, UiForm, UiFormGroup } from '@/components/ui'
import ExportIcon from '@/components/icons/Export.vue'
import ImportIcon from '@/components/icons/Import.vue'
import ViewIcon from '@/components/icons/View.vue'

const router = useRouter()
const profile = useProfileStore()
const saved = ref(false)
const importSuccess = ref(false)

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

// Import/Export functionality
const exportProfile = () => {
  const profileData = {
    first_name: profile.first_name,
    last_name: profile.last_name,
    email: profile.email,
    avatar: profile.avatar,
    dark_mode: profile.dark_mode,
    resumeEntries: profile.resumeEntries,
    educationEntries: profile.educationEntries,
    exportedAt: new Date().toISOString()
  }
  
  const dataStr = JSON.stringify(profileData, null, 2)
  const dataBlob = new Blob([dataStr], { type: 'application/json' })
  
  const link = document.createElement('a')
  link.href = URL.createObjectURL(dataBlob)
  link.download = `profile-${profile.first_name || 'user'}-${new Date().toISOString().split('T')[0]}.json`
  link.click()
  URL.revokeObjectURL(link.href)
}

const importProfile = (event) => {
  const file = event.target.files[0]
  if (!file) return
  
  const reader = new FileReader()
  reader.onload = async (e) => {
    try {
      const importedData = JSON.parse(e.target.result)
      
      // Validate the imported data structure
      if (typeof importedData !== 'object' || importedData === null) {
        throw new Error('Invalid JSON structure')
      }
      
      // Update profile with imported data
      if (importedData.first_name !== undefined) profile.first_name = importedData.first_name
      if (importedData.last_name !== undefined) profile.last_name = importedData.last_name
      if (importedData.email !== undefined) profile.email = importedData.email
      if (importedData.avatar !== undefined) profile.avatar = importedData.avatar
      if (importedData.dark_mode !== undefined) profile.dark_mode = importedData.dark_mode
      if (Array.isArray(importedData.resumeEntries)) profile.resumeEntries = importedData.resumeEntries
      if (Array.isArray(importedData.educationEntries)) profile.educationEntries = importedData.educationEntries
      
      // Save the imported profile
      await profile.saveProfile()
      
      // Show success message
      importSuccess.value = true
      setTimeout(() => (importSuccess.value = false), 3000)
      
    } catch (error) {
      console.error('Error importing profile:', error)
      alert('Error importing profile data. Please check the file format.')
    }
  }
  
  reader.readAsText(file)
  
  // Reset the file input
  event.target.value = ''
}

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

.header-actions {
  display: flex;
  gap: 0.5rem;
  align-items: center;
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

.save-indicator + .save-indicator {
  margin-top: 0.5rem;
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
  background: var(--color-surface-hover);
  border-radius: var(--border-radius-small);
  border: 1px solid var(--color-border);
}

[data-theme="dark"] .resume-form {
  background: var(--color-surface-dark);
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
  background: var(--color-surface-hover);
  border-radius: var(--border-radius-small);
  border: 1px solid var(--color-border);
}

[data-theme="dark"] .education-form {
  background: var(--color-surface-dark);
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

/* Add iOS-style popup button row styling */
.form-action-row {
  display: flex;
  flex-direction: row;
  gap: 1rem;
  margin-top: 1.25rem;
}
.form-action-btn {
  flex: 1 1 0;
  min-width: 0;
}

.header-actions {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 1rem;
  gap: 1rem;
  
}

.header-actions svg {
  position: relative;
  margin-right: 0.5em;
}

@media (max-width: 1023px) {
  .header-actions  {
    svg {
      margin-right: 0 !important;
    }
  }
  .button-tooltip {
    display: none !important;
  }
}


</style>

<template>
  <div class="profile-edit-view">
    <div class="profile-header">
      <h1>Edit Profile</h1>
    </div>
    <div class="header-actions">
      <UiButton @click="exportProfile" variant="success" aria-label="Export Profile">
          <ExportIcon  />
          <span class="button-tooltip">Export Profile</span>
        </UiButton>
        <UiButton @click="$refs.fileInput.click()" variant="info" aria-label="Import Profile">
          <ImportIcon  />
          <span class="button-tooltip">Import Profile</span>
        </UiButton>
        <UiButton @click="goToView" variant="primary" aria-label="View Profile">
          <ViewIcon  />
          <span class="button-tooltip">View Profile</span>
        </UiButton>

    </div>
    <input ref="fileInput" type="file" accept=".json" style="display: none" @change="importProfile" />
    <div class="profile-edit-content">
      <div class="user-info-section">
        <h2>Personal Information</h2>
        <div class="avatar-preview">
          <div class="avatar">
            <img v-if="profile.avatar" :src="profile.avatar" alt="Avatar preview" />
            <span v-else>{{ profile.first_name?.charAt(0) || 'U' }}</span>
          </div>
        </div>
        <UiForm>
          <UiFormGroup label="First Name">
            <UiTextInput v-model="profile.first_name" placeholder="Enter your first name" />
          </UiFormGroup>
          <UiFormGroup label="Last Name">
            <UiTextInput v-model="profile.last_name" placeholder="Enter your last name" />
          </UiFormGroup>
          <UiFormGroup label="Email">
            <UiTextInput v-model="profile.email" type="email" placeholder="Enter your email address" />
          </UiFormGroup>
          <UiFormGroup label="Avatar URL">
            <UiTextInput v-model="profile.avatar" placeholder="Enter avatar image URL" />
          </UiFormGroup>
          <div v-if="saved" class="save-indicator saved">✓ Changes saved automatically</div>
          <div v-if="importSuccess" class="save-indicator saved">✓ Profile imported successfully</div>
        </UiForm>
      </div>
      <div class="entries-section">
        <div class="resume-section">
          <div class="resume-header">
            <h2>Resume Entries</h2>
            <UiButton @click="startAddEntry" variant="primary">Add Entry</UiButton>
          </div>
          <div v-if="showEntryForm" class="resume-form">
            <UiTextInput v-model="newEntry.jobTitle" placeholder="Job Title" />
            <UiTextInput v-model="newEntry.company" placeholder="Company" />
            <UiDateInput v-model="newEntry.startDate" placeholder="Start Date" />
            <UiDateInput v-model="newEntry.endDate" placeholder="End Date" />
            <UiTextareaInput v-model="newEntry.description" placeholder="Description" />
            <div class="form-action-row">
              <UiButton @click="saveEntry" variant="primary" class="form-action-btn">{{ editingId ? 'Update' : 'Add' }} Entry</UiButton>
              <UiButton @click="cancelEntry" variant="secondary" class="form-action-btn">Cancel</UiButton>
            </div>
          </div>
          <div v-for="entry in profile.resumeEntries" :key="entry.id" class="resume-entry">
            <strong>{{ entry.jobTitle }}</strong> at <em>{{ entry.company }}</em>
            <span>{{ entry.startDate }} - {{ entry.endDate }}</span>
            <p>{{ entry.description }}</p>
            <div class="resume-entry-actions">
              <UiButton @click="startEditEntry(entry)" variant="primary" size="small">Edit</UiButton>
              <UiButton @click="removeEntry(entry.id)" variant="danger" size="small">Delete</UiButton>
            </div>
          </div>
          <div v-if="profile.resumeEntries.length === 0" class="empty-resume">
            <p>No resume entries yet. Click "Add Entry" to get started.</p>
          </div>
        </div>
        <div class="education-section">
          <div class="education-header">
            <h2>Education History</h2>
            <UiButton @click="startAddEducationEntry" variant="primary">Add Education</UiButton>
          </div>
          <div v-if="showEducationForm" class="education-form">
            <UiTextInput v-model="newEducationEntry.degree" placeholder="Degree" />
            <UiTextInput v-model="newEducationEntry.institution" placeholder="Institution" />
            <UiDateInput v-model="newEducationEntry.startDate" placeholder="Start Date" />
            <UiDateInput v-model="newEducationEntry.endDate" placeholder="End Date" />
            <UiTextareaInput v-model="newEducationEntry.description" placeholder="Description" />
            <div class="form-action-row">
              <UiButton @click="saveEducationEntry" variant="primary" class="form-action-btn">{{ editingEducationId ? 'Update' : 'Add' }} Education</UiButton>
              <UiButton @click="cancelEducationEntry" variant="secondary" class="form-action-btn">Cancel</UiButton>
            </div>
          </div>
          <div v-for="entry in profile.educationEntries" :key="entry.id" class="education-entry">
            <strong>{{ entry.degree }}</strong> from <em>{{ entry.institution }}</em>
            <span>{{ entry.startDate }} - {{ entry.endDate }}</span>
            <p>{{ entry.description }}</p>
            <div class="education-entry-actions">
              <UiButton @click="startEditEducationEntry(entry)" variant="primary" size="small">Edit</UiButton>
              <UiButton @click="removeEducationEntry(entry.id)" variant="danger" size="small">Delete</UiButton>
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
