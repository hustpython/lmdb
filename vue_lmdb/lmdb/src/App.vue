<template>
  <n-config-provider :theme="theme">
    <n-layout
      v-bind:class="{
        darkThemeBck: themeData === true,
        lightThemeBck: themeData === false,
      }"
    >
      <n-notification-provider>
        <router-view />
      </n-notification-provider>
    </n-layout>
  </n-config-provider>
</template>

<script setup>
import { computed } from "vue";
import { storeToRefs } from "pinia";
import { darkTheme } from "naive-ui";
import { useDarkTheme } from "@/store/themeData";

const darkThemeStore = useDarkTheme();
var { themeData } = storeToRefs(darkThemeStore);

const theme = computed(() => {
  if (themeData.value === true) {
    return darkTheme;
  } else {
    return null;
  }
});
</script>
