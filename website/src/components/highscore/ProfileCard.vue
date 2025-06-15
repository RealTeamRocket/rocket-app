<script setup lang="ts">
import type { RankedUser } from '@/types'

defineProps<{
  user: RankedUser
  onClose: () => void
  onAddFriend: (username: string) => void
  currentUsername: string
}>()
</script>

<template>
  <div class="profile-dialog-overlay" @click.self="onClose">
    <div class="profile-dialog">
      <button class="close-btn" @click="onClose">Close</button>
      <div class="profile-img-col">
        <img :src="user.imageUrl || '/src/assets/icons/user.svg'" class="profile-avatar" />
      </div>
      <div class="profile-info-col">
        <h2>{{ user.username }}</h2>
        <p>Punkte: {{ user.rocket_points }}</p>
        <div class="profile-actions">
          <template v-if="user.isFriend">
            <img src="/src/assets/icons/user.svg" alt="Friend Icon" class="friend-icon" />
          </template>
          <template v-else-if="user.username === currentUsername">
            <span></span>
          </template>
          <template v-else>
            <button class="add-btn" @click.stop="onAddFriend(user.username)">Add</button>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.profile-dialog-overlay {
position: fixed;
top: 0; left: 0; right: 0; bottom: 0;
background: rgba(0,0,0,0.4);
display: flex;
align-items: center;
justify-content: center;
z-index: 2000;
}

.profile-dialog {
display: flex;
flex-direction: row;
align-items: center;
justify-content: center;
min-width: 700px;
min-height: 400px;
padding: 3em 2em;
position: relative;
background: #fff;
border-radius: 16px;
box-shadow: 0 4px 32px rgba(0,0,0,0.2);
text-align: left;
}

.close-btn {
position: absolute;
top: 18px;
right: 18px;
background: #ff4d4f;
color: #fff;
border: none;
border-radius: 50px;
padding: 0.5em 1.5em;
font-size: 1em;
font-weight: bold;
cursor: pointer;
z-index: 10;
transition: background 0.2s;
}
.close-btn:hover {
background: #d32f2f;
}

.profile-img-col {
align-items: center;
justify-content: center;
flex: 0 0 200px;
}

.profile-avatar {
width: 180px;
height: 180px;
border-radius: 50%;
}

.profile-info-col {
align-items: flex-start;
justify-content: center;
padding: 0 2em;
}

.profile-actions {
margin-top: 2em;
}

.friend-icon {
width: 48px;
height: 48px;
filter: invert(41%) sepia(98%) saturate(1200%) hue-rotate(74deg) brightness(110%) contrast(120%);
}

.add-btn {
background: linear-gradient(90deg, #2196f3 0%, #00bcd4 100%);
color: #fff;
border: none;
border-radius: 24px;
padding: 0.5em 1.5em;
font-weight: bold;
font-size: 1em;
cursor: pointer;
box-shadow: 0 2px 8px rgba(33, 150, 243, 0.18);
transition: background 0.2s, transform 0.1s;
}
.add-btn:hover {
background: linear-gradient(90deg, #1976d2 0%, #0097a7 100%);
transform: translateY(-2px) scale(1.04);
}
</style>