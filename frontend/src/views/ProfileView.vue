<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useProfileStore } from '@/stores/profile'

const router = useRouter()
const profile = useProfileStore()

onMounted(() => {
  profile.loadProfile()
})

function goToEdit() {
  router.push('/profile/edit')
}
</script>

<style scoped>
.profile-view {
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  background: var(--background-color);
  border-radius: 8px;
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

.profile-info {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 1rem;
  margin-bottom: 2rem;
}

.profile-avatar {
  text-align: center;
}

.avatar {
  width: 120px;
  height: 120px;
  background: var(--color-surface-hover);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-subtle);
  border-radius: var(--border-radius-medium);
  margin: 0 auto 1rem;

  img {
    width: 120px;
    height: 120px;
    border-radius: var(--border-radius-medium);
    object-fit: cover;
  }
}

.profile-details {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.profile-field {
  margin-bottom: 0.75rem;
}

.profile-field label {
  font-weight: 600;
  color: var(--color-text-subtle);
  font-size: 0.9rem;
  margin-bottom: 0.25rem;
  display: block;
}

.profile-field .value {
  font-size: 1.1rem;
  color: var(--color-text);
}

.resume-section {
  margin-top: 2rem;
}

.resume-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.resume-entry {
  border: 1.5px solid var(--color-border, #d1d5db);
  border-radius: var(--border-radius-medium);
  padding: 1.25rem 1rem 1rem 1rem;
  margin-bottom: 1.25rem;
  background: var(--color-surface, #fff);
  box-shadow: 0 1px 4px rgba(79,140,255,0.06);
  transition: background 0.2s, border 0.2s;
}

:root {
  --resume-title-color: #2356c7;
  --resume-company-color: #4f8cff;
  --resume-date-color: #6c757d;
  --resume-desc-color: #222;
}

[data-theme="dark"] .resume-entry {
  background: var(--color-surface-dark, #23272f);
  border-color: var(--color-border-dark, #3a3f4b);
  box-shadow: 0 1px 6px rgba(79,140,255,0.10);
}

[data-theme="dark"] {
  --resume-title-color: #7daaff;
  --resume-company-color: #a3c9ff;
  --resume-date-color: #b3b8c5;
  --resume-desc-color: #e0e6f1;
}

.resume-entry strong {
  color: var(--resume-title-color);
  font-size: 1.1rem;
  font-weight: 700;
}
.resume-entry em {
  color: var(--resume-company-color);
  font-style: normal;
  font-weight: 600;
}
.resume-entry span {
  color: var(--resume-date-color);
  font-size: 0.95rem;
  display: block;
  margin-bottom: 0.5rem;
}
.resume-entry p {
  color: var(--resume-desc-color);
  margin: 0.5rem 0 0 0;
  font-size: 1rem;
}

.themed-btn {
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

.themed-btn:hover, .themed-btn:focus {
  background: linear-gradient(90deg, #2356c7 0%, #4f8cff 100%);
  outline: none;
}

.empty-resume {
  text-align: center;
  padding: 2rem;
  color: var(--color-text-subtle);
  background: var(--color-surface-hover);
  border-radius: var(--border-radius-medium);
  border: 1px dashed var(--color-border);
}

.empty-resume p {
  margin: 0;
  font-size: 1rem;
}

.empty-resume .themed-btn {
  margin-left: 0.5rem;
  padding: 0.25rem 0.75rem;
  font-size: 0.9rem;
}
</style>

<template>
  <div class="profile-view">
    <div class="profile-header">
      <h1>Profile</h1>
      <button @click="goToEdit" class="btn btn-primary">Edit Profile</button>
    </div>

    <div class="profile-info">
      <div class="profile-avatar">
        <div class="avatar">
          <img v-if="profile.avatar" :src="profile.avatar" alt="Avatar" />
          <span v-else>{{ profile.first_name?.charAt(0) || 'U' }}</span>
        </div>
      </div>

      <div class="profile-details">
        <div class="profile-field">
          <label>Name</label>
          <div class="value">{{ profile.first_name }} {{ profile.last_name }}</div>
        </div>
        <div class="profile-field">
          <label>Email</label>
          <div class="value">{{ profile.email }}</div>
        </div>
      </div>
    </div>

    <div class="resume-section">
      <div class="resume-header">
        <h2>Resume Entries</h2>
      </div>

      <div v-for="entry in profile.resumeEntries" :key="entry.id" class="resume-entry">
        <strong>{{ entry.jobTitle }}</strong> at <em>{{ entry.company }}</em>
        <span>{{ entry.startDate }} - {{ entry.endDate }}</span>
        <p>{{ entry.description }}</p>
      </div>

      <div v-if="profile.resumeEntries.length === 0" class="empty-resume">
        <p>No resume entries yet. <button @click="goToEdit" class="btn btn-primary">Add your first entry</button></p>
      </div>
    </div>
  </div>
</template>
