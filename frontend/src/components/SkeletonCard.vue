<template>
  <div class="skeleton-card" :class="{ 'skeleton-card--animated': animated }">
    <div v-if="showImage" class="skeleton-card-image">
      <SkeletonImage :width="imageWidth" :height="imageHeight" />
    </div>
    
    <div class="skeleton-card-content">
      <SkeletonText 
        v-if="showTitle" 
        variant="subtitle" 
        :lines="1" 
        :line-height="24"
      />
      
      <SkeletonText 
        v-if="showDescription" 
        :lines="descriptionLines" 
        :line-height="16"
        :last-line-width="lastLineWidth"
      />
      
      <div v-if="showButtons" class="skeleton-card-actions">
        <SkeletonButton 
          v-for="i in buttonCount" 
          :key="i"
          :width="buttonWidths[i - 1] || '80px'"
          height="32px"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import SkeletonText from './SkeletonText.vue'
import SkeletonImage from './SkeletonImage.vue'
import SkeletonButton from './SkeletonButton.vue'

defineProps({
  showImage: {
    type: Boolean,
    default: true
  },
  imageWidth: {
    type: String,
    default: '100%'
  },
  imageHeight: {
    type: String,
    default: '200px'
  },
  showTitle: {
    type: Boolean,
    default: true
  },
  showDescription: {
    type: Boolean,
    default: true
  },
  descriptionLines: {
    type: Number,
    default: 3
  },
  lastLineWidth: {
    type: String,
    default: '60%'
  },
  showButtons: {
    type: Boolean,
    default: false
  },
  buttonCount: {
    type: Number,
    default: 2
  },
  buttonWidths: {
    type: Array,
    default: () => ['80px', '100px']
  },
  animated: {
    type: Boolean,
    default: true
  }
})
</script>

<style scoped>
.skeleton-card {
  background: var(--color-background-soft);
  border-radius: var(--border-radius);
  padding: 20px;
  border: 1px solid var(--color-border);
}

.skeleton-card-image {
  margin-bottom: 16px;
}

.skeleton-card-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.skeleton-card-actions {
  display: flex;
  gap: 8px;
  margin-top: 8px;
}

.skeleton-card--animated {
  animation: skeleton-card-pulse 1.5s ease-in-out infinite;
}

@keyframes skeleton-card-pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.8;
  }
}
</style> 