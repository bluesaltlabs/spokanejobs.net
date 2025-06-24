<template>
  <textarea
    :value="modelValue"
    :placeholder="placeholder"
    :rows="rows"
    :disabled="disabled"
    class="ui-textarea ui-form-element"
    @input="$emit('update:modelValue', $event.target.value)"
    @focus="$emit('focus', $event)"
    @blur="$emit('blur', $event)"
  />
</template>

<script setup>
defineProps({
  modelValue: String,
  placeholder: String,
  rows: { type: Number, default: 3 },
  disabled: Boolean,
});
</script>

<style scoped>
.ui-textarea {
  width: 100%;
  max-width: 100%;
  box-sizing: border-box;
  padding: 0.5em 0.75em;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 1rem;
  outline: none;
  transition: border 0.2s;
  resize: vertical;
}
.ui-textarea:focus {
  border-color: #2563eb;
}
.ui-textarea:disabled {
  background: #f3f4f6;
  color: #9ca3af;
}
/* Utility class for consistent vertical spacing of form elements */
.ui-form-element {
  display: block;
  width: 100%;
  margin-bottom: 1rem !important;
}
</style> 