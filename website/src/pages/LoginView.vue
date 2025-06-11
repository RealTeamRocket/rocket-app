<template>
  <div class="login-bg d-flex align-items-center justify-content-center min-vh-100">
    <div class="login-card shadow p-4 rounded bg-white" style="min-width: 350px; max-width: 400px; width: 100%;">
      <div class="mb-3">
        <h2 class="mb-1">Rocket App</h2>
        <p class="text-muted mb-3">Sign in to your account</p>
      </div>
      <form @submit.prevent="handleLogin">
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
        <div v-if="error" class="alert alert-danger py-2 mb-3" role="alert">
          {{ error }}
        </div>
        <button class="btn btn-primary w-100" type="submit">Login</button>
        <div class="text-center mt-3">
          <span>Don't have an account?</span>
          <a href="/register" class="ms-1">Register</a>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '@/utils/useAuth'

const email = ref('')
const password = ref('')
const error = ref('')

const { login } = useAuth()

const handleLogin = async () => {
  error.value = ''
  try {
    await login(email.value, password.value)
  } catch (e: any) {
    error.value = e?.response?.data?.error || e.message || 'Login failed'
  }
}
</script>

<style scoped>
.login-bg {
  background: linear-gradient(135deg, #e0e7ff 0%, #f8fafc 100%);
}
.login-card {
  animation: fadeIn 0.7s;
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(30px);}
  to { opacity: 1; transform: translateY(0);}
}
</style>
