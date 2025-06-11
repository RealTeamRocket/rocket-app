<template>
  <Navbar />
  <div class="profile-page">
    <h1>Profile: {{ username }}</h1>
    <p>This is the profile page for <strong>{{ username }}</strong>.</p>
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
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
const username = route.params.username as string

const user = ref<User | null>(null)


import { useRouter } from 'vue-router'

const router = useRouter()

onMounted(async () => {
  if (!username) {
    router.replace('/404')
    return
  }
  try {
    const response = await api.getUser(username)
    user.value = response.data
    console.log('User data:', user.value)
    if (!user.value || !user.value.username) {
      router.replace('/404')
    }
  } catch (error) {
    router.replace('/404')
  }
})

</script>

<style scoped>
.profile-page {
  padding: 2rem;
}
</style>
