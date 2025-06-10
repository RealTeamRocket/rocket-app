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
        <div>
          <strong>{{ formatDate(run.created_at) }}</strong>
          <div>{{ run.distance?.toFixed(2) ?? '?' }} km, {{ run.duration ?? '?' }}</div>
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

function select(run: any) {
  emit('select', run)
}

function formatDate(dateStr: string) {
  if (!dateStr) return '?'
  return new Date(dateStr).toLocaleString()
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
  padding: 0.5rem;
  cursor: pointer;
  border-radius: 4px;
  margin-bottom: 0.5rem;
  transition: background 0.2s;
}
.run-sidebar li.selected,
.run-sidebar li:hover {
  background: #e0eaff;
}
</style>
