<script setup>
import { onMounted } from 'vue'
import { RouterLink, RouterView } from 'vue-router'
//import {storeToRefs } from 'pinia'
import { useCompany } from '@/stores/company'

const companyStore = useCompany()

onMounted(() => {
  companyStore.fetchCompany(this.$route.query.slug)
})

</script>

<template>
<div>
  <h1>
    <!-- <span>Company: {{ companiesStore.company?.name }}</span><br /> -->
    <!-- <small>{{ companiesStore.company?.slug }}</small> -->
  </h1>
  <!-- <RouterLink :to="{ name: 'companies-list' }"> -->
    <!-- < Companies -->
  <!-- </RouterLink> -->

  <code>
    {{ this.$route.query }}
  </code>
  <p>
  <!-- {{ companiesStore?.company.description }} -->
  </p>
</div>
</template>

<style scoped>
/* todo: styles */
</style>
