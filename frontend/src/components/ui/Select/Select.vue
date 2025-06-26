<template>
  <div
    class="ui-select"
    ref="root"
    @keydown="onKeyDown"
    @focusin="onFocusIn"
    @focusout="onFocusOut"
    :aria-expanded="dropdownOpen.toString()"
    :aria-haspopup="'listbox'"
    :aria-owns="dropdownId"
    role="combobox"
  >
    <input
      v-model="search"
      :placeholder="placeholder"
      class="ui-select__search"
      :disabled="disabled"
      @focus="onInputFocus"
      @click="openDropdown"
      :aria-controls="dropdownId"
      :aria-autocomplete="'list'"
      :aria-activedescendant="activeDescendantId"
      ref="input"
      role="searchbox"
      autocomplete="off"
      :readonly="!dropdownOpen"
    />
    <div
      v-show="dropdownOpen"
      class="ui-select__options"
      :id="dropdownId"
      role="listbox"
    >
      <div
        v-for="(option, idx) in filteredOptions"
        :key="option.value || option"
        class="ui-select__option"
        :id="optionId(idx)"
        :class="{ 'is-active': idx === highlightedIndex, 'is-selected': isSelected(option) }"
        role="option"
        :aria-selected="isSelected(option).toString()"
        tabindex="-1"
        @mouseenter="highlightedIndex = idx"
        @mousedown.prevent="select(option)"
      >
        {{ option.label || option }}
      </div>
      <div v-if="filteredOptions.length === 0" class="ui-select__no-options" role="option" aria-disabled="true">No options</div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount, nextTick } from 'vue';
const props = defineProps({
  modelValue: [String, Number],
  options: { type: Array, default: () => [] },
  placeholder: String,
  disabled: Boolean,
});
const emit = defineEmits(['update:modelValue']);
const search = ref('');
const dropdownOpen = ref(false);
const highlightedIndex = ref(-1);
const input = ref(null);
const root = ref(null);
const dropdownId = `ui-select-options-${Math.random().toString(36).slice(2)}`;
const activeDescendantId = computed(() =>
  highlightedIndex.value >= 0 ? optionId(highlightedIndex.value) : undefined
);

const filteredOptions = computed(() => {
  if (!search.value) return props.options;
  return props.options.filter(opt => (opt.label || opt).toLowerCase().includes(search.value.toLowerCase()));
});
function isSelected(option) {
  return (option.value || option) === props.modelValue;
}
function select(option) {
  emit('update:modelValue', option.value || option);
  closeDropdown();
}
function openDropdown() {
  if (!dropdownOpen.value && !props.disabled) {
    dropdownOpen.value = true;
    nextTick(() => {
      if (filteredOptions.value.length > 0) highlightedIndex.value = 0;
      search.value = '';
    });
  }
}
function closeDropdown() {
  dropdownOpen.value = false;
  highlightedIndex.value = -1;
  // Show selected label in input when closed
  const selected = props.options.find(opt => (opt.value || opt) === props.modelValue);
  search.value = selected ? (selected.label || selected) : '';
}
function onInputFocus() {
  openDropdown();
  // Clear search on focus to allow searching
  search.value = '';
}
function onKeyDown(e) {
  if (!dropdownOpen.value && (e.key === 'ArrowDown' || e.key === 'Enter' || e.key === ' ')) {
    openDropdown();
    e.preventDefault();
    return;
  }
  if (!dropdownOpen.value) return;
  if (e.key === 'ArrowDown') {
    highlightedIndex.value = (highlightedIndex.value + 1) % filteredOptions.value.length;
    e.preventDefault();
  } else if (e.key === 'ArrowUp') {
    highlightedIndex.value = (highlightedIndex.value - 1 + filteredOptions.value.length) % filteredOptions.value.length;
    e.preventDefault();
  } else if (e.key === 'Enter' || e.key === ' ') {
    if (highlightedIndex.value >= 0 && highlightedIndex.value < filteredOptions.value.length) {
      select(filteredOptions.value[highlightedIndex.value]);
    }
    e.preventDefault();
  } else if (e.key === 'Escape') {
    closeDropdown();
    e.preventDefault();
  } else if (e.key === 'Tab') {
    closeDropdown();
  }
}
function onFocusIn() {
  // No-op, but needed for focusout logic
}
function onFocusOut(e) {
  // Close dropdown if focus leaves the component
  setTimeout(() => {
    if (!root.value.contains(document.activeElement)) {
      closeDropdown();
    }
  }, 0);
}
function optionId(idx) {
  return `${dropdownId}-option-${idx}`;
}
function handleClickOutside(e) {
  if (root.value && !root.value.contains(e.target)) {
    closeDropdown();
  }
}
onMounted(() => {
  // Show selected label in input on mount
  const selected = props.options.find(opt => (opt.value || opt) === props.modelValue);
  search.value = selected ? (selected.label || selected) : '';
  document.addEventListener('mousedown', handleClickOutside);
});
onBeforeUnmount(() => {
  document.removeEventListener('mousedown', handleClickOutside);
});
watch(() => props.modelValue, (val) => {
  // Update input when modelValue changes externally
  const selected = props.options.find(opt => (opt.value || opt) === val);
  search.value = selected ? (selected.label || selected) : '';
});
watch(dropdownOpen, (open) => {
  if (open) {
    nextTick(() => {
      input.value && input.value.focus();
    });
  }
});
</script>

<style scoped>
.ui-select {
  width: 100%;
  position: relative;
}
.ui-select__search {
  width: 100%;
  max-width: 100%;
  box-sizing: border-box;
  padding: 0.5em 0.75em;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 1rem;
  background: #fff;
  color: #111;
  outline: none;
  transition: border 0.2s;
}
.ui-select__search:focus {
  border-color: #2563eb;
}
.ui-select__search:disabled {
  background: #f3f4f6;
  color: #9ca3af;
}
.ui-select__options {
  max-height: 160px;
  overflow-y: auto;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  background: #fff;
  position: absolute;
  width: 100%;
  z-index: 10;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
}
.ui-select__option {
  padding: 0.25em 0.75em;
  font-size: 1rem;
  cursor: pointer;
}
.ui-select__option.is-active {
  background: #e0e7ff;
}
.ui-select__option.is-selected {
  font-weight: bold;
  color: #2563eb;
}
.ui-select__no-options {
  padding: 0.5em 0.75em;
  color: #888;
}
</style> 