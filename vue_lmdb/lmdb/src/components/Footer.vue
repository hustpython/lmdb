<template>
  <n-layout-footer
    class="footer"
    v-bind:class="{
      darkThemeBck: themeData === true,
      lightThemeBck: themeData === false,
    }"
    bordered
  >
    版权所有 mxq</n-layout-footer
  >
</template>

<script setup>
import { useDarkTheme } from "@/store/themeData";
import { storeToRefs } from "pinia";
const darkThemeStore = useDarkTheme();
var { themeData } = storeToRefs(darkThemeStore);
</script>

<style>
.footer {
  text-align: center;
}
</style>
