<template>
  <div class="challenge-card">
    <h3>{{ challenge.text }}</h3>
    <p> 🚀Rocket-Points: {{ challenge.points }}</p>
    <!--  Later Buttons -->
  </div>
</template>

<script setup lang="ts">
defineProps<{ challenge: { id: string, text: string, points: number } }>();
</script>

<style scoped>
.challenge-card {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.07);
  padding: 1rem 1.5rem;
  margin-bottom: 1.2rem;
}
</style>