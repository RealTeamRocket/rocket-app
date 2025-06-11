<template>
  <router-view :key="$route.fullPath" />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useAuth } from './utils/useAuth';

const { checkAuth } = useAuth();

onMounted(() => {
  checkAuth()
})
</script>

<style scoped>
</style>
