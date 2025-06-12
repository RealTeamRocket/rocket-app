<template>
  <Navbar />
  <div class="friendlist-page">
    <FriendSearchBar v-model:search="search" />
    <div v-if="loading" class="loading-indicator">
      Loading...
    </div>
    <template v-else>
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
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import Navbar from '@/components/Navbar.vue';
import FriendSearchBar from '@/components/FriendSearchBar.vue';
import FriendCard from '@/components/FriendCard.vue';
import backendApi from '@/api/backend-api';

interface User {
  id: string;
  username: string;
  email: string;
  rocketPoints: number;
  steps?: number;
  image?: string;
}

const search = ref('');
const friends = ref<User[]>([]);
const allUsers = ref<User[]>([]);

const fetchFriends = async () => {
  try {
    const res = await backendApi.getFriends();
    friends.value = res.data.map((f: any) => ({
      id: f.id,
      username: f.username,
      email: f.email,
      rocketPoints: f.rocket_points,
      steps: f.steps,
      image: f.image_data ? `data:image/png;base64,${f.image_data}` : undefined,
    })) as User[];
  } catch (e) {
    console.error('Failed to fetch friends:', e);
    friends.value = [];
  }
  
};

const fetchAllUsers = async () => {
  try {
    const res = await backendApi.getAllUsers();
    allUsers.value = res.data.map((u: any) => ({
      id: u.id,
      username: u.username,
      email: u.email,
      rocketPoints: u.rocket_points,
      steps: u.steps,
      image: u.image_data ? `data:image/png;base64,${u.image_data}` : undefined,
    }));
  } catch (e) {
    console.error('Failed to fetch all users:', e);
    allUsers.value = [];
  }
  
};

const loading = ref(true);

onMounted(async () => {
  loading.value = true;
  await Promise.all([fetchFriends(), fetchAllUsers()]);
  loading.value = false;
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

const addFriend = async (user: User) => {
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
.loading-indicator {
  text-align: center;
  font-size: 1.2rem;
  margin: 2rem 0;
}
</style>