<script setup>
import CompaniesGrid from '../components/CompaniesGrid.vue'
import { useRouter } from 'vue-router'
const router = useRouter()
function goToCreate() {
  router.push({ name: 'company-create' })
}
</script>

<template>
  <main>
    <CompaniesGrid />
  </main>
</template>
