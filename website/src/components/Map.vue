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
let startMarker: leaflet.Marker | null = null;
let endMarker: leaflet.Marker | null = null;
let markerObjs: leaflet.Marker[] = [];

const parseRoute = (route: string | undefined): [number, number][] => {
  if (!route) return [];
  const match = route.match(/\((.*)\)/);
  if (!match) return [];
  return match[1].split(',').map(pair => {
    const [lng, lat] = pair.trim().split(' ').map(Number);
    return [lat, lng];
  });
}

const fitMapToRoute = (points: [number, number][]) => {
  if (map && points.length > 0) {
    const bounds = leaflet.latLngBounds(points);
    map.fitBounds(bounds, { padding: [30, 30] });
  }
}

const clearMap = () => {
  if (polyline && map) {
    map.removeLayer(polyline);
    polyline = null;
  }
  if (startMarker && map) {
    map.removeLayer(startMarker);
    startMarker = null;
  }
  if (endMarker && map) {
    map.removeLayer(endMarker);
    endMarker = null;
  }
  markerObjs.forEach(m => map && map.removeLayer(m));
  markerObjs = [];
}

const drawRoute = () => {
  if (!map) return;
  clearMap();
  const points = parseRoute(props.route);
  if (points.length > 0) {
    polyline = leaflet.polyline(points, { color: "red", weight: 5 }).addTo(map);
    fitMapToRoute(points);
    if (props.markers && props.markers.length) {
      props.markers.forEach((m, idx) => {
        // Use colored circle markers for start/end
        const color = idx === 0 ? "#1abc1a" : "#1a4abc";
        const marker = leaflet.circleMarker([m.latitude, m.longitude], {
          radius: 8,
          color,
          fillColor: color,
          fillOpacity: 1,
          weight: 2,
        })
          .addTo(map)
          .bindPopup(m.label || (idx === 0 ? "Start" : "End"));
        markerObjs.push(marker);
      });
    } else {
      startMarker = leaflet.circleMarker(points[0], {
        radius: 8,
        color: "#1abc1a",
        fillColor: "#1abc1a",
        fillOpacity: 1,
        weight: 2,
      }).addTo(map).bindPopup("Start");
      endMarker = leaflet.circleMarker(points[points.length - 1], {
        radius: 8,
        color: "#1a4abc",
        fillColor: "#1a4abc",
        fillOpacity: 1,
        weight: 2,
      }).addTo(map).bindPopup("End");
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
