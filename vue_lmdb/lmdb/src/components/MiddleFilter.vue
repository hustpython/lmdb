<template>
  <div class="videoTabs">
    <n-tabs type="bar" animated v-model:value="tabVal">
      <n-tab name="1"> 最近 </n-tab>
      <n-tab name="2"> 全部 </n-tab>
      <n-tab name="3"> 刷新 </n-tab>
    </n-tabs>
  </div>
  <n-space vertical v-show="tabVal.value === 'asas'">
    <n-select v-model:value="value" multiple :options="options" />
  </n-space>
</template>

<script setup>
import { NSpace, NSelect, NTabs, NTab } from "naive-ui";
import { ref } from "vue";
const value = ref(["song3"]);
const options = [
  {
    label: "Everybody's Got Something to Hide Except Me and My Monkey",
    value: "song0",
  },
  {
    label: "Drive My Car",
    value: "song1",
  },
];
var tabVal = ref("1");
</script>

<style></style>
