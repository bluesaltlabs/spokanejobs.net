<template>
  <div class="search-bar">
    <TextInput
      v-model="search"
      :placeholder="placeholder"
      class="search-input"
      @input="onInput"
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

function onInput() {
  emit('update:modelValue', search.value)
}
</script>
