<template>
  <div class="search-bar">
    <input
      type="text"
      v-model="search"
      placeholder="🔍 Search new friends..."
      @input="$emit('update:search', search)"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
const search = ref('');
</script>

<style scoped>
.search-bar {
  width: 100%;
  margin-bottom: 1.5rem;
  display: flex;
  justify-content: center;
}
input {
  width: 100%;
  max-width: 500px;
  padding: 0.7rem 1.2rem;
  border-radius: 1rem;
  border: 1px solid #c7d2fe;
  font-size: 1.1rem;
  background: #f3f8ff;
  outline: none;
  transition: border 0.2s;
}
input:focus {
  border: 1.5px solid #4f46e5;
}
</style>