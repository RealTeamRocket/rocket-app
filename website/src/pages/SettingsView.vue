<template>
  <Navbar :key="navbarKey" />
  <div class="settings-tabbar-layout">
    <div class="settings-tabbar">
      <button
        :class="{ active: section === 'profile' }"
        @click="section = 'profile'"
      >Profile</button>
      <button
        :class="{ active: section === 'security' }"
        @click="section = 'security'"
      >Security</button>
      <button
        :class="{ active: section === 'account' }"
        @click="section = 'account'"
      >Account</button>
    </div>
    <div class="settings-tab-content">
      <ProfileSettings
        v-if="section==='profile'"
        :userImage="userImage"
        :userColor="userColor"
        :userInitials="userInitials"
        :name="name"
        :dailyGoal="dailyGoal"
        @update-name="updateName"
        @update-goal="updateGoal"
        @delete-image="showDeleteImageConfirm = true"
        @trigger-file-input="triggerFileInput"
        @image-selected="onImageSelected"
      />
      <SecuritySettings
        v-if="section==='security'"
        :email="email"
        @update-email="updateEmail"
        @change-password="changePassword"
        :currentPassword="currentPassword"
        :newPassword="newPassword"
        :confirmPassword="confirmPassword"
        @update:currentPassword="val => currentPassword = val"
        @update:newPassword="val => newPassword = val"
        @update:confirmPassword="val => confirmPassword = val"
      />
      <AccountSettings
        v-if="section==='account'"
        :userImage="userImage"
        :userColor="userColor"
        :userInitials="userInitials"
        @logout="logout"
        @delete-account="showDeleteAccountConfirm = true"
      />
      <!-- Modals -->
      <ConfirmDialog
        :open="showDeleteImageConfirm"
        title="Delete Profile Image"
        message="Are you sure you want to delete your profile image?"
        confirmText="Delete"
        cancelText="Cancel"
        @confirm="deleteImage"
        @cancel="showDeleteImageConfirm = false"
      />
      <ConfirmDialog
        :open="showDeleteAccountConfirm"
        title="Delete Account"
        message="Are you sure you want to delete your account? This cannot be undone."
        confirmText="Delete"
        cancelText="Cancel"
        @confirm="deleteAccount"
        @cancel="showDeleteAccountConfirm = false"
      />
      <NotificationModal
        :open="notification.open"
        :message="notification.message"
        :type="notification.type"
        :autoClose="notification.autoClose"
        @close="notification.open = false"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import Navbar from '@/components/Navbar.vue'
import ProfileSettings from '@/components/settings/ProfileSettings.vue'
import SecuritySettings from '@/components/settings/SecuritySettings.vue'
import AccountSettings from '@/components/settings/AccountSettings.vue'
import api from '@/api/backend-api'
import { useRouter } from 'vue-router'
import ConfirmDialog from '@/components/modals/ConfirmDialog.vue'
import NotificationModal from '@/components/modals/NotificationModal.vue'
import { getColor, getInitials } from '@/utils/userUtils'


const router = useRouter()

const userImage = ref('https://via.placeholder.com/120')
const name = ref('')
const savedName = ref('')
const email = ref('')

const navbarKey = ref(0)
const section = ref<'profile'|'security'|'account'>('profile')

const currentPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const dailyGoal = ref(10000)
const fileInput = ref<HTMLInputElement | null>(null)

const userColor = computed(() => {
  return getColor(savedName.value || 'User')
})
const userInitials = computed(() => {
  return getInitials(savedName.value || 'User')
})

const showDeleteImageConfirm = ref(false)
const showDeleteAccountConfirm = ref(false)

const notification = reactive({
  open: false,
  message: '',
  type: 'info' as 'success' | 'error' | 'info',
  autoClose: 2200,
})



function showNotification(message: string, type: 'success' | 'error' | 'info' = 'info', autoClose = 2200) {
  notification.message = message
  notification.type = type
  notification.open = true
  notification.autoClose = autoClose
}

function triggerFileInput() {
  fileInput.value?.click()
}

async function onImageSelected(event: Event) {
  const file = (event.target as HTMLInputElement).files?.[0]
  if (file) {
    try {
      await api.uploadImage(file)
      // Fetch the new image from the backend to get the base64-encoded image
      const imageRes = await api.getUserImage()
      if (imageRes.data && imageRes.data.data && imageRes.data.mime_type) {
        userImage.value = `data:${imageRes.data.mime_type};base64,${imageRes.data.data}`
      } else {
        userImage.value = 'https://via.placeholder.com/120'
      }
      showNotification('Profile image updated!', 'success')
      navbarKey.value++
    } catch (error) {
      showNotification('Failed to upload image', 'error')
    }
  }
}

async function deleteImage() {
  showDeleteImageConfirm.value = false
  try {
    await api.deleteImage()
    userImage.value = 'https://via.placeholder.com/120'
    showNotification('Profile image deleted.', 'success')
    navbarKey.value++
  } catch (error) {
    showNotification('Failed to delete image', 'error')
  }
}

