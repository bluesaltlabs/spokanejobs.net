<script setup>
import { onMounted } from 'vue'
//import {storeToRefs } from 'pinia'
import { useJobs } from '@/stores/jobs'

const jobsStore = useJobs()

onMounted(() => {
  jobsStore.fetchJobs()
})

</script>

<template>
<div>
  <h1>Jobs</h1>
  <ul>
    <li v-for="job in jobsStore.sortedJobs" :key="job.id">

      {{ job.title }}

    </li>
    <!-- <li v-for="job in jobsStore.sortedJobs" :key="job.id">
      <RouterLink :to="{ name: 'Job', params: { id: job.id } }">
        {{ job.title }}
      </RouterLink>
    </li> -->
  </ul>
</div>
</template>

<style scoped>
/* todo: styles */
</style>
