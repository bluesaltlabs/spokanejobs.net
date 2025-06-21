<script setup>
import { ref, onMounted } from 'vue'
import { useProfileStore } from '@/stores/profile'

const profile = useProfileStore()
const saved = ref(false)

const save = async () => {
  await profile.saveProfile()
  saved.value = true
  setTimeout(() => (saved.value = false), 1500)
}

onMounted(() => {
  profile.loadProfile()
})
</script>

<style scoped>
.profile-view {
  max-width: 400px;
  margin: 0 auto;
  padding: 2rem;
  background: var(--background-color);
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}
.profile-view h1 {
  margin-bottom: 1.5rem;
}
.profile-view form > div {
  margin-bottom: 1rem;
}
.profile-view label {
  display: block;
  margin-bottom: 0.25rem;
}
.profile-view input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.profile-view .avatar {
  width: 100px;
  height: 100px;
  background: var(--color-surface-hover);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-subtle);

  img {
    max-width: 100px;
    max-height: 100px;
    border-radius: var(--border-radius-small);
    display: block;
  }
}
</style>

<template>
  <div class="profile-view">
    <h1>User Profile</h1>
    <form @submit.prevent="save">
      <div>
        <label>First Name:</label>
        <input v-model="profile.first_name" type="text" />
      </div>
      <div>
        <label>Last Name:</label>
        <input v-model="profile.last_name" type="text" />
      </div>
      <div>
        <label>Email:</label>
        <input v-model="profile.email" type="email" />
      </div>
      <div>
        <label>Avatar URL:</label>
        <input v-model="profile.avatar" type="text" />
      </div>
      <div class="avatar">
        <img v-if="profile.avatar" :src="profile.avatar" alt="Avatar" />
      </div>
      <button type="submit">Save</button>
      <span v-if="saved" style="margin-left:1rem;color:green;">Saved!</span>
    </form>
  </div>
</template>
