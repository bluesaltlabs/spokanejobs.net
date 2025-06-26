<script setup>
import navLinks from '@/config/navLinks.json'
import iconMap from '@/config/iconMap.js'
import { RouterLink } from 'vue-router'

</script>

<template>
  <aside class="sidebar">
    <nav class="sidebar-nav">
      <div v-for="section in navLinks" :key="section.section" class="nav-section">
        <h3 class="nav-title">{{ section.section }}</h3>
        <ul class="nav-list">
          <li v-for="item in section.items" :key="item.route">
            <RouterLink :to="{ name: item.route }" class="nav-link">
              <component :is="iconMap[item.icon]" />
              <span>{{ item.label }}</span>
            </RouterLink>
          </li>
        </ul>
      </div>
    </nav>
  </aside>
</template>

<style scoped>
.sidebar {
  width: 280px;
  background: var(--color-background-soft);
  border-right: 1px solid var(--color-border);
  height: 100%;
  transform: translateX(-100%);
  transition: transform 0.3s ease;
  position: relative;
  z-index: 50;
}


.sidebar-nav {
  padding: 1.5rem 0;
  height: 100%;
  overflow-y: auto;
}

.nav-section {
  margin-bottom: 2rem;
}

.nav-section:last-child {
  margin-bottom: 0;
}

.nav-title {
  font-size: 0.75rem;
  font-weight: 600;
  color: var(--color-text-subtle);
  text-transform: uppercase;
  letter-spacing: 0.05em;
  padding: 0 1.5rem;
  margin-bottom: 0.75rem;
}

.nav-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1.5rem;
  color: var(--color-text);
  text-decoration: none;
  font-weight: 500;
  transition: all 0.2s ease;
  border-left: 3px solid transparent;
}

.nav-link:hover {
  background: var(--color-surface-hover);
  color: var(--color-text);
  border-left-color: var(--color-border-hover);
}

.nav-link.router-link-active {
  background: var(--color-primary-50);
  color: var(--color-primary-600);
  border-left-color: var(--color-primary-600);
}

.nav-link svg {
  flex-shrink: 0;
}

.nav-link span {
  white-space: nowrap;
}

@media (max-width: 1200px) {
  .sidebar {
    position: fixed;
    top: 64px;
    left: 0;
    height: calc(100vh - 64px);
    box-shadow: 2px 0 8px var(--color-shadow-elevated);
  }
}

@media (min-width: 1201px) {
  .sidebar {
    transform: translateX(0);
  }
}
</style>
