<template>
  <div class="challenge-list-container">
    <div class="challenge-list-header">
      <h3>Active Challenges</h3>
      <button class="create-button" @click="$emit('addChallenge')">+</button>
    </div>
    <div>
      <ChallengeCard
        v-for="challenge in challenges"
        :key="challenge.id"
        :challenge="challenge"
        @complete="$emit('complete', $event)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import ChallengeCard from './ChallengeCard.vue';
defineProps<{ challenges: Array<{ id: string, text: string, points: number }> }>();
</script>

<style scoped>
.challenge-list-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}
h3 {
  text-align: center;
}
.create-button {
  padding: 0.4rem 1rem;
  border-radius: 0.5rem;
  border: none;
  background: #4f46e5;
  color: #fff;
  font-weight: 800;
  cursor: pointer;
  transition: background 0.2s;
  margin-left: 1rem;
}
.create-button:hover {
  background: #da13ab;
}
.challenge-list-header {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  margin-bottom: 1rem;
}
</style>
