<script setup>
import { onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useJob } from '@/stores/job'
import { useCompanies } from '@/stores/companies'
import { SkeletonText, SkeletonButton } from '@/components/skeleton'
import { Container } from '@/components/ui'
import { JobsIcon, CompaniesIcon, LocationIcon } from '@/components/icons'

// Setup
const jobStore = useJob()
const companiesStore = useCompanies()
const route = useRoute()

// Fetch job and companies on component mount
onMounted(() => {
  jobStore.fetchJob(route.params.id)
  companiesStore.fetchCompanies()
})

const job = computed(() => jobStore.job)
const loading = computed(() => jobStore.loading)

const companyMap = computed(() => {
  const map = {}
  for (const c of companiesStore.companies) {
    map[c.slug] = c
  }
  return map
})

// Helper function to format job description
function formatDescription(description) {
  if (!description) return ''

  // Convert markdown-style formatting
  return description
    .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
    .replace(/\*(.*?)\*/g, '<em>$1</em>')
    .replace(/\n/g, '<br>')
}

function getCompanyName(slug) {
  return companyMap.value[slug]?.name || 'Unknown Company'
}
</script>

<template>
  <Container>
    <!-- Back Navigation -->
    <div class="back-button">
      <router-link class="back-button-link" to="/jobs">
        <JobsIcon class="back-icon" />
        All Jobs
      </router-link>
    </div>

    <div v-if="loading">
      <!-- Loading skeleton -->
      <div class="skeleton-content">
        <SkeletonText variant="title" :lines="1" />
        <div class="skeleton-meta">
          <SkeletonText :lines="1" />
        </div>
        <div class="spacer"></div>
        <SkeletonText :lines="6" :line-height="20" />
        <SkeletonText :lines="4" :line-height="20" />
      </div>
    </div>

    <div v-else-if="job">
      <!-- Job Header -->
      <h1 class="job-title">{{ job.title }}</h1>

      <!-- Job Meta -->
      <div class="job-meta">
        <span v-if="job.city" class="job-location">
          <LocationIcon class="meta-icon" />
          {{ job.city }}{{ job.state ? ', ' + job.state : '' }}
        </span>
        <span v-if="job.company" class="job-company">
          <CompaniesIcon class="meta-icon" />
          <router-link
            v-if="companyMap[job.company]"
            :to="{ name: 'company-detail', params: { slug: job.company } }"
            class="company-name"
          >
            {{ getCompanyName(job.company) }}
          </router-link>
          <span v-else class="company-name">{{ getCompanyName(job.company) }}</span>
        </span>
      </div>

      <!-- Apply Button -->
      <div class="job-actions">
        <a
          v-if="job.url"
          :href="job.url"
          target="_blank"
          rel="noopener noreferrer"
          class="apply-button"
        >
          Apply for this job
        </a>
      </div>

      <div class="spacer"></div>

      <!-- Job Description -->
      <div v-if="job.description" class="job-description">
        <h3>Job Description</h3>
        <div class="description-content" v-html="formatDescription(job.description)"></div>
      </div>

      <!-- Job Metadata -->
      <div class="job-metadata">
        <div class="metadata-item">
          <strong>Job ID:</strong> {{ job.job_id || job.id }}
        </div>
        <div v-if="job.url" class="metadata-item">
          <strong>Original Posting:</strong>
          <a :href="job.url" target="_blank" rel="noopener noreferrer">
            View on {{ getCompanyName(job.company) }} Careers
          </a>
        </div>
      </div>
    </div>

    <div v-else>
      <h2>Job Not Found</h2>
      <p>The job you're looking for could not be found.</p>
      <router-link to="/jobs">Back to All Jobs</router-link>
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

.job-title {
  margin-top: 25px;
  margin-bottom: 15px;
  font-size: 3rem;
  font-weight: 700;
  color: var(--color-heading);
}

.job-meta {
  display: flex;
  gap: 2rem;
  margin-bottom: 1.5rem;
  font-size: 1.1rem;
  color: var(--color-text-muted);
}

.job-location,
.job-company {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.meta-icon {
  width: 18px;
  height: 18px;
  color: var(--color-text-muted);
}

.job-actions {
  margin-bottom: 2rem;
}

.apply-button {
  display: inline-block;
  background: var(--color-primary-600);
  color: white;
  padding: 0.75rem 1.5rem;
  border-radius: var(--border-radius-medium);
  text-decoration: none;
  font-weight: 600;
  font-size: 1.1rem;
  transition: background-color 0.2s ease;
}

.apply-button:hover {
  background: var(--color-primary-700);
}

.spacer {
  height: 30px;
}

.job-description {
  margin-bottom: 2rem;
}

.job-description h3 {
  font-size: 1.5rem;
  font-weight: 600;
  color: var(--color-heading);
  margin-bottom: 1rem;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid var(--color-border);
}

.description-content {
  line-height: 1.6;
  color: var(--color-text);
}

.description-content :deep(strong) {
  font-weight: 600;
  color: var(--color-heading);
}

.description-content :deep(em) {
  font-style: italic;
}

.job-metadata {
  padding: 1.5rem;
  background: var(--color-background-soft);
  border-radius: var(--border-radius-medium);
  border: 1px solid var(--color-border);
}

.metadata-item {
  margin-bottom: 0.75rem;
  font-size: 0.9rem;
  color: var(--color-text-muted);
}

.metadata-item:last-child {
  margin-bottom: 0;
}

.metadata-item a {
  color: var(--color-primary);
  text-decoration: none;
}

.metadata-item a:hover {
  text-decoration: underline;
}

.skeleton-content {
  margin-top: 1rem;
}

.skeleton-meta {
  margin-bottom: 1rem;
}
</style>
