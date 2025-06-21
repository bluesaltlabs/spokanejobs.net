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
  background: linear-gradient(90deg, var(--button-primary-start, #4f8cff) 0%, var(--button-primary-end, #2356c7) 100%);
  color: var(--button-primary-text, #fff);
  box-shadow: 0 1px 4px var(--button-primary-shadow, rgba(79,140,255,0.08));
}

.base-button--primary:hover:not(.base-button--disabled) {
  background: linear-gradient(90deg, var(--button-primary-end, #2356c7) 0%, var(--button-primary-start, #4f8cff) 100%);
  box-shadow: 0 2px 8px var(--button-primary-shadow-hover, rgba(79,140,255,0.15));
  transform: translateY(-1px);
}

.base-button--primary:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--button-primary-focus, rgba(79,140,255,0.25));
}

/* Secondary variant */
.base-button--secondary {
  background: var(--button-secondary-bg, #6c757d);
  color: var(--button-secondary-text, #fff);
  box-shadow: 0 1px 4px var(--button-secondary-shadow, rgba(108,117,125,0.08));
}

.base-button--secondary:hover:not(.base-button--disabled) {
  background: var(--button-secondary-hover, #5a6268);
  box-shadow: 0 2px 8px var(--button-secondary-shadow-hover, rgba(108,117,125,0.15));
  transform: translateY(-1px);
}

.base-button--secondary:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--button-secondary-focus, rgba(108,117,125,0.25));
}

/* Danger variant */
.base-button--danger {
  background: var(--button-danger-bg, #e74c3c);
  color: var(--button-danger-text, #fff);
  box-shadow: 0 1px 4px var(--button-danger-shadow, rgba(231,76,60,0.08));
}

.base-button--danger:hover:not(.base-button--disabled) {
  background: var(--button-danger-hover, #c0392b);
  box-shadow: 0 2px 8px var(--button-danger-shadow-hover, rgba(231,76,60,0.15));
  transform: translateY(-1px);
}

.base-button--danger:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--button-danger-focus, rgba(231,76,60,0.25));
}

/* Themed variant (gradient) */
.base-button--themed {
  background: linear-gradient(90deg, var(--button-themed-start, #4f8cff) 0%, var(--button-themed-end, #2356c7) 100%);
  color: var(--button-themed-text, #fff);
  box-shadow: 0 1px 4px var(--button-themed-shadow, rgba(79,140,255,0.08));
}

.base-button--themed:hover:not(.base-button--disabled) {
  background: linear-gradient(90deg, var(--button-themed-end, #2356c7) 0%, var(--button-themed-start, #4f8cff) 100%);
  box-shadow: 0 2px 8px var(--button-themed-shadow-hover, rgba(79,140,255,0.15));
  transform: translateY(-1px);
}

.base-button--themed:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--button-themed-focus, rgba(79,140,255,0.25));
}

/* Outline variant */
.base-button--outline {
  background: transparent;
  color: var(--button-outline-text, #4f8cff);
  border: 2px solid var(--button-outline-border, #4f8cff);
  box-shadow: none;
}

.base-button--outline:hover:not(.base-button--disabled) {
  background: var(--button-outline-hover-bg, #4f8cff);
  color: var(--button-outline-hover-text, #fff);
  transform: translateY(-1px);
  box-shadow: 0 2px 8px var(--button-outline-shadow-hover, rgba(79,140,255,0.15));
}

.base-button--outline:focus {
  outline: none;
  box-shadow: 0 0 0 3px var(--button-outline-focus, rgba(79,140,255,0.25));
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

/* Dark theme overrides */
[data-theme="dark"] .base-button--primary {
  --button-primary-start: #7daaff;
  --button-primary-end: #4f8cff;
  --button-primary-shadow: rgba(125,170,255,0.15);
  --button-primary-shadow-hover: rgba(125,170,255,0.25);
  --button-primary-focus: rgba(125,170,255,0.35);
}

[data-theme="dark"] .base-button--secondary {
  --button-secondary-bg: #495057;
  --button-secondary-hover: #343a40;
  --button-secondary-shadow: rgba(73,80,87,0.15);
  --button-secondary-shadow-hover: rgba(73,80,87,0.25);
  --button-secondary-focus: rgba(73,80,87,0.35);
}

[data-theme="dark"] .base-button--danger {
  --button-danger-bg: #dc3545;
  --button-danger-hover: #c82333;
  --button-danger-shadow: rgba(220,53,69,0.15);
  --button-danger-shadow-hover: rgba(220,53,69,0.25);
  --button-danger-focus: rgba(220,53,69,0.35);
}

[data-theme="dark"] .base-button--themed {
  --button-themed-start: #7daaff;
  --button-themed-end: #4f8cff;
  --button-themed-shadow: rgba(125,170,255,0.15);
  --button-themed-shadow-hover: rgba(125,170,255,0.25);
  --button-themed-focus: rgba(125,170,255,0.35);
}

[data-theme="dark"] .base-button--outline {
  --button-outline-text: #7daaff;
  --button-outline-border: #7daaff;
  --button-outline-hover-bg: #7daaff;
  --button-outline-hover-text: #1a1a1a;
  --button-outline-shadow-hover: rgba(125,170,255,0.25);
  --button-outline-focus: rgba(125,170,255,0.35);
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