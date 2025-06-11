<template>
  <div class="tab-bar-modern">
    <button
      v-for="tab in tabs"
      :key="tab"
      :class="{ active: tab === modelValue }"
      @click="$emit('update:modelValue', tab)"
      type="button"
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
.tab-bar-modern {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 2.5rem;
  background: #f3f8ff;
  border-radius: 2.5rem;
  box-shadow: 0 2px 12px rgba(30,60,114,0.07);
  padding: 0.4rem 0.5rem;
  width: fit-content;
  min-width: 320px;
  margin-left: auto;
  margin-right: auto;
}

.tab-bar-modern button {
  background: none;
  border: none;
  outline: none;
  padding: 0.7rem 2.2rem;
  border-radius: 2rem;
  cursor: pointer;
  font-weight: 600;
  font-size: 1.08rem;
  color: #1e3c72;
  letter-spacing: 0.02em;
  transition: 
    background 0.18s,
    color 0.18s,
    box-shadow 0.18s;
  position: relative;
}

.tab-bar-modern button.active {
  background: linear-gradient(90deg, #e0e7ff 0%, #b3c6ff 100%);
  color: #2a5298;
  box-shadow: 0 2px 8px rgba(30,60,114,0.10);
  font-weight: 700;
}

.tab-bar-modern button:not(.active):hover {
  background: #e0e7ff;
  color: #2a5298;
}

</style>
