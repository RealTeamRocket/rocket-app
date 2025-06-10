<template>
  <Navbar />
  <div class="runs-view">
    <RunSidebar
      :runs="runs"
      :selected-id="selectedRun?.id"
      @select="selectRun"
    />
    <main class="run-details">
      <h2 v-if="selectedRun">Run Details</h2>
      <div v-if="selectedRun">
        <strong>Date:</strong> {{ formatDate(selectedRun.created_at) }}<br>
        <strong>Distance:</strong> {{ selectedRun.distance?.toFixed(2) ?? '?' }} km<br>
        <strong>Duration:</strong> {{ selectedRun.duration ?? '?' }}
      </div>
      <div class="map-container">
        <Map
          v-if="selectedRun"
          :route="selectedRun.route"
          :markers="routeMarkers(selectedRun.route)"
        />
        <div v-else class="empty-map-msg">
          <p>Select a run to see its route.</p>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import RunSidebar from '@/components/RunSidebar.vue'
import backendApi from '@/api/backend-api'
import Navbar from '@/components/Navbar.vue'
import Map from '@/components/Map.vue'

const runs = ref<any[]>([])
const selectedRun = ref<any | null>(null)

onMounted(async () => {
  const res = await backendApi.getPastRuns()
  runs.value = res.data
  if (runs.value.length > 0) selectedRun.value = runs.value[0]
})

function selectRun(run: any) {
  selectedRun.value = run
}

function formatDate(dateStr: string) {
  if (!dateStr) return '?'
  return new Date(dateStr).toLocaleString()
}

// Parse WKT LINESTRING to array of marker objects for start/end
function routeMarkers(route: string) {
  const points = parseRoute(route)
  if (points.length === 0) return []
  return [
    { latitude: points[0][1], longitude: points[0][0], label: 'Start' },
    { latitude: points[points.length - 1][1], longitude: points[points.length - 1][0], label: 'End' }
  ]
}

// Parse WKT LINESTRING to array of [lng, lat]
function parseRoute(route: string): [number, number][] {
  if (!route) return []
  const match = route.match(/\((.*)\)/)
  if (!match) return []
  return match[1].split(',').map(pair => {
    const [lng, lat] = pair.trim().split(' ').map(Number)
    return [lng, lat]
  })
}
</script>

<style scoped>
.runs-view {
  display: flex;
  height: 100vh;
}
.run-details {
  flex: 1;
  padding: 2rem;
  display: flex;
  flex-direction: column;
}
.map-container {
  flex: 1;
  min-height: 400px;
  margin-top: 1.5rem;
  position: relative;
}
.empty-map-msg {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #888;
  font-size: 1.2rem;
}
</style>
