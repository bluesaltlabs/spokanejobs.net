<script setup>
import { ref, watch, onMounted } from 'vue'
import { useCompanies } from '@/stores/companies'
import { useRouter, useRoute } from 'vue-router'

const companiesStore = useCompanies()
const router = useRouter()
const route = useRoute()

const isEdit = route.name === 'company-edit'
const form = ref({
  name: '',
  slug: '',
  website: '',
  description: ''
})

onMounted(async () => {
  if (isEdit) {
    await companiesStore.fetchCompanies()
    const company = companiesStore.companies.find(c => c.slug === route.params.slug)
    if (company) {
      form.value = { ...company }
    }
  }
})

async function submit() {
  if (isEdit) {
    await companiesStore.updateCompany(form.value)
  } else {
    await companiesStore.createCompany(form.value)
  }
  router.push({ name: 'companies-list' })
}
</script>

<template>
  <form @submit.prevent="submit">
    <div>
      <label>Name</label>
      <input v-model="form.name" required />
    </div>
    <div>
      <label>Slug</label>
      <input v-model="form.slug" required />
    </div>
    <div>
      <label>Website</label>
      <input v-model="form.website" />
    </div>
    <div>
      <label>Description</label>
      <textarea v-model="form.description" />
    </div>
    <button type="submit" class="btn btn-primary">{{ isEdit ? 'Update' : 'Create' }}</button>
  </form>
</template>
