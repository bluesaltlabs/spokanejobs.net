<script setup>
import { RouterLink } from 'vue-router'

defineProps({
  isOpen: {
    type: Boolean,
    default: false
  }
})
</script>

<template>
  <aside class="sidebar" :class="{ 'sidebar-open': isOpen }">
    <nav class="sidebar-nav">
      <div class="nav-section">
        <h3 class="nav-title">Main</h3>
        <ul class="nav-list">
          <li>
            <RouterLink to="/" class="nav-link">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                <polyline points="9,22 9,12 15,12 15,22"></polyline>
              </svg>
              <span>Home</span>
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/about" class="nav-link">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"></circle>
                <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"></path>
                <line x1="12" y1="17" x2="12.01" y2="17"></line>
              </svg>
              <span>About</span>
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/profile" class="nav-link">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
              <span>Profile</span>
            </RouterLink>
          </li>
        </ul>
      </div>

      <div class="nav-section">
        <h3 class="nav-title">Data</h3>
        <ul class="nav-list">
          <li>
            <RouterLink to="/companies" class="nav-link">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                <polyline points="9,22 9,12 15,12 15,22"></polyline>
              </svg>
              <span>Companies</span>
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/jobs" class="nav-link">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="2" y="7" width="20" height="14" rx="2" ry="2"></rect>
                <path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"></path>
              </svg>
              <span>Jobs</span>
            </RouterLink>
          </li>
        </ul>
      </div>

      <!-- Future sections - commented out until routes are created -->
      <!--
      <div class="nav-section">
        <h3 class="nav-title">Additional Data</h3>
        <ul class="nav-list">
          <li>
            <RouterLink to="/contacts" class="nav-link">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                <circle cx="9" cy="7" r="4"></circle>
                <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
              </svg>
              <span>Contacts</span>
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/locations" class="nav-link">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path>
                <circle cx="12" cy="10" r="3"></circle>
              </svg>
              <span>Locations</span>
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/categories" class="nav-link">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="8" y1="6" x2="21" y2="6"></line>
                <line x1="8" y1="12" x2="21" y2="12"></line>
                <line x1="8" y1="18" x2="21" y2="18"></line>
                <line x1="3" y1="6" x2="3.01" y2="6"></line>
                <line x1="3" y1="12" x2="3.01" y2="12"></line>
                <line x1="3" y1="18" x2="3.01" y2="18"></line>
              </svg>
              <span>Categories</span>
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/tags" class="nav-link">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z"></path>
                <line x1="7" y1="7" x2="7.01" y2="7"></line>
              </svg>
              <span>Tags</span>
            </RouterLink>
          </li>
        </ul>
      </div>
      -->
    </nav>
  </aside>
</template>

<style scoped>
.sidebar {
  width: 280px;
  background: --color-background;
  border-right: 1px solid #e5e7eb;
  height: 100%;
  transform: translateX(-100%);
  transition: transform 0.3s ease;
  position: relative;
  z-index: 50;
}

.sidebar-open {
  transform: translateX(0);
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
  color: #6b7280;
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
  color: #374151;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.2s ease;
  border-left: 3px solid transparent;
}

.nav-link:hover {
  background: #01f5f9;
  color: #1f2937;
  border-left-color: #d1d5db;
}

.nav-link.router-link-active {
  background: #0ff6ff;
  color: #2563eb;
  border-left-color: #2563eb;
}

.nav-link svg {
  flex-shrink: 0;
}

.nav-link span {
  white-space: nowrap;
}

@media (max-width: 768px) {
  .sidebar {
    position: fixed;
    top: 64px;
    left: 0;
    height: calc(100vh - 64px);
    box-shadow: 2px 0 8px rgba(0, 0, 0, 0.1);
  }
}

@media (min-width: 769px) {
  .sidebar {
    transform: translateX(0);
  }
}
</style>
