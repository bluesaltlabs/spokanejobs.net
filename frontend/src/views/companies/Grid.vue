<script setup>
import { onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useCompanies } from '@/stores/companies'
import { SkeletonCard } from '@/components/skeleton'
import { Container, Button } from '@/components/ui'
import { CompaniesIcon } from '@/components/icons'
import { ItemContainer, ItemCard, ItemEmptyState, SearchBar } from '@/components/common'

// setup
const router = useRouter()
const companiesStore = useCompanies()

// Fetch companies on component mount
onMounted(() => {
  companiesStore.fetchCompanies()
})

const loading = computed(() => companiesStore.loading)
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
        @click="companiesStore.updateFilters({ search: '' })"
        aria-label="Clear all filters"
        type="button"
      >
        Clear Filters
      </Button>
    </div>

    <div v-if="loading" class="loading-state">
      <div class="loading-message">Loading companies...</div>
      <ItemContainer grid>
        <SkeletonCard v-for="i in 6" :key="`skeleton-${i}`" />
      </ItemContainer>
    </div>

    <div v-else-if="companiesStore.sortedCompanies.length > 0" class="companies-results">
      <ItemContainer grid>
        <ItemCard
          v-for="company in companiesStore.sortedCompanies"
          :key="company.id"
          :onClick="() => router.push({ name: 'company-detail', params: { slug: company.slug }})"
        >
          <template #header>
            <div class="company-header">
              <div class="company-logo">
                <img v-if="company?.logo_url" :src="company.logo_url" alt="Company Logo" class="company-logo-img"/>
                <CompaniesIcon v-else class="company-logo-icon" />
              </div>
              <div class="company-info">
                <h3 class="company-title">{{ company.name }}</h3>
              </div>
            </div>
          </template>
          <template #meta>
            <div class="company-meta">
              <div v-if="company.city || company.state" class="company-location">
                {{ company.city }}{{ company.state ? ', ' + company.state : '' }}
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
        <CompaniesIcon />
      </template>
      <template #title>No companies found</template>
      <template #desc>Companies will appear here once they are added to the system.</template>
    </ItemEmptyState>
  </Container>
</template>

<style scoped>

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


.company-header {
  display: flex;
  align-items: flex-start;
  gap: 0.75rem;
  margin-bottom: 0.75rem;
}

.company-logo {
  width: 40px;
  height: 48px;
  background: var(--color-primary-50);
  border-radius: var(--border-radius-medium);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-primary-600);
  flex-shrink: 0;
}

.company-logo-img {
  width: 28px;
  height: 32px;
  object-fit: contain;
  border-radius: var(--border-radius-small);
}

.company-logo-icon {
  width: 28px;
  height: 32px;
}

.company-info {
  flex: 1;
  min-width: 0;
}

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

.company-meta {
  margin-bottom: 0.75rem;
}

.company-location {
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

@media (max-width: 768px) {
  .company-title { font-size: 0.95rem; }
  .company-logo {
    width: 40px;
    height: 40px;
  }
  .company-logo-img {
    width: 28px;
    height: 28px;
  }
  .company-logo-icon {
    width: 28px;
    height: 28px;
  }
}
</style>
