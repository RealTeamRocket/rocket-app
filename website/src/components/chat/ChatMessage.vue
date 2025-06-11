<template>
  <div
    :class="['d-flex mb-2', mine ? 'justify-content-end' : 'justify-content-start']"
    @dblclick="handleReaction"
  >
    <div
      :class="[
        'p-2 rounded shadow-sm chat-message-bubble position-relative',
        mine ? 'bg-primary text-white align-self-end' : 'bg-light align-self-start'
      ]"
    >
      <!-- Reaction badge: left for own messages, right for others -->
      <div
        v-if="reactions > 0"
        class="reaction-badge"
        :class="[{ reacted: hasReacted }, mine ? 'left' : 'right']"
      >
        <img src="/src/assets/icons/rocket.svg" alt="Rocket" style="width:1em;height:1em;vertical-align:middle;" />
        <span>{{ reactions }}</span>
      </div>
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
import { getColor } from '@/utils/userUtils'

const props = defineProps<{
  username: string
  message: string
  reactions: number
  mine: boolean
  timestamp: string
  hasReacted: boolean
  onReact?: () => void
}>()

function formatTime(ts: string) {
  if (!ts) return ''
  const date = new Date(ts)
  return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}

function handleReaction() {
  if (props.mine) return
  if (typeof props.onReact === 'function' && !props.hasReacted) {
    props.onReact()
  }
}
</script>

<style scoped>
.chat-message-bubble {
  max-width: 70%;
  min-width: 48px;
  position: relative;
  padding-bottom: 0.3rem;
  font-size: 0.97rem;
  border-radius: 1.1em 1.1em 1.1em 0.3em;
}

.reaction-badge {
  position: absolute;
  background: #fff;
  border: 1px solid #e0e0e0;
  border-radius: 50%;
  padding: 0.04em 0.22em 0.04em 0.18em;
  font-size: 1em;
  font-weight: normal;
  display: flex;
  align-items: center;
  z-index: 2;
  color: #222;
  box-shadow: none;
  min-width: 22px;
  min-height: 22px;
  justify-content: center;
}
.reaction-badge.right {
  top: -10px;
  right: -10px;
}
.reaction-badge.left {
  top: -10px;
  left: -10px;
}
.reaction-badge span {
  font-size: 0.65em;
  color: #b0b8c9;
  margin-left: 0.12em;
  font-weight: 400;
}
.reaction-badge.reacted {
  border-color: #bdbdbd;
  background: #fafbfc;
}
.chat-username {
  font-weight: 600;
  font-size: 0.91em;
  opacity: 0.95;
}
.chat-message-text {
  word-break: break-word;
}
.chat-message-time {
  font-size: 0.68rem;
  color: #b0b8c9;
  align-self: flex-end;
  margin-bottom: -2px;
  margin-left: 0.4em;
  position: relative;
  top: 2px;
  white-space: nowrap;
}
</style>
