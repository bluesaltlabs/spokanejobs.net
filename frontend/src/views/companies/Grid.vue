<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCompanies } from '@/stores/companies'

// setup
const router = useRouter()
const companiesStore = useCompanies()

// Fetch companies on component mount
onMounted(() => {
  companiesStore.fetchCompanies()
})

// functions
function goToCreate() {
  router.push({ name: 'company-create' })
}
</script>

<template>
  <main>
    <h1>Companies</h1>
    <hr class="my-4" />

    <div class="companies-grid">
      <div class="grid-container">
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
                <div v-if="company?.logo_url">
                  <img :src="company.logo_url" alt="Company Logo" class="company-logo"/>
                </div>
                <div v-if="!company?.logo_url">
                  <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                    <polyline points="9,22 9,12 15,12 15,22"></polyline>
                  </svg>
                </div>
              </div>
              <div class="company-info">
                <h3 class="company-name">{{ company.name }}</h3>

              </div>
            </div>

            <div class="company-meta">
              <p v-if="company.website" class="company-website">{{ company.website }}</p>
            </div>

            <div class="company-actions">
              <span class="view-details">View Details →</span>
            </div>
          </router-link>
        </div>
      </div>

      <div v-if="companiesStore.sortedCompanies.length === 0" class="empty-state">
        <div class="empty-icon">
          <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
            <polyline points="9,22 9,12 15,12 15,22"></polyline>
          </svg>
        </div>
        <h3>No companies found</h3>
        <p>Companies will appear here once they are added to the system.</p>
      </div>
    </div>
  </main>
</template>


<style scoped>
img.company-logo {
  max-width: 32px;
  max-height: 32px;
  width: 100%;
  height: auto;
}
hr {
  margin-top: 1rem;
  margin-bottom: 2rem;
}

.companies-grid {
  width: 100%;
}

.grid-container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.company-card {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-medium);
  overflow: hidden;
  transition: all 0.2s ease;
  box-shadow: 0 1px 3px var(--color-shadow);
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
  padding: 1.5rem;
  height: 100%;
}

.company-header {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  margin-bottom: 1rem;
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

.company-info {
  flex: 1;
  min-width: 0;
}

.company-name {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--color-heading);
  margin: 0 0 0.25rem 0;
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.company-website {
  font-size: 0.875rem;
  color: var(--color-text-muted);
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.company-meta {
  margin-bottom: 1rem;
}

.meta-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem 0;
  border-bottom: 1px solid var(--color-border);
}

.meta-item:last-child {
  border-bottom: none;
}

.meta-label {
  font-size: 0.875rem;
  color: var(--color-text-subtle);
  font-weight: 500;
}

.meta-value {
  font-size: 0.875rem;
  color: var(--color-text);
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  background: var(--color-background-soft);
  padding: 0.25rem 0.5rem;
  border-radius: var(--border-radius-small);
  max-width: 120px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.company-actions {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  padding-top: 1rem;
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

/* Responsive adjustments */
@media (max-width: 768px) {
  .grid-container {
    grid-template-columns: 1fr;
    gap: 1rem;
  }

  .company-card {
    margin-bottom: 0;
  }

  .company-link {
    padding: 1rem;
  }

  .company-header {
    gap: 0.75rem;
  }

  .company-logo {
    width: 40px;
    height: 40px;
  }

  .company-name {
    font-size: 1rem;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .grid-container {
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 1.25rem;
  }
}

@media (min-width: 1025px) {
  .grid-container {
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    gap: 1.75rem;
  }
}
</style>
