<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useProfileStore } from '@/stores/profile'
import { Button } from '@/components/ui'
import { ProfileIcon } from '@/components/icons'

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
  max-width: 1200px;
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

.profile-content {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

/* Desktop 2-column layout */
@media (min-width: 768px) {
  .profile-content {
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

.profile-avatar {
  text-align: center;
  margin-bottom: 1.5rem;
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

.profile-field {
  margin-bottom: 1.5rem;
}

.profile-field label {
  font-weight: 600;
  color: var(--color-text-subtle);
  font-size: 0.9rem;
  margin-bottom: 0.5rem;
  display: block;
}

.profile-field .value {
  font-size: 1.1rem;
  color: var(--color-text);
  padding: 0.75rem;
  background: var(--color-surface-hover);
  border-radius: var(--border-radius-small);
  border: 1px solid var(--color-border);
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

/* Dark theme overrides for resume colors */
[data-theme="dark"] .resume-entry strong {
  color: var(--color-primary-400);
}

[data-theme="dark"] .resume-entry em {
  color: var(--color-primary-300);
}

[data-theme="dark"] .resume-entry span {
  color: var(--color-text-muted);
}

[data-theme="dark"] .resume-entry p {
  color: var(--color-text);
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

.empty-resume .base-button {
  margin-left: 0.5rem;
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

/* Dark theme overrides for education colors */
[data-theme="dark"] .education-entry strong {
  color: var(--color-primary-400);
}

[data-theme="dark"] .education-entry em {
  color: var(--color-primary-300);
}

[data-theme="dark"] .education-entry span {
  color: var(--color-text-muted);
}

[data-theme="dark"] .education-entry p {
  color: var(--color-text);
}

.empty-education {
  text-align: center;
  padding: 2rem;
  color: var(--color-text-subtle);
  background: var(--color-surface-hover);
  border-radius: var(--border-radius-medium);
  border: 1px dashed var(--color-border);
}

.empty-education p {
  margin: 0;
  font-size: 1rem;
}

.empty-education .base-button {
  margin-left: 0.5rem;
}

.profile-actions {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 1rem;
}
</style>

<template>
  <div class="profile-view">
    <div class="profile-header">
      <h1>Profile</h1>
    </div>
    <div class="profile-actions">
      <Button @click="goToEdit" variant="primary">
        <ProfileIcon style="margin-right: 0.5em; vertical-align: middle;" />
        Edit Profile
      </Button>
    </div>

    <div class="profile-content">

      <!-- Left Column - User Info -->
      <div class="user-info-section">
        <h2>Personal Information</h2>

        <div class="profile-avatar">
          <div class="avatar">
            <img v-if="profile.personal_information.avatar_url" :src="profile.personal_information.avatar_url" alt="Avatar" />
            <span v-else>{{ profile.personal_information.first_name?.charAt(0) || 'U' }}</span>
          </div>
        </div>

        <div class="profile-field">
          <label>First Name</label>
          <div class="value">{{ profile.personal_information.first_name || 'Not provided' }}</div>
        </div>

        <div class="profile-field">
          <label>Last Name</label>
          <div class="value">{{ profile.personal_information.last_name || 'Not provided' }}</div>
        </div>

        <div class="profile-field">
          <label>Email</label>
          <div class="value">{{ profile.personal_information.email || 'Not provided' }}</div>
        </div>
      </div>

      <!-- Right Column - Resume and Education Entries -->
      <div class="entries-section">

        <div class="resume-section">
          <div class="resume-header">
            <h2>Work Experience</h2>
          </div>

          <div v-for="entry in profile.work_experiences" :key="entry.id" class="resume-entry">
            <strong>{{ entry.job_title_start }}</strong> at <em>{{ entry.company }}</em>
            <span>{{ entry.start_date }} - {{ entry.end_date }}</span>
            <p>{{ entry.description }}</p>
          </div>

          <div v-if="profile.work_experiences.length === 0" class="empty-resume">
            <p>No work experience. </p>
          </div>
        </div>

        <div class="education-section">
          <div class="education-header">
            <h2>Education Experience</h2>
          </div>

          <div v-for="entry in profile.education_experiences" :key="entry.id" class="education-entry">
            <strong>{{ entry.degree }}</strong> from <em>{{ entry.institution }}</em>
            <span>{{ entry.start_date }} - {{ entry.end_date }}</span>
            <p>{{ entry.description }}</p>
          </div>

          <div v-if="profile.education_experiences.length === 0" class="empty-education">
            <p>No education experience. </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
