<script setup lang="ts">
import { defineProps } from 'vue'
import { getInitials, getColor } from '@/utils/userUtils'

interface PodiumUser {
  id: string
  username: string
  rocket_points: number
  imageData?: string
}

const props = defineProps<{
  users: PodiumUser[]
  openProfile: (user: PodiumUser) => void
}>()

function getPodiumUser(idx: number): PodiumUser | null {
  return props.users && props.users[idx] ? props.users[idx] : null
}
</script>

<template>
  <div class="podium-container">
    <!-- 2nd Place -->
    <div
      class="podium-slot second"
      :class="{ empty: !getPodiumUser(1) }"
      @click="getPodiumUser(1) && props.openProfile(getPodiumUser(1)!)"
    >
      <div v-if="getPodiumUser(1)" class="podium-card">
        <template v-if="getPodiumUser(1)?.imageData">
          <img
            :src="`data:image/*;base64,${getPodiumUser(1)?.imageData}`"
            class="podium-avatar"
            alt="User"
          />
        </template>
        <template v-else>
          <div
            class="podium-avatar initials-avatar"
            :style="{ backgroundColor: getColor(getPodiumUser(1)?.username || '') }"
          >
            {{ getInitials(getPodiumUser(1)?.username || '') }}
          </div>
        </template>
        <div class="podium-username">{{ getPodiumUser(1)?.username }}</div>
        <div class="podium-points">{{ getPodiumUser(1)?.rocket_points }} RP</div>
        <div class="podium-place-label">2</div>
      </div>
      <div v-else class="podium-placeholder">
        <img src="/src/assets/icons/user.svg" class="placeholder-avatar" alt="Empty" />
        <div class="podium-placeholder-label">2</div>
      </div>
    </div>
    <!-- 1st Place -->
    <div
      class="podium-slot first"
      :class="{ empty: !getPodiumUser(0) }"
      @click="getPodiumUser(0) && props.openProfile(getPodiumUser(0)!)"
    >
      <div v-if="getPodiumUser(0)" class="podium-card">
        <template v-if="getPodiumUser(0)?.imageData">
          <img
            :src="`data:image/*;base64,${getPodiumUser(0)?.imageData}`"
            class="podium-avatar"
            alt="User"
          />
        </template>
        <template v-else>
          <div
            class="podium-avatar initials-avatar"
            :style="{ backgroundColor: getColor(getPodiumUser(0)?.username || '') }"
          >
            {{ getInitials(getPodiumUser(0)?.username || '') }}
          </div>
        </template>
        <div class="podium-username">{{ getPodiumUser(0)?.username }}</div>
        <div class="podium-points">{{ getPodiumUser(0)?.rocket_points }} RP</div>
        <div class="podium-place-label">1</div>
      </div>
      <div v-else class="podium-placeholder">
        <img src="/src/assets/icons/user.svg" class="placeholder-avatar" alt="Empty" />
        <div class="podium-placeholder-label">1</div>
      </div>
    </div>
    <!-- 3rd Place -->
    <div
      class="podium-slot third"
      :class="{ empty: !getPodiumUser(2) }"
      @click="getPodiumUser(2) && props.openProfile(getPodiumUser(2)!)"
    >
      <div v-if="getPodiumUser(2)" class="podium-card">
        <template v-if="getPodiumUser(2)?.imageData">
          <img
            :src="`data:image/*;base64,${getPodiumUser(2)?.imageData}`"
            class="podium-avatar"
            alt="User"
          />
        </template>
        <template v-else>
          <div
            class="podium-avatar initials-avatar"
            :style="{ backgroundColor: getColor(getPodiumUser(2)?.username || '') }"
          >
            {{ getInitials(getPodiumUser(2)?.username || '') }}
          </div>
        </template>
        <div class="podium-username">{{ getPodiumUser(2)?.username }}</div>
        <div class="podium-points">{{ getPodiumUser(2)?.rocket_points }} RP</div>
        <div class="podium-place-label">3</div>
      </div>
      <div v-else class="podium-placeholder">
        <img src="/src/assets/icons/user.svg" class="placeholder-avatar" alt="Empty" />
        <div class="podium-placeholder-label">3</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.podium-container {
  display: flex;
  justify-content: center;
  align-items: flex-end;
  gap: 2.5rem;
  margin-bottom: 2.5rem;
  width: 100%;
  max-width: 700px;
  margin-left: auto;
  margin-right: auto;
}

