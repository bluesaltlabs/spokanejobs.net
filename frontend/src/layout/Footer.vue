<template>
  <footer class="mobile-footer">
    <nav class="footer-nav">
      <button class="footer-btn" @click="goHome">
        <span class="icon">🏠</span>
        <span class="label">Home</span>
      </button>
      <button class="footer-btn" @click="goSearch">
        <span class="icon">🔍</span>
        <span class="label">Search</span>
      </button>
      <button class="footer-btn" @click="goNotifications">
        <span class="icon">🔔</span>
        <span class="label">Alerts</span>
      </button>
      <button class="footer-btn" @click="goProfile">
        <span class="icon">👤</span>
        <span class="label">Profile</span>
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
function goSearch() {
  router.push({ name: 'search' })
}
function goNotifications() {
  router.push({ name: 'notifications' })
}
function goProfile() {
  router.push({ name: 'profile' })
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
