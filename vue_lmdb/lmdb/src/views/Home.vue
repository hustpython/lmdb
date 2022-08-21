<template>
  <Header />

  <MiddleSearch />

  <n-notification-provider>
    <MiddleList />
  </n-notification-provider>
</template>

<script>
import Header from "@/components/Header.vue";
import MiddleSearch from "@/components/MiddleSearch.vue";
import MiddleList from "@/components/MiddleList";
import { NNotificationProvider } from "naive-ui";
export default {
  components: {
    Header,
    MiddleSearch,
    MiddleList,
    NNotificationProvider,
  },
};
</script>

<style>
</style>