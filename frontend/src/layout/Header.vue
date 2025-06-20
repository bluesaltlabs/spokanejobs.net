<script setup>
import { RouterLink } from 'vue-router'

defineEmits(['toggle-sidebar'])
</script>

<template>
  <header class="header">
    <div class="header-content">
      <!-- Logo -->
      <div class="logo">
        <RouterLink to="/" class="logo-link">
          <div class="logo-placeholder">
            <span>BR</span>
          </div>
          <span class="logo-text">Bedrock</span>
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

      <!-- User Account Menu -->
      <div class="user-menu">
        <button class="user-menu-button" aria-label="User menu">
          <div class="user-avatar">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </div>
          <span class="user-name">User</span>
          <svg class="dropdown-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="6,9 12,15 18,9"></polyline>
          </svg>
        </button>
      </div>
    </div>
  </header>
</template>

<style scoped>
.header {
  background: #ffffff;
  border-bottom: 1px solid #e5e7eb;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
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
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
  font-size: 1.2rem;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: 700;
  color: #1f2937;
}

.sidebar-toggle {
  background: none;
  border: none;
  padding: 0.5rem;
  border-radius: 6px;
  cursor: pointer;
  color: #6b7280;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.sidebar-toggle:hover {
  background: #f3f4f6;
  color: #374151;
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
  border-radius: 6px;
  cursor: pointer;
  color: #6b7280;
  transition: all 0.2s ease;
}

.user-menu-button:hover {
  background: #f3f4f6;
  color: #374151;
}

.user-avatar {
  width: 32px;
  height: 32px;
  background: #f3f4f6;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
}

.user-name {
  font-weight: 500;
  color: #374151;
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
