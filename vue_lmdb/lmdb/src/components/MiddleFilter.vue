<template>
  <n-space class="videoTabs" justify="space-between">
    <n-tabs type="bar" animated v-model:value="tabVal">
      <n-tab name="1"> 最近 </n-tab>
      <n-tab name="2"> 全部 </n-tab>
      <n-tab name="3"> 刷新 </n-tab>
    </n-tabs>
    <n-pagination
      v-model:page="page"
      :page-count="pageCount"
      simple
      @update:page="handleUpdatePage"
    />
  </n-space>
</template>

<script setup>
import { NSpace, NSelect, NTabs, NTab, NPagination } from "naive-ui";
import { ref, defineEmits, computed } from "vue";
import { useVideoData } from "@/store/videoData";
import { storeToRefs } from "pinia";
const videoDataStore = useVideoData();
var tabVal = ref("1");
var { videoData } = storeToRefs(videoDataStore);
const emit = defineEmits(["videoPageChange"]);
const page = ref(1);
const pageWidth = ref(document.body.clientWidth);

const handleUpdatePage = (page) => {
  emit("videoPageChange", page);
};

const pageCount = computed(() => {
  handleUpdatePage(page.value);
  return Math.ceil(videoData.value.length / Math.floor(pageWidth.value / 400));
});

window.onresize = () => {
  pageWidth.value = document.body.clientWidth;
};
</script>

<style>
.videoTabs {
  padding: 7px 17px 7px 17px;
}
</style>