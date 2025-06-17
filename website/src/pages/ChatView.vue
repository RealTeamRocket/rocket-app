<template>
  <Navbar />
  <div class="chatview-wrapper d-flex flex-column align-items-center justify-content-center">
    <ChatRoom :user="user"/>
  </div>
  <Footer />
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
import ChatRoom from '@/components/chat/ChatRoom.vue'
import Footer from '@/components/footer/Footer.vue'
import { ref, onMounted } from 'vue'
import api from '@/api/backend-api'

const user = ref({
  id: '',
  username: '',
  rocket_points: 0
})

onMounted(async () => {
  const response = await api.getMyself()
  if (response.status === 200) {
    user.value = response.data
  } else {
    console.error('Failed to fetch user data:', response.statusText)
  }
})
</script>

<style scoped>
.chatview-wrapper {
  min-height: 100vh;
  width: 100vw;
  padding: 0;
  margin: 0;
  background: #f5f7fa;
}
</style>
