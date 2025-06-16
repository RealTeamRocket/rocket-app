<script setup lang="ts">
import { ref } from 'vue'
import { getInitials, getColor } from '@/utils/userUtils'
import ImageModal from '@/components/modals/ImageModal.vue'

interface RankedUser {
  id: string
  username: string
  rocket_points: number
  imageData?: string
  isFriend?: boolean
}

const props = defineProps<{
  user: RankedUser
  onClose: () => void
  onAddFriend: (username: string) => void
  currentUsername: string
}>()

const showImageModal = ref(false)
const emit = defineEmits(['goToProfile'])
function openImageModal() {
  if (props.user.imageData) showImageModal.value = true
}
function goToProfile() {
  emit('goToProfile', props.user.username)
}
</script>

<template>
  <div class="profile-dialog-overlay" @click.self="onClose">
    <div class="profile-dialog">
      <button class="close-btn" @click="onClose">Close</button>
      <div class="profile-img-col">
        <template v-if="user.imageData">
          <img
            :src="`data:image/*;base64,${user.imageData}`"
            class="profile-avatar"
            @click="openImageModal"
            style="cursor:pointer"
            title="Click to enlarge"
          />
        </template>
        <template v-else>
          <div
            class="profile-avatar initials-avatar"
            :style="{ background: getColor(user.username), cursor: 'default' }"
          >
            {{ getInitials(user.username) }}
          </div>
        </template>
      </div>
      <div class="profile-info-col">
        <h2 class="profile-username-link" @click="goToProfile" title="Go to profile">{{ user.username }}</h2>
        <p>Punkte: {{ user.rocket_points }}</p>
        <div class="profile-actions">
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
            <button class="add-btn" @click.stop="onAddFriend(user.username)">Add</button>
          </template>
        </div>
      </div>
    </div>
    <ImageModal
      v-if="user.imageData"
      :show="showImageModal"
      :src="`data:image/*;base64,${user.imageData}`"
      :alt="user.username"
      @close="showImageModal = false"
    />
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
  object-fit: cover;
  display: block;
}

.initials-avatar {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 3.2rem;
  font-weight: 700;
  color: #fff;
  text-transform: uppercase;
  letter-spacing: 0.04em;
  width: 180px;
  height: 180px;
  border-radius: 50%;
  user-select: none;
}

.profile-info-col {
  align-items: flex-start;
  justify-content: center;
  padding: 0 2em;
}
.profile-username-link {
  cursor: pointer;
  color: #1e3c72;
  text-decoration: none;
  transition: color 0.15s;
  font-size: 2rem;
  font-weight: 700;
  margin: 0 0 0.2em 0;
}
.profile-username-link:hover {
  color: #2196f3;
  text-decoration: none;
}

.profile-actions {
margin-top: 2em;
}

.friend-icon {
width: 48px;
height: 48px;
filter: invert(41%) sepia(98%) saturate(1200%) hue-rotate(74deg) brightness(110%) contrast(120%);
}

.friend-badge {
  display: inline-flex;
  align-items: center;
  gap: 0.5em;
  background: #eaffea;
  border-radius: 18px;
  padding: 0.4em 1.1em 0.4em 0.7em;
  font-weight: 700;
  color: #228c22;
  font-size: 1.1em;
  box-shadow: 0 2px 8px rgba(52,199,89,0.08);
  border: 1.5px solid #34c759;
}
.friend-badge-icon {
  width: 1.5em;
  height: 1.5em;
  margin-right: 0.2em;
  flex-shrink: 0;
}
.friend-badge-label {
  font-size: 1em;
  color: #228c22;
  letter-spacing: 0.01em;
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
