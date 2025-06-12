<template>
  <div class="friend-card">
    <img v-if="friend.image" :src="friend.image" class="friend-avatar" />
    <div v-else class="friend-avatar-placeholder">{{ initials }}</div>
    <div class="friend-info">
      <div class="friend-name">{{ friend.username }}</div>
      <div class="friend-email">{{ friend.email }}</div>
      <div v-if="isFriend">
        <div class="friend-points">🚀 {{ friend.rocketPoints ?? 0 }}</div>
        <div class="friend-steps">👣 {{ friend.steps ?? 0 }}</div>
      </div>
    </div>
    <button v-if="isFriend" class="unfollow-btn" @click="$emit('unfollow', friend.id)"> Unfollow </button>
    <button v-else class="follow-btn" @click="$emit('add-friend', friend)"> Follow </button>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  friend: { id: string, username: string, email: string, rocketPoints: number, image?: string, steps?: number },
  isFriend?: boolean
}>();

const initials = props.friend.username
  ? props.friend.username.split(' ').map(n => n[0]).join('').toUpperCase()
  : '';
</script>

<style scoped>
.friend-card {
  display: flex;
  align-items: center;
  background: linear-gradient(90deg, #e0e7ff 0%, #f3f8ff 100%);
  border-radius: 1rem;
  box-shadow: 0 2px 8px rgba(30,60,114,0.08);
  padding: 1rem 1.5rem;
  min-width: 600px;
  max-width: 700px;
  width: 100%;
  margin: 0.5rem;
  position: relative;
}
.friend-avatar, .friend-avatar-placeholder {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  background: #c7d2fe;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  color: #374151;
  margin-right: 1rem;
}
.friend-info {
  flex: 1;
}
.friend-name {
  font-size: 2.0rem;
  font-weight: 600;
  color: #2a5298;
}
.friend-email {
  font-size: 0.95rem;
  color: #64748b;
}
.friend-points {
  font-size: 0.95rem;
  color: #64748b;
}
.friend-steps{
  font-size: 0.95rem;
  color: #64748b;
}
.unfollow-btn {
  background: #fff;
  border: 1px solid #ef4444;
  color: #ef4444;
  border-radius: 0.5rem;
  padding: 0.4rem 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}
.unfollow-btn:hover {
  background: #ef4444;
  color: #fff;
}
.follow-btn {
  background: #22c55e;
  border: 1px solid #22c55e;
  color: #fff;
  border-radius: 0.5rem;
  padding: 0.4rem 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}
.follow-btn:hover {
  background: #16a34a;
  border-color: #16a34a;
}
</style>