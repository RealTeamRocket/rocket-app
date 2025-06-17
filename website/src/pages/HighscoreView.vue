<template>
  <Rocket/>
  <Navbar/>
  <div class="highscore-bg">
    <div class="highscore-page">
      <div class="header">
        <img src="/src/assets/icons/rocket.svg" alt="Rocket" class="header-rocket left" />
        <h1 class="highscore-title">Highscore</h1>
        <img src="/src/assets/icons/rocket.svg" alt="Rocket" class="header-rocket right" />
      </div>
      <div class="main">
        <ToggleSwitch v-model="isGlobal"/>
        <Podium :users="selectedList.slice(0,3)" :openProfile="openProfile" />
        <List
            :users="selectedList.slice(3, 28)"
            :openProfile="openProfile"
            :addFriend="addFriend"
            :currentUsername="currentUsername"
        />
        <ProfileCard
            v-if="showProfileDialog && selectedProfile"
            :user="selectedProfile"
            :onClose="closeProfile"
            :onAddFriend="addFriend"
            :currentUsername="currentUsername"
            @goToProfile="goToProfileFromDialog"
        />
      </div>
    </div>
    <NotificationModal
      :open="notification.open"
      :message="notification.message"
      :type="notification.type"
      :autoClose="notification.autoClose"
      @close="notification.open = false"
    />
  </div>
</template>
<script setup lang="ts">
import Podium from '@/components/highscore/Podium.vue'
import ToggleSwitch from '@/components/highscore/ToggleSwitch.vue'
import Navbar from '@/components/Navbar.vue'
import {ref, onMounted, computed, reactive} from 'vue'
import api from '@/api/backend-api'
import List from "@/components/highscore/List.vue";
import ProfileCard from '@/components/highscore/ProfileCard.vue'
import Rocket from '@/components/highscore/Rocket.vue'
import NotificationModal from '@/components/modals/NotificationModal.vue'

interface RankedUser {
  id: string
  username: string
  rocket_points: number
  isFriend?: boolean
  imageUrl?: string
}

const currentUsername = ref('')
const isGlobal = ref(true)
const rankedUsers = ref<RankedUser[]>([])
const rankedFriends = ref<RankedUser[]>([])

const loadRanking = async () => {
  try {
    const {data: userDataRaw} = await api.getRankedUsers()
    const {data: friendDataRaw} = await api.getRankedFriends()
    const userData = Array.isArray(userDataRaw) ? userDataRaw : []
    const friendData = Array.isArray(friendDataRaw) ? friendDataRaw : []
    const friendUsernames = new Set(friendData.map((f: RankedUser) => f.username))
    // Helper to convert backend user to frontend user with correct imageData mapping
    function mapUser(u: any, isFriend: boolean): RankedUser {
      let imageData = ''
      // Accept both camelCase and snake_case from backend
      if (u.imageData && typeof u.imageData === 'string' && u.imageData.length > 0) {
        imageData = u.imageData
      } else if (u.image_data && typeof u.image_data === 'string' && u.image_data.length > 0) {
        imageData = u.image_data
      }
      return {
        ...u,
        isFriend,
        imageData,
      }
    }

    rankedUsers.value = userData.map((user: any) =>
      mapUser(user, friendUsernames.has(user.username))
    )
    rankedFriends.value = friendData.map((user: any) =>
      mapUser(user, true)
    )


  } catch {
    console.error("Error fetching ranking")
  }
}

const selectedList = computed(() => isGlobal.value ? rankedUsers.value : rankedFriends.value)

import { useRouter } from 'vue-router'

onMounted(async () => {
  const res = await api.getMyself()
  currentUsername.value = res.data.username
  await loadRanking()
})

const router = useRouter()
function goToProfileFromDialog(username: string) {
  // Optionally close the dialog before navigating
  showProfileDialog.value = false
  selectedProfile.value = null
  router.push(`/profile/${encodeURIComponent(username)}`)
}

const notification = reactive({
  open: false,
  message: '',
  type: 'info' as 'success' | 'error' | 'info',
  autoClose: 2200,
})

function showNotification(message: string, type: 'success' | 'error' | 'info' = 'info', autoClose = 2200) {
  notification.message = message
  notification.type = type
  notification.open = true
  notification.autoClose = autoClose
}

function addFriend(friendName: string) {
  api.addFriend(friendName)
      .then(() => {
        showNotification('Freund hinzugefügt!', 'success')
        rankedUsers.value = rankedUsers.value.map(u =>
          u.username === friendName ? { ...u, isFriend: true } : u
        )
        rankedFriends.value = rankedFriends.value.map(u =>
          u.username === friendName ? { ...u, isFriend: true } : u
        )
        if (showProfileDialog.value && selectedProfile.value && selectedProfile.value.username === friendName) {
          selectedProfile.value = { ...selectedProfile.value, isFriend: true }
        }
        loadRanking()
      })
      .catch((err) => {
        showNotification('Fehler beim Hinzufügen: ' + (err.response?.data?.message || err.message), 'error')
      })
}

const showProfileDialog = ref(false)
const selectedProfile = ref<RankedUser | null>(null)

function openProfile(user: RankedUser) {
  selectedProfile.value = user
  showProfileDialog.value = true
}

function closeProfile() {
  showProfileDialog.value = false
  selectedProfile.value = null
}


</script>

<style scoped>
.highscore-bg {
  min-height: 100vh;
  width: 100vw;
  /* background: linear-gradient(120deg, #e0e7ff 0%, #f8fafc 100%); */
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding-top: 0;
}

.highscore-page {
  margin: 3rem auto 2rem auto;
  padding: 2.5rem 2rem 3rem 2rem;
  background: rgba(255,255,255,0.96);
  border-radius: 32px;
  box-shadow: 0 6px 32px rgba(30,60,114,0.10);
  max-width: 900px;
  width: 100%;
  min-height: 80vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.header {
  flex-direction: row;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  margin-bottom: 2.5rem;
  gap: 1.5rem;
}

.header-rocket {
  width: 40px;
  height: 40px;
  opacity: 0.85;
}

.highscore-title {
  font-size: 2.5rem;
  font-weight: 800;
  letter-spacing: 0.04em;
  color: #1e3c72;
  margin: 0 0.5em;
  text-shadow: 0 2px 8px rgba(30,60,114,0.08);
}

.main {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  width: 100%;
  min-height: 60vh;
  gap: 0.5rem;
}

@media (max-width: 1100px) {
  .highscore-page {
    max-width: 98vw;
    padding: 1.2rem 0.5rem 2rem 0.5rem;
  }
  .main {
    min-height: 40vh;
  }
}
@media (max-width: 700px) {
  .highscore-page {
    margin: 1rem 0 0 0;
    border-radius: 0;
    box-shadow: none;
    padding: 0.5rem 0.2rem 1rem 0.2rem;
  }
  .header-rocket {
    width: 28px;
    height: 28px;
  }
  .highscore-title {
    font-size: 1.5rem;
  }
}
</style>
