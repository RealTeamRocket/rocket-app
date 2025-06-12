<template>
  <img src="/src/assets/icons/rocket.svg" alt="Rocket" class="rocket-fly" />
  <Navbar/>
  <div class="highscore-page">
    <div class="header">
      <h1>
        <img src="/src/assets/icons/rocket.svg" alt="Rocket"
             style="width:1.3em;height:1.3em;vertical-align:middle;margin-right:0.35em;"/>
        Highscore
        <img src="/src/assets/icons/rocket.svg" alt="Rocket"
             style="width:1.3em;height:1.3em;vertical-align:middle;margin-right:0.35em;"/>
      </h1>
      <p>This is the Highscore page. Here you will see the top scores and rankings.</p>
    </div>
    <div class="main">
      <div class="switchBtn">
        <button :class="{active: isGlobal}" @click="isGlobal = true">Global</button>
        <button :class="{active: !isGlobal}" @click="isGlobal = false">Friends only</button>
      </div>
      <div class="podium">
        <div class="podium-row podium-row-top">
          <div class="podium-place first">
            <span v-if="selectedList[0]">{{ selectedList[0].username }}</span>
            <span v-else>—</span>
          </div>
        </div>
        <div class="podium-row podium-row-bottom">
          <div class="podium-place second">
            <span v-if="selectedList[1]">{{ selectedList[1].username }}</span>
            <span v-else>—</span>
          </div>
          <div class="podium-place third">
            <span v-if="selectedList[2]">{{ selectedList[2].username }}</span>
            <span v-else>—</span>
          </div>
        </div>
      </div>
      <div class="rest">
        <div
            class="rest-user"
            v-for="(user, idx) in selectedList.slice(3, 28)"
            :key="user.id || idx"
        >
          <img
              class="user-avatar"
              src="/src/assets/icons/rocket.svg"
              alt="User Icon"
          />
          {{ user.username }}
        </div>
      </div>
  </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
import {ref, onMounted, computed} from 'vue'
import api from '@/api/backend-api'

const isGlobal = ref(true)
const rankedUsers = ref([])
const rankedFriends = ref([])

const loadRanking = async () => {
  try {
    const {data: userData} = await api.getRankedUsers()
    rankedUsers.value = userData || []

    const {data: friendData} = await api.getRankedFriends()
    rankedFriends.value = friendData || []
  } catch {
    console.error("Error fetching ranking")
  }
}
const selectedList = computed(() => isGlobal.value ? rankedUsers.value : rankedFriends.value)

onMounted(async () => {
  await loadRanking()
  fillWithDummyUsers()
})

function fillWithDummyUsers() {
  if (selectedList.value.length <= 1) {
    const dummy = Array.from({ length: 25 }, (_, i) => ({
      id: `dummy-${i + 1}`,
      username: `User${i + 1}`
    }))
    if (isGlobal.value) {
      rankedUsers.value = dummy
    } else {
      rankedFriends.value = dummy
    }
  }
}

console.log(selectedList)








</script>

<style scoped>
.highscore-page {
  padding: 2rem;
}

.header {
  flex-direction: column;
  display: flex;
  justify-content: center;
  text-align: center;
  margin-bottom: 2rem;
}

.main {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  height: 100vh;
}

.switchBtn {
  display: flex;
  gap: 1rem;
  margin-bottom: 3rem;
  border: 2px solid #007bff;
  border-radius: 24px;
  background: #f5f5f5;
  padding: 0.5rem 1rem;
}

.switchBtn button:hover {
  background: lightgray;
}

.switchBtn button {
  padding: 0.5rem 1.5rem;
  border: none;
  background: #eee;
  color: #333;
  border-radius: 20px;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
}

.switchBtn button.active {
  background: linear-gradient(90deg, #1e3c72 0%, #2a5298 100%);
  box-shadow: 0 2px 8px rgba(30, 60, 114, 0.1);
  color: #fff;
}

.podium {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
  width: 100%;
  max-width: 900px;
  justify-content: center;
  padding: 2rem 0;
}

.podium-row {
  display: flex;
  width: 100%;
  justify-content: center;
}

.podium-row-top {
  margin-bottom: 1.5rem;
}

.podium-place {
  min-width: 120px;
  min-height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 1.2em;
  border-radius: 12px;
  background: #f5f5f5;
  box-shadow: 0 2px 8px rgba(30, 60, 114, 0.05);
}

.first {
  width: 400px;
  height: 400px;
  background: gold;
  color: #333;
  font-size: 1.4em;
  border: 2px solid #007bff;
  border-radius: 30%;
}

.second, .third {
  width: 400px;
  height: 400px;
  background: #e0e0e0;
  color: #333;
  margin: 0 2rem;
  border: 2px solid #007bff;
  border-radius: 30%;
}
.podium-place:hover {
  opacity: 0.7;
}
.rest {
  width: 100%;
  max-width: 34%;
  display: flex;
  flex-direction: column;
  align-items: stretch;
}

.rest-user:hover{
  opacity: 0.7;
}
.rest-user {
  background: #fff;
  border: 1px solid #007bff;
  border-radius: 24px;
  margin-bottom: 0.7rem;
  padding: 1em 1.5em;
  box-shadow: 0 1px 6px rgba(30,60,114,0.09);
  font-size: 1.1em;
  display: flex;
  align-items: center;
  width: 100%;
  min-width: 0;
  max-width: 100%;
  height: 100px;
  box-sizing: border-box;
}

.user-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 1em;
  background: #f5f5f5;
}










.rocket-fly {
  position: fixed;
  left: -100px;
  bottom: -100px;
  width: 150px;
  height: 150px;
  z-index: 1000;
  animation: rocket-move 2s cubic-bezier(0.6,0,0.4,1) forwards;
}

@keyframes rocket-move {
  to {
    left: calc(100vw + 100px);
    bottom: calc(100vh + 100px);
  }
}
</style>