<template>
  <label class="ui-switch">
    <input
      type="checkbox"
      :checked="modelValue"
      :disabled="disabled"
      @change="$emit('update:modelValue', $event.target.checked)"
    />
    <span class="ui-switch__slider">
      <span class="ui-switch__toggle">
        <slot />
      </span>
    </span>
    <span v-if="label" class="ui-switch__label">{{ label }}</span>
  </label>
</template>

<script setup>
defineProps({
  modelValue: Boolean,
  disabled: Boolean,
  label: String,
});
</script>

<style scoped>
.ui-switch {
  display: inline-flex;
  align-items: center;
  cursor: pointer;
}
.ui-switch input {
  display: none;
}
.ui-switch__slider {
  width: 3.5em;
  height: 2em;
  background: var(--color-gray-300);
  border-radius: 1em;
  position: relative;
  transition: background 0.2s;
  margin-right: 0.5em;
  display: flex;
  align-items: center;
  justify-content: center;
}
.ui-switch input:checked + .ui-switch__slider {
  background: var(--color-primary-600);
}
.ui-switch__toggle {
  position: absolute;
  left: 0.25em;
  top: 0.25em;
  width: 1.5em;
  height: 1.5em;
  background: var(--color-white);
  border-radius: 50%;
  transition: left 0.2s;
  box-shadow: var(--color-shadow);
  display: flex;
  align-items: center;
  justify-content: center;
}
.ui-switch input:checked + .ui-switch__slider .ui-switch__toggle {
  left: 1.75em;
}

/* Style SVG icons in the toggle circle */
.ui-switch__toggle :deep(svg) {
  width: 0.75em;
  height: 0.75em;
  color: var(--color-gray-600);
  transition: color 0.2s;
}

.ui-switch input:checked + .ui-switch__slider .ui-switch__toggle :deep(svg) {
  color: var(--color-primary-600);
}

.ui-switch__label {
  font-size: 1rem;
  color: var(--color-text);
}
</style>