async function updateName(newName: string) {
  try {
    console.log('Updating name to:', newName)
    await api.updateUserInfo({ name: newName })
    name.value = newName
    savedName.value = newName
    showNotification('Name updated!', 'success')
    navbarKey.value++ // Force Navbar reload
  } catch (e) {
    showNotification('Failed to update name', 'error')
  }
}

async function updateEmail(newEmail: string) {
  try {
    console.log('Updating email to:', newEmail)
    await api.updateUserInfo({ email: newEmail })
    email.value = newEmail
    showNotification('Email updated!', 'success')
  } catch (e) {
    showNotification('Failed to update email', 'error')
  }
}

async function changePassword({ currentPassword: curr, newPassword: next, confirmPassword: confirm }: { currentPassword: string, newPassword: string, confirmPassword: string }) {
  if (!curr || !next || !confirm) {
    showNotification('Please fill all password fields', 'error')
    return
  }
  if (next !== confirm) {
    showNotification('New passwords do not match', 'error')
    return
  }
  try {
    await api.updateUserInfo({
      currentPassword: curr,
      newPassword: next
    })
    showNotification('Password updated!', 'success')
    currentPassword.value = ''
    newPassword.value = ''
    confirmPassword.value = ''
  } catch (e) {
    showNotification('Failed to update password', 'error')
  }
}

async function updateGoal(newGoal: number) {
  try {
    await api.updateStepGoal(newGoal)
    dailyGoal.value = newGoal
    showNotification('Daily goal updated', 'success')
  } catch (error) {
    showNotification('Failed to update daily goal', 'error')
  }
}

async function logout() {
  try {
    await api.logout()
    router.push('/login')
  } catch (error) {
    alert('Failed to logout')
  }
}

async function deleteAccount() {
    try {
      await api.deleteUser()
      router.push('/login')
    } catch (error) {
      alert('Failed to delete account')
    }
}

// Optionally, fetch user info and settings on mount
onMounted(async () => {
  try {
    const [userRes, settingsRes, imageRes] = await Promise.all([
      api.getMyself(),
      api.getSettings ? api.getSettings() : Promise.resolve({ data: {} }),
      api.getUserImage ? api.getUserImage() : Promise.resolve({ data: {} })
    ])
    if (userRes.data) {
      name.value = userRes.data.username || ''
      savedName.value = userRes.data.username || ''
      email.value = userRes.data.email || ''
    }
    if (settingsRes.data && settingsRes.data.step_goal) {
      dailyGoal.value = settingsRes.data.step_goal
    }
    // Handle base64 user image like Navbar.vue
    if (imageRes.data && imageRes.data.data && imageRes.data.mime_type) {
      userImage.value = `data:${imageRes.data.mime_type};base64,${imageRes.data.data}`
    } else {
      userImage.value = 'https://via.placeholder.com/120'
    }
  } catch (e) {
    userImage.value = 'https://via.placeholder.com/120'
    // Ignore errors for optional fetches
  }
})
</script>

<style scoped>
.settings-tabbar-layout {
  min-height: 80vh;
  background: #f6f8fa;
  display: flex;
  flex-direction: column;
  align-items: stretch;
}

.settings-tabbar {
  display: flex;
  gap: 1rem;
  padding: 2.5rem 2rem 1.5rem 2rem;
  background: #f7fafd;
  border-bottom: 1px solid #e0eaff;
  justify-content: center;
}

.settings-tabbar button {
  background: none;
  border: none;
  font-size: 1.1rem;
  font-weight: 600;
  color: #4a90e2;
  padding: 0.5rem 1.5rem;
  border-radius: 6px 6px 0 0;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
}

.settings-tabbar button.active,
.settings-tabbar button:hover {
  background: #e0eaff;
  color: #222;
}

.settings-tab-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding: 3rem 2rem;
  min-width: 0;
}

.profile-section {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}
.profile-img {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #eee;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2.5rem;
  font-weight: 700;
  color: #fff;
  background: #2a5298;
  user-select: none;
  overflow: hidden;
}
.profile-initials {
  object-fit: none;
  border: 2px solid #eee;
}
.profile-actions {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}
.info-section label,
.goal-section label,
.password-section label {
  display: flex;
  flex-direction: column;
  margin-bottom: 1rem;
}
.info-section input,
.goal-section input,
.password-section input {
  padding: 0.5rem;
  border-radius: 6px;
  border: 1px solid #ccc;
  margin-top: 0.25rem;
}
button {
  padding: 0.5rem 1.2rem;
  border: none;
  border-radius: 6px;
  background: #1976d2;
  color: #fff;
  cursor: pointer;
  margin-top: 0.5rem;
  transition: background 0.2s;
}
button:hover {
  background: #1565c0;
}
button.danger {
  background: #e53935;
}
button.danger:hover {
  background: #b71c1c;
}
button.logout {
  background: #757575;
}
button.logout:hover {
  background: #424242;
}
.danger-zone {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
  margin-top: 2rem;
}
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.25s;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>
