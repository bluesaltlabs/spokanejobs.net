<script setup>
import { onMounted, computed } from 'vue'
import { useJobs } from '@/stores/jobs'
import { useCompanies } from '@/stores/companies'
import { SkeletonTableRow } from '@/components/skeleton'
import { Container } from '@/components/ui'
import { LocationIcon } from '@/components/icons'
import { ItemContainer, ItemCard, ItemTableRow, ItemEmptyState, SearchBar } from '@/components/common'

const jobsStore = useJobs()
const companiesStore = useCompanies()

onMounted(() => {
  jobsStore.fetchJobs()
  companiesStore.fetchCompanies()
})

const loading = computed(() => jobsStore.loading)

const companyMap = computed(() => {
  const map = {}
  for (const c of companiesStore.companies) {
    map[c.slug] = c
  }
  return map
})

function getCompanyName(slug) {
  return companyMap.value[slug]?.name || 'Unknown Company'
}
</script>

<template>
  <Container>
    <h1>Jobs</h1>
    <hr class="divider" />
    <SearchBar
      v-model="jobsStore.filters.search"
      placeholder="Search jobs..."
      @update:modelValue="val => jobsStore.updateFilters({ search: val })"
    />

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
      <ItemContainer grid customClass="mobile-only">
        <ItemCard v-for="i in 6" :key="`skeleton-card-${i}`" customClass="skeleton-card">
          <template #header>
            <div class="job-header">
              <div class="skeleton-title"></div>
              <div class="skeleton-company"></div>
            </div>
          </template>
          <template #meta>
            <div class="job-meta">
              <div class="skeleton-location"></div>
            </div>
          </template>
        </ItemCard>
      </ItemContainer>
    </div>

    <div v-else-if="jobsStore.sortedJobs.length > 0" class="jobs-results">
      <div class="results-count">{{ jobsStore.sortedJobs.length }} jobs found</div>
      <table class="jobs-table desktop-only">
        <thead>
          <tr>
            <th>Title</th>
            <th>Company</th>
            <th>Location</th>
          </tr>
        </thead>
        <tbody>
          <ItemTableRow v-for="job in jobsStore.sortedJobs" :key="job.id || job.job_id">
            <template #cells>
              <td class="job-title-cell">
                <router-link :to="{ name: 'job-detail', params: { id: job.id || job.job_id } }" class="job-title-link">
                  {{ job.title }}
                </router-link>
              </td>
              <td class="job-company-cell">
                <router-link
                  v-if="companyMap[job.company]"
                  :to="{ name: 'company-detail', params: { slug: job.company } }"
                  class="company-name"
                >
                  {{ getCompanyName(job.company) }}
                </router-link>
                <span v-else class="company-name">{{ getCompanyName(job.company) }}</span>
              </td>
              <td class="job-location-cell">
                <span class="location">
                  <LocationIcon class="location-icon" />
                  {{ job.city }}{{ job.state ? ', ' + job.state : '' }}
                </span>
              </td>
            </template>
          </ItemTableRow>
        </tbody>
      </table>
      <ItemContainer grid customClass="mobile-only">
        <ItemCard
          v-for="job in jobsStore.sortedJobs"
          :key="job.id || job.job_id"
          :onClick="() => $router.push({ name: 'job-detail', params: { id: job.id || job.job_id } })"
        >
          <template #header>
            <div class="job-header">
              <div class="job-info">
                <h3 class="job-title">{{ job.title }}</h3>
                <div class="job-company">
                  <router-link
                    v-if="companyMap[job.company]"
                    :to="{ name: 'company-detail', params: { slug: job.company } }"
                    class="company-name"
                  >
                    {{ getCompanyName(job.company) }}
                  </router-link>
                  <span v-else class="company-name">{{ getCompanyName(job.company) }}</span>
                </div>
              </div>
            </div>
          </template>
          <template #meta>
            <div class="job-meta">
              <div class="job-location">
                <LocationIcon class="location-icon" />
                {{ job.city }}{{ job.state ? ', ' + job.state : '' }}
              </div>
            </div>
          </template>
          <template #actions>
            <div class="job-actions">
              <span class="view-details">View Details →</span>
            </div>
          </template>
        </ItemCard>
      </ItemContainer>
    </div>

    <ItemEmptyState v-else>
      <template #icon>
        <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/>
        </svg>
      </template>
      <template #title>No jobs found</template>
      <template #desc>Check back later for new opportunities in the Spokane area.</template>
    </ItemEmptyState>
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

.job-header {
  margin-bottom: 0.75rem;
}

.job-info {
  flex: 1;
  min-width: 0;
}

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

@media (max-width: 768px) {
  .desktop-only { display: none !important; }
  .mobile-only { display: block !important; }
  .mobile-only .item-card {
    margin-bottom: 1rem;
  }
  .job-title { font-size: 0.95rem; }
}
@media (min-width: 769px) {
  .desktop-only { display: table; }
  .mobile-only { display: none; }
}
@media (min-width: 769px) and (max-width: 1024px) {
  .jobs-table { font-size: 0.9rem; }
  th, td { padding: 0.875rem 1.25rem; }
}
</style>
