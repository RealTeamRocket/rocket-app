<template>
  <div class="challenge-card">
    <div class="challenge-card-row">
      <div>
        <div class="challenge-card-title">{{ challenge.text }}</div>
        <div class="challenge-card-value">🚀Rocket-Points: {{ challenge.points }}</div>
      </div>
      <div class="challenge-card-actions">
        <button class="complete-button" @click="$emit('complete', { id: challenge.id, points: challenge.points })">
          Complete
        </button>
        <button class="complete-button" @click="$emit('invite', challenge.id)">
            Invite friend
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{ challenge: { id: string, text: string, points: number } }>();
</script>

<style scoped>
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
.challenge-card-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}
.challenge-card-title {
  font-size: 1.5rem;
  font-weight: 500;
  color: #2a5298;
  margin-bottom: 0.5rem;
  letter-spacing: 0.5px;
}
.challenge-card-value {
  font-size: 1rem;
  font-weight: bold;
  color: #1e3c72;
  margin-bottom: 0.25rem;
}
.challenge-card-actions {
  display: flex;
  flex-direction: row;
  gap: 0.5rem;
}
.complete-button {
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  border: none;
  background: #4f46e5;
  color: #fff;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s;
}
.complete-button:hover {
  background: #22c55e;
}
</style>