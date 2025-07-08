<script setup>
import { onMounted } from 'vue'
import { useProfileStore } from '@/stores/profile'
import { Container, Button } from '@/components/ui'
import { SkeletonImage, SkeletonText } from '@/components/skeleton'
import { EditIcon } from '@/components/icons'

const profile = useProfileStore()

onMounted(() => {
  profile.loadProfile()
})

function toggleEdit() {
  // todo: toggle edit mode.
}

</script>

<template>
<Container>
  <div class="page-header">
    <h1>Profile</h1>


    <div class="profile-actions">
      <Button @click="toggleEdit" variant="primary">
        <EditIcon />
      </Button>
    </div>
  </div>

  <hr class="divider" />

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
    margin: 0 auto;
    padding: 1rem;
    background: var(--color-background-elevated);
    border-radius: var(--border-radius-medium);
    border: 1px solid var(--color-border);
  }
</style>
