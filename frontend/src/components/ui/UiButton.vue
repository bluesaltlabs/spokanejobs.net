<template>
  <button
    :type="type"
    :class="['ui-btn', variant, { loading, disabled }]"
    :disabled="disabled || loading"
    @click="$emit('click', $event)"
  >
    <span v-if="loading" class="ui-btn__spinner"></span>
    <slot />
  </button>
</template>

<script setup>
defineProps({
  type: { type: String, default: 'button' },
  variant: { type: String, default: 'primary' },
  disabled: { type: Boolean, default: false },
  loading: { type: Boolean, default: false },
});
</script>

<style scoped>
.ui-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.5em 1.25em;
  border-radius: 4px;
  border: none;
  font-size: 1rem;
  cursor: pointer;
  background: #2563eb;
  color: #fff;
  transition: background 0.2s;
}
.ui-btn.secondary {
  background: #e5e7eb;
  color: #111;
}
.ui-btn:disabled,
.ui-btn.disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
.ui-btn__spinner {
  width: 1em;
  height: 1em;
  border: 2px solid #fff;
  border-top: 2px solid #2563eb;
  border-radius: 50%;
  margin-right: 0.5em;
  animation: spin 1s linear infinite;
}
@keyframes spin {
  to { transform: rotate(360deg); }
}
</style> 