<script setup>
import { onMounted, computed } from 'vue'
import { useCompanies } from '@/stores/companies'
import { SkeletonTableRow } from '@/components/skeleton'
import { Container } from '@/components/ui'
import { ItemContainer, ItemCard, ItemTableRow, ItemEmptyState, SearchBar } from '@/components/common'
import Button from '@/components/ui/Button/Button.vue'

// setup
const companiesStore = useCompanies()

// Fetch companies on component mount
onMounted(() => {
  companiesStore.fetchCompanies()
})

const loading = computed(() => companiesStore.loading)

function clearFilters() {
  companiesStore.updateFilters({ search: '' })
}
</script>

<template>
  <Container>
    <h1>Companies</h1>
    <hr class="divider" />
    <SearchBar
      v-model="companiesStore.filters.search"
      placeholder="Search companies..."
      @update:modelValue="val => companiesStore.updateFilters({ search: val })"
    />

    <div class="results-row">
      <div class="results-count">{{ companiesStore.sortedCompanies.length }} companies found</div>
      <Button
        variant="secondary"
        class="clear-filters-btn"
        @click="clearFilters"
        aria-label="Clear all filters"
        type="button"
      >
        Clear Filters
      </Button>
    </div>
    <div v-if="loading" class="loading-state">
      <div class="loading-message">Loading companies...</div>
      <table class="companies-table desktop-only">
        <thead>
          <tr>
            <th>Name</th>
            <th>Slug</th>
            <th>Website</th>
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
      <ItemContainer grid customClass="mobile-only">
        <ItemCard v-for="i in 6" :key="`skeleton-card-${i}`" customClass="skeleton-card">
          <template #header>
            <div class="company-header">
              <div class="skeleton-name"></div>
              <div class="skeleton-slug"></div>
            </div>
          </template>
          <template #meta>
            <div class="company-meta">
              <div class="skeleton-website"></div>
            </div>
          </template>
        </ItemCard>
      </ItemContainer>
    </div>

    <div v-else-if="companiesStore.sortedCompanies.length > 0" class="companies-results">
      <table class="companies-table desktop-only">
        <thead>
          <tr>
            <th>Name</th>
            <th>Slug</th>
            <th>Website</th>
          </tr>
        </thead>
        <tbody>
          <ItemTableRow v-for="company in companiesStore.sortedCompanies" :key="company.id">
            <template #cells>
              <td class="company-name-cell">
                <router-link :to="{ name: 'company-detail', params: { slug: company.slug } }" class="company-name-link">
                  {{ company.name }}
                </router-link>
              </td>
              <td class="company-slug-cell">{{ company.slug }}</td>
              <td class="company-website-cell">
                <a v-if="company.website" :href="company.website" target="_blank" rel="noopener noreferrer" class="website-link">
                  {{ company.website }}
                </a>
              </td>
            </template>
          </ItemTableRow>
        </tbody>
      </table>
      <ItemContainer grid customClass="mobile-only">
        <ItemCard v-for="company in companiesStore.sortedCompanies" :key="company.id">
          <template #header>
            <router-link :to="{ name: 'company-detail', params: { slug: company.slug } }" class="company-link">
              <div class="company-header">
                <h3 class="company-title">{{ company.name }}</h3>
                <div class="company-slug">{{ company.slug }}</div>
              </div>
            </router-link>
          </template>
          <template #meta>
            <div class="company-meta">
              <div v-if="company.website" class="company-website">
                <a :href="company.website" target="_blank" rel="noopener noreferrer" class="website-link">
                  {{ company.website }}
                </a>
              </div>
            </div>
          </template>
          <template #actions>
            <div class="company-actions">
              <span class="view-details">View Details →</span>
            </div>
          </template>
        </ItemCard>
      </ItemContainer>
    </div>

    <ItemEmptyState v-else>
      <template #icon>
        <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
          <circle cx="9" cy="7" r="4"/>
          <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
          <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
        </svg>
      </template>
      <template #title>No companies found</template>
      <template #desc>Check back later for new companies in the Spokane area.</template>
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

.companies-results {
  margin-bottom: 1rem;
}

.results-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  margin-bottom: 1rem;
}

.results-count {
  font-weight: 600;
  color: var(--color-heading);
  margin-bottom: 1rem;
}

.companies-table {
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

.company-name-cell { width: 40%; }
.company-slug-cell { width: 30%; }
.company-website-cell { width: 30%; }

.company-name-link {
  color: var(--color-link);
  text-decoration: none;
  font-weight: 600;
  font-size: 1rem;
  line-height: 1.4;
  transition: color 0.2s ease;
}

.company-name-link:hover {
  color: var(--color-link-hover);
  text-decoration: underline;
}

.website-link {
  color: var(--color-link);
  text-decoration: none;
  font-size: 0.9rem;
  transition: color 0.2s ease;
}

.website-link:hover {
  color: var(--color-link-hover);
  text-decoration: underline;
}

.company-link {
  display: block;
  text-decoration: none;
  color: inherit;
  padding: 1rem;
  height: 100%;
  box-sizing: border-box;
}

.company-header { margin-bottom: 0.75rem; }
.company-title {
  font-size: 1rem;
  font-weight: 600;
  color: var(--color-heading);
  margin: 0 0 0.25rem 0;
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.company-slug {
  margin-bottom: 0.5rem;
  color: var(--color-text-muted);
  font-size: 0.875rem;
  font-family: monospace;
}
.company-meta { margin-bottom: 0.75rem; }
.company-website {
  color: var(--color-text-muted);
  font-size: 0.875rem;
}
.company-actions {
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
.skeleton-name {
  height: 1.25rem;
  background: var(--color-background-soft);
  border-radius: var(--border-radius-small);
  margin-bottom: 0.5rem;
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
.skeleton-slug {
  height: 1rem;
  background: var(--color-background-soft);
  border-radius: var(--border-radius-small);
  margin-bottom: 0.5rem;
  width: 60%;
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
.skeleton-website {
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
  .company-link { padding: 0.875rem; }
  .company-title { font-size: 0.95rem; }
}
@media (min-width: 769px) {
  .desktop-only { display: table; }
  .mobile-only { display: none; }
}
@media (min-width: 769px) and (max-width: 1024px) {
  .companies-table { font-size: 0.9rem; }
  th, td { padding: 0.875rem 1.25rem; }
}
</style>
