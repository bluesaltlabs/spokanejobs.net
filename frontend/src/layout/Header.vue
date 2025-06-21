<script setup>
import { RouterLink } from 'vue-router'
import DarkModeToggle from '@/components/DarkModeToggle.vue'
import { useProfileStore } from '@/stores/profile'
import { onMounted } from 'vue'

const appName = import.meta.env.VITE_APP_NAME
const appAbbr = import.meta.env.VITE_APP_ABBR

// oh so fancy!
// todo: but this should be in the store and not here.
defineEmits(['toggle-sidebar'])

const profile = useProfileStore()
onMounted(() => {
  profile.loadProfile()
})
</script>

<template>
  <header class="header">
    <div class="header-content">
      <div class="header-content__left">

        <!-- Logo -->
        <div class="logo">
          <RouterLink to="/" class="logo-link">
            <div class="logo-placeholder">
              <span>{{ appAbbr }}</span>
            </div>
            <span class="logo-text">{{ appName }}</span>
          </RouterLink>
        </div>

        <!-- Sidebar Toggle Button -->
        <button
          class="sidebar-toggle"
          @click="$emit('toggle-sidebar')"
          aria-label="Toggle sidebar"
        >
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="3" y1="6" x2="21" y2="6"></line>
            <line x1="3" y1="12" x2="21" y2="12"></line>
            <line x1="3" y1="18" x2="21" y2="18"></line>
          </svg>
        </button>
      </div>

      <div class="header-content__right">
        <!-- Dark Mode Toggle -->
        <DarkModeToggle />

        <!-- User Account Menu -->
        <div class="user-menu">
          <button class="user-menu-button" aria-label="User menu">
            <div class="user-avatar">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
            </div>
            <span class="user-name">
              {{ profile.first_name || profile.last_name ? `${profile.first_name} ${profile.last_name}`.trim() : profile.email ? `${profile.email}` : 'User' }}
            </span>
            <svg class="dropdown-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="6,9 12,15 18,9"></polyline>
            </svg>
          </button>
        </div>
      </div>

    </div>
  </header>
</template>

<style scoped>
.header {
  background: var(--color-background-elevated);
  border-bottom: 1px solid var(--color-border);
  box-shadow: 0 1px 3px 0 var(--color-shadow);
  position: sticky;
  top: 0;
  z-index: 100;
  height: 64px;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 1.5rem;
  max-width: 100%;

  .header-content__left,
  .header-content__right {
    white-space: nowrap;
    display: flex;
    flex-wrap: nowrap;
    flex: grow;
  }
}

.logo {
  display: flex;
  align-items: center;
}

.logo-link {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: inherit;
  gap: 0.75rem;
}

.logo-placeholder {
  width: 40px;
  height: 40px;
  background: var(--gradient-primary);
  border-radius: var(--border-radius-medium);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-inverse);
  font-weight: bold;
  font-size: 1.2rem;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--color-heading);
}

.sidebar-toggle {
  background: none;
  border: none;
  padding: 0.5rem;
  border-radius: var(--border-radius-medium);
  cursor: pointer;
  color: var(--color-text-subtle);
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.sidebar-toggle:hover {
  background: var(--color-surface-hover);
  color: var(--color-text);
}

.user-menu {
  display: flex;
  align-items: center;
}

.user-menu-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: none;
  border: none;
  padding: 0.5rem 0.75rem;
  border-radius: var(--border-radius-medium);
  cursor: pointer;
  color: var(--color-text-subtle);
  transition: all 0.2s ease;
}

.user-menu-button:hover {
  background: var(--color-surface-hover);
  color: var(--color-text);
}

.user-avatar {
  width: 32px;
  height: 32px;
  background: var(--color-surface-hover);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-subtle);
}

.user-name {
  font-weight: 500;
  color: var(--color-text);
}

.dropdown-arrow {
  transition: transform 0.2s ease;
}

.user-menu-button:hover .dropdown-arrow {
  transform: rotate(180deg);
}

@media (max-width: 768px) {
  .logo-text {
    display: none;
  }

  .user-name {
    display: none;
  }

  .header-content {
    padding: 0 1rem;
  }
}
</style>
