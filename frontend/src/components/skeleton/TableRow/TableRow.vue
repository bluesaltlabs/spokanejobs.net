<script setup>
import SkeletonText from '../Text/Text.vue'

const props = defineProps({
  columns: {
    type: Number,
    default: 3
  },
  columnWidths: {
    type: Array,
    default: () => ['40%', '30%', '30%']
  }
})

const getColumnWidth = (columnIndex) => {
  return props.columnWidths[columnIndex - 1] || '100%'
}
</script>

<template>
  <tr class="skeleton-table-row">
    <td v-for="i in columns" :key="i" class="skeleton-table-cell">
      <SkeletonText
        :lines="1"
        :line-height="16"
        :last-line-width="getColumnWidth(i)"
      />
    </td>
  </tr>
</template>

<style scoped>
.skeleton-table-row {

  border-bottom: 1px solid var(--color-border);
}

.skeleton-table-cell {
  padding: 0.75rem 1rem;
}

.skeleton-table-row:hover {
  background: transparent;
}
</style>
