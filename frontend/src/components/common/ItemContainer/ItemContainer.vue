<template>
  <component :is="as" :class="[isGrid ? 'item-grid' : 'item-list', customClass]">
    <slot />
  </component>
</template>

<script setup>
import { computed } from 'vue'
const props = defineProps({
  as: { type: String, default: 'div' },
  grid: { type: Boolean, default: false },
  customClass: { type: String, default: '' }
})
const isGrid = computed(() => props.grid)
</script>

<style scoped>
.item-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 1rem;
  width: 100%;
}
.item-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  width: 100%;
}
</style>
