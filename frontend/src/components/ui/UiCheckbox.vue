<template>
  <label class="ui-checkbox">
    <input
      type="checkbox"
      :checked="modelValue"
      :disabled="disabled"
      @change="$emit('update:modelValue', $event.target.checked)"
    />
    <span class="ui-checkbox__box"></span>
    <span v-if="label" class="ui-checkbox__label">{{ label }}</span>
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
.ui-checkbox {
  display: inline-flex;
  align-items: center;
  cursor: pointer;
}
.ui-checkbox input {
  display: none;
}
.ui-checkbox__box {
  width: 1.1em;
  height: 1.1em;
  border: 1.5px solid #2563eb;
  border-radius: 3px;
  background: #fff;
  margin-right: 0.5em;
  position: relative;
}
.ui-checkbox input:checked + .ui-checkbox__box {
  background: #2563eb;
}
.ui-checkbox__label {
  font-size: 1rem;
  color: #111;
}
</style> 