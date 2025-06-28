<template>
  <component 
    :is="iconComponent" 
    v-if="iconComponent"
    :class="iconClass"
    :style="iconStyle"
  />
  <span v-else class="icon-placeholder">
    {{ slug }}
  </span>
</template>

<script>
import { computed } from 'vue'
import iconMap from '@/config/iconMap.js'

export default {
  name: 'Icon',
  props: {
    slug: {
      type: String,
      required: true,
      validator: (value) => {
        return Object.keys(iconMap).includes(value)
      }
    },
    size: {
      type: [String, Number],
      default: '20'
    },
    color: {
      type: String,
      default: 'currentColor'
    },
    class: {
      type: String,
      default: ''
    }
  },
  setup(props) {
    const iconComponent = computed(() => {
      return iconMap[props.slug] || null
    })

    const iconClass = computed(() => {
      return `icon icon-${props.slug} ${props.class}`.trim()
    })

    const iconStyle = computed(() => {
      return {
        width: typeof props.size === 'number' ? `${props.size}px` : props.size,
        height: typeof props.size === 'number' ? `${props.size}px` : props.size,
        color: props.color
      }
    })

    return {
      iconComponent,
      iconClass,
      iconStyle
    }
  }
}
</script>

<style scoped>
.icon {
  display: inline-block;
  vertical-align: middle;
}

.icon-placeholder {
  display: inline-block;
  padding: 2px 6px;
  background-color: #f0f0f0;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 12px;
  color: #666;
  font-family: monospace;
}
</style>
