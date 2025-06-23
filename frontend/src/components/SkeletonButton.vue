<template>
  <div 
    class="skeleton-button" 
    :class="{ 'skeleton-button--animated': animated }"
    :style="{ 
      width: width,
      height: height,
      borderRadius: borderRadius
    }"
  ></div>
</template>

<script setup>
defineProps({
  width: {
    type: String,
    default: '120px'
  },
  height: {
    type: String,
    default: '40px'
  },
  borderRadius: {
    type: String,
    default: 'var(--border-radius-small)'
  },
  animated: {
    type: Boolean,
    default: true
  }
})
</script>

<style scoped>
.skeleton-button {
  background: var(--color-background-mute);
  opacity: 0.7;
  display: inline-block;
}

.skeleton-button--animated {
  animation: skeleton-pulse 1.5s ease-in-out infinite;
}

@keyframes skeleton-pulse {
  0%, 100% {
    opacity: 0.7;
  }
  50% {
    opacity: 0.4;
  }
}
</style> 