<template>
  <div class="plan-run-map">
    <div id="plan-map" class="plan-map"></div>
    <div class="plan-controls">
      <input
        v-model="runName"
        class="run-name-input"
        type="text"
        placeholder="Enter run name"
        maxlength="60"
      />
      <div class="distance-label" v-if="points.length > 1">
        Total Distance: <strong>{{ totalDistance.toFixed(2) }}</strong> km
      </div>
      <div class="plan-buttons">
        <button class="clear-btn" @click="clearRoute" :disabled="points.length === 0">Clear</button>
        <button class="save-btn" @click="saveRun" :disabled="!canSave">Save</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import leaflet from 'leaflet'

const emit = defineEmits<{
  (e: 'save', payload: { name: string, points: [number, number][], distance: number }): void
}>()

const points = ref<[number, number][]>([])
const runName = ref('')
const totalDistance = ref(0)

let map: leaflet.Map | null = null
let polyline: leaflet.Polyline | null = null
let markerObjs: leaflet.Marker[] = []

const haversineDistance = (a: [number, number], b: [number, number]) => {
  // [lat, lng] in degrees
  const toRad = (x: number) => (x * Math.PI) / 180
  const R = 6371 // km
  const dLat = toRad(b[0] - a[0])
  const dLng = toRad(b[1] - a[1])
  const lat1 = toRad(a[0])
  const lat2 = toRad(b[0])
  const h =
    Math.sin(dLat / 2) ** 2 +
    Math.cos(lat1) * Math.cos(lat2) * Math.sin(dLng / 2) ** 2
  return 2 * R * Math.asin(Math.sqrt(h))
}

const updatePolyline = () => {
  if (polyline && map) {
    map.removeLayer(polyline)
    polyline = null
  }
  markerObjs.forEach(m => map && map.removeLayer(m))
  markerObjs = []
  if (points.value.length > 0 && map) {
    // Draw markers
    points.value.forEach((pt, idx) => {
      const marker = leaflet.circleMarker(pt, {
        radius: 7,
        color: idx === 0 ? '#1abc1a' : idx === points.value.length - 1 ? '#1a4abc' : '#4a90e2',
        fillColor: idx === 0 ? '#1abc1a' : idx === points.value.length - 1 ? '#1a4abc' : '#4a90e2',
        fillOpacity: 1,
        weight: 2,
      })
        .addTo(map)
        .bindPopup(idx === 0 ? 'Start' : idx === points.value.length - 1 ? 'End' : `Point ${idx + 1}`)
      markerObjs.push(marker)
    })
    // Draw polyline
    if (points.value.length > 1) {
      polyline = leaflet.polyline(points.value, {
        color: '#4a90e2',
        weight: 5,
        opacity: 0.9,
      }).addTo(map)
    }
    // Fit map to route
    if (points.value.length > 1) {
      const bounds = leaflet.latLngBounds(points.value)
      map.fitBounds(bounds, { padding: [30, 30] })
    } else if (points.value.length === 1) {
      map.setView(points.value[0], 15)
    }
  }
  // Calculate distance
  let dist = 0
  for (let i = 1; i < points.value.length; ++i) {
    dist += haversineDistance(points.value[i - 1], points.value[i])
  }
  totalDistance.value = dist
}

const onMapClick = (e: any) => {
  points.value.push([e.latlng.lat, e.latlng.lng])
  updatePolyline()
}

const clearRoute = () => {
  points.value = []
  updatePolyline()
}

const saveRun = () => {
  if (!canSave.value) return
  emit('save', { name: runName.value.trim(), points: points.value, distance: totalDistance.value })

  clearRoute()
  runName.value = ''
}

const canSave = ref(false)
watch([runName, points], () => {
  canSave.value = runName.value.trim().length > 0 && points.value.length > 1
})

onMounted(() => {
  map = leaflet.map('plan-map').setView([47.41322, -1.219482], 13)
  leaflet.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map)
  map.on('click', onMapClick)
})

</script>

<style scoped>
.plan-run-map {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  width: 100%;
  height: 100%;
}
.plan-map {
  width: 100%;
  height: 400px;
  min-height: 300px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
  margin-bottom: 1rem;
}
.plan-controls {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 1.5rem;
  padding: 0.5rem 0;
}
.run-name-input {
  font-size: 1.1rem;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  border: 1px solid #c0d6f7;
  min-width: 220px;
  max-width: 320px;
  outline: none;
}
.distance-label {
  font-size: 1.08rem;
  color: #4a90e2;
}
.plan-buttons {
  display: flex;
  gap: 0.8rem;
}
.save-btn, .clear-btn {
  font-size: 1rem;
  padding: 0.45rem 1.2rem;
  border-radius: 6px;
  border: none;
  cursor: pointer;
  font-weight: 600;
  transition: background 0.2s, color 0.2s;
}
.save-btn {
  background: #4a90e2;
  color: #fff;
}
.save-btn:disabled {
  background: #b3d1f7;
  color: #eee;
  cursor: not-allowed;
}
.clear-btn {
  background: #f7fafd;
  color: #4a90e2;
  border: 1px solid #c0d6f7;
}
.clear-btn:disabled {
  color: #bbb;
  border-color: #eee;
  cursor: not-allowed;
}
</style>
