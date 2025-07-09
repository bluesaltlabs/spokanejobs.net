<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useProfileStore } from '@/stores/profile'
import { Button, TextInput, TextareaInput, DateInput, Form, FormGroup } from '@/components/ui'
import { ExportIcon, ImportIcon, ViewIcon } from '@/components/icons'
import { PersonalInformation, WorkExperience as WorkExperienceModel, EducationExperience } from '@/models'
import WorkExperience from '@/components/profile/WorkExperience.vue'

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
  () => [
    profile.personal_information.first_name,
    profile.personal_information.last_name,
    profile.personal_information.email,
    profile.personal_information.avatar_url
  ],
  () => {
    autoSave() // todo: debounce this.
  },
  { deep: true }
)

onMounted(() => {
  profile.loadProfile()
})

// Import/Export functionality
const exportProfile = () => {
  const profileData = {
    dark_mode: profile.dark_mode,
    personal_information: profile.personal_information,
    work_experiences: profile.work_experiences,
    education_experiences: profile.education_experiences,
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

      if (typeof importedData?.personal_information === 'object' && importedData.personal_information !== null) {
        profile.personal_information = importedData.personal_information
      } else {
        profile.personal_information = new PersonalInformation()
      }
      if (Array.isArray(importedData.work_experiences)) profile.work_experiences = importedData.work_experiences
      if (Array.isArray(importedData.education_experiences)) profile.education_experiences = importedData.education_experiences

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

// Education Entry Management
const newEducationEntry = ref(new EducationExperience())
const editingEducationId = ref(null)
const showEducationForm = ref(false)

function resetNewEducationEntry() {
  newEducationEntry.value = {
    id: '',
    degree: '',
    institution: '',
    start_date: '',
    end_date: '',
    description: ''
  }
}

function startAddEducationEntry() {
  editingEducationId.value = null
  resetNewEducationEntry()
  showEducationForm.value = true
}

function startEditEducationEntry(entry) {
  editingEducationId.value = entry.id
  newEducationEntry.value = { ...entry }
  showEducationForm.value = true
}

async function saveEducationEntry() {
  if (!newEducationEntry.value.degree || !newEducationEntry.value.institution) return
  if (editingEducationId.value) {
    profile.editEducationExperience(editingEducationId.value, { ...newEducationEntry.value })
  } else {
   profile.addEducationExperience({ ...newEducationEntry.value, id: Date.now().toString() })
  }
  await profile.loadProfile()
  editingEducationId.value = null
  resetNewEducationEntry()
  showEducationForm.value = false
}

function cancelEducationEntry() {
  resetNewEducationEntry()
  editingEducationId.value = null
  showEducationForm.value = false
}

async function removeEducationEntry(id) {
  profile.removeEducationExperience(id)
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
      <div class="header-actions">
        <Button @click="goToView" variant="secondary">
          <ViewIcon style="margin-right: 0.5em; vertical-align: middle;" />
          View Profile
        </Button>
        <Button @click="exportProfile" variant="primary">
          <ExportIcon style="margin-right: 0.5em; vertical-align: middle;" />
          Export
        </Button>
        <label class="import-label">
          <ImportIcon style="margin-right: 0.5em; vertical-align: middle;" />
          Import
          <input type="file" accept="application/json" @change="importProfile" style="display: none;" />
        </label>
      </div>
    </div>
    <div class="profile-edit-content">
      <div class="user-info-section">
        <h2>Personal Information</h2>
        <div class="avatar-preview">
          <div class="avatar">
            <img v-if="profile.personal_information.avatar_url" :src="profile.personal_information.avatar_url" alt="Avatar preview" />
            <span v-else>{{ profile.personal_information.first_name?.charAt(0) || 'U' }}</span>
          </div>
        </div>
        <Form>
          <FormGroup label="First Name">
            <TextInput v-model="profile.personal_information.first_name" placeholder="Enter your first name" />
          </FormGroup>
          <FormGroup label="Last Name">
            <TextInput v-model="profile.personal_information.last_name" placeholder="Enter your last name" />
          </FormGroup>
          <FormGroup label="Email">
            <TextInput v-model="profile.personal_information.email" type="email" placeholder="Enter your email address" />
          </FormGroup>
          <FormGroup label="Avatar URL">
            <TextInput v-model="profile.personal_information.avatar_url" placeholder="Enter avatar image URL" />
          </FormGroup>
          <div v-if="saved" class="save-indicator saved">✓ Changes saved automatically</div>
          <div v-if="importSuccess" class="save-indicator saved">✓ Profile imported successfully</div>
        </Form>
      </div>
      <div class="entries-section">
        <!-- Replace Resume Entries section with WorkExperience component -->
        <WorkExperience :editable="true" />
        <!-- Education section remains unchanged -->
        <div class="education-section">
          <div class="education-header">
            <h2>Education History</h2>
            <Button @click="startAddEducationEntry" variant="primary">Add Education</Button>
          </div>
          <div v-if="showEducationForm" class="education-form">
            <TextInput v-model="newEducationEntry.degree" placeholder="Degree" />
            <TextInput v-model="newEducationEntry.institution" placeholder="Institution" />
            <DateInput v-model="newEducationEntry.start_date" placeholder="Start Date" />
            <DateInput v-model="newEducationEntry.end_date" placeholder="End Date" />
            <TextareaInput v-model="newEducationEntry.description" placeholder="Description" />
            <div class="form-action-row">
              <Button @click="saveEducationEntry" variant="primary" class="form-action-btn">{{ editingEducationId ? 'Update' : 'Add' }} Education</Button>
              <Button @click="cancelEducationEntry" variant="secondary" class="form-action-btn">Cancel</Button>
            </div>
          </div>
          <div v-for="entry in profile.education_experiences" :key="entry.id" class="education-entry">
            <strong>{{ entry.degree }}</strong> from <em>{{ entry.institution }}</em>
            <span>{{ entry.start_date }} - {{ entry.end_date }}</span>
            <p>{{ entry.description }}</p>
            <div class="education-entry-actions">
              <Button @click="startEditEducationEntry(entry)" variant="primary" size="small">Edit</Button>
              <Button @click="removeEducationEntry(entry.id)" variant="danger" size="small">Delete</Button>
            </div>
          </div>
          <div v-if="profile.education_experiences.length === 0" class="empty-education">
            <p>No education entries yet. Click "Add Education" to get started.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
