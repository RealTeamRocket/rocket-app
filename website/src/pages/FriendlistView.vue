<template>
  <Navbar />
  <div class="friendlist-page">
    <FriendSearchBar v-model:search="search" />
    <div v-if="search" class="search-results-section">
      <h3 class="result-headline">Search Results</h3>
      <div class="friend-grid">
        <FriendCard
          v-for="user in filteredUserResults"
          :key="user.id"
          :friend="user"
          :isFriend="false"
          @add-friend="addFriend"
        />
      </div>
    </div>
    <div class="friend-grid">
      <FriendCard
        v-for="friend in filteredFriends"
        :key="friend.id"
        :friend="friend"
        :isFriend="true"
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
const allUsers = ref<any[]>([]);

const fetchFriends = async () => {
  const res = await backendApi.getFriends();
  friends.value = res.data.map((f: any) => ({
    id: f.ID || f.id,
    username: f.Username || f.username,
    email: f.Email || f.email,
    rocketPoints: f.rocket_points,
    image: f.ImageData ? `data:image/png;base64,${f.ImageData}` : undefined,
  }));
};

const fetchAllUsers = async () => {
  const res = await backendApi.getAllUsers();
  allUsers.value = res.data.map((u: any) => ({
    id: u.ID || u.id,
    username: u.Username || u.username,
    email: u.Email || u.email,
    rocketPoints: u.rocket_points,
    image: u.ImageData ? `data:image/png;base64,${u.ImageData}` : undefined,
  }));
};

onMounted(() => {
  fetchFriends();
  fetchAllUsers();
});

const filteredFriends = computed(() =>
  friends.value.filter(f =>
    f.username.toLowerCase().includes(search.value.toLowerCase())
  )
);

const filteredUserResults = computed(() =>
  allUsers.value.filter(u =>
    u.username.toLowerCase().includes(search.value.toLowerCase()) &&
    !friends.value.some(f => f.id === u.id)
  )
);

const unfollowFriend = async (id: string) => {
  const friend = friends.value.find(f => f.id === id);
  if (!friend) return;
  await backendApi.deleteFriend(friend.username);
  friends.value = friends.value.filter(f => f.id !== id);
};

const addFriend = async (user: any) => {
  try {
    await backendApi.addFriend(user.username);
    await fetchFriends();
  } catch (e) {
    console.error('Failed to add friend:', e);
  }
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
.result-headline {
  text-align: center;
  margin-bottom: 1.5rem;
}
.search-results-section {
  margin-bottom: 6rem; 
}
</style>