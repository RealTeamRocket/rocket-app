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
    <button
      :class="{ active: tab === 'planned' }"
      @click="tab = 'planned'"
    >
      Planned Runs
    </button>
  </div>
  <div v-if="feedback" :class="['feedback-message', feedback.type]">
    {{ feedback.message }}
  </div>
  <div class="runs-view">
    <RunSidebar
      v-if="tab === 'past'"
      :runs="runs ?? []"
      :selected-id="selectedRun?.id"
      @select="selectRun"
    />
    <PlannedSidebar
      v-if="tab === 'planned'"
      :runs="plannedRuns ?? []"
      :selected-id="selectedPlannedRun?.id"
      @select="selectPlannedRun"
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
      <template v-else-if="tab === 'planned'">
        <h2 v-if="selectedPlannedRun">Planned Run Details</h2>
        <div v-if="selectedPlannedRun">
          <strong>Name:</strong> {{ selectedPlannedRun.name }}<br>
          <strong>Distance:</strong> {{ selectedPlannedRun.distance?.toFixed(2) ?? '?' }} km<br>
          <strong>Created:</strong>
          <span style="white-space:nowrap; font-size:0.97em;">
            {{ new Date(selectedPlannedRun.created_at).toLocaleString(undefined, { dateStyle: 'medium', timeStyle: 'short' }) }}
          </span>
        </div>
        <div class="map-container">
          <Map
            v-if="selectedPlannedRun"
            :route="selectedPlannedRun.route"
            :markers="routeMarkers(selectedPlannedRun.route)"
          />
          <div v-else class="empty-map-msg">
            <p>Select a planned run to see its route.</p>
          </div>
          <ElevationProfile
            v-if="selectedPlannedRun"
            :coordinates="parseRoute(selectedPlannedRun.route)"
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
import RunSidebar from '@/components/runs/RunSidebar.vue'
import PlannedSidebar from '@/components/runs/plan/PlannedSidebar.vue'
import backendApi from '@/api/backend-api'
import Navbar from '@/components/Navbar.vue'
import Map from '@/components/runs/Map.vue'
import ElevationProfile from '@/components/runs/ElevationProfile.vue'
import { parseRoute } from '@/utils/routes'
import PlanRunMap from '@/components/runs/plan/PlanRunMap.vue'

const tab = ref<'past' | 'plan' | 'planned'>('past')

const runs = ref<any[]>([])
const selectedRun = ref<any | null>(null)
const plannedRuns = ref<any[]>([])
const selectedPlannedRun = ref<any | null>(null)

onMounted(async () => {
  const res = await backendApi.getPastRuns()
  runs.value = Array.isArray(res.data) ? res.data : []
  if (runs.value.length > 0) selectedRun.value = runs.value[0]
  const plannedRes = await backendApi.getPlannedRuns()
  plannedRuns.value = Array.isArray(plannedRes.data) ? plannedRes.data : []
  if (plannedRuns.value.length > 0) selectedPlannedRun.value = plannedRuns.value[0]
})

const selectRun = (run: any) => {
  selectedRun.value = run
}

const selectPlannedRun = (run: any) => {
  selectedPlannedRun.value = run
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

const feedback = ref<{ type: 'success' | 'error', message: string } | null>(null);

const handlePlanSave = async (payload: { name: string, points: [number, number][], distance: number }) => {
  // Convert points to WKT LINESTRING
  if (!payload.name.trim() || payload.points.length < 2) return;
  const wkt = `LINESTRING(${payload.points.map(([lat, lng]) => `${lng} ${lat}`).join(', ')})`;
  try {
    await backendApi.savePlannedRun(wkt, payload.name.trim(), payload.distance);
    const res = await backendApi.getPlannedRuns();
    plannedRuns.value = res.data;
    tab.value = 'planned';
    if (plannedRuns.value.length > 0) {
      selectedPlannedRun.value = plannedRuns.value[0];
    }
    feedback.value = { type: 'success', message: 'Planned run saved successfully!' };
    setTimeout(() => { feedback.value = null }, 2500);
  } catch (e) {
    feedback.value = { type: 'error', message: 'Failed to save planned run.' };
    setTimeout(() => { feedback.value = null }, 3000);
    console.error('Failed to save planned run', e);
  }
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
.feedback-message {
  position: fixed;
  top: 80px;
  left: 50%;
  transform: translateX(-50%);
  z-index: 1000;
  padding: 1rem 2rem;
  border-radius: 8px;
  font-size: 1.1rem;
  font-weight: 600;
  background: #e0ffe0;
  color: #217a21;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
  transition: opacity 0.2s;
}
.feedback-message.error {
  background: #ffe0e0;
  color: #a12121;
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
