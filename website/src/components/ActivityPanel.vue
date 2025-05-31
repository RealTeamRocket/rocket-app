<template>
  <div class="activity-panel">
    <h2 class="panel-title">Activity Feed</h2>
    <ul class="activity-list">
      <li v-for="(activity, idx) in activities" :key="idx" class="activity-item" :class="{ 'user-activity': activity.isUser }">
        <div class="avatar" :style="{ backgroundColor: activity.color }">
          <span>{{ activity.initials }}</span>
        </div>
        <div class="activity-content">
          <span class="name">{{ activity.name }}</span>
          <span class="desc">{{ activity.description }}</span>
          <span class="time">{{ activity.time }}</span>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Activity {
  name: string
  initials: string
  color: string
  description: string
  time: string
  isUser?: boolean
}

const activities = ref<Activity[]>([
  {
    name: 'You',
    initials: 'YO',
    color: '#2a5298',
    description: 'completed a 5km run 🚀',
    time: 'just now',
    isUser: true,
  },
  {
    name: 'Alex Kim',
    initials: 'AK',
    color: '#f39c12',
    description: 'walked 8,000 steps',
    time: '5 min ago',
  },
  {
    name: 'Samira Lee',
    initials: 'SL',
    color: '#27ae60',
    description: 'achieved a new daily record!',
    time: '20 min ago',
  },
  {
    name: 'Jordan P.',
    initials: 'JP',
    color: '#8e44ad',
    description: 'shared a cycling route',
    time: '1 hr ago',
  },
  {
    name: 'You',
    initials: 'YO',
    color: '#2a5298',
    description: 'beat your weekly goal 🎉',
    time: '2 hrs ago',
    isUser: true,
  },
])
</script>

<style scoped>
.activity-panel {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(30,60,114,0.08);
  padding: 1.5rem 1rem;
  margin-bottom: 2rem;
}

.panel-title {
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 1rem;
  color: #2a5298;
  letter-spacing: 0.5px;
}

.activity-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.activity-item {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  padding: 0.75rem 0;
  border-bottom: 1px solid #f4f8fb;
  transition: background 0.2s;
}

.activity-item:last-child {
  border-bottom: none;
}

.activity-item.user-activity {
  background: #f4f8fb;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  color: #fff;
  font-size: 1.1rem;
  flex-shrink: 0;
  box-shadow: 0 1px 4px rgba(30,60,114,0.07);
}

.activity-content {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
}

.name {
  font-weight: 600;
  color: #1e3c72;
  font-size: 1rem;
}

.desc {
  color: #4a5874;
  font-size: 0.97rem;
}

.time {
  color: #b0b8c9;
  font-size: 0.85rem;
}
</style>
