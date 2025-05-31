<template>
  <Navbar />
  <div class="dashboard-container">
    <div class="dashboard-header">
      <h1>Welcome back, Astronaut!</h1>
      <p>Here’s your step dashboard. Ready to break your record?</p>
    </div>
    <div class="dashboard-main">
      <div class="dashboard-left">
        <StatsCards :stats="stats" />
        <StepChart :data="chartData" />
      </div>
      <div class="dashboard-right">
        <ActivityPanel />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from '../components/Navbar.vue'
import StatsCards from '../components/StatsCards.vue'
import StepChart from '../components/StepChart.vue'
import ActivityPanel from '../components/ActivityPanel.vue'

const stats = {
  totalSteps: 56000,
  avgSteps: 8000,
  bestDay: 'Wednesday',
  bestSteps: 12000
}
const chartData = [7000, 8000, 9000, 10000, 12000, 8000, 7000]
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
</style>
