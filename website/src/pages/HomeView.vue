<template>
  <Navbar />
  <div class="dashboard-container">
    <div class="dashboard-header">
      <h1>Welcome back, Astronaut!</h1>
      <p>Here’s your step dashboard. Ready to break your record?</p>
    </div>
    <div class="dashboard-main">
      <div class="dashboard-left">
        <template v-if="isLoggedIn">
          <StatsCards :stats="stats" />
          <StepChart :data="chartData" :labels="chartLabels"/>
        </template>
        <template v-else>
          <div class="skeleton-card" />
          <div class="skeleton-chart" />
        </template>
      </div>
      <div class="dashboard-right">
        <template v-if="isLoggedIn">
          <ActivityPanel />
        </template>
        <template v-else>
          <div class="skeleton-panel" />
        </template>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import Navbar from '../components/Navbar.vue'
import StatsCards from '../components/StatsCards.vue'
import StepChart from '../components/StepChart.vue'
import ActivityPanel from '../components/ActivityPanel.vue'
import { useAuth } from '../utils/useAuth'
import api from '../api/backend-api'

type StepStatistic = {
  day: string
  steps: number
}

const { isLoggedIn } = useAuth()

const stats = ref({
  totalSteps: 0,
  avgSteps: 0,
  bestDay: '',
  bestSteps: 0
})
const chartData = ref<number[]>([])
const chartLabels = ref<string[]>([])

const fetchStats = async () => {
  try {
    const res = await api.getUserStatistics()
    const dailyStats: StepStatistic[] = res.data // or res.data.daily_steps if wrapped

    // Calculate stats
    const totalSteps = dailyStats.reduce((sum, s) => sum + s.steps, 0)
    const avgSteps = Math.round(totalSteps / dailyStats.length)
    const best = dailyStats.reduce(
      (prev, curr) => (curr.steps > prev.steps ? curr : prev),
      dailyStats[0]
    )

    stats.value = {
      totalSteps,
      avgSteps,
      bestDay: best.day,
      bestSteps: best.steps
    }
    chartData.value = dailyStats.map((s) => s.steps)
    chartLabels.value = dailyStats.map((s) => s.day) // <-- dynamic labels
  } catch (e) {
    // Optionally handle error
    stats.value = { totalSteps: 0, avgSteps: 0, bestDay: '', bestSteps: 0 }
    chartData.value = []
    chartLabels.value = []
  }
}

onMounted(() => {
  if (isLoggedIn.value) fetchStats()
})
watch(isLoggedIn, (val) => {
  if (val) fetchStats()
})
</script>

<style scoped>
.dashboard-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1rem;
}
.dashboard-header {
  text-align: center;
  margin-bottom: 2rem;
}
.dashboard-main {
  display: flex;
  gap: 2rem;
}
.dashboard-left {
  flex: 2;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}
.dashboard-right {
  flex: 1;
}

.skeleton-card,
.skeleton-chart,
.skeleton-panel {
  background: #e5e7eb;
  border-radius: 8px;
  margin-bottom: 1rem;
  animation: pulse 1.5s infinite;
}
.skeleton-card {
  height: 120px;
}
.skeleton-chart {
  height: 240px;
}
.skeleton-panel {
  height: 400px;
}
@keyframes pulse {
  0% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
  100% {
    opacity: 1;
  }
}
</style>
