<template>
  <CodeMirror
    v-model="localValue"
    :lang="lang"
    :placeholder="placeholder"
    :disabled="disabled"
    basic
    class="ui-richtextinput"
    @update:modelValue="onInput"
  />
</template>

<script setup>
import { ref, watch, computed } from 'vue';
import CodeMirror from 'vue-codemirror6';
import { json } from '@codemirror/lang-json';
import { markdown } from '@codemirror/lang-markdown';
import { html } from '@codemirror/lang-html';

const props = defineProps({
  modelValue: String,
  mode: { type: String, default: 'markdown' }, // 'json', 'markdown', 'html'
  placeholder: String,
  disabled: Boolean,
});
const emit = defineEmits(['update:modelValue']);
const localValue = ref(props.modelValue || '');
watch(() => props.modelValue, v => { if (v !== localValue.value) localValue.value = v; });
function onInput(val) { emit('update:modelValue', val); }
const lang = computed(() => {
  if (props.mode === 'json') return json();
  if (props.mode === 'html') return html();
  return markdown();
});
</script>

<style scoped>
.ui-richtextinput {
  width: 100%;
  min-height: 180px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 1rem;
  background: #fff;
}
</style> 