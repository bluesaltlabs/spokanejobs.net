<template>
  <div
    class="ui-multiselect"
    ref="root"
    @keydown="onKeyDown"
    @focusin="onFocusIn"
    @focusout="onFocusOut"
    :aria-expanded="dropdownOpen.toString()"
    :aria-haspopup="'listbox'"
    :aria-owns="dropdownId"
    role="combobox"
  >
    <div class="ui-multiselect__input-wrapper">
      <div class="ui-multiselect__chips-inline">
        <span
          v-for="option in selectedOptions"
          :key="option.value || option"
          class="ui-multiselect__chip"
        >
          {{ option.label || option }}
          <button type="button" class="ui-multiselect__chip-remove" @click="remove(option.value || option)" aria-label="Remove option">&times;</button>
        </span>
        <input
          v-model="search"
          :placeholder="selectedOptions.length === 0 ? placeholder : ''"
          class="ui-multiselect__search-inline"
          :disabled="disabled"
          @focus="openDropdown"
          @click="openDropdown"
          :aria-controls="dropdownId"
          :aria-autocomplete="'list'"
          :aria-activedescendant="activeDescendantId"
          ref="input"
          role="searchbox"
          autocomplete="off"
        />
      </div>
    </div>
    <div
      v-show="dropdownOpen"
      class="ui-multiselect__options"
      :id="dropdownId"
      role="listbox"
      :aria-multiselectable="'true'"
    >
      <label
        v-for="(option, idx) in filteredOptions"
        :key="option.value || option"
        class="ui-multiselect__option"
        :id="optionId(idx)"
        :class="{ 'is-active': idx === highlightedIndex }"
        role="option"
        :aria-selected="modelValue.includes(option.value || option).toString()"
        tabindex="-1"
        @mouseenter="highlightedIndex = idx"
        @mousedown.prevent="toggle(option.value || option)"
      >
        <input
          type="checkbox"
          :value="option.value || option"
          :checked="modelValue.includes(option.value || option)"
          :disabled="disabled"
          @change.stop="toggle(option.value || option)"
          tabindex="-1"
        />
        {{ option.label || option }}
      </label>
      <div v-if="filteredOptions.length === 0" class="ui-multiselect__no-options" role="option" aria-disabled="true">No options</div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount, nextTick } from 'vue';
const props = defineProps({
  modelValue: { type: Array, default: () => [] },
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
const dropdownId = `ui-multiselect-options-${Math.random().toString(36).slice(2)}`;
const activeDescendantId = computed(() =>
  highlightedIndex.value >= 0 ? optionId(highlightedIndex.value) : undefined
);

const filteredOptions = computed(() => {
  if (!search.value) return props.options;
  return props.options.filter(opt => (opt.label || opt).toLowerCase().includes(search.value.toLowerCase()));
});
const selectedOptions = computed(() => {
  return props.options.filter(opt => props.modelValue.includes(opt.value || opt));
});
function toggle(val) {
  let newValue;
  if (props.modelValue.includes(val)) {
    newValue = props.modelValue.filter(v => v !== val);
  } else {
    newValue = [...props.modelValue, val];
  }
  emit('update:modelValue', newValue);
  search.value = '';
}
function remove(val) {
  const newValue = props.modelValue.filter(v => v !== val);
  emit('update:modelValue', newValue);
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
      toggle(filteredOptions.value[highlightedIndex.value].value || filteredOptions.value[highlightedIndex.value]);
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
  document.addEventListener('mousedown', handleClickOutside);
});
onBeforeUnmount(() => {
  document.removeEventListener('mousedown', handleClickOutside);
});
watch(dropdownOpen, (open) => {
  if (open) {
    nextTick(() => {
      input.value && input.value.focus();
    });
  }
});
watch(() => props.modelValue, () => {
  search.value = '';
});
</script>

<style scoped>
.ui-multiselect {
  width: 100%;
  position: relative;
}
.ui-multiselect__input-wrapper {
  width: 100%;
}
.ui-multiselect__chips-inline {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 0.25em;
  min-height: 2.5em;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  padding: 0.25em 0.5em;
  background: #fff;
  cursor: text;
}
.ui-multiselect__chip {
  display: inline-flex;
  align-items: center;
  background: #2563eb;
  color: #fff;
  border-radius: 16px;
  padding: 0.15em 0.6em 0.15em 0.6em;
  font-size: 0.95em;
  margin-right: 0.15em;
}
.ui-multiselect__chip-remove {
  background: none;
  border: none;
  color: #fff;
  margin-left: 0.4em;
  cursor: pointer;
  font-size: 1em;
  line-height: 1;
}
.ui-multiselect__search-inline {
  flex: 1 1 120px;
  min-width: 80px;
  border: none;
  outline: none;
  font-size: 1rem;
  padding: 0.4em 0.2em;
  background: transparent;
  color: #111;
}
.ui-multiselect__search-inline:disabled {
  background: #f3f4f6;
  color: #9ca3af;
}
.ui-multiselect__options {
  max-height: 160px;
  overflow-y: auto;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  background: #fff;
  position: absolute;
  width: 100%;
  z-index: 10;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
  margin-top: 0.25em;
}
.ui-multiselect__option {
  display: flex;
  align-items: center;
  padding: 0.25em 0.75em;
  font-size: 1rem;
  cursor: pointer;
}
.ui-multiselect__option.is-active {
  background: #e0e7ff;
}
.ui-multiselect__option input[type='checkbox'] {
  margin-right: 0.5em;
}
.ui-multiselect__no-options {
  padding: 0.5em 0.75em;
  color: #888;
}
</style> 