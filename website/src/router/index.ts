import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/pages/HomeView.vue'
import LoginView from '@/pages/LoginView.vue'
import RegisterView from '@/pages/RegisterView.vue'
import ChatView from '@/pages/ChatView.vue'
import HighscoreView from '@/pages/HighscoreView.vue'
import FriendlistView from '@/pages/FriendlistView.vue'
import ChallengesView from '@/pages/ChallengesView.vue'
import RunsView from '@/pages/RunsView.vue'
import ProfileView from '@/pages/ProfileView.vue'
import SettingsView from '@/pages/SettingsView.vue'
import { useAuth } from '@/utils/useAuth'

const routes = [
  { path: '/', component: HomeView },
  { path: '/login', component: LoginView },
  { path: '/register', component: RegisterView },
  { path: '/chat', component: ChatView },
  { path: '/highscore', component: HighscoreView },
  { path: '/friendlist', component: FriendlistView },
  { path: '/challenges', component: ChallengesView },
  { path: '/runs', component: RunsView },
  { path: '/settings', component: SettingsView },
  { path: '/profile', component: ProfileView },
  { path: '/:pathMatch(.*)*', redirect: '/' }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to, from, next) => {
  const { isLoggedIn, checkAuth } = useAuth()
  const publicPages = ['/', '/login', '/register']
  const authRequired = !publicPages.includes(to.path)

  // Always check auth before proceeding
  await checkAuth()

  if (authRequired && !isLoggedIn.value) {
    return next('/login')
  }
  next()
})

export default router
