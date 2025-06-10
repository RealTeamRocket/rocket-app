<template>
  <Navbar />
  <div class="challenges-view">
    <h1>Challenges</h1>
    <ChallengeList :challenges="challenges" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import Navbar from '@/components/Navbar.vue';
import ChallengeList from '@/components/ChallengeList.vue';
import backendApi from '@/api/backend-api';

const challenges = ref([]);

onMounted(async () => {
  try {
    const response = await backendApi.getChallenges();
    challenges.value = response.data; // response.data enthält das Array der Challenges
  } catch (e) {
    console.error('Fehler beim Laden der Challenges:', e);
  }
});
</script>

<style scoped>
.challenges-view {
  padding: 2rem;
}
</style>
