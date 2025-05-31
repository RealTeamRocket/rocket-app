<template>
  <nav class="rocket-navbar">
    <div class="navbar-content">
      <div class="navbar-left">
        <router-link to="/" class="navbar-brand">🚀 Rocket App</router-link>
        <router-link to="/chat" class="nav-link">Chat</router-link>
        <router-link to="/highscore" class="nav-link">Highscore</router-link>
        <router-link to="/friendlist" class="nav-link">Friendlist</router-link>
        <router-link to="/challenges" class="nav-link">Challenges</router-link>
        <router-link to="/runs" class="nav-link">Runs</router-link>
      </div>
      <div class="navbar-right">
        <template v-if="!isLoggedIn">
          <router-link to="/login" class="custom-btn custom-btn-outline">Login</router-link>
          <router-link to="/register" class="custom-btn custom-btn-primary">Register</router-link>
        </template>
        <template v-else>
          <div class="user-info dropdown" @click="toggleDropdown" @blur="closeDropdown" tabindex="0">
            <span class="user-icon">👤</span>
            <span class="user-name">User</span>
            <span class="dropdown-caret">&#9662;</span>
            <div v-if="dropdownOpen" class="dropdown-menu show">
              <router-link to="/profile" class="dropdown-item">Profile</router-link>
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
import { ref, onMounted } from 'vue'
import { useAuth } from '../utils/useAuth'

const { isLoggedIn, checkAuth, logout } = useAuth()

const dropdownOpen = ref(false)

const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value
}
const closeDropdown = () => {
  setTimeout(() => {
    dropdownOpen.value = false
  }, 100)
}
const handleLogout = async () => {
  dropdownOpen.value = false
  await logout()
}

onMounted(() => {
  checkAuth()
})
</script>

<style scoped>
.rocket-navbar {
  width: 100%;
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
  padding-left: 2rem;  /* Add left padding */
}

.navbar-right {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding-right: 2rem; /* Add right padding */
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

.user-icon {
  font-size: 1.3rem;
}

.user-name {
  margin-left: 0.2rem;
}

.btn-outline-light {
  border-color: #fff;
  color: #fff;
}

.btn-outline-light:hover {
  background: #fff;
  color: #1e3c72;
}

.btn-primary {
  background: #ffb347;
  border: none;
  color: #1e3c72;
}

.btn-primary:hover {
  background: #ffd580;
  color: #1e3c72;
}
</style>
