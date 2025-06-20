<script setup>
import { onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useCompany } from '@/stores/company'

const companyStore = useCompany()
const route = useRoute()

onMounted(() => {
  companyStore.fetchCompany(route.params.slug)
})

const company = computed(() => companyStore.company)
</script>

<template>
  <div v-if="company">
    <h1>{{ company.name }}</h1>
    <p><strong>Slug:</strong> {{ company.slug }}</p>
    <p v-if="company.website"><strong>Website:</strong> <a :href="company.website" target="_blank">{{ company.website }}</a></p>
    <p v-if="company.description"><strong>Description:</strong> {{ company.description }}</p>
  </div>
  <div v-else>
    <p>Loading...</p>
  </div>
</template>

<style scoped>
/* todo: styles */
</style>
