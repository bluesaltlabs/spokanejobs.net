<script setup>
import { onMounted, computed } from 'vue'
import { useJobs } from '@/stores/jobs'
import { SkeletonTableRow } from '@/components/skeleton'
import { Container } from '@/components/ui'
import { LocationIcon } from '@/components/icons'

// setup
const jobsStore = useJobs()

// Fetch jobs on component mount
onMounted(() => {
  jobsStore.fetchJobs()
})

const loading = computed(() => jobsStore.loading)

// Helper function to format company name
function formatCompanyName(companySlug) {
  if (!companySlug) return 'Unknown Company'
  
  // Convert slug to readable name
  return companySlug
    .split('_')
    .map(word => word.charAt(0).toUpperCase() + word.slice(1))
    .join(' ')
}
</script>

<template>
  <Container>
    <h1>Jobs</h1>
    <hr class="my-4" />

    <div v-if="loading" class="loading-state">
      <div class="loading-message">Loading jobs...</div>
      <table class="jobs-table">
        <thead>
          <tr>
            <th>Title</th>
            <th>Company</th>
            <th>Location</th>
          </tr>
        </thead>
        <tbody>
          <SkeletonTableRow
            v-for="i in 8"
            :key="`skeleton-${i}`"
            :columns="3"
            :column-widths="['40%', '30%', '30%']"
          />
        </tbody>
      </table>
    </div>

    <div v-else-if="jobsStore.sortedJobs.length > 0" class="jobs-results">
      <div class="results-count">{{ jobsStore.countJobs }} jobs found</div>
      
      <table class="jobs-table">
        <thead>
          <tr>
            <th>Title</th>
            <th>Company</th>
            <th>Location</th>
          </tr>
        </thead>
        <tbody>
          <tr 
            v-for="job in jobsStore.sortedJobs" 
            :key="job.id || job.job_id"
            class="job-row"
          >
            <td class="job-title-cell">
              <router-link 
                :to="{ name: 'job-detail', params: { id: job.id || job.job_id } }"
                class="job-title-link"
              >
                {{ job.title }}
              </router-link>
            </td>
            <td class="job-company-cell">
              <span class="company-name">{{ formatCompanyName(job.company) }}</span>
            </td>
            <td class="job-location-cell">
              <span class="location">
                <LocationIcon class="location-icon" />
                {{ job.city }}{{ job.state ? ', ' + job.state : '' }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-else class="empty-state">
      <div class="empty-icon">
        <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/>
        </svg>
      </div>
      <h3>No jobs found</h3>
      <p>Check back later for new opportunities in the Spokane area.</p>
    </div>
  </Container>
</template>

<style scoped>
hr {
  margin-top: 1rem;
  margin-bottom: 2rem;
}

.loading-state {
  margin-bottom: 1rem;
}

.loading-message {
  text-align: center;
  color: var(--color-text-muted);
  margin-bottom: 1rem;
  font-weight: 500;
}

.jobs-results {
  margin-bottom: 1rem;
}

.results-count {
  font-weight: 600;
  color: var(--color-heading);
  margin-bottom: 1rem;
}

.jobs-table {
  width: 100%;
  border-collapse: collapse;
  background: var(--color-surface);
  border-radius: var(--border-radius-medium);
  overflow: hidden;
  box-shadow: 0 1px 3px var(--color-shadow);
}

th, td {
  padding: 1rem 1.5rem;
  text-align: left;
  border-bottom: 1px solid var(--color-border);
}

th {
  background: var(--color-background-soft);
  font-weight: 600;
  color: var(--color-heading);
  font-size: 0.9rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

td {
  color: var(--color-text);
  vertical-align: top;
}

.job-row:hover {
  background: var(--color-surface-hover);
}

.job-title-cell {
  width: 40%;
}

.job-title-link {
  color: var(--color-link);
  text-decoration: none;
  font-weight: 600;
  font-size: 1rem;
  line-height: 1.4;
  transition: color 0.2s ease;
}

.job-title-link:hover {
  color: var(--color-link-hover);
  text-decoration: underline;
}

.job-company-cell {
  width: 30%;
}

.company-name {
  font-weight: 500;
  color: var(--color-text);
}

.job-location-cell {
  width: 30%;
}

.location {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--color-text-muted);
  font-size: 0.9rem;
}

.location-icon {
  width: 16px;
  height: 16px;
  color: var(--color-text-muted);
}

.empty-state {
  text-align: center;
  padding: 4rem 2rem;
  color: var(--color-text-muted);
}

.empty-icon {
  margin-bottom: 1rem;
  color: var(--color-text-subtle);
}

.empty-state h3 {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-heading);
  margin: 0 0 0.5rem 0;
}

.empty-state p {
  font-size: 1rem;
  margin: 0;
  line-height: 1.5;
}
</style>
