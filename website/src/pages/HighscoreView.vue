<template>
  <Navbar />
  <div class="highscore-page">
    <h1>Highscore</h1>
    <p>This is the Highscore page. Here you will see the top scores and rankings.</p>
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
</script>

<style scoped>
.highscore-page {
  padding: 2rem;
}
</style>
