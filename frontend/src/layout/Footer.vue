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
  background: #fff;
  border-top: 1px solid #eee;
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
  color: #333;
  flex: 1;
  padding: 0;
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