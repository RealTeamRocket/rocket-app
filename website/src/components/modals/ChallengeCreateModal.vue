<template>
  <div class="modal-overlay" @click.self="close">
    <div class="modal-content">
      <button class="close-btn" @click="close" aria-label="Close">&times;</button>
      <h2>Create New Challenge</h2>
      <form @submit.prevent="onSubmit">
        <div class="form-group">
          <label for="title">Title</label>
          <input
            id="title"
            v-model="form.title"
            type="text"
            placeholder="Enter challenge title"
            required
            maxlength="60"
          />
        </div>
        <div class="form-group">
          <label for="description">Description</label>
          <textarea
            id="description"
            v-model="form.description"
            placeholder="Describe the challenge"
            required
            maxlength="300"
            rows="3"
          ></textarea>
        </div>
        <div class="form-group">
          <label for="points">Points</label>
          <input
            id="points"
            v-model.number="form.points"
            type="number"
            min="1"
            max="100"
            required
          />
        </div>
        <div class="modal-actions">
          <button type="submit" class="submit-btn">Create</button>
          <button type="button" class="cancel-btn" @click="close">Cancel</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, defineEmits, defineProps } from 'vue';

const emit = defineEmits(['submit', 'close']);
const props = defineProps<{
  show: boolean
}>();

const form = ref({
  title: '',
  description: '',
  points: 1,
});

watch(() => props.show, (val) => {
  if (val) {
    form.value = { title: '', description: '', points: 1 };
  }
});

function onSubmit() {
  emit('submit', { ...form.value });
}
function close() {
  emit('close');
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(30, 27, 75, 0.55);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
  animation: fadeIn 0.2s;
}
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
.modal-content {
  background: linear-gradient(135deg, #f8fafc 70%, #e0e7ff 100%);
  border-radius: 1.25rem;
  box-shadow: 0 8px 32px rgba(80, 0, 120, 0.18);
  padding: 2.5rem 2rem 2rem 2rem;
  min-width: 340px;
  max-width: 95vw;
  position: relative;
  animation: popIn 0.18s;
}
@keyframes popIn {
  from { transform: scale(0.96); opacity: 0.7; }
  to { transform: scale(1); opacity: 1; }
}
.close-btn {
  position: absolute;
  top: 1.1rem;
  right: 1.3rem;
  background: none;
  border: none;
  font-size: 2rem;
  color: #7c3aed;
  cursor: pointer;
  transition: color 0.2s;
  z-index: 10;
}
.close-btn:hover {
  color: #da13ab;
}
h2 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #4f46e5;
  font-weight: 800;
  letter-spacing: 0.01em;
}
.form-group {
  margin-bottom: 1.2rem;
  display: flex;
  flex-direction: column;
}
label {
  font-weight: 600;
  margin-bottom: 0.4rem;
  color: #3730a3;
}
input[type="text"],
input[type="number"],
textarea {
  border: 1.5px solid #a5b4fc;
  border-radius: 0.5rem;
  padding: 0.6rem 0.9rem;
  font-size: 1rem;
  background: #fff;
  transition: border 0.2s;
  outline: none;
}
input:focus,
textarea:focus {
  border-color: #7c3aed;
}
textarea {
  resize: vertical;
  min-height: 2.5rem;
  max-height: 8rem;
}
.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;
}
.submit-btn {
  background: linear-gradient(90deg, #7c3aed 60%, #da13ab 100%);
  color: #fff;
  font-weight: 700;
  border: none;
  border-radius: 0.5rem;
  padding: 0.6rem 1.6rem;
  cursor: pointer;
  font-size: 1rem;
  box-shadow: 0 2px 8px rgba(124, 58, 237, 0.08);
  transition: background 0.18s, box-shadow 0.18s;
}
.submit-btn:hover {
  background: linear-gradient(90deg, #da13ab 60%, #7c3aed 100%);
  box-shadow: 0 4px 16px rgba(218, 19, 171, 0.13);
}
.cancel-btn {
  background: #e0e7ff;
  color: #4f46e5;
  font-weight: 600;
  border: none;
  border-radius: 0.5rem;
  padding: 0.6rem 1.2rem;
  cursor: pointer;
  font-size: 1rem;
  transition: background 0.18s, color 0.18s;
}
.cancel-btn:hover {
  background: #c7d2fe;
  color: #da13ab;
}
</style>
