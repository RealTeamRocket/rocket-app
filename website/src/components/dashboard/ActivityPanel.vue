<template>
  <div class="activity-panel">
    <h2 class="panel-title">Activity Feed</h2>
    <ul class="activity-list">
      <li v-for="(activity, idx) in activities" :key="idx" class="activity-item" :class="{ 'user-activity': activity.isUser }">
        <span
          v-if="activity.image_data"
          class="avatar avatar-img"
          @click.stop="openImageModal(activity)"
          style="cursor:pointer;"
        >
          <img
            :src="`data:${activity.image_type || 'image/jpeg'};base64,${activity.image_data}`"
            alt="User"
            style="width: 40px; height: 40px; border-radius: 50%; object-fit: cover;"
          />
        </span>
        <div v-else class="avatar" :style="{ backgroundColor: activity.color }">
          <span>{{ activity.initials }}</span>
        </div>
        <div class="activity-content">
          <span
            class="name clickable"
            @click.stop="goToProfile(activity.name)"
            style="cursor:pointer;"
          >{{ activity.name }}</span>
          <span class="desc">{{ activity.description }}</span>
          <span class="time">{{ activity.time }}</span>
        </div>
      </li>
    </ul>
    <!-- Image Popup Modal -->
    <ImageModal
      v-if="showImageModal"
      :show="showImageModal"
      :src="modalImageSrc"
      alt="Profile Full"
      @close="showImageModal = false"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api/backend-api'
import { getColor, getInitials } from '@/utils/userUtils'
import ImageModal from '@/components/modals/ImageModal.vue'

interface Activity {
  name: string
  initials: string
  color: string
  description: string
  time: string
  isUser?: boolean
  image_data?: string | null
  image_type?: string
}

function formatRelativeTime(iso: string): string {
  const now = new Date()
  const then = new Date(iso)
  const diff = Math.floor((now.getTime() - then.getTime()) / 1000)
  if (diff < 60) return 'just now'
  if (diff < 3600) return `${Math.floor(diff / 60)} min ago`
  if (diff < 86400) return `${Math.floor(diff / 3600)} hr ago`
  const days = Math.floor(diff / 86400)
  if (days === 1) return 'yesterday'
  return `${days} days ago`
}

const activities = ref<Activity[]>([])

const showImageModal = ref(false)
const modalImageSrc = ref<string | undefined>(undefined)
const router = useRouter()

const goToProfile = (username: string) => {
  // Route to the user's profile page
  router.push(`/profile/${username}`)
}

const getUserImage = (image_data: string | null | undefined, image_type?: string) => {
  return image_data ? `data:${image_type || 'image/jpeg'};base64,${image_data}` : undefined
}

const openImageModal = (activity: { image_data?: string | null, image_type?: string }) => {
  if (activity.image_data) {
    modalImageSrc.value = getUserImage(activity.image_data, activity.image_type)
    showImageModal.value = true
  }
}

import { onMounted } from 'vue'
onMounted(async () => {
  try {
    const { username, activities: backendActivities } = (await api.getActivityFeed()).data
    activities.value = backendActivities.map(({ name, time, message, image_data, image_type }) => {
      const displayName = name === 'You' ? username : name
      return {
        name: displayName,
        initials: getInitials(displayName),
        color: getColor(displayName),
        description: message,
        time: formatRelativeTime(time),
        isUser: name === 'You',
        image_data,
        image_type
      }
    })
  } catch {
    activities.value = []
  }
})
</script>

<style scoped>
.activity-panel {
    height: 580px;
    display: flex;
    flex-direction: column;
    background: #fff;
    border-radius: 12px;
    box-shadow: 0 2px 8px rgba(42, 82, 152, 0.06);
    padding: 1.5rem 1rem;
    overflow: hidden;
}

.panel-title {
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 1rem;
  color: #2a5298;
  letter-spacing: 0.5px;
}

.activity-list {
  list-style: none;
  padding: 0;
  margin: 0;
  flex: 1 1 auto;
  overflow-y: auto;
  min-height: 0;
  max-height: 100%;
}

.activity-item {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  padding: 0.75rem 0;
  border-bottom: 1px solid #f4f8fb;
  transition: background 0.2s, box-shadow 0.2s;
  border-radius: 10px;
}

.activity-item:hover {
  background: #f0f4fa;
  box-shadow: 0 2px 8px rgba(42, 82, 152, 0.04);
}

.activity-item:last-child {
  border-bottom: none;
}

.activity-item.user-activity {
  background: #eaf3ff;
  position: relative;
}

.activity-item.user-activity::after {
  content: "You";
  position: absolute;
  top: 1rem;
  right: 1rem;
  background: #2a5298;
  color: #fff;
  font-size: 0.7rem;
  padding: 0.15rem 0.5rem;
  border-radius: 8px;
  font-weight: 600;
  letter-spacing: 0.5px;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  color: #fff;
  font-size: 1.1rem;
  flex-shrink: 0;
  box-shadow: 0 1px 4px rgba(30,60,114,0.07);
}

.activity-content {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
}

.name {
  font-weight: 600;
  color: #1e3c72;
  font-size: 1rem;
}

.desc {
  color: #4a5874;
  font-size: 0.97rem;
}

.time {
  color: #b0b8c9;
  font-size: 0.85rem;
}
.username.clickable {
  text-decoration: underline;
  color: #1e3c72;
}
</style>
