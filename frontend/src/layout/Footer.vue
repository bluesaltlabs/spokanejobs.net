<template>
  <footer class="mobile-footer">
    <nav class="footer-nav">
      <button class="footer-btn" @click="goHome">
        <span class="icon">🏠</span>
        <span class="label">Home</span>
      </button>
      <button class="footer-btn" @click="goCompanies">
        <span class="icon">🏢</span>
        <span class="label">Companies</span>
      </button>
      <button class="footer-btn" @click="goJobs">
        <span class="icon">💼</span>
        <span class="label">Jobs</span>
      </button>
      <button class="footer-btn" @click="goAbout">
        <span class="icon">ℹ️</span>
        <span class="label">About</span>
      </button>
    </nav>
  </footer>
</template>

<script setup>
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

@media (max-width: 768px) {
  .mobile-footer {
    display: block;
  }
}
</style>
