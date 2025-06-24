<template>
  <label class="ui-radio">
    <input
      type="radio"
      :checked="modelValue"
      :name="name"
      :disabled="disabled"
      @change="$emit('update:modelValue', true)"
    />
    <span class="ui-radio__circle"></span>
    <span v-if="label" class="ui-radio__label">{{ label }}</span>
  </label>
</template>

<script setup>
defineProps({
  modelValue: Boolean,
  name: String,
  disabled: Boolean,
  label: String,
});
</script>

<style scoped>
.ui-radio {
  display: inline-flex;
  align-items: center;
  cursor: pointer;
}
.ui-radio input {
  display: none;
}
.ui-radio__circle {
  width: 1.1em;
  height: 1.1em;
  border: 1.5px solid #2563eb;
  border-radius: 50%;
  background: #fff;
  margin-right: 0.5em;
  position: relative;
}
.ui-radio input:checked + .ui-radio__circle {
  background: #2563eb;
}
.ui-radio__label {
  font-size: 1rem;
  color: #111;
}
</style> 