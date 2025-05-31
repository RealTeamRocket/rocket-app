<template>
  <Navbar />
  <div class="challenges-view">
    <h1>Challenges</h1>
    <p>This is the Challenges page. Here you will see all your active and completed challenges in the future.</p>
  </div>
</template>

<script setup lang="ts">
import Navbar from '../components/Navbar.vue'
</script>

<style scoped>
.challenges-view {
  padding: 2rem;
}
</style>
