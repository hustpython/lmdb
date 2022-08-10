<template>
  <video :src="videourl" controls="controls"></video>
</template>

<script setup>
import { useRoute } from "vue-router";
import {ref} from 'vue';
const route = useRoute();
const ls = ref("")
ls.value = route.query.id;
const videourl = require([ls.value]);
</script>
<style>
video {
  text-align: center;
  margin: 0 auto;
  position: relative;
  align-items: center;
  height: 0 auto;
  max-height: 100%;
  max-width: 100%;
}
</style>