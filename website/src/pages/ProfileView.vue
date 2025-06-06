<template>
  <Navbar />
  <div class="profile-page">
    <h1>Profile</h1>
    <p>This is the profile page. Here you will see the stats and infos about a specific user.</p>
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
</script>

<style scoped>
.profile-page {
  padding: 2rem;
}
</style>
