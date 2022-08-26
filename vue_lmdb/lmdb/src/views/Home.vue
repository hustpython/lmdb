<template>
  <Header />
  <n-layout>
    <MiddleSearch />
    <n-notification-provider>
      <MiddleList />
    </n-notification-provider>
  </n-layout>
</template>

<script setup>
import Header from "@/components/Header.vue";
import MiddleSearch from "@/components/MiddleSearch.vue";
import MiddleList from "@/components/MiddleList";
import { NNotificationProvider } from "naive-ui";
</script>
