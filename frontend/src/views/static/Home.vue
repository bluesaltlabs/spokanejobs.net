<script setup>
const appName = import.meta.env.VITE_APP_NAME
const appDescription = import.meta.env.VITE_APP_DESCRIPTION

import { CompaniesIcon, JobsIcon } from '@/components/icons';
</script>

<template>
  <main class="static home landing-container">
    <h1 class="landing-title">{{ appName }}</h1>
    <p class="landing-description">
      {{ appDescription ?? `${appName} description`}}
    </p>
    <nav class="landing-nav">
      <router-link to="/companies" class="landing-card-link">
        <CompaniesIcon class="landing-card-icon" />
        <span>View Companies</span>
      </router-link>
      <router-link to="/jobs" class="landing-card-link">
        <JobsIcon class="landing-card-icon" />
        <span>View Jobs</span>
      </router-link>
    </nav>
  </main>
</template>

<style scoped>
.landing-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 60vh;
}
.landing-title {
  font-size: 2.5rem;
  margin-bottom: 0.5em;
}
.landing-description {
  max-width: 500px;
  text-align: center;
  margin-bottom: 2em;
}
.landing-nav {
  display: flex;
  gap: 2rem;
}
.landing-card-link {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-medium);
  box-shadow: 0 1px 3px var(--color-shadow);
  padding: 2rem 2.5rem;
  min-width: 200px;
  text-decoration: none;
  color: inherit;
  transition: all 0.2s;
  font-weight: 600;
  font-size: 1.1rem;
}
.landing-card-link:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px var(--color-shadow-elevated);
  border-color: var(--color-border-hover);
}
.landing-card-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: var(--color-primary-600);
}
@media (max-width: 600px) {
  .landing-nav {
    flex-direction: column;
    gap: 1.25rem;
    width: 100%;
  }
  .landing-card-link {
    width: 100%;
    min-width: unset;
    padding: 1.5rem 1rem;
  }
}
</style>
