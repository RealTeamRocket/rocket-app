<script setup lang="ts">
defineProps<{
  isFriend: boolean,
  username: string
}>()
  defineEmits(['add-friend'])
</script>

<template>
  <button
      v-if="!isFriend"
      @click="$emit('add-friend', username)"
      class="add-friend-btn"
  >
    Freund hinzufügen
  </button>
  <span v-else class="friend-icon">
    <img src="/src/assets/icons/user.svg" alt="Friend Icon" />
  </span>
</template>

<style scoped>

</style>