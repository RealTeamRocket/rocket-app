<template>
  <Navbar />
  <div class="challenges-view">
    <h1>Challenges</h1>
    <ChallengeList :challenges="challenges" @complete="handleCompleteChallenge" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import Navbar from '@/components/Navbar.vue';
import ChallengeList from '@/components/ChallengeList.vue';
import backendApi from '@/api/backend-api';

const challenges = ref([]);

const fetchChallenges = async () => {
  try {
    const response = await backendApi.getChallenges();
    challenges.value = response.data;
  } catch (e) {
    console.error('Failed to load challenges', e);
  }
};

const handleCompleteChallenge = async (payload: { id: string, points: number }) => {
  try {
    await backendApi.completeChallenge(payload.id, payload.points);
    await fetchChallenges();
  } catch (e) {
    console.error('Failed to complete challenge', e);
  }
};

onMounted(fetchChallenges);
</script>

<style scoped>
.challenges-view {
  padding: 2rem;
}
.challenges-view h1 {
  text-align: center;
}
</style>
