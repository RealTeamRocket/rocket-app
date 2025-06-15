<script setup lang="ts">
import type { RankedUser } from '@/types' // Passe den Import ggf. an

defineProps<{
  users: RankedUser[]
  openProfile: (user: RankedUser) => void
  addFriend: (username: string) => void
}>()
</script>
<template>
  <div class="rest">
    <div
        class="rest-user"
        v-for="(user, idx) in users"
        :key="user.id || idx"
        @click="openProfile(user)"
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
            style="color: lightgray"
        />
      </div>
      <div class="rest-username">
        {{ user.username }}
      </div>
      <div class="rest-rocketpoints">
        {{ user.rocket_points }}
        <img
            src="/src/assets/icons/rocket.svg"
            alt="Rocket"
            style="width:1.3em;height:1.3em;vertical-align:middle;margin-right:0.35em;"
        />
      </div>
      <div class="rest-action">
        <template v-if="user.isFriend">
          <img src="/src/assets/icons/user.svg" alt="Friend Icon" class="friend-icon" />
        </template>
        <template v-else>
          <button class="add-btn" @click.stop="addFriend(user.username)">Add</button>
        </template>
      </div>
    </div>
  </div>
</template>

<style scoped>
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
.rest-user:hover {
  opacity: 0.8;
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
.user-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 1em;
  background: #f5f5f5;
}
.rest-username,
.rest-rocketpoints,
.rest-action {
  min-width: 100px;
}
.rest-rocketpoints {
  margin-left: 4rem;
  font-weight: bold;
}
.rest-username {
  margin-left: 2rem;
  font-weight: bold;
}
.rest-action {
  width: 56px;
  text-align: center;
  justify-content: center;
  padding-left: 0;
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
.friend-icon {
  width: 48px;
  height: 48px;
  filter: invert(41%) sepia(98%) saturate(1200%) hue-rotate(74deg) brightness(110%) contrast(120%);
}
</style>