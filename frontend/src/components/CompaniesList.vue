<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCompanies } from '@/stores/companies'

const companiesStore = useCompanies()
const router = useRouter()

onMounted(() => {
  companiesStore.fetchCompanies()
})

function editCompany(slug) {
  router.push({ name: 'company-edit', params: { slug } })
}

async function deleteCompany(id) {
  if (confirm('Delete this company?')) {
    await companiesStore.deleteCompany(id)
    await companiesStore.fetchCompanies()
  }
}
</script>

<template>
  <div>
    <h1>Companies</h1>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Slug</th>
          <th>Website</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="company in companiesStore.sortedCompanies" :key="company.id">
          <td>
            <router-link :to="{ name: 'company-detail', params: { slug: company.slug } }">
              {{ company.name }}
            </router-link>
          </td>
          <td>{{ company.slug }}</td>
          <td>
            <a :href="company.website" target="_blank" v-if="company.website">{{ company.website }}</a>
          </td>
          <td>
            <button @click="editCompany(company.slug)" class="btn btn-secondary btn-sm">Edit</button>
            <button @click="deleteCompany(company.id)" class="btn btn-danger btn-sm">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
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

button {
  margin-right: 0.5rem;
  padding: 0.5rem 1rem;
  border: 1px solid var(--color-border);
  background: var(--color-button-secondary);
  color: var(--color-button-secondary-text);
  border-radius: var(--border-radius-small);
  cursor: pointer;
  transition: all 0.2s ease;
}

button:hover {
  background: var(--color-button-secondary-hover);
  border-color: var(--color-border-hover);
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
