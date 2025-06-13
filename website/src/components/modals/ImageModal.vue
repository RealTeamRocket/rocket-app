<template>
  <div v-if="show" class="image-modal-overlay" @click.self="close">
    <div class="image-modal-content">
      <img :src="src" :alt="alt" />
      <button class="close-btn" @click="close">&times;</button>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  show: boolean
  src: string
  alt?: string
}>()
const emit = defineEmits(['close'])
function close() {
  emit('close')
}
</script>

<style scoped>
.image-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(30, 60, 114, 0.55);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}
.image-modal-content {
  position: relative;
  background: #fff;
  border-radius: 1.2rem;
  box-shadow: 0 4px 24px rgba(30,60,114,0.18);
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 90vw;
  max-height: 90vh;
}
.image-modal-content img {
  max-width: 70vw;
  max-height: 70vh;
  border-radius: 1rem;
  box-shadow: 0 2px 8px rgba(30,60,114,0.10);
}
.close-btn {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  background: #e53935;
  color: #fff;
  border: none;
  border-radius: 50%;
  width: 2.2rem;
  height: 2.2rem;
  font-size: 1.5rem;
  font-weight: bold;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 1px 4px rgba(30,60,114,0.10);
  transition: background 0.18s;
}
.close-btn:hover {
  background: #b71c1c;
}
</style>