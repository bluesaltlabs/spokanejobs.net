<template>
  <div class="ui-input-wrapper">
    <input
      v-if="editable"
      :type="type"
      :value="modelValue"
      :placeholder="placeholder"
      :disabled="disabled"
      :editable="editable"
      class="ui-input ui-form-element"
      @input="$emit('update:modelValue', $event.target.value)"
      @focus="$emit('focus', $event)"
      @blur="$emit('blur', $event)"
      :style="clearable && modelValue ? 'padding-right:2em' : ''"
    />
    <button
      v-if="editable && clearable && modelValue"
      class="clear-btn"
      type="button"
      aria-label="Clear input"
      @click="$emit('update:modelValue', '')"
    >&times;</button>
    <span
      v-if="!editable"
      class="ui-input ui-form-element"
    >{{ modelValue }}</span>
  </div>
</template>

<script setup>
defineProps({
  modelValue: String,
  placeholder: String,
  type: { type: String, default: 'text' },
  disabled: Boolean,
  editable: Boolean,
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
  border: 1px solid var(--color-border);
  border-radius: 4px;
  font-size: 1rem;
  outline: none;
  transition: border 0.2s;
}
.ui-input:focus {
  border-color: var(--color-border-focus);
}
.ui-input:disabled {
  background-color: var(--color-background-mute);
  color: var(--color-text-muted);
}
/* Utility class for consistent vertical spacing of form elements */
.ui-form-element {
  display: block;
  width: 100%;
  margin-bottom: 1rem !important;
}
.clear-btn {
  position: absolute;
  right: 0.5em;
  top: 10%;
  background: none;
  border: none;
  font-size: 1.25rem;
  color: var(--color-primary);
  cursor: pointer;
  padding: 0.5em 0.75em;
  line-height: 1;
  z-index: 2;
  height: 1.5em;
  width: 1.5em;
  display: flex;
  align-items: center;
  justify-content: center;
}
.clear-btn:focus {
  outline: 2px solid var(--color-primary);
  outline-offset: 2px;
  color: var(--color-border-focus);
}
</style>
