<template>
  <video :src="videourl" controls="controls"></video>
</template>

<script setup>
import { useRoute } from "vue-router";
const route = useRoute();
const videourl = "http://localhost:9091/" + route.query.id;
</script>
<style>
video {
  position: relative;
  height: 0 auto;
  max-height: 100%;
  max-width: 100%;
  width: 1000px;
  display: block;
  margin: auto;
  margin-top: 20px;
}
</style>