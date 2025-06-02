<template>
  <div class="chat-room">
    <div class="chat-messages" ref="messagesContainer">
      <div
        v-for="(msg, idx) in messages"
        :key="idx"
        :class="['chat-message', msg.mine ? 'mine' : '']"
      >
        <span class="chat-username">{{ msg.username }}:</span>
        <span class="chat-text">{{ msg.message }}</span>
      </div>
    </div>
    <form class="chat-input-form" @submit.prevent="sendMessage">
      <input
        v-model="input"
        class="chat-input"
        type="text"
        placeholder="Type your message..."
        autocomplete="off"
        @keydown.enter.exact.prevent="sendMessage"
      />
      <button class="chat-send-btn" type="submit" :disabled="!input.trim()">Send</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { ChatWebSocket, getChatWebSocketURL } from '@/api/chat-ws'

// Message type for local state
type LocalMessage = {
  username: string
  message: string
  mine: boolean
}

const messages = ref<LocalMessage[]>([])
const input = ref('')
const ws = ref<ChatWebSocket | null>(null)
const messagesContainer = ref<HTMLElement | null>(null)

// Get username from somewhere (e.g., user profile, localStorage, etc.)
// For demo, fallback to "Me"
function getUsername(): string {
  // Replace with actual user logic
  const user = localStorage.getItem('username')
  return user || 'Me'
}

function scrollToBottom() {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
    }
  })
}

function handleIncomingMessage(msg: { username: string; message: string }) {
  messages.value.push({
    username: msg.username,
    message: msg.message,
    mine: msg.username === getUsername(),
  })
  scrollToBottom()
}

function sendMessage() {
  const text = input.value.trim()
  if (!text || !ws.value) return
  ws.value.sendMessage(text)
  // Optimistically add message to UI
  messages.value.push({
    username: getUsername(),
    message: text,
    mine: true,
  })
  input.value = ''
  scrollToBottom()
}

onMounted(() => {
  const chatWS = new ChatWebSocket(getChatWebSocketURL())
  ws.value = chatWS
  chatWS.connect(handleIncomingMessage)
})

onBeforeUnmount(() => {
  ws.value?.close()
})
</script>

<style scoped>
.chat-room {
  display: flex;
  flex-direction: column;
  height: 100%;
  max-height: 500px;
  border: 1px solid #ddd;
  border-radius: 8px;
  background: #fff;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  background: #f9f9f9;
}

.chat-message {
  margin-bottom: 8px;
  display: flex;
  align-items: baseline;
}

.chat-message.mine .chat-username {
  color: #1976d2;
  font-weight: bold;
}

.chat-username {
  margin-right: 6px;
  font-weight: 600;
}

.chat-text {
  word-break: break-word;
}

.chat-input-form {
  display: flex;
  border-top: 1px solid #eee;
  padding: 8px;
  background: #fff;
}

.chat-input {
  flex: 1;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  margin-right: 8px;
  font-size: 1rem;
}

.chat-send-btn {
  padding: 8px 16px;
  background: #1976d2;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-weight: bold;
  cursor: pointer;
  transition: background 0.2s;
}

.chat-send-btn:disabled {
  background: #aaa;
  cursor: not-allowed;
}
</style>
