<template>
  <form class="ui-form" @submit.prevent="$emit('submit', $event)">
    <slot />
  </form>
</template>

<style scoped>
.ui-form {
  width: 100%;
  max-width: 600px;
  margin: 0 auto;
}
</style> 