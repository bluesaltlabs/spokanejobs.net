<script setup>
import { onMounted, computed } from 'vue'
import { useCompanies } from '@/stores/companies'
import { SkeletonTableRow } from '@/components/skeleton'

// setup
const companiesStore = useCompanies()

// Fetch companies on component mount
onMounted(() => {
  companiesStore.fetchCompanies()
})

const loading = computed(() => companiesStore.loading)
</script>

<template>
  <main>
    <h1>Companies</h1>
    <hr />

    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Slug</th>
          <th>Website</th>
        </tr>
      </thead>
      <tbody>
        <!-- Loading skeleton rows -->
        <SkeletonTableRow
          v-if="loading"
          v-for="i in 8"
          :key="`skeleton-${i}`"
          :columns="3"
          :column-widths="['40%', '30%', '30%']"
        />

        <!-- Actual company data -->
        <tr v-else v-for="company in companiesStore.sortedCompanies" :key="company.id">
          <td>
            <router-link :to="{ name: 'company-detail', params: { slug: company.slug } }">
              {{ company.name }}
            </router-link>
          </td>
          <td>{{ company.slug }}</td>
          <td>
            <a :href="company.website" target="_blank" v-if="company.website">{{ company.website }}</a>
          </td>
        </tr>
      </tbody>
    </table>
  </main>
</template>


<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
  background: var(--color-surface);
  border-radius: var(--border-radius-medium);
  overflow: hidden;
  box-shadow: 0 1px 3px var(--color-shadow);
}

th, td {
  padding: 0.75rem 1rem;
  border-bottom: 1px solid var(--color-border);
}

th {
  text-align: left;
  background: var(--color-background-soft);
  font-weight: 600;
  color: var(--color-heading);
}

td {
  color: var(--color-text);
}

tr:hover {
  background: var(--color-surface-hover);
}

a {
  color: var(--color-link);
  text-decoration: none;
}

a:hover {
  color: var(--color-link-hover);
  text-decoration: underline;
}
</style>
