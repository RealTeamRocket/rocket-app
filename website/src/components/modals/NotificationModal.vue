<template>
  <div v-if="open" class="notification-modal-overlay" @click.self="close">
    <div class="notification-modal-content" :class="type">
      <div class="notification-message">{{ message }}</div>
      <button class="close-btn" @click="close">&times;</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount } from 'vue'

const props = defineProps<{
  open: boolean
  message: string
  type?: 'success' | 'error' | 'info'
  autoClose?: number // ms, optional
}>()
const emit = defineEmits(['close'])

let timer: ReturnType<typeof setTimeout> | null = null

function close() {
  emit('close')
}

onMounted(() => {
  if (props.autoClose && props.open) {
    timer = setTimeout(() => {
      close()
    }, props.autoClose)
  }
})

onBeforeUnmount(() => {
  if (timer) clearTimeout(timer)
})
</script>

<style scoped>
.notification-modal-overlay {
  position: fixed;
  z-index: 2100;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.18);
  display: flex;
  align-items: center;
  justify-content: center;
}
.notification-modal-content {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 24px rgba(0,0,0,0.18);
  padding: 1.5rem 2.5rem 1.5rem 2rem;
  min-width: 280px;
  max-width: 90vw;
  display: flex;
  align-items: center;
  gap: 1rem;
  position: relative;
  font-size: 1.08rem;
}
.notification-modal-content.success {
  border-left: 6px solid #43a047;
}
.notification-modal-content.error {
  border-left: 6px solid #e53935;
}
.notification-modal-content.info {
  border-left: 6px solid #1976d2;
}
.notification-icon {
  font-size: 1.6rem;
  margin-right: 0.5rem;
}
.notification-message {
  flex: 1;
  color: #222;
}
.close-btn {
  position: absolute;
  top: 0.6rem;
  right: 0.7rem;
  background: transparent;
  color: #888;
  border: none;
  font-size: 1.5rem;
  font-weight: bold;
  cursor: pointer;
  transition: color 0.18s;
}
.close-btn:hover {
  color: #e53935;
}
</style>
