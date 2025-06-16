<script setup lang="ts">

</script>

<template>
  <img src="/src/assets/icons/rocket.svg" alt="Rocket" class="rocket-fly" />
</template>

<style scoped>
.rocket-fly {
  position: fixed;
  left: -100px;
  bottom: -100px;
  width: 150px;
  height: 150px;
  z-index: 1000;
  animation: rocket-move 2s cubic-bezier(0.6,0,0.4,1) forwards;
}

@keyframes rocket-move {
  to {
    left: calc(100vw + 100px);
    bottom: calc(100vh + 100px);
  }
}
</style>