.podium-slot {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-end;
  width: 140px;
  height: 200px;
  border-radius: 18px 18px 12px 12px;
  box-shadow: 0 2px 12px rgba(30,60,114,0.10);
  background: #f8fafc;
  position: relative;
  cursor: pointer;
  transition: transform 0.15s, box-shadow 0.15s;
  overflow: visible;
}
.podium-slot:hover:not(.empty) {
  transform: translateY(-8px) scale(1.04);
  box-shadow: 0 8px 32px rgba(30,60,114,0.18);
  z-index: 2;
}
.podium-slot.empty {
  cursor: default;
  opacity: 0.5;
  background: repeating-linear-gradient(135deg, #e0e0e0 0 10px, #f8fafc 10px 20px);
  box-shadow: none;
}

.podium-slot.first {
  background: linear-gradient(135deg, #fffbe6 0%, #ffe066 60%, #ffd700 100%);
  border: 2.5px solid #e6c200;
  height: 240px;
  z-index: 2;
  margin-bottom: 0;
}
.podium-slot.second {
  background: linear-gradient(135deg, #f8f8f8 0%, #e5e4e2 60%, #bfc1c2 100%);
  border: 2.5px solid #bfc1c2;
  height: 190px;
  z-index: 1;
  margin-bottom: 20px;
}
.podium-slot.third {
  background: linear-gradient(135deg, #fbeee6 0%, #c97e4e 60%, #ad6c2d 100%);
  border: 2.5px solid #ad6c2d;
  height: 170px;
  z-index: 1;
  margin-bottom: 30px;
}

.podium-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 1.2em;
  width: 100%;
  padding-top: 1.2em;
  position: relative;
}

.podium-avatar {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  margin-bottom: 0.7em;
  background: #f5f5f5;
  object-fit: cover;
  box-shadow: 0 2px 8px rgba(30,60,114,0.08);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2em;
  font-weight: 700;
  color: #fff;
  text-transform: uppercase;
  letter-spacing: 0.02em;
}
.initials-avatar {
  background: #bdbdbd;
}

.podium-username {
  font-size: 1.08em;
  font-weight: 600;
  color: #222;
  margin-bottom: 0.18em;
  text-align: center;
  word-break: break-word;
  max-width: 110px;
}

.podium-points {
  font-size: 0.98em;
  color: #2a5298;
  font-weight: 500;
  margin-bottom: 0.3em;
}

.podium-place-label {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  background: #fff;
  color: #222;
  font-size: 1.25em;
  font-weight: 700;
  border-radius: 50%;
  border: 2px solid #e3e3e3;
  width: 38px;
  height: 38px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(30,60,114,0.10);
  z-index: 3;
  bottom: -30px;
}

.podium-slot.first .podium-place-label {
  background: #ffe066;
  color: #bfa600;
  border-color: #e6c200;
}
.podium-slot.second .podium-place-label {
  background: #e5e4e2;
  color: #888;
  border-color: #bfc1c2;
}
.podium-slot.third .podium-place-label {
  background: #c97e4e;
  color: #fff;
  border-color: #ad6c2d;
}

.podium-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-end;
  width: 100%;
  height: 100%;
  padding-top: 1.2em;
  position: relative;
}
.placeholder-avatar {
  width: 54px;
  height: 54px;
  border-radius: 50%;
  opacity: 0.5;
  margin-bottom: 0.7em;
  background: #e0e0e0;
  object-fit: cover;
}
.podium-placeholder-label {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  background: #fff;
  color: #bbb;
  font-size: 1.25em;
  font-weight: 700;
  border-radius: 50%;
  border: 2px solid #e3e3e3;
  width: 38px;
  height: 38px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(30,60,114,0.10);
  z-index: 3;
  bottom: -30px;
}
</style>
