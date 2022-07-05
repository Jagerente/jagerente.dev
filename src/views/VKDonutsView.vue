<template>
  <div class="container">
    <h1 class="container__header">VK Donuts</h1>
    <div class="container__body">
      <p
        v-for="donut in donuts"
        class="donut"
      >{{ donut }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import type { Ref } from 'vue'
import axios from 'axios';
import { useRoute } from 'vue-router'
const route = useRoute();

type GetMembers = {
  access_token,
  group_id,
}

let donuts: Ref<string[]> = ref([])


async function fetchLinks() {
  const access_token = route.query.access_token;
  const group_id = route.query.group_id;

  if (!access_token || !group_id)
    return;

  const getMembers: GetMembers = {
    access_token: access_token,
    group_id: group_id,
  }

  await axios.get(`https://jagerente.xyz/api/getDons`, { params: getMembers })
    .then(response => {
      donuts.value = response.data;
    })
    .catch(error => {
      console.log(error);
    })
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