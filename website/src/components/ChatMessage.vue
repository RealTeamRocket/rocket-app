<template>
  <div :class="['d-flex mb-2', mine ? 'justify-content-end' : 'justify-content-start']">
    <div :class="['p-2 rounded', mine ? 'bg-primary text-white' : 'bg-light']" style="max-width: 70%;">
      <div class="small fw-bold">{{ username }}</div>
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
