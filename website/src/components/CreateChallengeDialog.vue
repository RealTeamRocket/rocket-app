<template>
  <div class="modal-backdrop">
    <div class="modal">
      <h2>Neue Challenge erstellen</h2>
      <form @submit.prevent="submit">
        <input v-model="title" placeholder="Name" required />
        <textarea v-model="description" placeholder="Beschreibung" required />
        <input v-model.number="points" type="number" min="1" placeholder="Rocketpoints" required />
        <div class="modal-actions">
          <button type="submit">Erstellen</button>
          <button type="button" @click="$emit('close')">Abbrechen</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
const emit = defineEmits(['submit', 'close']);
const title = ref('');
const description = ref('');
const points = ref(1);

function submit() {
  emit('submit', { title: title.value, description: description.value, points: points.value });
  title.value = '';
  description.value = '';
  points.value = 1;
}
</script>

<style scoped>
.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 200;
}
.modal {
  background: #fff;
  border-radius: 1rem;
  padding: 2rem;
  min-width: 320px;
  box-shadow: 0 2px 16px rgba(30,60,114,0.15);
}
.modal-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
}
input, textarea {
  width: 100%;
  margin-bottom: 1rem;
  padding: 0.5rem;
  border-radius: 0.5rem;
  border: 1px solid #e0e7ff;
  font-size: 1rem;
}
button[type="submit"] {
  background: #4f46e5;
  color: #fff;
  border: none;
  padding: 0.5rem 1.2rem;
  border-radius: 0.5rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s;
}
button[type="submit"]:hover {
  background: #22c55e;
}
button[type="button"] {
  background: #eee;
  color: #222;
  border: none;
  padding: 0.5rem 1.2rem;
  border-radius: 0.5rem;
  font-weight: 600;
  cursor: pointer;
}
</style>