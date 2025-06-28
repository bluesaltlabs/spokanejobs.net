<template>
  <tr class="item-row">
    <slot name="cells" />
    <slot name="actions" />
  </tr>
</template>

<script setup>
// No props needed for now
</script>

<style scoped>
.item-row {
  background: var(--color-surface);
  border-bottom: 1px solid var(--color-border);
  transition: background 0.2s;
}
.item-row:hover {
  background: var(--color-background-soft);
}
</style>
