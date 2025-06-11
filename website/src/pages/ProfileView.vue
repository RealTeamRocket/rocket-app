<template>
  <Navbar />
  <div class="profile-page">
    <TabBar :tabs="tabs" v-model="activeTab" />
    <div v-if="activeTab === 'Profile'">
      <ProfileTab :user="user" />
    </div>
    <div v-else-if="activeTab === 'Followed'">
      <FollowListTab :users="followedUsers" title="Followed" />
    </div>
    <div v-else-if="activeTab === 'Following'">
      <FollowListTab :users="followingUsers" title="Following" />
    </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
import TabBar from '@/components/profile/TabBar.vue'
import ProfileTab from '@/components/profile/ProfileTab.vue'
import FollowListTab from '@/components/profile/FollowListTab.vue'
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import api from '@/api/backend-api'

interface User {
   id: string
   username: string
   email: string
   rocket_points: number
   image_name: string
   image_data: string | null
}

const route = useRoute()
const router = useRouter()
const username = route.params.username as string

const user = ref<User | null>(null)
const followedUsers = ref([])
const followingUsers = ref([])

const tabs = ['Profile', 'Followed', 'Following']
const activeTab = ref('Profile')

const loadProfile = async (username: string) => {
  if (!username) {
    router.replace('/404')
    return
  }
  try {
    const response = await api.getUser(username)
    user.value = response.data
    // TODO: Replace with real API calls for followed/following users
    followedUsers.value = [] // await api.getFollowedUsers(username)
    followingUsers.value = [] // await api.getFollowingUsers(username)
    if (!user.value || !user.value.username) {
      router.replace('/404')
    }
  } catch (error) {
    router.replace('/404')
  }
}

onMounted(() => {
  loadProfile(username)
})

watch(
  () => route.params.username,
  (newUsername) => {
    loadProfile(newUsername as string)
  }
)
</script>

<style scoped>
.profile-page {
  padding: 2rem;
}
</style>
