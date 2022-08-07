<template>
  <v-container>
    <Header/>
    <Middle/>
  </v-container>
</template>

<script>
import Header from "../components/Header.vue";
import Middle from '@/components/Middle.vue';
export default {
  components: {
    Header,
    Middle,
  },
};
</script>

<style>
</style>