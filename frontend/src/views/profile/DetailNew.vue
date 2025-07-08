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

  <!-- todo: User Avatar -->


  <p>
    todo
  </p>
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
</style>
