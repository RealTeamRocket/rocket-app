<template>
  <div class="popup-backdrop" @click.self="$emit('close')">
    <div class="popup-content">
      <h2>Invite a Friend</h2>
      <div v-if="loading">Loading friends...</div>
      <div v-if="successMessage" class="success-message">{{ successMessage }}</div>
      <div v-else>
        <div v-if="friends.length === 0">No friends found.</div>
        <div class="friend-list">
          <div
            v-for="friend in friends"
            :key="friend.id"
            class="friend-item"
            @click="invite(friend)"
          >
            {{ friend.username }}
          </div>
        </div>
      </div>
      <button class="cancel-button" @click="$emit('close')">Cancel</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import backendApi from '@/api/backend-api';

const props = defineProps<{ challengeId: string }>();
const emit = defineEmits(['close', 'invited']);

const friends = ref<{ id: string; username: string; image_name: string }[]>([]);
const loading = ref(true);

const successMessage = ref('');

const fetchFriends = async () => {
  loading.value = true;
  try {
    const res = await backendApi.getFriends();
    friends.value = res.data.map((f: any) => ({
      id: f.id,
      username: f.username,
      image_name: f.image_name
        ? `data:image/png;base64,${f.image_name}`
        : '/default-avatar.png',
    }));
  } catch (e) {
    friends.value = [];
  }
  loading.value = false;
};

const invite = async (friend: { id: string }) => {
  try {
    await backendApi.inviteFriendToChallenge(props.challengeId, friend.id);
    successMessage.value = 'Invitation sent!';
    emit('invited', friend);
    setTimeout(() => {
      emit('close');
      successMessage.value = '';
    }, 1200);
  } catch (e) {
    alert('Failed to invite friend.');
  }
};

onMounted(fetchFriends);
</script>

<style scoped>
.popup-backdrop {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.3); display: flex; align-items: center; justify-content: center;
  z-index: 1000;
}
.popup-content {
  background: #fff;
  padding: 2rem;
  border-radius: 1rem;
  min-width: 300px;
  display: flex;
  flex-direction: column;
  max-height: 80vh;
  overflow-y: auto;
}
.friend-list {
  display: flex;
  flex-direction: column;
  flex-wrap: wrap; gap: 0.5rem;
  margin: 1rem 0;
  width: 100%;
}
.friend-item {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 0.5rem;
  transition: background 0.2s;
  border: 1px solid #e5e7eb;
  width: 100%;
  box-sizing: border-box;
}
.friend-item:hover { background: #f3f4f6; }
.friend-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%; object-fit: cover; }
.cancel-button {
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  border: none;
  background: #4f46e5;
  color: #fff;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s;
  margin-top: 1rem;

}
.cancel-button:hover {
  background: #ef4444;
}
.success-message {
  color: #22c55e;
  font-weight: bold;
  margin: 1rem 0;
}
</style>
