<script setup>
import { onMounted, computed } from 'vue'
import { useJobs } from '@/stores/jobs'
import { SkeletonTableRow } from '@/components/skeleton'
import { Container } from '@/components/ui'
import { LocationIcon } from '@/components/icons'

const jobsStore = useJobs()

onMounted(() => {
  jobsStore.fetchJobs()
})

const loading = computed(() => jobsStore.loading)

function formatCompanyName(companySlug) {
  if (!companySlug) return 'Unknown Company'
  return companySlug
    .split('_')
    .map(word => word.charAt(0).toUpperCase() + word.slice(1))
    .join(' ')
}
</script>

<template>
  <Container>
    <h1>Jobs</h1>
    <hr class="divider" />

    <div v-if="loading" class="loading-state">
      <div class="loading-message">Loading jobs...</div>
      <table class="jobs-table desktop-only">
        <thead>
          <tr>
            <th>Title</th>
            <th>Company</th>
            <th>Location</th>
          </tr>
        </thead>
        <tbody>
          <SkeletonTableRow v-for="i in 8" :key="`skeleton-${i}`" :columns="3" :column-widths="['40%', '30%', '30%']" />
        </tbody>
      </table>
      <div class="jobs-grid mobile-only">
        <div v-for="i in 6" :key="`skeleton-card-${i}`" class="job-card skeleton-card">
          <div class="job-header">
            <div class="skeleton-title"></div>
            <div class="skeleton-company"></div>
          </div>
          <div class="job-meta">
            <div class="skeleton-location"></div>
          </div>
        </div>
      </div>
    </div>

    <div v-else-if="jobsStore.sortedJobs.length > 0" class="jobs-results">
      <div class="results-count">{{ jobsStore.countJobs }} jobs found</div>
      <table class="jobs-table desktop-only">
        <thead>
          <tr>
            <th>Title</th>
            <th>Company</th>
            <th>Location</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="job in jobsStore.sortedJobs" :key="job.id || job.job_id" class="job-row">
            <td class="job-title-cell">
              <router-link :to="{ name: 'job-detail', params: { id: job.id || job.job_id } }" class="job-title-link">
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
      <div class="jobs-grid mobile-only">
        <div v-for="job in jobsStore.sortedJobs" :key="job.id || job.job_id" class="job-card">
          <router-link :to="{ name: 'job-detail', params: { id: job.id || job.job_id } }" class="job-link">
            <div class="job-header">
              <h3 class="job-title">{{ job.title }}</h3>
              <div class="job-company">
                <span class="company-name">{{ formatCompanyName(job.company) }}</span>
              </div>
            </div>
            <div class="job-meta">
              <div class="job-location">
                <LocationIcon class="location-icon" />
                {{ job.city }}{{ job.state ? ', ' + job.state : '' }}
              </div>
            </div>
            <div class="job-actions">
              <span class="view-details">View Details →</span>
            </div>
          </router-link>
        </div>
      </div>
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
.divider {
  margin: 1.5rem 0 2rem 0;
  border: none;
  border-top: 1px solid var(--color-border);
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

.job-title-cell { width: 40%; }
.job-company-cell { width: 30%; }
.job-location-cell { width: 30%; }

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

.company-name {
  font-weight: 500;
  color: var(--color-text);
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

.jobs-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1rem;
  margin-bottom: 2rem;
  width: 100%;
}

.job-card {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-medium);
  overflow: hidden;
  transition: all 0.2s ease;
  box-shadow: 0 1px 3px var(--color-shadow);
  margin-bottom: 0;
  width: 100%;
  min-height: 120px;
}

.job-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px var(--color-shadow-elevated);
  border-color: var(--color-border-hover);
}

.job-link {
  display: block;
  text-decoration: none;
  color: inherit;
  padding: 1rem;
  height: 100%;
  box-sizing: border-box;
}

.job-header { margin-bottom: 0.75rem; }
.job-title {
  font-size: 1rem;
  font-weight: 600;
  color: var(--color-heading);
  margin: 0 0 0.25rem 0;
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.job-company { margin-bottom: 0.5rem; }
.job-meta { margin-bottom: 0.75rem; }
.job-location {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--color-text-muted);
  font-size: 0.875rem;
}
.job-actions {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  padding-top: 0.75rem;
  border-top: 1px solid var(--color-border);
}
.view-details {
  font-size: 0.875rem;
  color: var(--color-primary-600);
  font-weight: 500;
  transition: color 0.2s ease;
}
.job-card:hover .view-details { color: var(--color-primary-700); }
.skeleton-card { pointer-events: none; }
.skeleton-title {
  height: 1.25rem;
  background: var(--color-background-soft);
  border-radius: var(--border-radius-small);
  margin-bottom: 0.5rem;
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
.skeleton-company {
  height: 1rem;
  background: var(--color-background-soft);
  border-radius: var(--border-radius-small);
  margin-bottom: 0.5rem;
  width: 60%;
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
.skeleton-location {
  height: 1rem;
  background: var(--color-background-soft);
  border-radius: var(--border-radius-small);
  width: 40%;
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
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
@media (max-width: 768px) {
  .desktop-only { display: none !important; }
  .mobile-only { display: block !important; }
  .jobs-results .jobs-grid.mobile-only {
    display: grid !important;
    grid-template-columns: 1fr;
    gap: 1rem;
    margin-bottom: 2rem;
    width: 100%;
  }
  .jobs-results .job-card {
    margin-bottom: 0;
    width: 100%;
    display: block;
  }
  .job-link { padding: 0.875rem; }
  .job-title { font-size: 0.95rem; }
}
@media (min-width: 769px) {
  .desktop-only { display: table; }
  .mobile-only { display: none; }
}
@media (min-width: 769px) and (max-width: 1024px) {
  .jobs-table { font-size: 0.9rem; }
  th, td { padding: 0.875rem 1.25rem; }
  .jobs-grid { gap: 1.25rem; }
}
@media (min-width: 1025px) {
  .jobs-grid { gap: 1.75rem; }
}
</style>
