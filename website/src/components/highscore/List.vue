<script setup lang="ts">
import { getInitials, getColor } from '@/utils/userUtils'

interface RankedUser {
  id: string
  username: string
  rocket_points: number
  imageData?: string
  isFriend?: boolean
}

defineProps<{
  users: RankedUser[]
  openProfile: (user: RankedUser) => void
  addFriend: (username: string) => void
  currentUsername: string
}>()
</script>
<template>
  <div class="list-container">
    <div
      class="list-user-card"
      v-for="(user, idx) in users"
      :key="user.id || idx"
      @click="openProfile(user)"
    >
      <div class="list-avatar-col">
        <img
          v-if="user.imageData"
          :src="`data:image/*;base64,${user.imageData}`"
          alt="User Icon"
          class="list-avatar"
        />
        <div
          v-else
          class="list-avatar list-avatar-initials"
          :style="{ backgroundColor: getColor(user.username) }"
        >
          {{ getInitials(user.username) }}
        </div>
      </div>
      <div class="list-info-col">
        <div class="list-username">{{ user.username }}</div>
        <div class="list-rocketpoints">
          <img
            src="/src/assets/icons/rocket.svg"
            alt="Rocket"
            class="rocket-icon"
          />
          {{ user.rocket_points }}
        </div>
      </div>
      <div class="list-action-col">
        <template v-if="user.isFriend">
          <span class="friend-badge">
            <svg class="friend-badge-icon" viewBox="0 0 20 20" fill="currentColor">
              <circle cx="10" cy="10" r="10" fill="#34c759"/>
              <path d="M7.5 10.5l2 2 3-3" stroke="#fff" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
            <span class="friend-badge-label">Friend</span>
          </span>
        </template>
        <template v-else-if="user.username === currentUsername">
          <span></span>
        </template>
        <template v-else>
          <button class="add-btn" @click.stop="addFriend(user.username)">Add</button>
        </template>
      </div>
    </div>
  </div>
</template>

<style scoped>
.list-container {
  width: 100%;
  max-width: 420px;
  margin-top: 2rem;
  display: flex;
  flex-direction: column;
  gap: 1.1rem;
}

.list-user-card {
  display: flex;
  align-items: center;
  background: #f8fafc;
  border-radius: 18px;
  border: 1.5px solid #e3eaf3;
  box-shadow: 0 1px 6px rgba(30,60,114,0.06);
  padding: 0.7em 1.2em;
  transition: box-shadow 0.18s, transform 0.13s;
  cursor: pointer;
  min-height: 60px;
  gap: 1.2rem;
}
.list-user-card:hover {
  box-shadow: 0 4px 16px rgba(30,60,114,0.13);
  transform: translateY(-2px) scale(1.02);
  background: #f0f6fa;
}

.list-avatar-col {
  flex: 0 0 44px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.list-avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  object-fit: cover;
  background: #f5f5f5;
  border: 1.5px solid #e3eaf3;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 1.15em;
  color: #fff;
  text-transform: uppercase;
  letter-spacing: 0.02em;
  user-select: none;
}
.list-avatar-initials {
  background: #bfc1c2;
  border: 1.5px solid #e3eaf3;
}

.list-info-col {
  flex: 1 1 auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 0.2em;
}
.list-username {
  font-size: 1.08em;
  font-weight: 600;
  color: #223;
  margin-bottom: 0.1em;
}
.list-rocketpoints {
  display: flex;
  align-items: center;
  font-size: 0.98em;
  color: #2a5298;
  font-weight: 500;
  gap: 0.3em;
}
.rocket-icon {
  width: 1.1em;
  height: 1.1em;
  margin-right: 0.15em;
  vertical-align: middle;
}

.list-action-col {
  flex: 0 0 48px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.add-btn {
  background: linear-gradient(90deg, #2196f3 0%, #00bcd4 100%);
  color: #fff;
  border: none;
  border-radius: 16px;
  padding: 0.35em 1.1em;
  font-weight: 600;
  font-size: 0.98em;
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(33, 150, 243, 0.10);
  transition: background 0.2s, transform 0.1s;
}
.add-btn:hover {
  background: linear-gradient(90deg, #1976d2 0%, #0097a7 100%);
  transform: translateY(-1px) scale(1.03);
}
.friend-icon {
  display: none;
}

.friend-badge {
  display: flex;
  align-items: center;
  gap: 0.3em;
  background: #eaffea;
  border-radius: 16px;
  padding: 0.2em 0.8em 0.2em 0.4em;
  font-size: 0.98em;
  font-weight: 600;
  color: #218838;
  border: 1.5px solid #34c759;
  box-shadow: 0 1px 4px rgba(52,199,89,0.08);
}

.friend-badge-icon {
  width: 20px;
  height: 20px;
  margin-right: 0.1em;
  flex-shrink: 0;
}

.friend-badge-label {
  font-size: 1em;
  font-weight: 600;
  color: #218838;
  letter-spacing: 0.01em;
}
</style>
