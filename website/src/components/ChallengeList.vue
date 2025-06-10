<template>
  <div>
    <h2>Active Challenges</h2>
    <ul>
      <li v-for="challenge in challenges" :key="challenge.id">
        <strong>{{ challenge.text }}</strong> – {{ challenge.points }} Punkte
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
defineProps<{ challenges: Array<{ id: string, text: string, points: number }> }>();
</script>