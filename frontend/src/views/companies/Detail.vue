<script setup>
import { onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useCompany } from '@/stores/company'
import { SkeletonText, SkeletonImage, SkeletonButton } from '@/components/skeleton'
import { Container } from '@/components/ui'
import { CompaniesIcon, ViewIcon } from '@/components/icons'
import Button from '@/components/ui/Button/Button.vue'
import ScrapingStatus from '@/components/common/ScrapingStatus/ScrapingStatus.vue'

// Setup
const companyStore = useCompany()
const route = useRoute()

// Fetch company on component mount
onMounted(() => {
  companyStore.fetchCompany(route.params.slug)
})

// todo: update this page so that it loads from JSON and not supabase.
const company = computed(() => companyStore.company)
const loading = computed(() => companyStore.loading)
</script>

<template>
  <Container>
    <div>
      <!-- todo: back to companies grid header -->
      <div class="back-button">
        <router-link class="back-button-link" to="/companies">
          <CompaniesIcon class="back-icon" />
          All Companies
        </router-link>
      </div>
    </div>

    <div v-if="loading">
      <!-- Loading skeleton -->
      <div class="skeleton-content">
        <SkeletonText variant="title" :lines="1" />
        <div class="skeleton-links">
          <SkeletonButton width="100px" height="35px" />
          <SkeletonButton width="80px" height="35px" />
        </div>
        <div class="spacer"></div>
        <SkeletonImage width="425px" height="425px" />
        <div class="spacer"></div>
        <SkeletonText :lines="4" :line-height="20" />
        <SkeletonText :lines="3" :line-height="20" last-line-width="60%" />
      </div>
    </div>

    <div v-else-if="company">
      <h1 class="company-name">{{ company.name }}</h1>
      <div class="company-links-container">
        <span v-if="company.website">
          <Button
            variant="info"
            class="company-link-btn"
            as="a"
            :href="company.website"
            target="_blank"
            :aria-label="company.website"
          >
            <ViewIcon class="link-icon" />
            Website
          </Button>
        </span>
        <span v-if="company.job_board_url">
          <Button
            variant="info"
            class="company-link-btn"
            as="a"
            :href="company.job_board_url"
            target="_blank"
            :aria-label="company.job_board_url"
          >
            <ViewIcon class="link-icon" />
            Jobs
          </Button>
        </span>
      </div>

      <div class="spacer"></div>
      <img v-if="company?.logo_url" :src="company.logo_url" :alt="company.name + ' Logo'" class="company-logo" />

      <div class="spacer"></div>
      <ScrapingStatus
        v-if="company.scraper_status"
        :status="company.scraper_status"
        :last-scraped-at="company.last_scraped_at"
        :scraper-created-at="company.scraper_created_at"
        :active-jobs-count="company.active_jobs_count"
        :scraper-error="company.scraper_error"
      />

      <div class="spacer"></div>
      <p v-if="company.description">{{ company.description }}</p>
    </div>

    <div v-else>
      <p>Company not found.</p>
    </div>
  </Container>
</template>

<style scoped>
.back-button {
  margin-bottom: 2rem;
}

.back-button-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--color-text);
  text-decoration: none;
  font-weight: bold;
  font-size: 1.2rem;
}

.back-button-link:hover {
  color: var(--color-primary);
}

.back-icon {
  width: 20px;
  height: 20px;
}

.company-name {
  margin-top: 25px;
  margin-bottom: 15px;
  font-size: 3.5rem;
}

img.company-logo {
  max-height: 425px;
  max-width: 425px;
  width: 100%;
  height: auto;
  /* image rounding, padding, box-shadow */
  border-radius: var(--border-radius);
  padding: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.spacer {
  height: 30px;
}

.skeleton-content {
  margin-top: 25px;
}

.skeleton-links {
  margin-top: 15px;
  margin-bottom: 15px;
  display: flex;
  gap: 8px;
}

.company-links-container {
  margin-top: 15px;
  margin-bottom: 15px;
}

.company-links-container a {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  border-radius: var(--border-radius-small);
  padding: 5px 10px;
  margin: 5px 8px;
  background: var(--color-primary-700);
  color: var(--color-text);
  text-decoration: none;
  font-weight: bold;
  font-size: 1.2rem;
}

.company-links-container a:hover {
  background: var(--color-primary-500);
  color: var(--color-white);
}

.company-link-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  border-radius: var(--border-radius-small);
  padding: 5px 10px;
  margin: 5px 8px;
  font-size: 1.2rem;
  text-decoration: none;
}

.link-icon {
  width: 16px;
  height: 16px;
}
</style>
