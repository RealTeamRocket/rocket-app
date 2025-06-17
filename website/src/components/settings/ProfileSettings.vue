<template>
  <div class="profile-settings">
    <div class="profile-img-row">
      <span v-if="userImage && userImage !== 'https://via.placeholder.com/120'" class="profile-img">
        <img
          :src="userImage"
          alt="Profile"
          class="profile-img"
          @click="showImageModal = true"
          style="cursor:pointer;"
        />
      </span>
      <span v-else class="profile-img profile-initials" :style="{ backgroundColor: userColor }">
        {{ userInitials }}
      </span>
      <div class="profile-actions">
        <input type="file" ref="fileInput" @change="onImageSelected" accept="image/*" hidden />
        <button @click="triggerFileInput">Upload</button>
        <button @click="$emit('delete-image')" class="danger">Delete</button>
      </div>
    </div>
    <label>
      Name:
      <input v-model="localName" type="text" />
      <button @click="emitUpdateName">Save</button>
    </label>
    <label>
      Daily Step Goal:
      <input v-model.number="localDailyGoal" type="number" min="0" />
      <button @click="emitUpdateGoal">Save</button>
    </label>
    <!-- Place modal outside the avatar span, only show if image is not placeholder and showImageModal is true -->
    <ImageModal
      v-if="userImage && userImage !== 'https://via.placeholder.com/120' && showImageModal"
      :show="showImageModal"
      :src="userImage"
      alt="Profile Image"
      @close="showImageModal = false"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import ImageModal from '../modals/ImageModal.vue'

const props = defineProps<{
  userImage: string
  userColor: string
  userInitials: string
  name: string
  dailyGoal: number
}>()
const emit = defineEmits(['update-name', 'update-goal', 'delete-image', 'trigger-file-input', 'image-selected'])

const fileInput = ref<HTMLInputElement | null>(null)
const localName = ref(props.name)
const localDailyGoal = ref(props.dailyGoal)
const showImageModal = ref(false)

watch(() => props.name, (val) => { localName.value = val })
watch(() => props.dailyGoal, (val) => { localDailyGoal.value = val })

function triggerFileInput() {
  emit('trigger-file-input')
  fileInput.value?.click()
}
function onImageSelected(e: Event) {
  emit('image-selected', e)
}
function emitUpdateName() {
  emit('update-name', localName.value)
}
function emitUpdateGoal() {
  emit('update-goal', localDailyGoal.value)
}
</script>

<style scoped>
.profile-settings {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-width: 280px;
  max-width: 350px;
  margin: 0 auto;
  gap: 2rem;
}
.profile-img-row {
  display: flex;
  align-items: center;
  gap: 2rem;
}

.profile-img {
  width: 220px;
  height: 220px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 3rem;
  font-weight: 700;
  color: #fff;
  background: #2a5298;
  user-select: none;
  overflow: hidden;
  box-shadow: 0 2px 16px rgba(0,0,0,0.08);
}
.profile-img img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
  border: 6px solid #bbb;
  box-shadow: 0 2px 16px rgba(0,0,0,0.08);
}
.profile-initials {
  object-fit: none;
  border: 6px solid #bbb;
}
.profile-actions {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
}
label {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin: 0.5rem 0;
  font-size: 1rem;
  width: 100%;
}
input[type="text"], input[type="number"] {
  margin-top: 0.3rem;
  padding: 0.4rem 0.7rem;
  border-radius: 6px;
  border: 1px solid #ccc;
  width: 100%;
  font-size: 1rem;
}
button {
  margin-top: 0.4rem;
  padding: 0.3rem 1.1rem;
  border: none;
  border-radius: 6px;
  background: #1976d2;
  color: #fff;
  cursor: pointer;
  font-size: 1rem;
  transition: background 0.2s;
}
button:hover {
  background: #1565c0;
}
button.danger {
  background: none;
  color: #e53935;
  border: 1px solid #e53935;
}
button.danger:hover {
  background: #ffeaea;
}
</style>
