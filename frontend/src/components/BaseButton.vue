<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'primary',
    validator: (value) => ['primary', 'secondary', 'danger', 'themed', 'outline'].includes(value)
  },
  size: {
    type: String,
    default: 'medium',
    validator: (value) => ['small', 'medium', 'large'].includes(value)
  },
  disabled: {
    type: Boolean,
    default: false
  },
  type: {
    type: String,
    default: 'button',
    validator: (value) => ['button', 'submit', 'reset'].includes(value)
  }
})

const emit = defineEmits(['click'])

const buttonClasses = computed(() => {
  return [
    'base-button',
    `base-button--${props.variant}`,
    `base-button--${props.size}`,
    {
      'base-button--disabled': props.disabled
    }
  ]
})

const handleClick = (event) => {
  if (!props.disabled) {
    emit('click', event)
  }
}
</script>

<style scoped>
.base-button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border: none;
  border-radius: var(--border-radius-small, 6px);
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  text-decoration: none;
  font-family: inherit;
  line-height: 1;
  white-space: nowrap;
  user-select: none;
  position: relative;
  overflow: hidden;
}

/* Size variants */
.base-button--small {
  padding: 0.25rem 0.75rem;
  font-size: 0.875rem;
  min-height: 32px;
}

.base-button--medium {
  padding: 0.5rem 1.2rem;
  font-size: 1rem;
  min-height: 40px;
}

.base-button--large {
  padding: 0.75rem 1.5rem;
  font-size: 1.125rem;
  min-height: 48px;
}

/* Primary variant */
.base-button--primary {
  background: linear-gradient(90deg, var(--color-primary-500) 0%, var(--color-primary-700) 100%);
  color: var(--color-button-primary-text);
  box-shadow: 0 1px 4px var(--color-button-primary-shadow);
}

.base-button--primary:hover:not(.base-button--disabled) {
  background: linear-gradient(90deg, var(--color-primary-700) 0%, var(--color-primary-500) 100%);
  box-shadow: 0 2px 8px var(--color-button-primary-shadow);
  transform: translateY(-1px);
}

.base-button--primary:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--color-border-focus);
}

/* Secondary variant */
.base-button--secondary {
  background: var(--color-button-secondary);
  color: var(--color-button-secondary-text);
  box-shadow: 0 1px 4px var(--color-button-secondary-shadow);
}

.base-button--secondary:hover:not(.base-button--disabled) {
  background: var(--color-button-secondary-hover);
  box-shadow: 0 2px 8px var(--color-button-secondary-shadow);
  transform: translateY(-1px);
}

.base-button--secondary:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--color-border-focus);
}

/* Danger variant */
.base-button--danger {
  background: var(--color-button-danger);
  color: var(--color-button-danger-text);
  box-shadow: 0 1px 4px var(--color-button-danger-shadow);
}

.base-button--danger:hover:not(.base-button--disabled) {
  background: var(--color-button-danger-hover);
  box-shadow: 0 2px 8px var(--color-button-danger-shadow);
  transform: translateY(-1px);
}

.base-button--danger:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--color-border-focus);
}

/* Themed variant (gradient) */
.base-button--themed {
  background: linear-gradient(90deg, var(--color-primary-500) 0%, var(--color-primary-700) 100%);
  color: var(--color-button-primary-text);
  box-shadow: 0 1px 4px var(--color-button-primary-shadow);
}

.base-button--themed:hover:not(.base-button--disabled) {
  background: linear-gradient(90deg, var(--color-primary-700) 0%, var(--color-primary-500) 100%);
  box-shadow: 0 2px 8px var(--color-button-primary-shadow);
  transform: translateY(-1px);
}

.base-button--themed:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--color-border-focus);
}

/* Outline variant */
.base-button--outline {
  background: transparent;
  color: var(--color-primary-600);
  border: 2px solid var(--color-primary-600);
  box-shadow: none;
}

.base-button--outline:hover:not(.base-button--disabled) {
  background: var(--color-primary-600);
  color: var(--color-white);
  transform: translateY(-1px);
  box-shadow: 0 2px 8px var(--color-button-primary-shadow);
}

.base-button--outline:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--color-border-focus);
}

/* Disabled state */
.base-button--disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none !important;
  box-shadow: none !important;
}

.base-button--disabled:hover {
  transform: none !important;
  box-shadow: none !important;
}

/* Loading state */
.base-button--loading {
  position: relative;
  color: transparent;
}

.base-button--loading::after {
  content: '';
  position: absolute;
  width: 16px;
  height: 16px;
  top: 50%;
  left: 50%;
  margin-left: -8px;
  margin-top: -8px;
  border: 2px solid transparent;
  border-top: 2px solid currentColor;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* Icon support */
.base-button .icon {
  margin-right: 0.5rem;
}

.base-button .icon:last-child {
  margin-right: 0;
  margin-left: 0.5rem;
}
</style>

<template>
  <button
    :class="buttonClasses"
    :type="type"
    :disabled="disabled"
    @click="handleClick"
  >
    <slot />
  </button>
</template> 