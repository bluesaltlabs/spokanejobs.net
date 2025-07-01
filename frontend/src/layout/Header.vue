<script setup>
import { RouterLink } from 'vue-router'
import { DarkModeToggle } from '@/components/common'
import { ProfileIcon } from '@/components/icons'
import { useProfileStore } from '@/stores/profile'

const appName = import.meta.env.VITE_APP_NAME
const appAbbr = import.meta.env.VITE_APP_ABBR

const profile = useProfileStore()
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
      </div>

      <div class="header-content__right">
        <!-- Dark Mode Toggle -->
        <DarkModeToggle />

        <!-- User Account Menu -->
        <div class="user-menu">
          <RouterLink to="/profile" class="btn btn-ghost user-menu-button" aria-label="User menu">
            <div class="user-avatar">
              <div v-if="profile.avatar">
                <img :src="profile.avatar" alt="Avatar" />
              </div>
              <ProfileIcon v-if="!profile.avatar" />
            </div>
            <span class="user-name">
              {{ profile.first_name || profile.last_name ? `${profile.first_name} ${profile.last_name}`.trim() : profile.email ? `${profile.email}` : 'User' }}
            </span>
          </RouterLink>
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

  img {
    max-width: 32px;
    max-height: 32px;
    border-radius: var(--border-radius-small);
  }
}

.user-name {
  font-weight: 500;
  color: var(--color-text);
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
