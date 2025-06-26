<script setup>
import { onMounted, computed } from 'vue'
import { useJobs } from '@/stores/jobs'
import { SkeletonTableRow } from '@/components/skeleton'

// setup
const jobsStore = useJobs()

// Fetch jobs on component mount
onMounted(() => {
  jobsStore.fetchJobs()
})

const loading = computed(() => jobsStore.loading)
</script>

<template>
  <main>
    <h1>Jobs</h1>
    <hr />

    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Company</th>
        </tr>
      </thead>
      <tbody>
        <!-- Loading skeleton rows -->
        <SkeletonTableRow
          v-if="loading"
          v-for="i in 8"
          :key="`skeleton-${i}`"
          :columns="2"
          :column-widths="['50%', '50%']"
        />

        <!-- Actual company data -->
        <tr v-else v-for="job in jobsStore.sortedJobs" :key="job.id">
          <td>
            <router-link :to="{ name: 'job-detail', params: { id: job.id } }">
              {{ job.title }}
            </router-link>
          </td>
          <td>{{ job.company }}</td>
        </tr>

        <tr v-if="jobsStore.sortedJobs.length === 0" class="empty-state">
          <td colspan="2">No jobs found</td>
        </tr>
      </tbody>
    </table>



      <!-- <li v-for="job in jobsStore.sortedJobs" :key="job.id">
        <RouterLink :to="{ name: 'job-detail', params: { id: job.id } }">
          {{ job.title }}
        </RouterLink>
      </li> -->

  </main>
</template>



<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
  background: var(--color-surface);
  border-radius: var(--border-radius-medium);
  overflow: hidden;
  box-shadow: 0 1px 3px var(--color-shadow);
}

th, td {
  padding: 0.75rem 1rem;
  border-bottom: 1px solid var(--color-border);
}

th {
  text-align: left;
  background: var(--color-background-soft);
  font-weight: 600;
  color: var(--color-heading);
}

td {
  color: var(--color-text);
}

tr:hover {
  background: var(--color-surface-hover);
}

a {
  color: var(--color-link);
  text-decoration: none;
}

a:hover {
  color: var(--color-link-hover);
  text-decoration: underline;
}
</style>
