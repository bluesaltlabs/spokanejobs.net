<template>
  <div class="item-empty-state">
    <slot name="icon" />
    <h3><slot name="title" /></h3>
    <p><slot name="desc" /></p>
  </div>
</template>

<script setup>
// No props needed
</script>

<style scoped>
.item-empty-state {
  text-align: center;
  color: var(--color-text-muted);
  margin: 2rem 0;
}
.item-empty-state h3 {
  margin: 1rem 0 0.5rem 0;
}
</style>
