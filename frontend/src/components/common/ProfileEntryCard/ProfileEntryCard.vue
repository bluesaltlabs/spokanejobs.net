<template>
  <div class="profile-entry-card">
    <div class="profile-entry-content">
      <slot name="title" />
      <slot name="subtitle" />
      <slot name="description" />
      <slot name="details" />
      <slot name="metadata" />
    </div>
    <div v-if="$slots.actions" class="profile-entry-actions">
      <div class="actions-container">
        <slot name="actions" />
      </div>
    </div>
  </div>
</template>

<script setup>
// No props needed - this is a pure layout component
</script>

<style scoped>
.profile-entry-card {
  border: 1.5px solid var(--color-border);
  border-radius: var(--border-radius-medium);
  padding: 1.25rem 1rem 1rem 1rem;
  margin-bottom: 1.25rem;
  background: var(--color-surface);
  box-shadow: 0 1px 4px var(--color-shadow);
  transition: border 0.2s;
}

.profile-entry-content {
  /* Content styling will be handled by slots */
}

.profile-entry-actions {
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: 1px solid var(--color-border);
}

.actions-container {
  display: flex;
  gap: 0.5rem;
  justify-content: flex-end;
}

/* Common text styling for profile entries */
.profile-entry-card :deep(strong) {
  color: var(--color-primary-600);
  font-size: 1.1rem;
  font-weight: 700;
}

.profile-entry-card :deep(em) {
  color: var(--color-primary-500);
  font-style: normal;
  font-weight: 600;
}

.profile-entry-card :deep(span) {
  color: var(--color-text-muted);
  font-size: 0.95rem;
  display: block;
  margin-bottom: 0.5rem;
}

.profile-entry-card :deep(p) {
  color: var(--color-text);
  margin: 0.5rem 0 0 0;
  font-size: 1rem;
}

/* Info box styling for details like supervisor info, education details, etc. */
.profile-entry-card :deep(.info-box) {
  margin-top: 0.5rem;
  padding: 0.5rem;
  background: var(--color-surface-hover);
  border-radius: var(--border-radius-small);
  border-left: 3px solid var(--color-primary-500);
}

.profile-entry-card :deep(.info-box strong) {
  color: var(--color-text);
  font-weight: 600;
  font-size: 0.95rem;
}

/* Success info box for honors/recognition */
.profile-entry-card :deep(.success-info-box) {
  margin-top: 0.5rem;
  padding: 0.5rem;
  background: var(--color-surface-hover);
  border-radius: var(--border-radius-small);
  border-left: 3px solid var(--color-success-500);
}

.profile-entry-card :deep(.success-info-box ul) {
  margin: 0.5rem 0 0 0;
  padding-left: 1.5rem;
}

.profile-entry-card :deep(.success-info-box li) {
  margin-bottom: 0.25rem;
}
</style> 