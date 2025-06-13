<script setup>
import { onMounted } from 'vue'
import { RouterLink } from 'vue-router'
//import {storeToRefs } from 'pinia'
import { useCompanies } from '@/stores/companies'

const companiesStore = useCompanies()

onMounted(() => {
  companiesStore.fetchCompanies()
})

</script>

<template>
<div>
  <h1>Companies</h1>
  <ul>
    <li v-for="company in companiesStore.sortedCompanies" :key="company.id">

      <RouterLink :to="{ name: 'company-detail', params: { slug: company.slug } }">
        {{ company.name }}
      </RouterLink>
    </li>
    <!-- <li v-for="company in companies" :key="company.id">
      <RouterLink :to="{ name: 'Company', params: { id: company.id } }">
        {{ company.name }}
      </RouterLink>
    </li> -->
  </ul>
</div>
</template>

<style scoped>
/* todo: styles */
</style>
