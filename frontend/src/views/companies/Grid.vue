<script setup>
import { onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useCompanies } from '@/stores/companies'
import { SkeletonCard } from '@/components/skeleton'
import { Container } from '@/components/ui'
import { CompaniesIcon } from '@/components/icons'

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

    <div v-if="loading" class="loading-state">
      <div class="loading-message">Loading companies...</div>
      <div class="companies-grid">
        <SkeletonCard v-for="i in 6" :key="`skeleton-${i}`" />
      </div>
    </div>

    <div v-else-if="companiesStore.sortedCompanies.length > 0" class="companies-results">
      <div class="results-count">{{ companiesStore.sortedCompanies.length }} companies found</div>
      <div class="companies-grid">
        <div
          v-for="company in companiesStore.sortedCompanies"
          :key="company.id"
          class="company-card"
        >
          <router-link
            :to="{ name: 'company-detail', params: { slug: company.slug }}"
            class="company-link"
          >
            <div class="company-header">
              <div class="company-logo">
                <img v-if="company?.logo_url" :src="company.logo_url" alt="Company Logo" class="company-logo-img"/>
                <CompaniesIcon v-else class="company-logo-icon" />
              </div>
              <div class="company-info">
                <h3 class="company-title">{{ company.name }}</h3>
              </div>
            </div>

            <div class="company-meta">
              <div v-if="company.website" class="company-website">
                <a :href="company.website" target="_blank" rel="noopener noreferrer" class="website-link">
                  {{ company.website }}
                </a>
              </div>
              <div v-if="company.city || company.state" class="company-location">
                {{ company.city }}{{ company.state ? ', ' + company.state : '' }}
              </div>
            </div>

            <div class="company-actions">
              <span class="view-details">View Details →</span>
            </div>
          </router-link>
        </div>
      </div>
    </div>

    <div v-else class="empty-state">
      <div class="empty-icon">
        <CompaniesIcon />
      </div>
      <h3>No companies found</h3>
      <p>Companies will appear here once they are added to the system.</p>
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

.companies-results {
  margin-bottom: 1rem;
}

.results-count {
  font-weight: 600;
  color: var(--color-heading);
  margin-bottom: 1rem;
}

.companies-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
  width: 100%;
}

.company-card {
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

.company-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px var(--color-shadow-elevated);
  border-color: var(--color-border-hover);
}

.company-link {
  display: block;
  text-decoration: none;
  color: inherit;
  padding: 1rem;
  height: 100%;
  box-sizing: border-box;
}

.company-header {
  display: flex;
  align-items: flex-start;
  gap: 0.75rem;
  margin-bottom: 0.75rem;
}

.company-logo {
  width: 48px;
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
  max-width: 32px;
  max-height: 32px;
  width: 100%;
  height: auto;
  border-radius: var(--border-radius-small);
}

.company-logo-icon {
  width: 24px;
  height: 24px;
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

.company-website {
  color: var(--color-text-muted);
  font-size: 0.875rem;
  margin-bottom: 0.25rem;
}

.website-link {
  color: var(--color-link);
  text-decoration: none;
  transition: color 0.2s ease;
}

.website-link:hover {
  color: var(--color-link-hover);
  text-decoration: underline;
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

.company-card:hover .view-details {
  color: var(--color-primary-700);
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
  .companies-grid {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
  .company-link { padding: 0.875rem; }
  .company-title { font-size: 0.95rem; }
  .company-logo {
    width: 40px;
    height: 40px;
  }
  .company-logo-img {
    max-width: 28px;
    max-height: 28px;
  }
  .company-logo-icon {
    width: 20px;
    height: 20px;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .companies-grid {
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 1.25rem;
  }
}

@media (min-width: 1025px) {
  .companies-grid {
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    gap: 1.75rem;
  }
}
</style>
