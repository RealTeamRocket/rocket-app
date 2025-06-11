<template>
  <div class="follow-list-tab">
    <h3 class="tab-title">{{ title }}</h3>
    <div v-if="users && users.length > 0" class="user-list">
      <div
        v-for="u in users"
        :key="u.id"
        class="user-card"
        @click="goToProfile(u.username)"
        tabindex="0"
        @keydown.enter="goToProfile(u.username)"
      >
        <span v-if="u.image_data" class="user-avatar-img">
          <img :src="getUserImage(u.image_data)" alt="User" />
        </span>
        <span v-else class="user-avatar-initials" :style="{ backgroundColor: getColor(u.username), color: '#fff' }">
          {{ getInitials(u.username) }}
        </span>
        <div class="user-info">
          <span class="username">{{ u.username }}</span>
          <span class="rocket-points">
            <span class="points-icon">
              <img src="/src/assets/icons/rocket.svg" alt="Rocket" style="width:1.1em;height:1.1em;vertical-align:middle;" />
            </span>
            {{ u.rocket_points ?? 0 }}
          </span>
        </div>
      </div>
    </div>
    <div v-else class="no-users">
      <p>No users to show.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { getColor, getInitials } from '@/utils/userUtils'

const props = defineProps<{
  users: Array<{
    id: string
    username: string
    image_data?: string | null
    rocket_points?: number
  }>
  title: string
}>()

const router = useRouter()
const route = useRoute()

const goToProfile = (username: string) => {
  if (route.params.username === username) {
    // If already on this profile, force reload with dummy query param
    router.replace({ path: `/profile/${username}`, query: { t: Date.now() } })
  } else {
    router.push(`/profile/${username}`)
  }
}

const getUserImage = (image_data: string | null | undefined) => {
  return image_data ? `data:image/jpeg;base64,${image_data}` : undefined
}
</script>

<style scoped>
.follow-list-tab {
  width: 100%;
  max-width: 520px;
  margin: 0 auto;
  padding: 0 0.5rem;
}
.tab-title {
  font-size: 1.25rem;
  font-weight: 700;
  color: #1e3c72;
  margin-bottom: 1.5rem;
  text-align: center;
  letter-spacing: 0.5px;
}
.user-list {
  display: flex;
  flex-direction: column;
  gap: 1.1rem;
}
.user-card {
  display: flex;
  align-items: center;
  gap: 1.1rem;
  background: #fff;
  border-radius: 1.2rem;
  box-shadow: 0 2px 12px rgba(30,60,114,0.07);
  padding: 1rem 1.5rem;
  cursor: pointer;
  transition: box-shadow 0.18s, background 0.18s;
  outline: none;
}
.user-card:hover, .user-card:focus {
  background: #f3f8ff;
  box-shadow: 0 4px 18px rgba(30,60,114,0.13);
}
.user-avatar-img img {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  object-fit: cover;
  background: #fff;
  box-shadow: 0 2px 8px rgba(30,60,114,0.10);
  border: 2.5px solid #e0e7ff;
}
.user-avatar-initials {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 56px;
  height: 56px;
  border-radius: 50%;
  color: #fff;
  font-weight: 700;
  font-size: 1.5rem;
  background: #2a5298;
  user-select: none;
  box-shadow: 0 2px 8px rgba(30,60,114,0.10);
  border: 2.5px solid #e0e7ff;
}
.user-info {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
}
.username {
  font-size: 1.15rem;
  font-weight: 600;
  color: #1e3c72;
}
.rocket-points {
  font-size: 1rem;
  color: #ffb347;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 0.3rem;
}
.points-icon {
  font-size: 1.1rem;
}
.no-users {
  text-align: center;
  color: #b0b8c9;
  font-size: 1.1rem;
  margin-top: 2rem;
}
</style>
