<template>
  <Navbar />
  <div class="chatview-wrapper d-flex flex-column align-items-center justify-content-center" style="min-height: 100vh; background: #f5f7fa;">
    <ChatRoom />
  </div>
</template>

<script setup lang="ts">
import Navbar from '../components/Navbar.vue'
import ChatRoom from '../components/ChatRoom.vue'
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
