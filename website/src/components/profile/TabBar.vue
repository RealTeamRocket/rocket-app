<template>
  <div class="tab-bar">
    <button
      v-for="tab in tabs"
      :key="tab"
      :class="{ active: tab === modelValue }"
      @click="$emit('update:modelValue', tab)"
    >
      {{ tab }}
    </button>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  tabs: string[]
  modelValue: string
}>()
defineEmits(['update:modelValue'])
</script>

<style scoped>
.tab-bar {
  display: flex;
  gap: 1rem;
  margin-bottom: 2rem;
}
button {
  background: none;
  border: none;
  padding: 0.5rem 1rem;
  cursor: pointer;
  font-weight: bold;
  font-size: 1rem;
  transition: color 0.2s, border-bottom 0.2s;
}
button.active {
  border-bottom: 2px solid #007bff;
  color: #007bff;
}
</style>
