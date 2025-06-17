<template>
  <div class="account-settings">
    <span v-if="!hasImage" class="account-profile-initials" :style="{ backgroundColor: userColor }">
      {{ userInitials }}
    </span>
    <img v-else :src="userImage" alt="Profile" class="account-profile-img" />
    <button @click="$emit('logout')" class="logout-btn">Logout</button>
    <button @click="$emit('delete-account')" class="danger-btn">Delete Account</button>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{ userImage: string; userColor: string; userInitials: string }>()
defineEmits(['logout', 'delete-account'])

const hasImage = computed(
  () => props.userImage && props.userImage !== 'https://via.placeholder.com/120'
)
</script>

<style scoped>
.account-settings {
  width: 100%;
  max-width: 100vw;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2rem;
  margin-top: 2rem;
}

.account-profile-img,
.account-profile-initials {
  width: 220px;
  height: 220px;
  max-width: 90vw;
  max-height: 60vh;
  border-radius: 50%;
  display: block;
  margin-bottom: 2.5rem;
  box-shadow: 0 2px 16px rgba(0, 0, 0, 0.08);
}

.account-profile-img {
  object-fit: cover;
  border: 6px solid #bbb;
}

.account-profile-initials {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 3rem;
  font-weight: 700;
  color: #fff;
  border: 6px solid #bbb;
  background: #2a5298;
  user-select: none;
  overflow: hidden;
}

.logout-btn,
.danger-btn {
  min-width: 160px;
  padding: 0.7rem 1.2rem;
  border-radius: 6px;
  border: none;
  font-size: 1rem;
  cursor: pointer;
  background: #f5f5f5;
  color: #232946;
  transition:
    background 0.18s,
    color 0.18s;
}

.logout-btn:hover {
  background: #e0e0e0;
}

.danger-btn {
  background: #fff0f0;
  color: #e53935;
  border: 1px solid #e53935;
}

.danger-btn:hover {
  background: #e53935;
  color: #fff;
}
</style>
