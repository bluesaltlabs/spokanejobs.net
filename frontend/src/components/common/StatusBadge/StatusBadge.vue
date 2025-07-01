<template>
  <span :class="['status-badge', statusClass]">
    <slot name="icon">
      <span v-if="icon" class="status-badge__icon">
        <component :is="icon" />
      </span>
    </slot>
    <span class="status-badge__label">{{ label || status }}</span>
  </span>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  status: { type: String, required: true },
  label: { type: String, default: '' },
  icon: { type: [String, Object], default: null },
});

const statusClass = computed(() => {
  switch (props.status) {
    case 'built':
      return 'status-badge--success';
    case 'in progress':
      return 'status-badge--info';
    case 'error':
      return 'status-badge--error';
    case 'not built':
      return 'status-badge--warning';
    default:
      return 'status-badge--default';
  }
});
</script>

<style scoped>
.status-badge {
  display: inline-flex;
  align-items: center;
  padding: 0.2em 0.6em;
  border-radius: 1em;
  font-size: 0.9em;
  font-weight: 500;
  background: #f3f4f6;
  color: #333;
}
.status-badge__icon {
  margin-right: 0.4em;
  display: flex;
  align-items: center;
}
.status-badge--success {
  background: #e6f9ed;
  color: #1a7f37;
}
.status-badge--info {
  background: #e6f0fa;
  color: #2563eb;
}
.status-badge--error {
  background: #fde8e8;
  color: #b91c1c;
}
.status-badge--warning {
  background: #fff7e6;
  color: #b45309;
}
.status-badge--default {
  background: #f3f4f6;
  color: #333;
}
</style>
