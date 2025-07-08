<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useProfileStore } from '@/stores/profile'
import { Container, Button } from '@/components/ui'
import { SkeletonImage, SkeletonText } from '@/components/skeleton'
import { EditIcon } from '@/components/icons'

const route = useRoute()
const router = useRouter()
const profile = useProfileStore()

const is_editing = computed(() => route.query.edit === '1')


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

    <!-- todo: User Avatar    (avatar_url)        -->
    <SkeletonImage style="margin-bottom: 1rem;" />

    <!-- todo: First Name     (first_name)        -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: Middle Name    (middle_name)       -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: Last Name      (last_name)         -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: Preferred Name (preferred_name)    -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: email          (email)             -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: phone          (phone)             -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: Gravatar       (gravatar_username) -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: Dark Mode      (gravatar_username) --> <!-- Should this be here? -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: Created        (created_at)        -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: Updated        (updated_at)        -->
    <SkeletonText style="margin-bottom: 1rem;" />

    <!-- todo: URLS           (urls)              --> <!-- figure out how to make this an array of values-->
    <SkeletonText style="margin-bottom: 1rem;" />

  </div>


  <!-- Work Experience -->
  <div class="profile-section">
    <h2>Work Experience</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Education Experience -->
  <div class="profile-section">
    <h2>Education Experience</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Licenses & Certifications -->
  <div class="profile-section">
    <h2>Licenses & Certifications</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Memberships -->
  <div class="profile-section">
    <h2>Memberships</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Skills -->
  <div class="profile-section">
    <h2>Skills</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Interests -->
  <div class="profile-section">
    <h2>Interests</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Projects -->
  <div class="profile-section">
    <h2>Projects</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- References -->
  <div class="profile-section">
    <h2>References</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>

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
</style>
