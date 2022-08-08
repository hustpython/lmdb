<template>
  <Header />
  <MiddleSearch />
  <MiddleList/>
</template>

<script>
import Header from "../components/Header.vue";
import MiddleSearch from "@/components/MiddleSearch.vue";
import MiddleList from '@/components/MiddleList';
export default {
  components: {
    Header,
    MiddleSearch,
    MiddleList,
  },
};
</script>

<style>
</style>