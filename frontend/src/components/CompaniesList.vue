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
            <button @click="editCompany(company.slug)">Edit</button>
            <button @click="deleteCompany(company.id)">Delete</button>
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
}
th, td {
  padding: 0.5rem 1rem;
  border-bottom: 1px solid #eee;
}
th {
  text-align: left;
  background: #f8fafc;
}
button {
  margin-right: 0.5rem;
}
</style>
