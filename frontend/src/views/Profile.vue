<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useProfileStore } from '@/stores/profile'
import {
  EditIcon,
  Container, Button, Form, FormGroup, FormRow, TextInput,
  WorkExperience, EducationExperience, LicensesCertifications, Memberships,
  Skills, Interests, Projects, References
} from '@/components'

const route = useRoute()
const router = useRouter()
const profile = useProfileStore()
const saved = ref(false)

const is_editing = computed(() => route.query.edit === '1')

// Auto-save functionality
const autoSave = async () => {
  await profile.saveProfile()
  saved.value = true
  setTimeout(() => (saved.value = false), 1500)
}

// Watch for changes in profile fields and auto-save
watch(
  () => [
    profile.personal_information.first_name,
    profile.personal_information.middle_name,
    profile.personal_information.last_name,
    profile.personal_information.email,
    profile.personal_information.phone,
    profile.personal_information.avatar_url
  ],
  () => {
    autoSave() // todo: debounce this.
  },
  { deep: true }
)

onMounted(() => {
  profile.loadProfile()
})

async function updateEditQueryParam(value) {
  const newQuery = { ...route.query }
  if (value) {
    newQuery.edit = '1'
  } else {
    delete newQuery.edit
  }
  await router.replace({ query: newQuery })
}

async function onEditButtonClick() {
  const newValue = !is_editing.value
  await updateEditQueryParam(newValue)
}

</script>

<template>
<Container>
  <div class="page-header">
    <h1>Profile</h1>

    <div class="profile-actions">
      <Button @click="onEditButtonClick" variant="primary">
        <EditIcon />
      </Button>
    </div>
  </div>

  <hr class="divider" />

  <!-- todo: add some kind of sidebar to scroll more quickly to each section as they grow? -->


  <!-- Personal Information -->
  <div class="profile-section">
    <h2>Personal Information</h2>
    <hr class="divider" />

    <!-- Avatar -->
    <div class="avatar-preview">
      <div class="avatar">
        <img v-if="profile.personal_information.avatar_url" :src="profile.personal_information.avatar_url" alt="Avatar preview" />
        <span v-else>{{ profile.personal_information.first_name?.charAt(0) || 'U' }}</span>
      </div>
    </div>


    <Form>

      <!-- Avatar URL -->
      <FormGroup style="flex-grow: 1;" label="Avatar URL" v-if="is_editing">
        <TextInput
          v-model="profile.personal_information.avatar_url"
          placeholder="avatar url"
          :editable="is_editing"
          />
      </FormGroup>

      <FormRow>
        <!-- First Name -->
        <FormGroup style="flex-grow: 1;" label="First Name" v-if="profile.personal_information.first_name || is_editing">
          <TextInput
            v-model="profile.personal_information.first_name"
            placeholder="first name"
            :editable="is_editing"
            />
        </FormGroup>

        <!-- Middle Name -->
        <FormGroup style="flex-shrink: 1;" label="Middle Name" v-if="profile.personal_information.middle_name || is_editing">
          <TextInput
            v-model="profile.personal_information.middle_name"
            placeholder="middle name"
            :editable="is_editing"
            />
        </FormGroup>

        <!-- Last Name -->
        <FormGroup style="flex-grow: 1;" label="Last Name" v-if="profile.personal_information.last_name || is_editing">
          <TextInput
            v-model="profile.personal_information.last_name"
            placeholder="last name"
            :editable="is_editing"
            />
        </FormGroup>
      </FormRow>

      <!-- Preferred Name -->
      <FormGroup label="Preferred Name" v-if="profile.personal_information.preferred_name || is_editing">
        <TextInput
          v-model="profile.personal_information.preferred_name"
          placeholder="preferred name"
          :editable="is_editing"
          />
      </FormGroup>

      <FormRow>
        <!-- Email -->
        <FormGroup style="flex-grow: 1;" label="Email" v-if="profile.personal_information.email || is_editing">
          <TextInput
            v-model="profile.personal_information.email"
            placeholder="email"
            :editable="is_editing"
            />
        </FormGroup>

        <!-- Phone -->
        <FormGroup style="flex-grow: 1;" label="Phone" v-if="profile.personal_information.phone || is_editing">
          <TextInput
            v-model="profile.personal_information.phone"
            placeholder="phone"
            :editable="is_editing"
            />
        </FormGroup>
      </FormRow>

        <FormRow>
          <!-- Created At -->
          <FormGroup style="flex-grow: 1;" label="Created">
            <TextInput
              v-model="profile.personal_information.created_at"
              placeholder="yyyy/mm/dd hh:mm:ss"
              :editable="false"
              />
          </FormGroup>


          <!-- Updated At -->
          <FormGroup style="flex-grow: 1;" label="Updated">
            <TextInput
              v-model="profile.personal_information.updated_at"
              placeholder="yyyy/mm/dd hh:mm:ss"
              :editable="false"
            />
          </FormGroup>

        <!-- todo: URLS           (urls)              --> <!-- figure out how to make this an array of values-->
        </FormRow>
      </Form>

    </div>


    <!-- Work Experience -->
    <WorkExperience :is-editing="is_editing" />

    <!-- Education Experience -->
    <EducationExperience />


    <!-- Licenses & Certifications -->
    <LicensesCertifications />


    <!-- Memberships -->
    <Memberships />


    <!-- Skills -->
    <Skills />


    <!-- Interests -->
    <Interests />


    <!-- Projects -->
    <Projects />


    <!-- References -->
    <References />

</Container>
</template>

<style scoped>
  .page-header {
    display: flex;
    justify-content: space-between;
    flex-grow: 1;
    margin-bottom: 1rem;
  }

  .profile-actions {
    display: flex;
    justify-content: end;
    padding: 0.5em 0;
  }

  .profile-section {
    max-width: 960px;
    margin-bottom: 1.5rem;
    padding: 1rem;
    background: var(--color-background-elevated);
    border-radius: var(--border-radius-medium);
    border: 1px solid var(--color-border);
  }


  .avatar-preview {
    text-align: center;
    margin-top: 1rem;
    margin-bottom: 1rem;
  }

  .avatar {
    width: 100px;
    height: 100px;
    background: var(--color-surface-hover);
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--color-text-subtle);
    border-radius: var(--border-radius-medium);
    margin: 0 auto;

    img {
      width: 100px;
      height: 100px;
      border-radius: var(--border-radius-small);
      object-fit: cover;
    }
  }


</style>
