<template>
  <div>
    <Header />
    <Bkg />
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Bkg from "./components/Bkgwithsearch.vue";

export default {
  name: "App",
  components: {
    Header,
    Bkg,
  },
};
</script>

<style>
</style>
