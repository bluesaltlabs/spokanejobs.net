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
    <router-link class="back-button-link" to="/companies">🏢 All Companies</router-link>
  </div>
</div>
  <div v-if="company">
    <h1 class="company-name">{{ company.name }}</h1>
    <div class="company-links-container">
      <span v-if="company.website"><a :href="company.website" target="_blank" :alt="company.website">🌐 Website</a></span>
      <span v-if="company.job_board_url"><a :href="company.job_board_url" target="_blank" :alt="company.job_board_url">💼 Jobs</a></span>
    </div>

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

.company-name {
  margin-top: 25px;
  margin-bottom: 15px;
  font-size: 3.5rem;
}

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

.company-links-container {
  margin-top: 15px;
  margin-bottom: 15px;


  a {
    border-radius: var(--border-radius-small);
    padding: 5px 10px;
    margin: 5px 8px;
    background: var(--color-primary-700);
    color: var(--color-text);
    text-decoration: none;
    font-weight: bold;
    font-size: 1.2rem;

    &:hover {
      background: var(--color-primary-500);
      color: var(--color-white);
    }

  }
}

</style>
