<template>
  <div class="info-row">
    <span v-if="icon" class="info-row__icon">
      <component :is="icon" />
    </span>
    <span class="info-row__label">{{ label }}</span>
    <span class="info-row__value">
      <slot>{{ value }}</slot>
    </span>
  </div>
</template>

<script setup>
const props = defineProps({
  label: { type: String, required: true },
  value: { type: [String, Number], default: '' },
  icon: { type: [String, Object], default: null },
});
</script>

<style scoped>
.info-row {
  display: flex;
  align-items: center;
  margin-bottom: 0.5em;
  font-size: 1em;
}
.info-row__icon {
  margin-right: 0.5em;
  display: flex;
  align-items: center;
}
.info-row__label {
  font-weight: 500;
  margin-right: 0.5em;
  color: #555;
  min-width: 120px;
}
.info-row__value {
  color: #222;
}
</style>
