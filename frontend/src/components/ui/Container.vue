<script setup>
defineProps({
  // Optional size prop for different container widths
  size: {
    type: String,
    default: 'default',
    validator: (value) => ['small', 'default', 'large', 'full'].includes(value)
  },
  // Optional padding prop
  padding: {
    type: String,
    default: 'default',
    validator: (value) => ['none', 'small', 'default', 'large'].includes(value)
  }
})
</script>

<template>
  <div class="container" :class="[`container--${size}`, `container--padding-${padding}`]">
    <slot />
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  margin: 0 auto;
}

/* Container sizes */
.container--small {
  max-width: 640px; /* sm */
}

.container--default {
  max-width: 1200px; /* lg */
}

.container--large {
  max-width: 1400px; /* xl */
}

.container--full {
  max-width: none;
}

/* Padding variants */
.container--padding-none {
  padding: 0;
}

.container--padding-small {
  padding: 1rem;
}

.container--padding-default {
  padding: 2rem;
}

.container--padding-large {
  padding: 3rem;
}

/* Responsive adjustments */
@media (max-width: 640px) {
  .container--padding-default {
    padding: 1rem;
  }
  
  .container--padding-large {
    padding: 1.5rem;
  }
}

@media (min-width: 641px) and (max-width: 1024px) {
  .container--padding-default {
    padding: 1.5rem;
  }
  
  .container--padding-large {
    padding: 2rem;
  }
}

@media (min-width: 1025px) {
  .container--default {
    padding-left: 2rem;
    padding-right: 2rem;
  }
  
  .container--large {
    padding-left: 2rem;
    padding-right: 2rem;
  }
}
</style> 