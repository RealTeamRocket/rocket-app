<template>
  <Navbar :key="navbarKey" />
  <div class="settings-page">
    <h1>Settings</h1>
    <div class="settings-card">
      <!-- Profile Image Section -->
      <div class="profile-section">
        <span
          v-if="userImage && userImage !== 'https://via.placeholder.com/120'"
          class="profile-img"
        >
          <img :src="userImage" alt="Profile" class="profile-img" />
        </span>
        <span v-else class="profile-img profile-initials" :style="{ backgroundColor: userColor }">
          {{ userInitials }}
        </span>
        <div class="profile-actions">
          <input type="file" ref="fileInput" @change="onImageSelected" accept="image/*" hidden />
          <button @click="triggerFileInput">Upload New Image</button>
          <button @click="showDeleteImageConfirm = true" class="danger">Delete Image</button>
        </div>
      </div>

      <!-- User Info Section -->
      <div class="info-section">
        <label>
          Name:
          <input v-model="name" type="text" />
          <button @click="updateName" style="margin-left: 0.5rem">Save Name</button>
        </label>
        <label>
          Email:
          <input v-model="email" type="email" />
          <button @click="updateEmail" style="margin-left: 0.5rem">Save Email</button>
        </label>
      </div>

      <!-- Password Change Section -->
      <div class="password-section">
        <label>
          Current Password:
          <input v-model="currentPassword" type="password" />
        </label>
        <label>
          New Password:
          <input v-model="newPassword" type="password" />
        </label>
        <label>
          Confirm New Password:
          <input v-model="confirmPassword" type="password" />
        </label>
        <button @click="changePassword">Change Password</button>
      </div>

      <!-- Step Goal Section -->
      <div class="goal-section">
        <label>
          Daily Step Goal:
          <input v-model.number="dailyGoal" type="number" min="0" />
        </label>
        <button @click="updateGoal">Update Goal</button>
      </div>

      <!-- Danger Zone -->
      <div class="danger-zone">
        <button @click="logout" class="logout">Logout</button>
        <button @click="showDeleteAccountConfirm = true" class="danger">Delete Account</button>
      </div>
    </div>
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
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import Navbar from '@/components/Navbar.vue'
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
  autoClose: 2200
})

function showNotification(
  message: string,
  type: 'success' | 'error' | 'info' = 'info',
  autoClose = 2200
) {
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
  } catch (error) {
    showNotification('Failed to delete image', 'error')
  }
}

async function updateName() {
  try {
    await api.updateUserInfo({ name: name.value })
    savedName.value = name.value
    showNotification('Name updated!', 'success')
    navbarKey.value++ // Force Navbar reload
  } catch (e) {
    showNotification('Failed to update name', 'error')
  }
}

async function updateEmail() {
  try {
    await api.updateUserInfo({ email: email.value })
    showNotification('Email updated!', 'success')
  } catch (e) {
    showNotification('Failed to update email', 'error')
  }
}

async function changePassword() {
  if (!currentPassword.value || !newPassword.value || !confirmPassword.value) {
    showNotification('Please fill all password fields', 'error')
    return
  }
  if (newPassword.value !== confirmPassword.value) {
    showNotification('New passwords do not match', 'error')
    return
  }
  try {
    await api.updateUserInfo({
      currentPassword: currentPassword.value,
      newPassword: newPassword.value
    })
    showNotification('Password updated!', 'success')
    currentPassword.value = ''
    newPassword.value = ''
    confirmPassword.value = ''
  } catch (e) {
    showNotification('Failed to update password', 'error')
  }
}

async function updateGoal() {
  try {
    await api.updateStepGoal(dailyGoal.value)
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
  if (confirm('Are you sure you want to delete your account? This cannot be undone.')) {
    try {
      await api.deleteUser()
      router.push
    } catch (error) {
      alert('Failed to delete account')
    }
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
.settings-page {
  max-width: 500px;
  margin: 2rem auto;
  padding: 1.5rem;
}
.settings-card {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 16px rgba(0, 0, 0, 0.07);
  padding: 2rem;
  display: flex;
  flex-direction: column;
  gap: 2rem;
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
.goal-section label {
  display: flex;
  flex-direction: column;
  margin-bottom: 1rem;
}
.info-section input,
.goal-section input {
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
</style>
