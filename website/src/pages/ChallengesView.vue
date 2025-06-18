<template>
  <div class="page-wrapper">
    <Navbar />
    <div class="challenges-view">
      <h1>Challenges</h1>
      <ChallengeList :challenges="challenges"
        @complete="handleCompleteChallenge"
      @invite="handleInvite"/>
    <FriendInvitePopup
      v-if="showInvitePopup"
      :challenge-id="inviteChallengeId!"
      @close="showInvitePopup = false"/> 
  </div>
  <Footer />
  <DailyChallengeProgress :completed="completedCount" :total="totalCount" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import Navbar from '@/components/Navbar.vue';
import ChallengeList from '@/components/challenges/ChallengeList.vue';
import DailyChallengeProgress from '@/components/challenges/DailyChallengeProgress.vue';
import FriendInvitePopup from '@/components/challenges/FriendInvitePopup.vue';
import backendApi from '@/api/backend-api';
import Footer from '@/components/footer/Footer.vue'

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
.page-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  /* Make space for the fixed progress bar */
  padding-bottom: 3.5rem;
}
.challenges-view {
  flex: 1;
  padding: 2rem;
}
.challenges-view h1 {
  text-align: center;
}
.footer {
  margin-top: auto;
  /* Add extra bottom padding so it's scrollable above the progress bar */
  padding-bottom: 2.5rem;
}
</style>
