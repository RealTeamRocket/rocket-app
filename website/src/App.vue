<template>
  <router-view />
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
