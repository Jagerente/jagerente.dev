<template>
  <div class="container">
    <h1 class="container__header">Jagerente</h1>
    <div class="container__body">
      <button
        v-for="link in links"
        class="button"
        @click="openUrl(link.url)"
      >{{ link.title }}</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import type { Ref } from 'vue'
import axios from 'axios';

type Link = { title: string, url: string }

let links: Ref<Link[]> = ref([])

async function fetchLinks() {
  await axios.get(`/links.json`)
    .then(response => {
      links.value = response.data;
    })
    .catch(error => {
      console.log(error);
    })
}

function openUrl(url) {
  window.open(url, 'blank')
}

onMounted(() => {
  fetchLinks();
})
</script>

<style lang="scss" scoped>
.container {
  padding: 1rem;
  box-shadow: inset 0 0 5rem rgba(0, 0, 0, .5);
  box-sizing: border-box;
  height: 100vh;

  &__header {
    text-align: center;
    font-size: 3rem;
  }

  &__body {
    max-width: $lg;
    margin-left: auto;
    margin-right: auto;
    flex-direction: column;
    padding: 50px 0;
  }
}

.button {

  display: block;
  margin: 0 auto;
  background-color: white;
  padding: 0.7rem 1rem;
  border-radius: 5px;

  font-family: $mainFontFamily;
  font-weight: bold;
  font-size: 1.25rem;
  color: #333;
  box-shadow: 0 0 15px black;

  transition: box-shadow .2s ease;

  &:hover {
    box-shadow: 0 0 15px*2 black;
  }

  &+& {
    margin-top: 2rem;
  }
}
</style>