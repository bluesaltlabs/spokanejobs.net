<template>
  <div class="ui-dateinput ui-form-element">
    <input
      :type="withTime ? 'datetime-local' : 'date'"
      :value="modelValue"
      :min="min"
      :max="max"
      :disabled="disabled"
      class="ui-dateinput__input"
      @input="$emit('update:modelValue', $event.target.value)"
    />
    <span class="calendar-icon">
      <CalendarIcon :width="20" :height="20" :color="iconColor" />
    </span>
  </div>
</template>

<script setup>
import { CalendarIcon } from '@/components/icons';
import { computed } from 'vue';

defineProps({
  modelValue: String,
  withTime: Boolean,
  min: String,
  max: String,
  disabled: Boolean,
});

// Use CSS variable for icon color
const iconColor = computed(() => getComputedStyle(document.documentElement).getPropertyValue('--color-text') || '#111');
</script>

<style scoped>
.ui-dateinput {
  position: relative;
}
.ui-dateinput__input {
  width: 100%;
  padding: 0.5em 2.5em 0.5em 0.75em; /* right padding for icon */
  border: 1px solid var(--color-border);
  border-radius: 4px;
  font-size: 1rem;
  outline: none;
  transition: border 0.2s;
  background-color: var(--color-surface);
  color: var(--color-text);
}
.ui-dateinput__input:focus {
  border-color: var(--color-border-focus);
}
.ui-dateinput__input:disabled {
  background-color: var(--color-background-mute);
  color: var(--color-text-muted);
}
/* Hide native calendar icon */
.ui-dateinput__input::-webkit-calendar-picker-indicator {
  opacity: 0;
  display: none;
}
.calendar-icon {
  position: absolute;
  right: 0.75em;
  top: 50%;
  transform: translateY(-50%);
  pointer-events: none;
  display: flex;
  align-items: center;
}

input::-webkit-calendar-picker-indicator,
::-webkit-input-placeholder,
::-ms-input-placeholder,
::-o-input-placeholder {
  filter: invert(1) brightness(1.5);
    color: var(--color-text);
}


/* Utility class for consistent vertical spacing of form elements */
.ui-form-element {
  display: block;
  width: 100%;
  margin-bottom: 1rem !important;
}
</style>

<!-- Non-scoped style for calendar icon color fix in dark mode for both date and datetime-local -->
<style>
[data-theme="dark"] .ui-dateinput__input[type="date"]::-webkit-calendar-picker-indicator,
[data-theme="dark"] .ui-dateinput__input[type="datetime-local"]::-webkit-calendar-picker-indicator {
  filter: invert(1) brightness(2) grayscale(1) contrast(2);
}
</style>
