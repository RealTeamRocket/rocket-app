<template>
  <aside class="run-sidebar">
    <h2>Your Runs</h2>
    <ul>
      <li
        v-for="run in runs"
        :key="run.id"
        :class="{ selected: run.id === selectedId }"
        @click="select(run)"
      >
        <div class="run-info">
          <div class="run-date">
            <span class="icon-calendar" title="Date"></span>
            <span class="run-date-value">{{ formatDate(run.created_at) }}</span>
          </div>
          <div class="run-detail-row">
            <span class="icon-distance" title="Distance"></span>
            <span class="run-detail-label">Distance:</span>
            <span class="run-detail-value">{{ run.distance?.toFixed(2) ?? '?' }} km</span>
          </div>
          <div class="run-detail-row">
            <span class="icon-duration" title="Duration"></span>
            <span class="run-detail-label">Duration:</span>
            <span class="run-detail-value">{{ run.duration ?? '?' }}</span>
          </div>
        </div>
      </li>
    </ul>
  </aside>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'

const props = defineProps<{
  runs: any[],
  selectedId?: string
}>()

const emit = defineEmits<{
  (e: 'select', run: any): void
}>()

const select = (run: any) => {
  emit('select', run)
}

const formatDate = (dateStr: string) => {
  if (!dateStr) return '?'
  return new Date(dateStr).toLocaleString(undefined, { dateStyle: 'medium', timeStyle: 'short' })
}
</script>

<style scoped>
.run-sidebar {
  width: 300px;
  background: #f7f7f7;
  border-right: 1px solid #ddd;
  overflow-y: auto;
  padding: 1rem;
  height: 100%;
}
.run-sidebar ul {
  list-style: none;
  padding: 0;
  margin: 0;
}
.run-sidebar li {
  padding: 1rem 0.75rem;
  cursor: pointer;
  border-radius: 8px;
  margin-bottom: 0.75rem;
  transition: background 0.2s, box-shadow 0.2s;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 0.5rem;
  background: #f9fafd;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
  border: 1.5px solid transparent;
}
.run-sidebar li.selected,
.run-sidebar li:hover {
  background: #e0eaff;
  border: 1.5px solid #4a90e2;
  box-shadow: 0 2px 8px rgba(74,144,226,0.08);
}
.run-info {
  flex: 1;
  min-width: 0;
  width: 100%;
  text-align: left;
  margin-top: 0.2rem;
  color: #222;
}
.run-date {
  font-size: 1.05rem;
  font-weight: 600;
  color: #4a90e2;
  margin-bottom: 0.3rem;
  display: flex;
  align-items: center;
  gap: 0.4em;
}
.run-date-value {
  white-space: nowrap;
  font-size: 0.97em;
}
.run-detail-row {
  display: flex;
  align-items: center;
  gap: 0.4em;
  font-size: 0.98rem;
  margin-bottom: 0.15rem;
}
.run-detail-label {
  color: #888;
  font-weight: 500;
}
.run-detail-value {
  color: #222;
  font-weight: 600;
}
.icon-calendar::before {
  content: "📅";
  margin-right: 0.2em;
}
.icon-distance::before {
  content: "📏";
  margin-right: 0.2em;
}
.icon-duration::before {
  content: "⏱️";
  margin-right: 0.2em;
}
</style>
