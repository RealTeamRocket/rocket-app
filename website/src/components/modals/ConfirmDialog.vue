<template>
  <div v-if="open" class="confirm-dialog-overlay">
    <div class="confirm-dialog" @click.stop>
      <h3 class="confirm-dialog-title">{{ title }}</h3>
      <p class="confirm-dialog-message">{{ message }}</p>
      <div class="confirm-dialog-actions">
        <button class="cancel-btn" @click="$emit('cancel')">{{ cancelText || 'Cancel' }}</button>
        <button class="confirm-btn" @click="$emit('confirm')">{{ confirmText || 'Delete' }}</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  open: boolean,
  title?: string,
  message?: string,
  confirmText?: string,
  cancelText?: string,
}>()
</script>

<style scoped>
.confirm-dialog-overlay {
  position: fixed;
  z-index: 2000;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.25);
  display: flex;
  align-items: center;
  justify-content: center;
}
.confirm-dialog {
  background: #fff;
  border-radius: 10px;
  padding: 2rem 2.5rem;
  box-shadow: 0 4px 24px rgba(0,0,0,0.18);
  min-width: 320px;
  max-width: 90vw;
  text-align: center;
  position: relative;
}
.confirm-dialog-title {
  margin: 0 0 0.7em 0;
  font-size: 1.25rem;
  font-weight: 700;
  color: #222;
}
.confirm-dialog-message {
  margin: 0 0 1.5em 0;
  color: #444;
  font-size: 1.07rem;
}
.confirm-dialog-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1.2rem;
  margin-top: 2rem;
}
.cancel-btn {
  background: #f7fafd;
  color: #4a90e2;
  border: 1.5px solid #c0d6f7;
  border-radius: 6px;
  padding: 0.5rem 1.2rem;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.15s, color 0.15s, border 0.15s;
}
.cancel-btn:hover {
  background: #e0eaff;
  color: #222;
  border-color: #4a90e2;
}
.confirm-btn {
  background: #e74c3c;
  color: #fff;
  border: none;
  border-radius: 6px;
  padding: 0.5rem 1.2rem;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.15s;
}
.confirm-btn:hover {
  background: #c0392b;
}
</style>
