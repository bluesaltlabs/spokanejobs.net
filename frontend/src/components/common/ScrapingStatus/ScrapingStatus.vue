<template>
  <div class="scraping-status">
    <StatusBadge :status="status" :label="statusLabel" :icon="statusIcon" />
    <div class="scraping-status__details">
      <InfoRow label="Last Scraped" :value="formatDate(lastScrapedAt)" v-if="lastScrapedAt" icon="mdi:clock-outline" />
      <InfoRow label="Scraper Created" :value="formatDate(scraperCreatedAt)" v-if="scraperCreatedAt" icon="mdi:calendar-plus" />
      <InfoRow label="Active Jobs" :value="activeJobsCount" v-if="activeJobsCount !== undefined" icon="mdi:briefcase-outline" />
      <InfoRow label="Error" :value="scraperError" v-if="scraperError" icon="mdi:alert-circle-outline" />
      <slot />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import StatusBadge from '../StatusBadge/StatusBadge.vue';
import InfoRow from '../InfoRow/InfoRow.vue';

const props = defineProps({
  status: { type: String, required: true },
  lastScrapedAt: { type: String, default: '' },
  scraperCreatedAt: { type: String, default: '' },
  activeJobsCount: { type: [Number, String], default: undefined },
  scraperError: { type: String, default: '' },
});

const statusLabel = computed(() => {
  switch (props.status) {
    case 'built': return 'Scraper Built';
    case 'in progress': return 'Scraping In Progress';
    case 'error': return 'Scraper Error';
    case 'not built': return 'No Scraper';
    default: return props.status;
  }
});

const statusIcon = computed(() => {
  switch (props.status) {
    case 'built': return 'mdi:check-circle-outline';
    case 'in progress': return 'mdi:progress-clock';
    case 'error': return 'mdi:alert-circle-outline';
    case 'not built': return 'mdi:close-circle-outline';
    default: return null;
  }
});

function formatDate(dateStr) {
  if (!dateStr) return '';
  const date = new Date(dateStr);
  if (isNaN(date)) return dateStr;
  return date.toLocaleString();
}
</script>

<style scoped>
.scraping-status {
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 1em;
  background: #fafbfc;
  max-width: 400px;
}
.scraping-status__details {
  margin-top: 0.8em;
}
</style>
