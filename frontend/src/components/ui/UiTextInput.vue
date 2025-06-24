<template>
  <input
    :type="type"
    :value="modelValue"
    :placeholder="placeholder"
    :disabled="disabled"
    class="ui-input"
    @input="$emit('update:modelValue', $event.target.value)"
    @focus="$emit('focus', $event)"
    @blur="$emit('blur', $event)"
  />
</template>

<script setup>
defineProps({
  modelValue: String,
  placeholder: String,
  type: { type: String, default: 'text' },
  disabled: Boolean,
});
</script>

<style scoped>
.ui-input {
  width: 100%;
  max-width: 100%;
  box-sizing: border-box;
  padding: 0.5em 0.75em;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 1rem;
  outline: none;
  transition: border 0.2s;
}
.ui-input:focus {
  border-color: #2563eb;
}
.ui-input:disabled {
  background: #f3f4f6;
  color: #9ca3af;
}
</style> 