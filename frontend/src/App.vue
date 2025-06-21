<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { RouterView } from 'vue-router'
import Header from '@/layout/Header.vue'
import Sidebar from '@/layout/Sidebar.vue'
import Footer from '@/layout/Footer.vue'
import { useProfileStore } from '@/stores/profile'

const isSidebarOpen = ref(false)
const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}

const profile = useProfileStore()
const themeClass = computed(() => (profile.dark_mode ? 'dark' : 'light'))

const applyDarkMode = (val) => {
  if (val) {
    document.documentElement.classList.add('dark')
    document.documentElement.classList.remove('light')
    localStorage.setItem('theme', 'dark')
  } else {
    document.documentElement.classList.remove('dark')
    document.documentElement.classList.add('light')
    localStorage.setItem('theme', 'light')
  }
}

onMounted(async () => {
  // Load profile data first
  await profile.loadProfile()
  
  // Apply the saved theme
  applyDarkMode(profile.dark_mode)
})

watch(() => profile.dark_mode, (val) => {
  applyDarkMode(val)
})
</script>

<template>
  <div :class="['app-layout', themeClass]">
    <!-- todo: update this so it uses state pinia app state store instead of this inline thing  -->
    <Header @toggle-sidebar="toggleSidebar" />

    <div class="main-content">
      <Sidebar :is-open="isSidebarOpen" />
      <main class="content-area">
        <RouterView />
      </main>
    </div>
  </div>

  <Footer />
</template>

<style scoped>
.app-layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.main-content {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.content-area {
  flex: 1;
  overflow-y: auto;
  padding: 2rem;
}
</style>
