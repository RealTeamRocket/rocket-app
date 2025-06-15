<script setup lang="ts">
  import AddFriendButton from './AddFriendButton.vue'
  defineProps<{ user: any }>()
  defineEmits(['close', 'add-friend'])
</script>

<template>
  <div class="profile-dialog-overlay" @click.self="$emit('close')">
    <div class="profile-dialog">
      <button class="close-btn" @click="$emit('close')">Close</button>
      <div class="profile-img-col">
        <img :src="user.imageUrl || '/src/assets/icons/user.svg'" class="profile-avatar" />
      </div>
      <div class="profile-info-col">
        <h2>{{ user.username }}</h2>
        <p>Punkte: {{ user.rocket_points }}</p>
        <div class="profile-actions">
          <template v-if="user.isFriend">
            <img src="/src/assets/icons/user.svg" alt="Friend Icon" class="friend-icon" />
          </template>
          <template v-else>
            <AddFriendButton :username="user.username" @add-friend="$emit('add-friend', $event)" />
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>