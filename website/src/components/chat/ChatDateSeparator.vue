<template>
  <div class="chat-date-separator my-2 d-flex justify-content-center">
    <span class="date-badge px-3 py-1 rounded shadow-sm">{{ formattedDate }}</span>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{ date: string }>()
function formatDate(dateStr: string) {
  const date = new Date(dateStr)
  return date.toLocaleDateString(undefined, { year: 'numeric', month: 'short', day: 'numeric' })
}
const formattedDate = formatDate(props.date)
</script>

<style scoped>
.date-badge {
  background: #e9ecef;
  color: #495057;
  font-size: 0.92em;
  font-weight: 500;
  border-radius: 1em;
}
</style>
