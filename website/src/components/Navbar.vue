<template>
  <nav class="rocket-navbar">
    <div class="navbar-content">
      <div class="navbar-left">
        <router-link to="/" class="navbar-brand">
          <img src="/src/assets/icons/rocket.svg" alt="Rocket" style="width:1.3em;height:1.3em;vertical-align:middle;margin-right:0.35em;" />
          <span>Rocket App</span>
        </router-link>
        <router-link to="/chat" class="nav-link">Chat</router-link>
        <router-link to="/highscore" class="nav-link">Highscore</router-link>
        <router-link to="/friendlist" class="nav-link">Friendlist</router-link>
        <router-link to="/challenges" class="nav-link">Challenges</router-link>
        <router-link to="/runs" class="nav-link">Runs</router-link>
        <router-link to="/download" class="nav-link">Download</router-link>
      </div>
      <div class="navbar-right">
        <template v-if="!isLoggedIn">
          <router-link to="/login" class="nav-auth-btn nav-auth-login">Login</router-link>
          <router-link to="/register" class="nav-auth-btn nav-auth-register">Register</router-link>
        </template>
        <template v-else>
          <div
            class="user-info dropdown"
            @click="toggleDropdown"
            tabindex="0"
            ref="dropdownRef"
          >
            <span v-if="userImage" class="user-avatar-img">
              <img :src="userImage" alt="User" style="width:1.7em;height:1.7em;vertical-align:middle;border-radius:50%;" />
            </span>
            <span v-else class="user-avatar-initials" :style="{ backgroundColor: userColor }">
              {{ userInitials }}
            </span>
            <span class="user-name">{{ user?.username || 'User' }}</span>
            <span class="dropdown-caret">&#9662;</span>
            <div v-if="dropdownOpen" class="dropdown-menu show">
              <router-link
                v-if="user?.username"
                :to="`/profile/${user.username}`"
                class="dropdown-item"
              >Profile</router-link>
              <router-link
                v-else
                to="/profile"
                class="dropdown-item"
              >Profile</router-link>
              <router-link to="/settings" class="dropdown-item">Settings</router-link>
              <button class="dropdown-item" @click.stop="handleLogout">Logout</button>
            </div>
          </div>
        </template>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import { useAuth } from '@/utils/useAuth'
import { useRouter } from 'vue-router'
import api from '@/api/backend-api'
import { getColor, getInitials } from '@/utils/userUtils'

interface User {
  username: string
  name: string
  mime_type: string
  data: string | null
}

const router = useRouter()

const { isLoggedIn, checkAuth, logout } = useAuth()

const dropdownOpen = ref(false)
const dropdownRef = ref<HTMLElement | null>(null)

const user = ref<User | null>(null)

const userImage = computed(() => {
  if (user.value && user.value.data && user.value.mime_type) {
    return `data:${user.value.mime_type};base64,${user.value.data}`
  }
  return null
})

const userColor = computed(() => {
  if (user.value && user.value.username) {
    return getColor(user.value.username)
  }
  return '#2a5298'
})

const userInitials = computed(() => {
  if (user.value && user.value.username) {
    return getInitials(user.value.username)
  }
  return ''
})

const toggleDropdown = (event: MouseEvent) => {
  dropdownOpen.value = !dropdownOpen.value
}

const handleClickOutside = (event: MouseEvent) => {
  if (
    dropdownRef.value &&
    !dropdownRef.value.contains(event.target as Node)
  ) {
    dropdownOpen.value = false
  }
}
const handleLogout = async () => {
  dropdownOpen.value = false
  await logout()
  router.push('/')
  window.location.reload()
}

onMounted(async () => {
  document.addEventListener('click', handleClickOutside)
  checkAuth()
  try {
    const response = await api.getUserImage()
    user.value = response.data
  } catch (e) {
    user.value = null
  }
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.rocket-navbar {
  min-width: 100vw;
  width: max-content;
  background: linear-gradient(90deg, #1e3c72 0%, #2a5298 100%);
  box-shadow: 0 2px 8px rgba(30,60,114,0.1);
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}

.navbar-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  padding: 0;
  min-height: 56px;
}

.navbar-left {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  padding-left: 2rem;
}

.navbar-right {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding-right: 2rem;
}

.navbar-brand {
  font-weight: bold;
  color: #fff;
  text-decoration: none;
  font-size: 1.3rem;
  margin-right: 2rem;
}

.nav-link {
  color: #e0e0e0;
  text-decoration: none;
  font-size: 1rem;
  padding: 0.25rem 0.5rem;
  transition: color 0.2s;
}

.nav-link:hover {
  color: #fff;
  text-decoration: underline;
}

.user-info {
  display: flex;
  align-items: center;
  color: #fff;
  font-weight: 500;
  gap: 0.5rem;
  position: relative;
  cursor: pointer;
  outline: none;
}

.dropdown-caret {
  margin-left: 0.3rem;
  font-size: 0.9rem;
}

.dropdown-menu {
  position: absolute;
  top: 120%;
  right: 0;
  background: #fff;
  color: #1e3c72;
  border-radius: 0.4rem;
  box-shadow: 0 2px 8px rgba(30,60,114,0.15);
  min-width: 140px;
  z-index: 100;
  display: flex;
  flex-direction: column;
  padding: 0.3rem 0;
  animation: fadeIn 0.2s;
}

.dropdown-item {
  padding: 0.5rem 1rem;
  background: none;
  border: none;
  color: #1e3c72;
  text-align: left;
  text-decoration: none;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.15s;
}

.dropdown-item:hover {
  background: #e0e7ff;
  color: #2a5298;
}

.user-avatar-img img {
  display: inline-block;
  border-radius: 50%;
  object-fit: cover;
  background: #fff;
}

.user-avatar-initials {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 1.7em;
  height: 1.7em;
  border-radius: 50%;
  color: #fff;
  font-weight: 700;
  font-size: 1.1em;
  margin-right: 0.2em;
  background: #2a5298;
  vertical-align: middle;
  user-select: none;
}

.user-icon {
  font-size: 1.3rem;
}

.user-name {
  margin-left: 0.2rem;
}

.nav-auth-btn {
  padding: 0.35rem 1.1rem;
  border-radius: 6px;
  font-size: 1rem;
  font-weight: 500;
  border: none;
  outline: none;
  text-decoration: none;
  transition: background 0.18s, color 0.18s, border 0.18s;
  margin-left: 0.3rem;
  margin-right: 0.3rem;
  cursor: pointer;
  display: inline-block;
}

.nav-auth-login {
  background: transparent;
  color: #fff;
  border: 1.5px solid #fff;
}

.nav-auth-login:hover {
  background: #fff;
  color: #1e3c72;
  border-color: #fff;
  text-decoration: none;
}

.nav-auth-register {
  background: #ffb347;
  color: #1e3c72;
  border: 1.5px solid #ffb347;
}

.nav-auth-register:hover {
  background: #ffd580;
  color: #1e3c72;
  border-color: #ffd580;
  text-decoration: none;
}
</style>
