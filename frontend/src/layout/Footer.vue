<script setup>
import IconAbout from '@/components/icons/About.vue'
import IconCompanies from '@/components/icons/Companies.vue'
import IconHome from '@/components/icons/Home.vue'
import IconJobs from '@/components/icons/Jobs.vue'
//import IconProfile from '@/components/icons/Profile.vue'

import { useRouter } from 'vue-router'
const router = useRouter()

function goHome() {
  router.push({ name: 'home' })
}
function goCompanies() {
  router.push({ name: 'companies-grid' })
}
function goJobs() {
  router.push({ name: 'jobs-list' })
}
function goAbout() {
  router.push({ name: 'about' })
}
</script>

<template>
  <footer class="mobile-footer">
    <nav class="footer-nav">
      <button class="btn btn-ghost footer-btn" @click="goHome">
        <IconHome />
        <span class="label">Home</span>
      </button>
      <button class="btn btn-ghost footer-btn" @click="goCompanies">
        <IconCompanies />
        <span class="label">Companies</span>
      </button>
      <button class="btn btn-ghost footer-btn" @click="goJobs">
        <IconJobs />
        <span class="label">Jobs</span>
      </button>
      <button class="btn btn-ghost footer-btn" @click="goAbout">
        <IconAbout />
        <span class="label">About</span>
      </button>
    </nav>
  </footer>
</template>


<style scoped>
.mobile-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: var(--color-background-elevated);
  border-top: 1px solid var(--color-border);
  z-index: 100;
  display: none;
}

.footer-nav {
  display: flex;
  justify-content: space-around;
  align-items: center;
  height: 56px;
}

.footer-btn {
  background: none;
  border: none;
  outline: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 14px;
  color: var(--color-text);
  flex: 1;
  padding: 0;
  transition: color 0.2s ease;
}

.footer-btn:hover {
  color: var(--color-primary-600);
}

.footer-btn .icon {
  font-size: 22px;
}

.footer-btn .label {
  font-size: 12px;
  margin-top: 2px;
}

@media (max-width: 1200px) {
  .mobile-footer {
    display: block;
  }
}
</style>
