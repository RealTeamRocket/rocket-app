<template>
  <Navbar />
  <div class="friendlist-page">
    <FriendSearchBar v-model:search="search" />
    <div class="friend-grid">
      <FriendCard
        v-for="friend in filteredFriends"
        :key="friend.id"
        :friend="friend"
        @unfollow="unfollowFriend"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import Navbar from '@/components/Navbar.vue';
import FriendSearchBar from '@/components/FriendSearchBar.vue';
import FriendCard from '@/components/FriendCard.vue';
import backendApi from '@/api/backend-api';

const search = ref('');
const friends = ref<any[]>([]);

const fetchFriends = async () => {
  const res = await backendApi.getFriends();
  friends.value = res.data.map((f: any) => ({
    id: f.ID || f.id,
    username: f.Username || f.username,
    email: f.Email || f.email,
    rocketPoints: f.RocketPoints || f.rocketPoints,
    image: f.ImageData ? `data:image/png;base64,${f.ImageData}` : undefined,
  }));
};
onMounted(fetchFriends);

const filteredFriends = computed(() =>
  friends.value.filter(f =>
    f.username.toLowerCase().includes(search.value.toLowerCase())
  )
);

const unfollowFriend = async (id: string) => {
  const friend = friends.value.find(f => f.id === id);
  if (!friend) return;
  await backendApi.deleteFriend(friend.username);
  friends.value = friends.value.filter(f => f.id !== id);
};
</script>

<style scoped>
.friendlist-page {
  padding: 2rem;
  max-width: 900px;
  margin: 0 auto;
}
.friend-grid {
  display: grid;
  grid-template-columns: repeat(2, 700px); 
  justify-content: center;                
  gap: 1.5rem;
  width: 100%;
}
@media (max-width: 950px) {
  .friend-grid {
    grid-template-columns: 1fr;
  }
}
</style>