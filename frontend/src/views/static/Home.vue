<script setup>
const appName = import.meta.env.VITE_APP_NAME
const appDescription = import.meta.env.VITE_APP_DESCRIPTION

import { CompaniesIcon, JobsIcon } from '@/components/icons';
import { Container } from '@/components/ui';
</script>

<template>
  <Container>
    <h1>Home</h1>
    <hr class="divider" />

    <p class="home-description">
      {{ appDescription ?? `${appName} description`}}
    </p>

    <nav class="home-nav">
      <router-link to="/companies" class="home-card-link">
        <CompaniesIcon class="home-card-icon" />
        <span>View Companies</span>
      </router-link>
      <router-link to="/jobs" class="home-card-link">
        <JobsIcon class="home-card-icon" />
        <span>View Jobs</span>
      </router-link>
    </nav>
  </Container>
</template>

<style scoped>
.divider {
  margin: 1.5rem 0 2rem 0;
  border: none;
  border-top: 1px solid var(--color-border);
}

.home-description {
  max-width: 500px;
  margin-bottom: 2em;
  color: var(--color-text-secondary);
  line-height: 1.6;
}

.home-nav {
  display: flex;
  gap: 2rem;
}

.home-card-link {
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

.home-card-link:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px var(--color-shadow-elevated);
  border-color: var(--color-border-hover);
}

.home-card-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: var(--color-primary-600);
}

@media (max-width: 600px) {
  .home-nav {
    flex-direction: column;
    gap: 1.25rem;
    width: 100%;
  }
  
  .home-card-link {
    width: 100%;
    min-width: unset;
    padding: 1.5rem 1rem;
  }
}
</style>
