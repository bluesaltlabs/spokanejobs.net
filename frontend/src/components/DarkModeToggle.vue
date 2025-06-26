<script setup>
import { computed } from 'vue'
import { useProfileStore } from '@/stores/profile'
import { SunIcon, MoonIcon } from '@/components/icons'

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
    <SunIcon v-if="!isDark" width="20" height="20" />
    <MoonIcon v-else width="20" height="20" />
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
