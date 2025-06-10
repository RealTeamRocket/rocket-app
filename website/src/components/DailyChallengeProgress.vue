<template>
  <div class="progressbar-container">
    <div class="progressbar-label">
      Daily Challenge Progress: {{ props.completed }}/{{ props.total }}
    </div>
    <div class="progressbar-bg">
      <div class="progressbar-fill" :style="{ width: percent + '%' }"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{ completed: number, total: number }>();
const percent = computed(() => props.total > 0 ? Math.round((props.completed / props.total) * 100) : 0);
</script>

<style scoped>
.progressbar-container {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  background: #fff;
  padding: 0.5rem 1.5rem 1rem 1.5rem;
  box-shadow: 0 -2px 8px rgba(30,60,114,0.08);
  z-index: 100;
}
.progressbar-label {
  font-size: 1rem;
  margin-bottom: 0.3rem;
  color: #1e3c72;
  text-align: center;
}
.progressbar-bg {
  width: 100%;
  height: 16px;
  background: #e0e7ff;
  border-radius: 8px;
  overflow: hidden;
}
.progressbar-fill {
  height: 100%;
  background: linear-gradient(90deg, #4f46e5 0%, #22c55e 100%);
  transition: width 0.3s;
}
</style>