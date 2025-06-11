<template>
  <Navbar />
  <div class="challenges-view">
    <h1>Challenges</h1>
    <ChallengeList :challenges="challenges"
      @complete="handleCompleteChallenge"
      @invite="handleInvite"/>
    <FriendInvitePopup
      v-if="showInvitePopup"
      :challenge-id="inviteChallengeId"
      @close="showInvitePopup = false"/>
  </div>
  <DailyChallengeProgress :completed="completedCount" :total="totalCount" />
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import Navbar from '@/components/Navbar.vue';
import ChallengeList from '@/components/ChallengeList.vue';
import DailyChallengeProgress from '@/components/DailyChallengeProgress.vue';
import FriendInvitePopup from '@/components/FriendInvitePopup.vue';
import backendApi from '@/api/backend-api';

const challenges = ref<{ id: string; text: string; points: number }[]>([]);
const completedCount = ref(0);
const totalCount = ref(0);
const showInvitePopup = ref(false);
const inviteChallengeId = ref<string | null>(null);

const fetchChallenges = async () => {
  try {
    const response = await backendApi.getChallenges();
    challenges.value = response.data;
  } catch (e: any) {
    if (e.response && e.response.status === 404) {
      challenges.value = [];
    } else {
      console.error('Failed to load challenges', e);
    }
  }
};

const fetchProgress = async () => {
  try {
    const response = await backendApi.getChallengeProgress();
    completedCount.value = response.data.completed;
    totalCount.value = response.data.total;
  } catch (e: any) {
    completedCount.value = 0;
    totalCount.value = 0;
    console.error('Failed to load challenge progress', e);
  }
};

const handleCompleteChallenge = async (payload: { id: string, points: number }) => {
  try {
    await backendApi.completeChallenge(payload.id, payload.points);
    challenges.value = challenges.value.filter(c => c.id !== payload.id);
    await fetchChallenges();
    await fetchProgress();
  } catch (e) {
    console.error('Failed to complete challenge', e);
  }
};

const handleInvite = (challengeId: string) => {
  inviteChallengeId.value = challengeId;
  showInvitePopup.value = true;
};

onMounted(async () => {
  await fetchChallenges();
  await fetchProgress();
});

</script>

<style scoped>
.challenges-view {
  padding: 2rem;
}
.challenges-view h1 {
  text-align: center;
}
</style>
