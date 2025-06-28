<script setup>
const appName = import.meta.env.VITE_APP_NAME
import { Container } from '@/components/ui';
</script>

<template>
  <Container>
    <h1>Resources</h1>
    <hr class="divider" />

    <p>
      todo
    </p>
  </Container>
</template>

<style scoped>
.divider {
  margin: 1.5rem 0 2rem 0;
  border: none;
  border-top: 1px solid var(--color-border);
}
</style>
