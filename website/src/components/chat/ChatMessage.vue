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
        <span
          class="chat-message-text"
          v-html="linkify(message)"
        ></span>
        <span class="chat-message-time ms-2">{{ formatTime(timestamp) }}</span>
      </div>
      <!-- YouTube preview -->
      <div v-if="youtubeId" class="youtube-preview mt-2 d-flex justify-content-center">
        <iframe
          title="youtube link"
          :src="`https://www.youtube.com/embed/${youtubeId}`"
          width="320"
          height="180"
          frameborder="0"
          allowfullscreen
          style="border-radius: 12px; box-shadow: 0 2px 12px rgba(0,0,0,0.10);"
        ></iframe>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { getColor } from '@/utils/userUtils'
import { computed } from 'vue'

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

function extractYouTubeId(url: string): string | null {
  // Handles both youtu.be and youtube.com/watch?v= links
  const match = url.match(
    /(?:youtube\.com\/watch\?v=|youtu\.be\/)([A-Za-z0-9_-]{11})/
  )
  return match ? match[1] : null
}

const youtubeId = computed(() => {
  const urlMatch = props.message.match(
    /(https?:\/\/(?:www\.)?(?:youtube\.com\/watch\?v=|youtu\.be\/)[A-Za-z0-9_-]{11})/
  )
  if (urlMatch) {
    return extractYouTubeId(urlMatch[0])
  }
  return null
})

// Replace URLs in message with clickable links (except YouTube, which gets preview)
function linkify(text: string) {
  const ytRegex = /(https?:\/\/(?:www\.)?(?:youtube\.com\/watch\?v=|youtu\.be\/)[A-Za-z0-9_-]{11})/
  if (ytRegex.test(text)) {
    text = text.replace(ytRegex, '')
  }
  // Linkify other URLs
  return text.replace(
    /(https?:\/\/[^\s]+)/g,
    '<a href="$1" target="_blank" rel="noopener noreferrer">$1</a>'
  )
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

.youtube-preview {
  width: 100%;
  margin-top: 0.5em;
  margin-bottom: 0.2em;
  display: flex;
  justify-content: center;
}
.youtube-preview iframe {
  max-width: 100%;
  border: none;
  background: #000;
}
.chat-message-text a {
  text-decoration: underline;
  word-break: break-all;
  transition: color 0.2s;
}

:deep(.chat-message-text a){
  color: #ffffff !important;
}
:deep(.chat-message-text a:hover) {
  color: #ffd600 !important;
  background: transparent;
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
