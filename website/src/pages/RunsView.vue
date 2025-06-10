<template>
  <Navbar />
  <div class="tab-switcher">
    <button
      :class="{ active: tab === 'past' }"
      @click="tab = 'past'"
    >
      Past Runs
    </button>
    <button
      :class="{ active: tab === 'plan' }"
      @click="tab = 'plan'"
    >
      Plan a Run
    </button>
  </div>
  <div class="runs-view">
    <RunSidebar
      v-if="tab === 'past'"
      :runs="runs"
      :selected-id="selectedRun?.id"
      @select="selectRun"
    />
    <main class="run-details">
      <template v-if="tab === 'past'">
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
          <ElevationProfile
            v-if="selectedRun"
            :coordinates="parseRoute(selectedRun.route)"
          />
        </div>
      </template>
      <template v-else>
        <PlanRunMap @save="handlePlanSave" />
      </template>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import RunSidebar from '@/components/RunSidebar.vue'
import backendApi from '@/api/backend-api'
import Navbar from '@/components/Navbar.vue'
import Map from '@/components/Map.vue'
import ElevationProfile from '@/components/ElevationProfile.vue'
import { parseRoute } from '@/utils/routes'
import PlanRunMap from '@/components/plan/PlanRunMap.vue'
import PlannedSidebar from '@/components/plan/PlannedSidebar.vue'

const tab = ref<'past' | 'plan' | 'planned'>('past')

const runs = ref<any[]>([])
const selectedRun = ref<any | null>(null)
const plannedRuns = ref<any[]>([])
const selectedPlannedRun = ref<any | null>(null)

onMounted(async () => {
  const res = await backendApi.getPastRuns()
  runs.value = res.data
  if (runs.value.length > 0) selectedRun.value = runs.value[0]
})

const selectRun = (run: any) => {
  selectedRun.value = run
}

const formatDate = (dateStr: string) => {
  if (!dateStr) return '?'
  return new Date(dateStr).toLocaleString()
}

// Parse WKT LINESTRING to array of marker objects for start/end
const routeMarkers = (route: string) => {
  const points = parseRoute(route)
  if (points.length === 0) return []
  return [
    { latitude: points[0][0], longitude: points[0][1], label: 'Start' },
    { latitude: points[points.length - 1][0], longitude: points[points.length - 1][1], label: 'End' }
  ]
}

const handlePlanSave = (payload: { name: string, points: [number, number][] }) => {
 // Here you can send the planned run to the backend or store locally
 // Example: convert points to WKT LINESTRING and send to backend
 // For now, just log it
 console.log('Planned run saved:', payload)
}
</script>

<style scoped>
.tab-switcher {
  display: flex;
  gap: 1rem;
  padding: 1.5rem 2rem 0.5rem 2rem;
  background: #f7fafd;
  border-bottom: 1px solid #e0eaff;
}
.tab-switcher button {
  background: none;
  border: none;
  font-size: 1.1rem;
  font-weight: 600;
  color: #4a90e2;
  padding: 0.5rem 1.5rem;
  border-radius: 6px 6px 0 0;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
}
.tab-switcher button.active,
.tab-switcher button:hover {
  background: #e0eaff;
  color: #222;
}

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
.plan-run-placeholder {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: flex-start;
  padding: 2rem;
  color: #888;
}
</style>
