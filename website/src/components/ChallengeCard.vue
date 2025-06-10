<template>
  <div class="challenge-card">
    <div class="challenge-card-title">{{ challenge.text }}</div>
    <div class="challenge-card-value">🚀Rocket-Points: {{ challenge.points }}</div>
    <button class="complete-button" @click="$emit('complete', { id: challenge.id, points: challenge.points })">Complete</button>
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
.challenge-card-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: #2a5298;
  margin-bottom: 0.5rem;
  font-weight: 500;
  letter-spacing: 0.5px;
}
.challenge-card-value {
  font-size: 1rem;
  font-weight: bold;
  color: #1e3c72;
  margin-bottom: 0.25rem;
}
.challenge-card {
  background: linear-gradient(90deg, #e0e7ff 0%, #f3f8ff 100%);
  border-radius: 1rem;
  box-shadow: 0 2px 8px rgba(30,60,114,0.08);
  padding: 1.5rem 2rem;
  min-width: 180px;
  display: flex;
  flex-direction: column;
  margin-bottom: 1.2rem;
}
</style>