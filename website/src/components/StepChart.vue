<template>
  <div class="step-chart-card">
    <h2 class="chart-title">Steps in the Last 7 Days</h2>
    <canvas ref="chartRef"></canvas>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, onBeforeUnmount } from 'vue'
import Chart from 'chart.js/auto'

const props = defineProps<{
  data: number[]
}>()

const chartRef = ref<HTMLCanvasElement | null>(null)
let chartInstance: any = null

const labels = [
  'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'
]

const renderChart = () => {
  if (chartInstance) {
    chartInstance.destroy()
  }
  if (!chartRef.value) return

  chartInstance = new Chart(chartRef.value, {
    type: 'bar',
    data: {
      labels,
      datasets: [
        {
          label: 'Steps',
          data: props.data,
          backgroundColor: 'rgba(30, 60, 114, 0.7)',
          borderRadius: 6,
          maxBarThickness: 40
        }
      ]
    },
    options: {
      responsive: true,
      plugins: {
        legend: { display: false },
        tooltip: { enabled: true }
      },
      scales: {
        y: {
          beginAtZero: true,
          ticks: {
            color: '#1e3c72',
            stepSize: 1000
          },
          grid: {
            color: '#e0e7ff'
          }
        },
        x: {
          ticks: {
            color: '#1e3c72'
          },
          grid: {
            display: false
          }
        }
      }
    }
  })
}

onMounted(() => {
  renderChart()
})

watch(() => props.data, () => {
  renderChart()
})

onBeforeUnmount(() => {
  if (chartInstance) {
    chartInstance.destroy()
  }
})
</script>

<style scoped>
.step-chart-card {
  background: #fff;
  border-radius: 1rem;
  box-shadow: 0 2px 8px rgba(30,60,114,0.08);
  padding: 2rem 1.5rem 1.5rem 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.chart-title {
  margin-bottom: 1rem;
  color: #1e3c72;
  font-size: 1.2rem;
  font-weight: 600;
}
canvas {
  max-width: 100%;
  height: 260px !important;
}
</style>
