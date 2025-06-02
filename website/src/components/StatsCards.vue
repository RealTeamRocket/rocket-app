<template>
  <div class="stats-cards">
    <div class="stat-card">
      <div class="stat-label">Total Steps</div>
      <div class="stat-value">{{ stats.totalSteps.toLocaleString() }}</div>
    </div>
    <div class="stat-card">
      <div class="stat-label">Average Steps/Day</div>
      <div class="stat-value">{{ stats.avgSteps.toLocaleString() }}</div>
    </div>
    <div class="stat-card">
      <div class="stat-label">Best Day</div>
      <div class="stat-value">{{ stats.bestDay }}</div>
      <div class="stat-subvalue">{{ stats.bestSteps.toLocaleString() }} steps</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

interface Stats {
  totalSteps: number
  avgSteps: number
  bestDay: string
  bestSteps: number
}

defineProps<{ stats: Stats }>()
</script>

<style scoped>
.stats-cards {
  display: flex;
  gap: 1.5rem;
  justify-content: space-between;
  margin-bottom: 2rem;
  flex-wrap: wrap;
}

.stat-card {
  background: linear-gradient(90deg, #e0e7ff 0%, #f3f8ff 100%);
  border-radius: 1rem;
  box-shadow: 0 2px 8px rgba(30,60,114,0.08);
  padding: 1.5rem 2rem;
  min-width: 180px;
  flex: 1 1 200px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-label {
  font-size: 1rem;
  color: #2a5298;
  margin-bottom: 0.5rem;
  font-weight: 500;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 2rem;
  font-weight: bold;
  color: #1e3c72;
  margin-bottom: 0.25rem;
}

.stat-subvalue {
  font-size: 0.95rem;
  color: #4b5563;
  margin-top: 0.2rem;
}
</style>
