import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import api from '../api/backend-api'

const loggedIn = ref(false)

export function useAuth() {
  const router = useRouter()

  async function login(email: string, password: string) {
    try {
      const response = await api.login(email, password)
      if (response.status === 200) {
        loggedIn.value = true
        router.push('/')
      } else {
        throw new Error('Login failed')
      }
    } catch (error) {
      throw error
    }
  }

  async function logout() {
    try {
      await api.logout()
    } catch (e) {
      console.error("eyy yo somethign went wrong")
    }
    loggedIn.value = false
    router.push('/')
    window.location.reload()
  }

  async function checkAuth() {
    try {
      const response = await api.checkAuthenticated()
      loggedIn.value = response.status === 200 && response.data.authenticated === "true"
    } catch {
      loggedIn.value = false
    }
  }

  return {
    isLoggedIn: computed(() => loggedIn.value),
    login,
    logout,
    checkAuth,
  }
}
