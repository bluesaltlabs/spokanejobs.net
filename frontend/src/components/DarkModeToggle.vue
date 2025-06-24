<script setup>
import { computed } from 'vue'
import { useProfileStore } from '@/stores/profile'
import BaseButton from '@/components/BaseButton.vue'

const profile = useProfileStore()

const isDark = computed({
  get: () => profile.dark_mode,
  set: (val) => profile.setDarkMode(val)
})

const toggleDarkMode = () => {
  isDark.value = !isDark.value
}
</script>

<template>
  <button
    @click="toggleDarkMode"
    class="btn btn-ghost dark-mode-toggle"
    :title="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
  >
    <svg v-if="!isDark" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
      <circle cx="12" cy="12" r="5"></circle>
      <line x1="12" y1="1" x2="12" y2="3"></line>
      <line x1="12" y1="21" x2="12" y2="23"></line>
      <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
      <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
      <line x1="1" y1="12" x2="3" y2="12"></line>
      <line x1="21" y1="12" x2="23" y2="12"></line>
      <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
      <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
    </svg>
    <svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
      <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
    </svg>
  </button>
</template>

<style scoped>

/* Base Button Styles */
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: var(--button-padding-medium);
  font-size: var(--button-font-size-medium);
  font-weight: var(--button-font-weight);
  line-height: 1.5;
  text-align: center;
  text-decoration: none;
  vertical-align: middle;
  cursor: pointer;
  user-select: none;
  border: 1px solid transparent;
  border-radius: var(--border-radius-small);
  transition: all 0.2s ease-in-out;
  box-shadow: 0 1px 3px var(--color-shadow);
  min-height: 2.5rem;
}

.btn:focus {
  outline: 2px solid var(--color-border-focus);
  outline-offset: 2px;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  pointer-events: none;
}


.btn-ghost {
  background-color: var(--color-button-ghost);
  border-color: var(--color-button-ghost-border);
  color: var(--color-button-ghost-text);
  box-shadow: var(--color-button-ghost-shadow);
}

.btn-ghost:hover:not(:disabled) {
  background-color: var(--color-button-ghost-hover);
  border-color: var(--color-button-ghost-hover);
}

</style>
