<template>
  <div class="ui-input-wrapper">
    <input
      :type="type"
      :value="modelValue"
      :placeholder="placeholder"
      :disabled="disabled"
      class="ui-input ui-form-element"
      @input="$emit('update:modelValue', $event.target.value)"
      @focus="$emit('focus', $event)"
      @blur="$emit('blur', $event)"
      :style="clearable && modelValue ? 'padding-right:2em' : ''"
    />
    <button
      v-if="clearable && modelValue"
      class="clear-btn"
      type="button"
      aria-label="Clear input"
      @click="$emit('update:modelValue', '')"
      tabindex="-1"
    >&times;</button>
  </div>
</template>

<script setup>
defineProps({
  modelValue: String,
  placeholder: String,
  type: { type: String, default: 'text' },
  disabled: Boolean,
  clearable: Boolean,
});
</script>

<style scoped>
.ui-input-wrapper {
  position: relative;
  width: 100%;
  display: flex;
  align-items: center;
}
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
/* Utility class for consistent vertical spacing of form elements */
.ui-form-element {
  display: block;
  width: 100%;
  margin-bottom: 1rem !important;
}
.clear-btn {
  position: absolute;
  right: 0.75em;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  font-size: 1.25em;
  color: #9ca3af;
  cursor: pointer;
  padding: 0;
  line-height: 1;
  z-index: 2;
  height: 1.5em;
  width: 1.5em;
  display: flex;
  align-items: center;
  justify-content: center;
}
.clear-btn:focus {
  outline: none;
  color: #2563eb;
}
</style> 