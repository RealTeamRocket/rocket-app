<script setup lang="ts">
import { defineProps } from 'vue'

interface PodiumUser {
  id: number | string
  username: string
  rocket_points: number
  imageUrl?: string
}

const props = defineProps<{
  users: PodiumUser[]
  openProfile: (user: PodiumUser) => void
}>()
</script>

<template>
  <div class="podium">
    <div class="podium-row podium-row-top">
      <div class="podium-place first" @click="users[0] && openProfile(users[0])">
        <template v-if="users[0]">
          <div class="podium-user">
            <img
                class="podium-img"
                v-if="users[0]?.imageUrl"
                :src="users[0].imageUrl"
                alt="User Icon"
            />
            <img
                class="podium-img"
                v-else
                src="/src/assets/icons/rocket.svg"
                alt="Default User Icon"
            />
            <div class="podium-username">{{ users[0].username }}</div>
            <div class="podium-points">{{ users[0].rocket_points }} RP</div>
          </div>
        </template>
        <span v-else>—</span>
      </div>
    </div>
    <div class="podium-row podium-row-bottom">
      <div class="podium-place second" @click="users[1] && openProfile(users[1])">
        <template v-if="users[1]">
          <div class="podium-user">
            <img
                class="podium-img"
                v-if="users[1]?.imageUrl"
                :src="users[1].imageUrl"
                alt="User Icon"
            />
            <img
                class="podium-img"
                v-else
                src="/src/assets/icons/rocket.svg"
                alt="Default User Icon"
            />
            <div class="podium-username">{{ users[1].username }}</div>
            <div class="podium-points">{{ users[1].rocket_points }} RP</div>
          </div>
        </template>
        <span v-else>—</span>
      </div>
      <div class="podium-place third" @click="users[2] && openProfile(users[2])">
        <template v-if="users[2]">
          <div class="podium-user">
            <img
                class="podium-img"
                v-if="users[2]?.imageUrl"
                :src="users[2].imageUrl"
                alt="User Icon"
            />
            <img
                class="podium-img"
                v-else
                src="/src/assets/icons/rocket.svg"
                alt="Default User Icon"
            />
            <div class="podium-username">{{ users[2].username }}</div>
            <div class="podium-points">{{ users[2].rocket_points }} RP</div>
          </div>
        </template>
        <span v-else>—</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
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
.podium-place:hover {
  opacity: 0.8;
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
.podium-img{
  width: 120px;
  height: 120px;
  margin-bottom: 5rem;
  border-radius: 50%;
  border: 2px solid #000000;
  background: lightgray;
}
</style>