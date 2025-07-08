<script setup>
import { onMounted } from 'vue'
import { useProfileStore } from '@/stores/profile'
import { Container, Button } from '@/components/ui'
import { EditIcon } from '@/components/icons'

const profile = useProfileStore()

onMounted(() => {
  profile.loadProfile()
})

</script>

<style scoped>
  .profile-actions {
    display: flex;
    justify-content: end;
    margin-bottom: 1rem;
  }
</style>

<template>
<Container>
  <h1>Profile</h1>
  <hr class="divider" />

  <div class="profile-actions">
    <Button @click="goToEdit" variant="primary">
      <ProfileIcon style="margin-right: 0.5em; vertical-align: middle;" />
      Edit Profile
    </Button>
  </div>
  <!-- todo: User Avatar -->
  <p>
    todo
  </p>
</Container>
</template>
