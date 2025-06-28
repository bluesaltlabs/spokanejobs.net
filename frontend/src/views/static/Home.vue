<script setup>
const appName = import.meta.env.VITE_APP_NAME
const appDescription = import.meta.env.VITE_APP_DESCRIPTION

import { CompaniesIcon, JobsIcon } from '@/components/icons';
import { Container } from '@/components/ui';
import { ItemContainer, ItemCard } from '@/components/common';
</script>

<template>
  <Container>
    <h1>Home</h1>
    <hr class="divider" />

    <p class="home-description">
      {{ appDescription ?? `${appName} description`}}
    </p>

    <ItemContainer grid customClass="home-nav">
      <ItemCard
        :onClick="() => $router.push('/companies')"
        customClass="home-card-link"
      >
        <template #header>
          <div class="home-card-content">
            <CompaniesIcon class="home-card-icon" />
            <span>View Companies</span>
          </div>
        </template>
      </ItemCard>
      <ItemCard
        :onClick="() => $router.push('/jobs')"
        customClass="home-card-link"
      >
        <template #header>
          <div class="home-card-content">
            <JobsIcon class="home-card-icon" />
            <span>View Jobs</span>
          </div>
        </template>
      </ItemCard>
    </ItemContainer>
  </Container>
</template>

<style scoped>
.divider {
  margin: 1.5rem 0 2rem 0;
  border: none;
  border-top: 1px solid var(--color-border);
}

.home-description {

  margin-bottom: 2em;
  color: var(--color-text-secondary);
  line-height: 1.6;
}

.home-nav {
  gap: 2rem !important;
}

.home-card-link {
  padding: 2rem 2.5rem !important;
  min-width: 200px;
  font-weight: 600;
  font-size: 1.1rem;
  text-align: center;
}

.home-card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.home-card-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: var(--color-primary-600);
}

@media (max-width: 600px) {
  .home-nav {
    flex-direction: column !important;
    gap: 1.25rem !important;
    width: 100%;
  }

  .home-card-link {
    width: 100%;
    min-width: unset;
    padding: 1.5rem 1rem !important;
  }
}
</style>
