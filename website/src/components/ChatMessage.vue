<template>
  <div :class="['d-flex mb-2', mine ? 'justify-content-end' : 'justify-content-start']">
    <div
      :class="[
        'p-2 rounded shadow-sm chat-message-bubble position-relative',
        mine ? 'bg-primary text-white align-self-end' : 'bg-light align-self-start'
      ]"
    >
      <div class="d-flex align-items-end">
        <span
          v-if="!mine"
          class="chat-username me-2"
          :style="{ color: getColor(username) }"
        >{{ username }}</span>
        <span class="chat-message-text">{{ message }}</span>
        <span class="chat-message-time ms-2">{{ formatTime(timestamp) }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { getColor } from '../utils/colorUtils'

defineProps<{
  username: string
  message: string
  reactions: number
  mine: boolean
  timestamp: string
}>()

function formatTime(ts: string) {
  if (!ts) return ''
  const date = new Date(ts)
  return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}
</script>

<style scoped>
.chat-message-bubble {
  max-width: 75%;
  min-width: 60px;
  position: relative;
  padding-bottom: 0.5rem;
  font-size: 1.05rem;
  border-radius: 1.1em 1.1em 1.1em 0.3em;
}

.chat-username {
  font-weight: 600;
  font-size: 0.97em;
  opacity: 0.95;
}

.chat-message-text {
  word-break: break-word;
}

.chat-message-time {
  font-size: 0.75rem;
  color: #b0b8c9;
  align-self: flex-end;
  margin-bottom: -2px;
  margin-left: 0.5em;
  position: relative;
  top: 2px;
}
</style>
