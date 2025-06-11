<template>
  <div>
    <h3>{{ title }}</h3>
    <ul>
      <li v-for="u in users" :key="u.id" class="user-list-item">
        <img
          v-if="u.image_data"
          :src="u.image_data"
          alt="User Image"
          width="40"
          height="40"
          class="user-avatar"
        />
        <span>{{ u.username }}</span>
      </li>
    </ul>
    <div v-if="!users || users.length === 0">
      <p>No users to show.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  users: Array<{
    id: string
    username: string
    image_data?: string | null
  }>
  title: string
}>()
</script>

<style scoped>
.user-list-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.5rem;
}
.user-avatar {
  border-radius: 50%;
  object-fit: cover;
  background: #eee;
}
</style>
