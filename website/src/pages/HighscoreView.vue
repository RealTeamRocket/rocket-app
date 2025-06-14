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
            <template v-if="selectedList[0]">
              <div class="podium-user">
                <img
                    class="podium-img"
                    v-if="selectedList[0]?.imageUrl"
                    :src="selectedList[0].imageUrl"
                    alt="User Icon"
                />
                <img
                    class="podium-img"
                    v-else
                    src="/src/assets/icons/rocket.svg"
                    alt="Default User Icon"
                />
                <div class="podium-username">{{ selectedList[0].username }}</div>
                <div class="podium-points">{{ selectedList[0].rocket_points }} RP</div>
              </div>
            </template>
            <span v-else>—</span>
          </div>
        </div>
        <div class="podium-row podium-row-bottom">
          <div class="podium-place second">
            <template v-if="selectedList[1]">
              <div class="podium-user">
                <img
                    class="podium-img"
                    v-if="selectedList[1]?.imageUrl"
                    :src="selectedList[1].imageUrl"
                    alt="User Icon"
                />
                <img
                    class="podium-img"
                    v-else
                    src="/src/assets/icons/rocket.svg"
                    alt="Default User Icon"
                />
                <div class="podium-username">{{ selectedList[1].username }}</div>
                <div class="podium-point">{{ selectedList[1].rocket_points }} RP</div>
              </div>
            </template>
            <span v-else>—</span>
          </div>
          <div class="podium-place third">
            <template v-if="selectedList[2]">
              <div class="podium-user">
                <img
                    class="podium-img"
                    v-if="selectedList[2]?.imageUrl"
                    :src="selectedList[2].imageUrl"
                    alt="User Icon"
                />
                <img
                    class="podium-img"
                    v-else
                    src="/src/assets/icons/rocket.svg"
                    alt="Default User Icon"
                />
                <div class="podium-username">{{ selectedList[2].username }}</div>
                <div class="podium-point">{{ selectedList[2].rocket_points }} RP</div>
              </div>
            </template>
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
          <div class="rest-img">
            <img
                v-if="user.imageUrl"
                :src="user.imageUrl"
                alt="User Icon"
                class="user-avatar"
            />
            <img
                v-else
                src="/src/assets/icons/user.svg"
                alt="Default User Icon"
                class="user-avatar"
                style="color: lightgray"/>
          </div>
          <div class="rest-username">
            {{ user.username }}
          </div>
          <div class="rest-rocketpoints">
            {{ user.rocket_points }}  <img src="/src/assets/icons/rocket.svg" alt="Rocket"
                 style="width:1.3em;height:1.3em;vertical-align:middle;margin-right:0.35em;"/>
          </div>
          <div class="rest-action">
            <template v-if="user.isFriend">
              <img src="/src/assets/icons/user.svg" alt="Friend Icon" class="friend-icon" />
            </template>
            <template v-else><button class="add-btn" @click="addFriend(user.username)">Add</button>
            </template>
          </div>
        </div>
      </div>
  </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
import {ref, onMounted, computed} from 'vue'
import api from '@/api/backend-api'

interface RankedUser {
  id: number | string
  username: string
  rocket_points: number
  isFriend?: boolean
  imageUrl?: string
}

const currentUsername = ref('')
const isGlobal = ref(true)
const rankedUsers = ref<RankedUser[]>([])
const rankedFriends = ref<RankedUser[]>([])

async function loadUserImages(list: RankedUser[]) {
  for (const user of list) {
    try {
      const res = await api.getUserImage(user.id)
      // Annahme: Das Bild kommt als Blob zurück
      const blob = res.data
      user.imageUrl = URL.createObjectURL(blob)
    } catch {
      user.imageUrl = '' // Kein Bild vorhanden
    }
  }
}

const loadRanking = async () => {
  try {
    const {data: userData} = await api.getRankedUsers()
    const {data: friendData} = await api.getRankedFriends()
    const friendUsernames = new Set(friendData.map((f: RankedUser) => f.username))
    rankedUsers.value = (userData || []).map((user: RankedUser) => ({
      ...user,
      isFriend: friendUsernames.has(user.username)
    }))
    rankedFriends.value = (friendData || []).map((user: RankedUser) => ({
      ...user,
      isFriend: true
    }))
    // Bilder laden
    await loadUserImages(rankedUsers.value)
    await loadUserImages(rankedFriends.value)
  } catch {
    console.error("Error fetching ranking")
  }
}

