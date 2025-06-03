<template>
  <div class="register-bg d-flex align-items-center justify-content-center min-vh-100">
    <div class="register-card shadow p-4 rounded bg-white" style="min-width: 350px; max-width: 400px; width: 100%;">
      <div class="mb-3">
        <img src="https://rocket.emoji.gg/🚀" alt="Rocket Logo" style="width: 48px; height: 48px;" class="mb-2" />
        <h2 class="mb-1">Create Account</h2>
        <p class="text-muted mb-3">Sign up for Rocket App</p>
      </div>
      <form @submit.prevent="handleRegister">
        <div class="form-group mb-3 text-start">
          <label for="username" class="form-label">Username</label>
          <div class="input-group">
            <span class="input-group-text"><i class="bi bi-person"></i></span>
            <input v-model="username" type="text" id="username" class="form-control" placeholder="Username" required />
          </div>
        </div>
        <div class="form-group mb-3 text-start">
          <label for="email" class="form-label">Email</label>
          <div class="input-group">
            <span class="input-group-text"><i class="bi bi-envelope"></i></span>
            <input v-model="email" type="email" id="email" class="form-control" placeholder="you@email.com" required />
          </div>
        </div>
        <div class="form-group mb-3 text-start">
          <label for="password" class="form-label">Password</label>
          <div class="input-group">
            <span class="input-group-text"><i class="bi bi-lock"></i></span>
            <input v-model="password" type="password" id="password" class="form-control" placeholder="Password" required />
          </div>
        </div>
        <div class="form-group mb-3 text-start">
          <label for="repeatPassword" class="form-label">Repeat Password</label>
          <div class="input-group">
            <span class="input-group-text"><i class="bi bi-lock"></i></span>
            <input v-model="repeatPassword" type="password" id="repeatPassword" class="form-control" placeholder="Repeat Password" required />
          </div>
        </div>
        <div v-if="error" class="alert alert-danger py-2 mb-3" role="alert">
          {{ error }}
        </div>
        <button class="btn btn-success w-100" type="submit">Register</button>
        <div class="text-center mt-3">
          <span>Already have an account?</span>
          <a href="/login" class="ms-1">Login</a>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api/backend-api'

const username = ref('')
const email = ref('')
const password = ref('')
const repeatPassword = ref('')
const error = ref('')
const router = useRouter()

async function handleRegister() {
  error.value = ''
  if (password.value !== repeatPassword.value) {
    error.value = "Passwords do not match."
    return
  }
  try {
    const response = await api.register(username.value, email.value, password.value)
    if (response.status === 200 || response.status === 201) {
      router.push('/login')
    } else {
      error.value = response.data?.message || 'Registration failed.'
    }
  } catch (err: any) {
    error.value = err.response?.data?.message || err.message || 'Registration failed.'
  }
}
</script>

<style scoped>
.register-bg {
  background: linear-gradient(135deg, #e0ffe7 0%, #f8fafc 100%);
}
.register-card {
  animation: fadeIn 0.7s;
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(30px);}
  to { opacity: 1; transform: translateY(0);}
}
</style>
