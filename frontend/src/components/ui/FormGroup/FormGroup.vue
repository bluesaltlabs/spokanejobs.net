<template>
  <div class="ui-form-group">
    <label v-if="label" class="ui-form-group__label">
      {{ label }}<span v-if="required" class="ui-form-group__required">*</span>
    </label>
    <slot />
    <div v-if="error" class="ui-form-group__error">{{ error }}</div>
  </div>
</template>

<script setup>
defineProps({
  label: String,
  error: String,
  required: Boolean,
});
</script>

<style scoped>
.ui-form-group {
  margin-bottom: 1.25em;
}
.ui-form-group__label {
  display: block;
  font-weight: 500;
  margin-bottom: 0.25em;
}
.ui-form-group__required {
  color: #dc2626;
  margin-left: 0.25em;
}
.ui-form-group__error {
  color: #dc2626;
  font-size: 0.95em;
  margin-top: 0.25em;
}
</style> 