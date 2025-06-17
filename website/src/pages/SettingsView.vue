<template>
  <div class="page-wrapper">
    <Navbar />
    <div class="settings-page">
      <h1>Settings</h1>
      <p>This is the Settings page. Here you will see settings and can change them.</p>
    </div>
  </div>
  <Footer />
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
import Footer from '@/components/footer/Footer.vue'
</script>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
.settings-page {
  flex: 1;
  padding: 2rem;
}
.footer {
  margin-top: auto;
}
</style>
