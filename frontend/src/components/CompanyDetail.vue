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
<div>
  <!-- todo: back to companies grid header -->
  <div class="back-button">
    <router-link class="back-button-link" to="/companies"> < Companies</router-link>
  </div>
</div>
  <div v-if="company">
    <h1>{{ company.name }}</h1>
    <p v-if="company.website"><a :href="company.website" target="_blank">{{ company.website }}</a></p>
    <div class="spacer"></div>
    <img v-if="company?.logo_url" :src="company.logo_url" :alt="company.name + ' Logo'" class="company-logo" />


    <!-- add a spacer here, or add padding to the paragraphs -->
    <p v-if="company.description">{{ company.description }}</p>
  </div>
  <div v-else>
    <p>Loading...</p>
    <!-- todo: add loading shadow elements for better UX -->
  </div>
</template>

<style scoped>

img.company-logo {
  max-height: 425px;
  max-width: 425px;
  width: 100%;
  height: auto;
  /* image rounding, padding, box-shadow */
  border-radius: var(--border-radius);
  padding: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.back-button-link {
  color: var(--color-text);
  text-decoration: none;
  font-weight: bold;
  font-size: 1.2rem;
}

.back-button-link:hover {
  color: var(--color-primary);
}

.spacer {
  height: 30px;
}

</style>
