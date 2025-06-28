<template>
  <component 
    :is="as" 
    :class="['item-card', customClass, { 'clickable': onClick }]" 
    @click="onClick"
  >
    <slot name="header" />
    <slot name="meta" />
    <slot name="actions" />
  </component>
</template>

<script setup>
const props = defineProps({
  as: { type: String, default: 'div' },
  customClass: { type: String, default: '' },
  onClick: { type: Function, default: null }
})
</script>

<style scoped>
.item-card {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-medium);
  overflow: hidden;
  transition: all 0.2s ease;
  box-shadow: 0 1px 3px var(--color-shadow);
  width: 100%;
  min-height: 120px;
  padding: 1rem;
  box-sizing: border-box;
}

.item-card.clickable {
  cursor: pointer;
}

.item-card.clickable:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px var(--color-shadow-elevated);
  border-color: var(--color-border-hover);
  background: var(--color-background-soft);
}
</style>
