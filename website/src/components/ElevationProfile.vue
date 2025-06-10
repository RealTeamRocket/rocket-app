<template>
  <div class="elevation-profile">
    <canvas ref="canvas"></canvas>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from "vue";
import { Chart } from "chart.js/auto";
import { getElevations } from "../api/service-api";

const props = defineProps<{
  coordinates: [number, number][];
}>();

const canvas = ref<HTMLCanvasElement | null>(null);
let chart: any = null;

const fetchElevations = async (coords: [number, number][]) => {
  return await getElevations(coords);
}

const drawChart = async () => {
  if (!canvas.value || !props.coordinates.length) return;
  let coords = props.coordinates;
  coords = coords.map(([a, b]) => [a, b]); // shallow copy
  // Downsample if too many points
  const maxPoints = 100;
  if (coords.length > maxPoints) {
    const step = Math.ceil(coords.length / maxPoints);
    coords = coords.filter((_, i) => i % step === 0);
  }
  const elevations = await fetchElevations(coords);
  if (chart) chart.destroy();
  chart = new Chart(canvas.value, {
    type: "line",
    data: {
      labels: coords.map((_, i) => i + 1),
      datasets: [
        {
          label: "Elevation (m)",
          data: elevations,
          borderColor: "#4a90e2",
          backgroundColor: "rgba(74,144,226,0.1)",
          fill: true,
          tension: 0.3,
          pointRadius: 0,
        },
      ],
    },
    options: {
      scales: {
        x: { display: false },
        y: { title: { display: true, text: "m" } },
      },
      plugins: {
        legend: { display: false },
        tooltip: { enabled: true },
      },
      responsive: true,
      maintainAspectRatio: false,
    },
  });
}

watch(() => props.coordinates, drawChart, { immediate: true });
onMounted(drawChart);
</script>

<style scoped>
.elevation-profile {
  width: 100%;
  height: 180px;
  margin-top: 1.5rem;
  background: #f7fafd;
  border-radius: 8px;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
  padding: 1rem 0.5rem 0.5rem 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
}
canvas {
  width: 100% !important;
  height: 150px !important;
}
</style>
