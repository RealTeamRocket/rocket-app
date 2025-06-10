<template>
  <Navbar />
  <div class="challenges-view">
    <h1>Challenges</h1>
    <ChallengeList :challenges="challenges" 
    @complete="handleCompleteChallenge" 
    @addChallenge="showCreate" />
  </div>
    <DailyChallengeProgress :completed="completedCount" :total="totalCount" />
    <CreateChallengeDialog v-if="openDialog" @close="openDialog = false" @submit="submitChallenge" />
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import Navbar from '@/components/Navbar.vue';
import ChallengeList from '@/components/ChallengeList.vue';
import DailyChallengeProgress from '@/components/DailyChallengeProgress.vue';
import CreateChallengeDialog from '@/components/CreateChallengeDialog.vue';
import backendApi from '@/api/backend-api';

const MAX_DAILY_CHALLENGES = 5;
const openDialog = ref(false);
const challenges = ref<{ id: string; text: string; points: number }[]>([]);

const fetchChallenges = async () => {
  try {
    const response = await backendApi.getChallenges();
    challenges.value = response.data;
  }  catch (e: any) {
    // 404 means no challenges available
    if (e.response && e.response.status === 404) {
      challenges.value = [];
    } else {
      console.error('Failed to load challenges', e);
    }
  }
};

const handleCompleteChallenge = async (payload: { id: string, points: number }) => {
  try {
    await backendApi.completeChallenge(payload.id, payload.points);
    challenges.value = challenges.value.filter(c => c.id !== payload.id);
    await fetchChallenges();
  } catch (e) {
    console.error('Failed to complete challenge', e);
  }
};

const showCreate = () => {
  openDialog.value = true;
};

const submitChallenge = async (challenge: { title: string, description: string, points: number }) => {
  console.log('Submitting challenge:', challenge);
};

onMounted(fetchChallenges);

const completedCount = computed(() => MAX_DAILY_CHALLENGES - challenges.value.length);
const totalCount = computed(() => MAX_DAILY_CHALLENGES);

</script>

<style scoped>
.challenges-view {
  padding: 2rem;
}
.challenges-view h1 {
  text-align: center;
}
</style>
