<script setup>
import { ref, onMounted } from 'vue'
import { supabase } from '../lib/supabaseClient'

const companies = ref([])

async function getCompanies() {
  const { data, error } = await supabase.from('companies').select('*')
  if (error) throw error
  companies.value = data
}

onMounted(() => {
  getCompanies()
})
</script>

<template>
<div>
  <h1>Companies</h1>
  <ul>
    <li v-for="company in companies" :key="company.id">

      {{ company.name }}

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
