<template>
  <div class="ui-dateinput">
    <input
      :type="withTime ? 'datetime-local' : 'date'"
      :value="modelValue"
      :min="min"
      :max="max"
      :disabled="disabled"
      class="ui-dateinput__input"
      @input="$emit('update:modelValue', $event.target.value)"
    />
  </div>
</template>

<script setup>
defineProps({
  modelValue: String,
  withTime: Boolean,
  min: String,
  max: String,
  disabled: Boolean,
});
</script>

<style scoped>
.ui-dateinput__input {
  width: 100%;
  padding: 0.5em 0.75em;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 1rem;
  outline: none;
  transition: border 0.2s;
}
.ui-dateinput__input:focus {
  border-color: #2563eb;
}
.ui-dateinput__input:disabled {
  background: #f3f4f6;
  color: #9ca3af;
}
</style> 