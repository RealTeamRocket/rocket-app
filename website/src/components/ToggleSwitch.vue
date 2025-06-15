<script setup lang="ts">
  defineProps<{ modelValue: boolean }>()
  defineEmits(['update:modelValue'])
</script>

<template>
  <div class="switchBtn">
    <button
        :class="{active: modelValue}"
        @click="$emit('update:modelValue', true)"
    >Global</button>
    <button
        :class="{active: !modelValue}"
        @click="$emit('update:modelValue', false)"
    >Friends only</button>
  </div>
</template>

<style scoped>
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
</style>