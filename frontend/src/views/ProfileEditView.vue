<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useProfileStore } from '@/stores/profile'

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

function startAddEntry() {
  editingId.value = null
  resetNewEntry()
  showEntryForm.value = true
}

function startEditEntry(entry) {
  editingId.value = entry.id
  newEntry.value = { ...entry }
  showEntryForm.value = true
}

async function saveEntry() {
  if (!newEntry.value.jobTitle || !newEntry.value.company) return
  if (editingId.value) {
    await profile.editResumeEntry(editingId.value, { ...newEntry.value })
  } else {
    await profile.addResumeEntry({ ...newEntry.value, id: Date.now().toString() })
  }
  await profile.loadProfile()
  editingId.value = null
  resetNewEntry()
  showEntryForm.value = false
}

function cancelEntry() {
  resetNewEntry()
  editingId.value = null
  showEntryForm.value = false
}

async function removeEntry(id) {
  await profile.removeResumeEntry(id)
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

  .resume-section {
    flex: 1;
    margin-top: 0;
    padding-top: 0;
    border-top: none;
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

.btn-primary {
  background: linear-gradient(90deg, #4f8cff 0%, #2356c7 100%);
  color: #fff;
  border: none;
  border-radius: var(--border-radius-small);
  padding: 0.5rem 1.2rem;
  font-weight: 600;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.2s, box-shadow 0.2s;
  box-shadow: 0 1px 4px rgba(79,140,255,0.08);
  margin-right: 0.5rem;
}

.btn-primary:hover, .btn-primary:focus {
  background: linear-gradient(90deg, #2356c7 0%, #4f8cff 100%);
  outline: none;
}

.btn-secondary {
  background: #6c757d;
}

.btn-secondary:hover {
  background: #5a6268;
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
  border: 1.5px solid var(--color-border, #d1d5db);
  border-radius: 8px;
  padding: 1.25rem 1rem 1rem 1rem;
  margin-bottom: 1.25rem;
  background: var(--color-surface, #fff);
  box-shadow: 0 1px 4px rgba(79,140,255,0.06);
  transition: background 0.2s, border 0.2s;
}

[data-theme="dark"] .resume-entry {
  background: var(--color-surface-dark, #23272f);
  border-color: var(--color-border-dark, #3a3f4b);
  box-shadow: 0 1px 6px rgba(79,140,255,0.10);
}

.resume-entry strong {
  color: var(--resume-title-color, #2356c7);
  font-size: 1.1rem;
  font-weight: 700;
}

.resume-entry em {
  color: var(--resume-company-color, #4f8cff);
  font-style: normal;
  font-weight: 600;
}

.resume-entry span {
  color: var(--resume-date-color, #6c757d);
  font-size: 0.95rem;
  display: block;
  margin-bottom: 0.5rem;
}

.resume-entry p {
  color: var(--resume-desc-color, #222);
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

.btn-danger {
  background: #e74c3c;
}

.btn-danger:hover {
  background: #c0392b;
}
</style>

<template>
  <div class="profile-edit-view">
    <div class="profile-header">
      <h1>Edit Profile</h1>
      <button @click="goToView" class="btn btn-secondary">View Profile</button>
    </div>

    <div class="profile-edit-content">
      <!-- Left Column - User Info -->
      <div class="user-info-section">
        <h2>Personal Information</h2>
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
            <div class="avatar-preview">
              <div class="avatar">
                <img v-if="profile.avatar" :src="profile.avatar" alt="Avatar preview" />
                <span v-else>{{ profile.first_name?.charAt(0) || 'U' }}</span>
              </div>
            </div>
          </div>

          <div v-if="saved" class="save-indicator saved">
            ✓ Changes saved automatically
          </div>
        </form>
      </div>

      <!-- Right Column - Resume Entries -->
      <div class="resume-section">
        <div class="resume-header">
          <h2>Resume Entries</h2>
          <button @click="startAddEntry" class="btn btn-primary">Add Entry</button>
        </div>

        <div v-if="showEntryForm" class="resume-form">
          <input v-model="newEntry.jobTitle" placeholder="Job Title" class="themed-input" />
          <input v-model="newEntry.company" placeholder="Company" class="themed-input" />
          <input v-model="newEntry.startDate" placeholder="Start Date" type="date" class="themed-input" />
          <input v-model="newEntry.endDate" placeholder="End Date" type="date" class="themed-input" />
          <textarea v-model="newEntry.description" placeholder="Description" class="themed-input"></textarea>
          <button @click="saveEntry" class="btn btn-primary">{{ editingId ? 'Update' : 'Add' }} Entry</button>
          <button @click="cancelEntry" type="button" class="btn btn-secondary">Cancel</button>
        </div>

        <div v-for="entry in profile.resumeEntries" :key="entry.id" class="resume-entry">
          <strong>{{ entry.jobTitle }}</strong> at <em>{{ entry.company }}</em>
          <span>{{ entry.startDate }} - {{ entry.endDate }}</span>
          <p>{{ entry.description }}</p>
          <div class="resume-entry-actions">
            <button @click="startEditEntry(entry)" class="btn btn-primary btn-sm">Edit</button>
            <button @click="removeEntry(entry.id)" class="btn btn-danger btn-sm">Delete</button>
          </div>
        </div>

        <div v-if="profile.resumeEntries.length === 0" class="empty-resume">
          <p>No resume entries yet. Click "Add Entry" to get started.</p>
        </div>
      </div>
    </div>
  </div>
</template>
