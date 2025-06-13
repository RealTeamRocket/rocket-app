<template>
  <Navbar />
  <div class="download-view">
    <h1>Download the Rocket App</h1>
    <p>You can download the latest version of the Rocket App for Android below:</p>
    <a
      class="download-btn"
      :href="apkUrl"
      download
    >
      Download APK
    </a>
    <p class="note">If you have trouble installing, make sure to allow installation from unknown sources in your device settings.</p>
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
const apkUrl = '/apk/rocket-app-latest.apk'
</script>

<style scoped>
.download-view {
  max-width: 500px;
  margin: 2rem auto;
  padding: 2rem;
  background: #fff;
  border-radius: 1rem;
  box-shadow: 0 2px 12px rgba(30,60,114,0.08);
  text-align: center;
}
.download-btn {
  display: inline-block;
  margin: 1.5rem 0;
  padding: 0.8rem 2.2rem;
  background: #4a90e2;
  color: #fff;
  font-size: 1.2rem;
  font-weight: 600;
  border-radius: 0.6rem;
  text-decoration: none;
  transition: background 0.18s;
}
.download-btn:hover {
  background: #2a5298;
}
.note {
  color: #64748b;
  font-size: 0.95rem;
  margin-top: 1.2rem;
}
</style>