const selectedList = computed(() => isGlobal.value ? rankedUsers.value : rankedFriends.value)

onMounted(async () => {
  await loadRanking()
})

function addFriend(friendName: string) {
  api.addFriend(friendName)
      .then(() => {
        alert('Freund hinzugefügt!')
        loadRanking()
      })
      .catch((err) => {
        alert('Fehler beim Hinzufügen: ' + (err.response?.data?.message || err.message))
      })
}




</script>

<style scoped>
.highscore-page {
  padding: 2rem;
}
.rest-rocketpoints{
  margin-left: 4rem;
  font-weight: bold;
}
.rest-username{
  margin-left: 2rem;
  font-weight: bold;
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
.podium-user {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-end;
  height: 100%;
  padding-bottom: 2em;
  box-sizing: border-box;
}

.podium-username {
  font-size: 1.3em;
  font-weight: bold;
  color: #333;
}
.podium-points {
  font-size: 1em;
  color: #2a5298;
  margin-top: 0.2em;
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
  background: linear-gradient(135deg, #fffbe6 0%, #ffe066 60%, #ffd700 100%);
  color: #333;
  font-size: 1.4em;
  border: 3px solid #2196f3;
  border-radius: 30%;
}
.second {
  width: 400px;
  height: 400px;
  background: linear-gradient(135deg, #f8f8f8 0%, #e5e4e2 60%, #bfc1c2 100%);
  color: #333;
  margin: 0 2rem;
  border: 3px solid #2196f3;
  border-radius: 30%;
}

.third {
  width: 400px;
  height: 400px;
  background: linear-gradient(135deg, #fbeee6 0%, #c97e4e 60%, #ad6c2d 100%);
  color: #333;
  margin: 0 2rem;
  border: 3px solid #2196f3;
  border-radius: 30%;
}
.friend-icon {
  width: 48px;
  height: 48px;
  filter: invert(41%) sepia(98%) saturate(1200%) hue-rotate(74deg) brightness(110%) contrast(120%);
}
.rest {
  width: 100%;
  max-width: 34%;
  margin-top: 2rem;
}

.rest-user {
  display: flex;
  align-items: center;
  width: 100%;
  background: linear-gradient(#f6fcff 0%, #eafeea 60%, #e6f7fa 100%);
  border-radius: 24px;
  border: 2px solid #e3f0fb;
  box-shadow: 0 1px 6px rgba(30,60,114,0.06);
  margin-bottom: 1.5rem;
  height: 72px;
  overflow: hidden;
  gap: 12%;
}

.rest-img,
.rest-username,
.rest-rocketpoints,
.rest-action {
  display: flex;
  align-items: center;
  height: 100%;
  padding: 0.7em 1.2em;
}
.add-btn {
  background: linear-gradient(90deg, #2196f3 0%, #00bcd4 100%);
  color: #fff;
  border: none;
  border-radius: 24px;
  padding: 0.5em 1.5em;
  font-weight: bold;
  font-size: 1em;
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(33, 150, 243, 0.18);
  transition: background 0.2s, transform 0.1s;
}
.add-btn:hover {
  background: linear-gradient(90deg, #1976d2 0%, #0097a7 100%);
  transform: translateY(-2px) scale(1.04);
}
.rest-username,
.rest-rocketpoints,
.rest-action {
  min-width: 100px;
}

.rest-action {
  width: 56px;
  text-align: center;
  justify-content: center;
  padding-left: 0; /* Kein extra Padding mehr */
}

.action-icon {
  width: 36px;
  height: 36px;
  cursor: pointer;
}

.user-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 1em;
  background: #f5f5f5;
}

.podium-img{
  width: 120px;
  height: 120px;
  margin-bottom: 5rem;
  border-radius: 50%;
  border: 2px solid #000000;
  background: lightgray;
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