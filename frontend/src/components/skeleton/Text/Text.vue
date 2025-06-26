<script setup>
defineProps({
  lines: {
    type: Number,
    default: 1
  },
  lineHeight: {
    type: Number,
    default: 16
  },
  variant: {
    type: String,
    default: 'default', // default, title, subtitle
    validator: (value) => ['default', 'title', 'subtitle'].includes(value)
  },
  lastLineWidth: {
    type: String,
    default: null
  },
  animated: {
    type: Boolean,
    default: true
  }
})
</script>

<template>
  <div class="skeleton-text" :class="[`skeleton-text--${variant}`, { 'skeleton-text--animated': animated }]">
    <div
      v-for="i in lines"
      :key="i"
      class="skeleton-line"
      :style="{
        width: i === lines && lastLineWidth ? lastLineWidth : '100%',
        height: `${lineHeight}px`
      }"
    ></div>
  </div>
</template>

<style scoped>
.skeleton-text {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.skeleton-line {
  background: var(--color-background-mute);
  border-radius: 4px;
  opacity: 0.7;
}

.skeleton-text--animated .skeleton-line {
  animation: skeleton-pulse 1.5s ease-in-out infinite;
}

.skeleton-text--title .skeleton-line {
  height: 32px !important;
  border-radius: 6px;
}

.skeleton-text--subtitle .skeleton-line {
  height: 24px !important;
  border-radius: 5px;
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
