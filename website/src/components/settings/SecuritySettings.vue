<template>
  <div class="security-settings">
    <form @submit.prevent>
      <div class="form-group">
        <label>Email</label>
        <div class="input-row">
          <input v-model="localEmail" type="email" autocomplete="email" />
          <button type="button" @click="emitUpdateEmail">Save</button>
        </div>
      </div>
      <div class="form-group">
        <label>Change Password</label>
        <input
          v-model="localCurrentPassword"
          type="password"
          placeholder="Current password"
          autocomplete="current-password"
          @input="emitCurrentPassword"
        />
        <input
          v-model="localNewPassword"
          type="password"
          placeholder="New password"
          autocomplete="new-password"
          @input="emitNewPassword"
        />
        <input
          v-model="localConfirmPassword"
          type="password"
          placeholder="Confirm new password"
          autocomplete="new-password"
          @input="emitConfirmPassword"
        />
        <button type="button" @click="emitChangePassword">Change Password</button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps<{
  email: string
  currentPassword: string
  newPassword: string
  confirmPassword: string
}>()

const emit = defineEmits<{
  (e: 'update-email', value: string): void
  (e: 'change-password', value: { currentPassword: string; newPassword: string; confirmPassword: string }): void
  (e: 'update:currentPassword', value: string): void
  (e: 'update:newPassword', value: string): void
  (e: 'update:confirmPassword', value: string): void
}>()

const localEmail = ref(props.email)
const localCurrentPassword = ref(props.currentPassword)
const localNewPassword = ref(props.newPassword)
const localConfirmPassword = ref(props.confirmPassword)

watch(() => props.email, val => { localEmail.value = val })
watch(() => props.currentPassword, val => { localCurrentPassword.value = val })
watch(() => props.newPassword, val => { localNewPassword.value = val })
watch(() => props.confirmPassword, val => { localConfirmPassword.value = val })

function emitUpdateEmail() {
  emit('update-email', localEmail.value)
}
function emitChangePassword() {
  emit('change-password', {
    currentPassword: localCurrentPassword.value,
    newPassword: localNewPassword.value,
    confirmPassword: localConfirmPassword.value
  })
}
function emitCurrentPassword() {
  emit('update:currentPassword', localCurrentPassword.value)
}
function emitNewPassword() {
  emit('update:newPassword', localNewPassword.value)
}
function emitConfirmPassword() {
  emit('update:confirmPassword', localConfirmPassword.value)
}
</script>

<style scoped>
.security-settings {
  width: 100%;
  max-width: 340px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 2.5rem;
}
form {
  width: 100%;
}
.form-group {
  margin-bottom: 2rem;
  display: flex;
  flex-direction: column;
  gap: 0.7rem;
}
label {
  font-size: 1rem;
  font-weight: 500;
  margin-bottom: 0.2rem;
  color: #232946;
}
.input-row {
  display: flex;
  gap: 0.5rem;
}
input[type="email"],
input[type="password"] {
  flex: 1;
  padding: 0.5rem 0.7rem;
  border-radius: 6px;
  border: 1px solid #ccc;
  font-size: 1rem;
  margin-bottom: 0.3rem;
}
button {
  padding: 0.45rem 1.1rem;
  border: none;
  border-radius: 6px;
  background: #1976d2;
  color: #fff;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.18s;
}
button:hover {
  background: #1565c0;
}
</style>