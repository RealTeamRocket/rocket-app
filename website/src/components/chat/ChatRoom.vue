<template>
  <div class="chat-room d-flex flex-column border rounded bg-white shadow">
    <div
      class="chat-messages flex-grow-1 overflow-auto px-2 py-3 d-flex flex-column-reverse"
      ref="messagesContainer"
    >
      <template v-for="(item, idx) in messagesWithDates" :key="item.type === 'date' ? 'date-' + item.date + '-' + idx : item.msg.id || idx">
        <ChatDateSeparator
          v-if="item.type === 'date'"
          :date="item.date"
        />
        <ChatMessage
          v-else
          :username="item.msg.mine ? '' : item.msg.username"
          :message="item.msg.message"
          :mine="item.msg.mine"
          :reactions="item.msg.reactions || 0"
          :timestamp="item.msg.timestamp"
          :hasReacted="item.msg.hasReacted || false"
          :onReact="() => handleReact(item.msg)"
        />
      </template>
    </div>
    <form class="chat-input-form d-flex border-top p-2 bg-white" @submit.prevent="sendMessage">
      <input
        v-model="input"
        class="form-control me-2 chat-input"
        type="text"
        placeholder="Type your message..."
        autocomplete="off"
        @keydown.enter.exact.prevent="sendMessage"
      />
      <button class="btn btn-primary" type="submit" :disabled="!input.trim()">Send</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, onMounted, onBeforeUnmount, nextTick, computed } from 'vue'
import ChatMessage from './ChatMessage.vue'
import ChatDateSeparator from './ChatDateSeparator.vue'
import { ChatWebSocket, getChatWebSocketURL } from '@/api/chat-ws'
import api from '@/api/backend-api'

type LocalMessage = {
  id?: string
  username: string
  message: string
  mine: boolean
  timestamp: string
  reactions: number
  hasReacted: boolean
}

const props = defineProps<{
  user: {
    id: string
    username: string
    rocket_points: number
  }
}>()

const messages = ref<LocalMessage[]>([])
const input = ref('')
const ws = ref<ChatWebSocket | null>(null)
const messagesContainer = ref<HTMLElement | null>(null)

function getUsername(): string {
  return props.user?.username || 'Me'
}

const messagesWithDates = computed(() => {
  const result: Array<{ type: 'date', date: string } | { type: 'msg', msg: LocalMessage }> = []
  let lastDate = ''
  for (let i = 0; i < messages.value.length; i++) {
    const msg = messages.value[i]
    const msgDate = msg.timestamp ? msg.timestamp.split('T')[0] : ''
    if (msgDate && msgDate !== lastDate) {
      result.push({ type: 'date', date: msgDate })
      lastDate = msgDate
    }
    result.push({ type: 'msg', msg })
  }
  return result.reverse()
})

function scrollToTop() {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = 0
    }
  })
}

function handleIncomingWS(data: any) {
  if (data.type === "reaction" && data.messageId) {
    // Find the message and update its reactions/count
    const msg = messages.value.find(m => m.id === data.messageId)
    if (msg) {
      msg.reactions = data.reactions
      // Optionally, set hasReacted if the reaction is from this user
      if (data.username === getUsername()) {
        msg.hasReacted = true
      }
    }
  } else if (data.username && data.message) {
    // Handle new chat message as before
    messages.value.push({
      id: data.id, // If backend provides it
      username: data.username,
      message: data.message,
      mine: data.username === 'You' || data.username === getUsername(),
      timestamp: data.timestamp,
      reactions: data.reactions ?? 0,
      hasReacted: false,
    })
    scrollToTop()
  }
}

function sendMessage() {
  const text = input.value.trim()
  if (!text || !ws.value) return
  ws.value.sendMessage(text)
  input.value = ''
  scrollToTop()
}

function handleReact(msg: LocalMessage) {
  if (msg.hasReacted || msg.mine || !msg.id || !ws.value) return
  ws.value.sendReaction(msg.id)
  // Optimistic UI update (optional)
  msg.hasReacted = true
}

onMounted(async () => {
  // 1. Load chat history from backend
  try {
    const response = await api.getChatHistory()
    if (response.status === 200 && Array.isArray(response.data.messages)) {
      messages.value = response.data.messages.map((msg: any) => ({
        id: msg.id,
        username: msg.username,
        message: msg.message,
        mine: msg.username === 'You' || msg.username === getUsername(),
        timestamp: msg.timestamp,
        reactions: msg.reactions ?? 0,
        hasReacted: msg.hasReacted ?? false,
      }))
      scrollToTop()
    }
  } catch (e) {
    console.error('Failed to load chat history', e)
  }

  // 2. Setup websocket for live chat and reactions
  const chatWS = new ChatWebSocket(getChatWebSocketURL())
  ws.value = chatWS
  chatWS.connect(handleIncomingWS)
  scrollToTop()
})

onBeforeUnmount(() => {
  ws.value?.close()
})
</script>

<style scoped>
.chat-room {
  width: 90vw;
  max-width: 900px;
  height: 80vh;
  max-height: 900px;
  min-width: 320px;
  min-height: 400px;
  box-shadow: 0 2px 16px rgba(0,0,0,0.07);
}
.chat-messages {
  background: #f9f9f9;
  min-height: 0;
}
.chat-input-form {
  flex-shrink: 0;
}
</style>
