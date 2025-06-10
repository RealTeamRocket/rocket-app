<template>
  <div class="challenge-list-container">
    <h3>Active Challenges</h3>
    <div>
      <ChallengeCard
        v-for="challenge in challenges"
        :key="challenge.id"
        :challenge="challenge"
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
</style>