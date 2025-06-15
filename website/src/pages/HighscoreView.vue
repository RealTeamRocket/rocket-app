<template>
  <Rocket/>
  <Navbar/>
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
      />
    </div>
  </div>
</template>
<script setup lang="ts">
import Podium from '@/components/highscore/Podium.vue'
import ToggleSwitch from '@/components/highscore/ToggleSwitch.vue'
import Navbar from '@/components/Navbar.vue'
import {ref, onMounted, computed} from 'vue'
import api from '@/api/backend-api'
import List from "@/components/highscore/List.vue";
import ProfileCard from '@/components/highscore/ProfileCard.vue'
import Rocket from '@/components/highscore/Rocket.vue'

interface RankedUser {
  id: number | string
  username: string
  rocket_points: number
  isFriend?: boolean
  imageUrl?: string
}

const currentUsername = ref('')
const isGlobal = ref(true)
const rankedUsers = ref<RankedUser[]>([])
const rankedFriends = ref<RankedUser[]>([])

async function loadUserImages(list: RankedUser[]) {
  for (const user of list) {
    try {
      const res = await api.getUserImage(user.id)
      // Annahme: Das Bild kommt als Blob zurück
      const blob = res.data
      user.imageUrl = URL.createObjectURL(blob)
    } catch {
      user.imageUrl = '' // Kein Bild vorhanden
    }
  }
}

const loadRanking = async () => {
  try {
    const {data: userData} = await api.getRankedUsers()
    const {data: friendData} = await api.getRankedFriends()
    const friendUsernames = new Set(friendData.map((f: RankedUser) => f.username))
    rankedUsers.value = (userData || []).map((user: RankedUser) => ({
      ...user,
      isFriend: friendUsernames.has(user.username)
    }))
    rankedFriends.value = (friendData || []).map((user: RankedUser) => ({
      ...user,
      isFriend: true
    }))
    // Bilder laden
    await loadUserImages(rankedUsers.value)
    await loadUserImages(rankedFriends.value)
  } catch {
    console.error("Error fetching ranking")
  }
}

const selectedList = computed(() => isGlobal.value ? rankedUsers.value : rankedFriends.value)

onMounted(async () => {
  const res = await api.getMyself()
  currentUsername.value = res.data.username
  await loadRanking()
})

function addFriend(friendName: string) {
  api.addFriend(friendName)
      .then(() => {
        alert('Freund hinzugefügt!')
        loadRanking()
      })
      .catch((err) => {
        alert('Fehler beim Hinzufügen: ' + (err.response?.data?.message || err.message))
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
.highscore-page {
  padding: 2rem;
}
.header {
  flex-direction: row;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  margin-bottom: 4rem;
  gap: 2rem;
}
.header-rocket {
  width: 48px;
  height: 48px;
}
.main {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  height: 100vh;
}
</style>