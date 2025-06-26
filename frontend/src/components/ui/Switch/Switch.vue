<template>
  <label class="ui-switch">
    <input
      type="checkbox"
      :checked="modelValue"
      :disabled="disabled"
      @change="$emit('update:modelValue', $event.target.checked)"
    />
    <span class="ui-switch__slider"></span>
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
  width: 2em;
  height: 1em;
  background: #d1d5db;
  border-radius: 1em;
  position: relative;
  transition: background 0.2s;
  margin-right: 0.5em;
}
.ui-switch input:checked + .ui-switch__slider {
  background: #2563eb;
}
.ui-switch__slider::after {
  content: '';
  position: absolute;
  left: 0.2em;
  top: 0.2em;
  width: 0.6em;
  height: 0.6em;
  background: #fff;
  border-radius: 50%;
  transition: left 0.2s;
}
.ui-switch input:checked + .ui-switch__slider::after {
  left: 1.2em;
}
.ui-switch__label {
  font-size: 1rem;
  color: #111;
}
</style> 