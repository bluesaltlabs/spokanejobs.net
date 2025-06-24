<script setup>
import { RouterLink } from 'vue-router'
import IconAbout from '@/components/icons/About.vue'
import IconCompanies from '@/components/icons/Companies.vue'
import IconHome from '@/components/icons/Home.vue'
import IconJobs from '@/components/icons/Jobs.vue'
import IconProfile from '@/components/icons/Profile.vue'
</script>

<template>
  <aside class="sidebar">
    <nav class="sidebar-nav">
      <div class="nav-section">
        <h3 class="nav-title">Main</h3>
        <ul class="nav-list">
          <li>
            <RouterLink to="/" class="nav-link">
              <IconHome />
              <span>Home</span>
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/about" class="nav-link">
              <IconAbout />
              <span>About</span>
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/profile" class="nav-link">
              <IconProfile />
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
              <IconCompanies />
              <span>Companies</span>
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/jobs" class="nav-link">
              <IconJobs />
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
