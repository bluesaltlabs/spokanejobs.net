<template>
  <div 
    class="skeleton-image" 
    :class="{ 'skeleton-image--animated': animated }"
    :style="{ 
      width: width,
      height: height,
      borderRadius: borderRadius
    }"
  >
    <div class="skeleton-image-placeholder">
      <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M21 19V5C21 3.9 20.1 3 19 3H5C3.9 3 3 3.9 3 5V19C3 20.1 3.9 21 5 21H19C20.1 21 21 20.1 21 19ZM8.5 13.5L11 16.51L14.5 12L19 18H5L8.5 13.5Z" fill="currentColor"/>
      </svg>
    </div>
  </div>
</template>

<script setup>
defineProps({
  width: {
    type: String,
    default: '100%'
  },
  height: {
    type: String,
    default: '200px'
  },
  borderRadius: {
    type: String,
    default: 'var(--border-radius)'
  },
  animated: {
    type: Boolean,
    default: true
  }
})
</script>

<style scoped>
.skeleton-image {
  background: var(--color-background-mute);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0.7;
  position: relative;
  overflow: hidden;
}

.skeleton-image--animated {
  animation: skeleton-pulse 1.5s ease-in-out infinite;
}

.skeleton-image-placeholder {
  color: var(--color-text-muted);
  opacity: 0.5;
}

.skeleton-image-placeholder svg {
  width: 48px;
  height: 48px;
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