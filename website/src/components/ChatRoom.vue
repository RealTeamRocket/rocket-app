<template>
  <div class="chat-room d-flex flex-column border rounded bg-white shadow">
    <div
      class="chat-messages flex-grow-1 overflow-auto px-2 py-3 d-flex flex-column-reverse"
      ref="messagesContainer"
    >
      <ChatMessage
        v-for="(msg, idx) in reversedMessages"
        :key="idx"
        :username="msg.mine ? '' : msg.username"
        :message="msg.message"
        :mine="msg.mine"
        :reactions="0"
        :timestamp="msg.timestamp"
      />
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
import { ChatWebSocket, getChatWebSocketURL } from '@/api/chat-ws'
import ChatMessage from './ChatMessage.vue'

type LocalMessage = {
  username: string
  message: string
  mine: boolean
  timestamp: string
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

const reversedMessages = computed(() => [...messages.value].reverse())

function scrollToTop() {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = 0
    }
  })
}

function handleIncomingMessage(msg: { username: string; message: string; timestamp: string }) {
  if (msg.username === getUsername()) return
  messages.value.push({
    username: msg.username,
    message: msg.message,
    mine: msg.username === getUsername(),
    timestamp: msg.timestamp, // Add this
  })
  scrollToTop()
}

function sendMessage() {
  const text = input.value.trim()
  if (!text || !ws.value) return
  ws.value.sendMessage(text)
  messages.value.push({
    username: getUsername(),
    message: text,
    mine: true,
    timestamp: new Date().toISOString(),
  })
  input.value = ''
  scrollToTop()
}

onMounted(() => {
  const chatWS = new ChatWebSocket(getChatWebSocketURL())
  ws.value = chatWS
  chatWS.connect(handleIncomingMessage)
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
