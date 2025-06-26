<script setup>
import navLinks from '@/config/navLinks.json'
import iconMap from '@/config/iconMap.js'
import { useRouter, useRoute } from 'vue-router'
import { computed } from 'vue'

const router = useRouter()
const route = useRoute()

// Flatten nav items for the footer (no sections)
const footerLinks = computed(() => {
  return navLinks.flatMap(section => section.items)
})

function goTo(routeName) {
  router.push({ name: routeName })
}
</script>

<template>
  <footer class="mobile-footer">
    <nav class="footer-nav">
      <button
        v-for="item in footerLinks"
        :key="item.route"
        class="btn btn-ghost footer-btn"
        :class="{ active: route.name === item.route }"
        @click="goTo(item.route)"
      >
        <component :is="iconMap[item.icon]" />
        <span class="label">{{ item.label }}</span>
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
  justify-content: stretch;
  align-items: stretch;
  height: 56px;
  padding: 0;
}

.footer-btn {
  background: none;
  border: none;
  outline: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  color: var(--color-text);
  flex: 1 1 0%;
  min-width: 0;
  min-height: 0;
  height: 100%;
  width: 100%;
  padding: 0;
  margin: 0;
  transition: color 0.2s ease, background 0.2s ease;
  border-radius: 0;
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

.footer-btn.active {
  background: var(--color-primary-50);
  color: var(--color-primary-600);
  border-top: 3px solid var(--color-primary-600);
}

@media (max-width: 1200px) {
  .mobile-footer {
    display: block;
  }
}
</style>
