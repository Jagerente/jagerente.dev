<template>
  <Particles
    id="tsparticles"
    :particlesInit="particlesInit"
    url="/particles.json"
  />
  <div class="container">
    <Header class="container__header" />
    <router-view />
  </div>

</template>

<script setup>
import Header from "@/components/Header.vue";

import { loadFull } from "tsparticles";

const particlesInit = async engine => {
  await loadFull(engine);
};
</script>

<style lang="scss">
* {
  user-select: none;
}

body {
  font-family: $mainFontFamily;
  background-color: $background;
  color: $text-color;
  padding: 1rem;
  box-shadow: inset 0 0 2.5rem rgba(0, 0, 0, .5);
  box-sizing: border-box;
  height: 100vh;
}
</style>