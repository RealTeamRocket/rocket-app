<template>
  <div :class="['d-flex mb-2', mine ? 'justify-content-end' : 'justify-content-start']">
    <div
      :class="[
        'p-2 rounded shadow-sm',
        mine ? 'bg-primary text-white align-self-end' : 'bg-light align-self-start'
      ]"
      style="max-width: 75%; min-width: 60px;"
    >
      <div v-if="!mine" class="small fw-bold mb-1">{{ username }}</div>
      <div>{{ message }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  username: string
  message: string
  reactions: number
  mine: boolean
}>()
</script>
