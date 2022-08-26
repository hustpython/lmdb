<template>
  <n-config-provider :theme="theme">
    <router-view />
  </n-config-provider>
</template>

<script setup>
import { computed } from "vue";
import { storeToRefs } from "pinia";
import { useDarkTheme } from "@/store/themeData";
import { darkTheme } from "naive-ui";
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
