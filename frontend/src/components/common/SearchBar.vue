<template>
  <div class="search-bar">
    <TextInput
      v-model="search"
      :placeholder="placeholder"
      class="search-input"
      @update:modelValue="onInput"
      clearable
    />
  </div>
</template>

<script setup>
import { ref, watch, defineProps, defineEmits } from 'vue'
import TextInput from '@/components/ui/TextInput/TextInput.vue'

const props = defineProps({
  modelValue: String,
  placeholder: {
    type: String,
    default: 'Search...'
  }
})
const emit = defineEmits(['update:modelValue'])

const search = ref(props.modelValue || '')

watch(() => props.modelValue, (val) => {
  if (val !== search.value) search.value = val || ''
})

function onInput(val) {
  search.value = val
  emit('update:modelValue', val)
}
</script>

<style scoped>
.search-bar {
  display: flex;
  align-items: center;
  margin-bottom: 1.5rem;
  position: relative;
}
.search-input {
  width: 100%;
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-medium);
  font-size: 1rem;
  outline: none;
}
</style> 