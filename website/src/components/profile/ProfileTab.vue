<template>
  <div class="profile-tab" v-if="user">
    <div class="profile-header">
      <span v-if="userImage" class="user-avatar-img" @click="showImageModal = true" style="cursor:pointer;">
        <img :src="userImage" alt="User" />
      </span>
      <span v-else class="user-avatar-initials" :style="{ backgroundColor: userColor, color: '#fff' }">
        {{ userInitials }}
      </span>
      <div class="profile-info">
        <h2 class="username">{{ user.username }}</h2>
        <div class="email">{{ user.email }}</div>
        <div class=" ocket-points">
          <span class="points-label">Rocket Points:</span>
          <span class="points-value">{{ user.rocket_points }}</span>
        </div>
      </div>
    </div>
    <div class="profile-stats">
      <StatsCards :stats="stats" />
    </div>
    <div class="profile-chart">
      <StepChart :data="chartData" :labels="chartLabels" />
    </div>
    <!-- Image Popup Modal -->
    <ImageModal
      v-if="userImage"
      :show="showImageModal"
      :src="userImage"
      alt="Profile Full"
      @close="showImageModal = false"
    />
  </div>
  <div v-else class="profile-loading">
    <p>Loading profile...</p>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, onMounted, watch } from 'vue'
import StatsCards from '@/components/dashboard/StatsCards.vue'
import StepChart from '@/components/dashboard/StepChart.vue'
import { getColor, getInitials } from '@/utils/userUtils'
import api from '@/api/backend-api'
import ImageModal from '@/components/modals/ImageModal.vue'

const props = defineProps<{
  user: {
    id: string
    username: string
    email: string
    rocket_points: number
    image_name: string
    image_data: string | null
  } | null
}>()

const user = computed(() => props.user)

const userImage = computed<string | undefined>(() => {
  if (user.value && user.value.image_data) {
    return `data:image/jpeg;base64,${user.value.image_data}`
  }
  return undefined
})

const userInitials = computed(() => {
  if (user.value && user.value.username) {
    return getInitials(user.value.username)
  }
  return ''
})

const userColor = computed(() => {
  if (user.value && user.value.username) {
    return getColor(user.value.username)
  }
  return '#2a5298'
})

const showImageModal = ref(false)

const stats = ref({
  totalSteps: 0,
  avgSteps: 0,
  bestDay: '',
  bestSteps: 0
})
const chartData = ref<number[]>([])
const chartLabels = ref<string[]>([])

const fetchStats = async () => {
  if (!user.value) return
  try {
    const res = await api.getUserStatistics(user.value.id)
    const dailyStats = res.data
    const totalSteps = dailyStats.reduce((sum: any, s: any) => sum + s.steps, 0)
    const avgSteps = dailyStats.length ? Math.round(totalSteps / dailyStats.length) : 0
    const best = dailyStats.reduce(
      (prev: any, curr: any) => (curr.steps > prev.steps ? curr : prev),
      dailyStats[0] || { day: '', steps: 0 }
    )
    stats.value = {
      totalSteps,
      avgSteps,
      bestDay: best.day,
      bestSteps: best.steps
    }
    chartData.value = dailyStats.map((s: any) => s.steps)
    chartLabels.value = dailyStats.map((s: any) => s.day)
  } catch (e) {
    stats.value = { totalSteps: 0, avgSteps: 0, bestDay: '', bestSteps: 0 }
    chartData.value = []
    chartLabels.value = []
  }
}

onMounted(() => {
  if (user.value) fetchStats()
})

watch(
  () => user.value,
  (newUser, oldUser) => {
    if (newUser && (!oldUser || newUser.id !== oldUser.id)) {
      fetchStats()
    }
  }
)
</script>

<style scoped>
.profile-tab {
  max-width: 700px;
  margin: 0 auto;
  background: #fff;
  border-radius: 1.2rem;
  box-shadow: 0 2px 12px rgba(30,60,114,0.08);
  padding: 2.5rem 2rem 2rem 2rem;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.profile-header {
  display: flex;
  align-items: center;
  gap: 2rem;
  margin-bottom: 2.5rem;
  width: 100%;
  justify-content: center;
}

.user-avatar-img img {
  width: 110px;
  height: 110px;
  border-radius: 50%;
  object-fit: cover;
  background: #fff;
  box-shadow: 0 2px 8px rgba(30,60,114,0.10);
  border: 3px solid #e0e7ff;
}

.user-avatar-initials {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 110px;
  height: 110px;
  border-radius: 50%;
  color: #fff;
  font-weight: 700;
  font-size: 2.5rem;
  background: #2a5298;
  user-select: none;
  box-shadow: 0 2px 8px rgba(30,60,114,0.10);
  border: 3px solid #e0e7ff;
}

.profile-info {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 0.5rem;
}

.username {
  font-size: 2rem;
  font-weight: 700;
  color: #1e3c72;
  margin-bottom: 0.2rem;
}

.email {
  font-size: 1.1rem;
  color: #4a5874;
  margin-bottom: 0.2rem;
}

.rocket-points {
  font-size: 1.1rem;
  color: #2a5298;
  margin-top: 0.2rem;
}

.points-label {
  font-weight: 500;
  margin-right: 0.3rem;
}

.points-value {
  font-weight: bold;
  color: #ffb347;
  font-size: 1.2rem;
}

.profile-stats {
  width: 100%;
  margin-top: 1.5rem;
  margin-bottom: 2rem;
}

.profile-chart {
  width: 100%;
  max-width: 520px;
  margin: 0 auto;
}

.profile-loading {
  text-align: center;
  padding: 2rem;
}

</style>
