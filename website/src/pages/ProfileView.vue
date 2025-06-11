<template>
  <Navbar />
  <div class="profile-page">
    <TabBar :tabs="tabs" v-model="activeTab" />
    <div v-if="activeTab === 'Profile'">
      <ProfileTab :user="user" />
    </div>
    <div v-else-if="activeTab === 'Followers'">
      <FollowListTab :users="followedUsers" title="Followers" />
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
import { useRouter } from 'vue-router'
import api from '@/api/backend-api'

interface User {
   id: string
   username: string
   email: string
   rocket_points: number
   image_name: string
   image_data: string | null
}

const props = defineProps<{ username: string }>()
const router = useRouter()
const username = props.username

const user = ref<User | null>(null)
const followedUsers = ref([])
const followingUsers = ref([])

const tabs = ['Profile', 'Followers', 'Following']
const activeTab = ref('Profile')

const loadProfile = async (username: string) => {
  if (!username) {
    router.replace('/404')
    return
  }
  try {
    const { data: userData } = await api.getUser(username)
    if (!userData || !userData.username) {
      router.replace('/404')
      return
    }
    user.value = userData

    if (!user.value) {
      router.replace('/404')
      return
    }
    const { data: followingData } = await api.getFollowing(user.value.id)
    followingUsers.value = followingData || []

    const { data: followersData } = await api.getFollowers(user.value.id)
    followedUsers.value = followersData || []
  } catch {
    router.replace('/404')
  }
}

onMounted(() => {
  loadProfile(username)
})

watch(
  () => props.username,
  (newUsername) => {
    loadProfile(newUsername)
  }
)
</script>

<style scoped>
.profile-page {
  padding: 2rem;
}
</style>
