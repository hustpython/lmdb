<template>
  <video :src="videourl" controls="controls" poster="../file/one.jpg">
  </video>
</template>

<script setup>
import { useRoute } from "vue-router";
const route = useRoute();
const videourl = "http://localhost:9091/"+route.query.id;
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