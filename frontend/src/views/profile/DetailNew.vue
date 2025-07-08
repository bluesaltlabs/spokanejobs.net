<script setup>
import { onMounted } from 'vue'
import { useProfileStore } from '@/stores/profile'
import { Container, Button } from '@/components/ui'
import { EditIcon } from '@/components/icons'

const profile = useProfileStore()

onMounted(() => {
  profile.loadProfile()
})

function toggleEdit() {
  // todo: toggle edit mode.
}

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
    <Button @click="toggleEdit" variant="primary">
      <EditIcon />
    </Button>
  </div>
  <!-- todo: User Avatar -->
  <p>
    todo
  </p>
</Container>
</template>
