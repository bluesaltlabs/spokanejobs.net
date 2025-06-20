<script setup>
import CompaniesList from '../components/CompaniesList.vue'
import { useRouter } from 'vue-router'
const router = useRouter()
function goToCreate() {
  router.push({ name: 'company-create' })
}
</script>

<template>
  <main>
    <div style="display: flex; justify-content: space-between; align-items: center;">
      <h1>Companies</h1>
      <button @click="goToCreate">Add Company</button>
    </div>
    <CompaniesList />
  </main>
</template>
