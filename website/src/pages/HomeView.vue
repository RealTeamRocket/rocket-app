<template>
    <Navbar />
  <div class="container mt-5 text-center">
    <div v-if="!isLoggedIn">
      <h1>Welcome to Rocket App</h1>
      <p>Track your steps, reach your goals, and blast off to a healthier you!</p>
      <router-link to="/register" class="btn btn-primary me-2">Get Started</router-link>
      <router-link to="/login" class="btn btn-outline-primary">Login</router-link>
    </div>
    <div v-else>
      <h1>Welcome back, Astronaut!</h1>
      <p>Here’s your step dashboard. Ready to break your record?</p>
      <!-- Add dashboard widgets here -->
    </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from '../components/Navbar.vue'
import { onMounted } from 'vue'
import { useAuth } from '../utils/useAuth'

const { isLoggedIn, checkAuth } = useAuth()

onMounted(() => {
  checkAuth()
})
</script>


<style scoped>
/* Add custom styles here */
</style>
