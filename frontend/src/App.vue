<script setup>
import { ref } from 'vue'
import { RouterView } from 'vue-router'
import Header from '@/layout/Header.vue'
import Sidebar from '@/layout/Sidebar.vue'
import Footer from '@/layout/Footer.vue'

const isSidebarOpen = ref(false)

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}
</script>

<template>
  <div class="app-layout">
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
