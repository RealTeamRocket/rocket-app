<template>
  <div id="map"></div>
</template>

<script setup lang="ts">
import { onMounted, watch } from "vue";
import leaflet from "leaflet";

const props = defineProps<{
  route?: string,
  markers?: { latitude: number, longitude: number, label?: string }[]
}>();

let map: leaflet.Map | null = null;
let polyline: leaflet.Polyline | null = null;
let markerObjs: leaflet.Marker[] = [];

function parseRoute(route: string | undefined): [number, number][] {
  if (!route) return [];
  const match = route.match(/\((.*)\)/);
  if (!match) return [];
  return match[1].split(',').map(pair => {
    const [lng, lat] = pair.trim().split(' ').map(Number);
    return [lat, lng];
  });
}

function fitMapToRoute(points: [number, number][]) {
  if (map && points.length > 0) {
    const bounds = leaflet.latLngBounds(points);
    map.fitBounds(bounds, { padding: [30, 30] });
  }
}

function clearMap() {
  if (polyline && map) {
    map.removeLayer(polyline);
    polyline = null;
  }
  markerObjs.forEach(m => map && map.removeLayer(m));
  markerObjs = [];
}

function drawRoute() {
  if (!map) return;
  clearMap();
  const points = parseRoute(props.route);
  if (points.length > 0) {
    polyline = leaflet.polyline(points, { color: "blue", weight: 5 }).addTo(map);
    fitMapToRoute(points);
    // Add start/end markers if available
    if (props.markers && props.markers.length) {
      props.markers.forEach((m, idx) => {
        const marker = leaflet.marker([m.latitude, m.longitude])
          .addTo(map)
          .bindPopup(m.label || (idx === 0 ? "Start" : "End"));
        markerObjs.push(marker);
      });
    } else {
      // Default: add start/end markers
      leaflet.marker(points[0]).addTo(map).bindPopup("Start");
      leaflet.marker(points[points.length - 1]).addTo(map).bindPopup("End");
    }
  }
}

onMounted(() => {
  map = leaflet.map("map").setView([47.41322, -1.219482], 13);
  leaflet
    .tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
      maxZoom: 19,
      attribution:
        '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
    })
    .addTo(map);
  drawRoute();
});

watch(
  () => [props.route, props.markers],
  () => {
    drawRoute();
  }
);
</script>

<style scoped>
#map {
  width: 100%;
  height: 400px;
  min-height: 300px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
}
</style>
