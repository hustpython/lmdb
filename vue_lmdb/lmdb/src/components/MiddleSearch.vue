<template>
  <div class="imgbkg">
    <div class="content_wrap">
      <div class="title">
        <h2 style="color: white; font-size: 2.2em; font-weight: 600; line-height: 2">
          欢迎！
        </h2>
        <h3 style="color: white; font-size: 1.8em; font-weight: 500; margin: 0">
          这里是你本地所有的影视，快来探索与创作！
        </h3>
      </div>

      <div class="search" style="width: 100%; margin-top: 30px">
        <n-select
          v-bind:class="{
            darkThemeBck: themeData === true,
            lightThemeBck: themeData === false,
          }"
          style="
            width: 80%;
            height: 32px;
            line-height: 30px;
            font-size: 1em;
            border: 0;
            outline: none;
            border-radius: 30px;
            padding: 10px 20px;
          "
          filterable
          :options="options"
          :on-update:value="handleSeleteed"
          placeholder="搜索电影、剧集、影评..."
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { storeToRefs } from "pinia";
import { useVideoData } from "@/store/videoData";
import { useRouter } from "vue-router";
import { computed } from "vue";
import { useDarkTheme } from "@/store/themeData";

const imgurl = require("../assets/backimg.jpg");

const videoDataStore = useVideoData();
var { videoData } = storeToRefs(videoDataStore);

const darkThemeStore = useDarkTheme();
var { themeData } = storeToRefs(darkThemeStore);

const options = computed(() => {
  let tmpOptions = [];
  for (var i = 0; i < videoData.value.length; i++) {
    tmpOptions.push({
      label: videoData.value[i].Title,
      value: i,
    });
  }
  return tmpOptions;
});

const router = useRouter();
const handleSeleteed = (value) => {
  let routeData = router.resolve({ path: "video", query: { id: value } });
  window.open(routeData.href, "_blank");
};
</script>

<style>
.imgbkg {
  width: 100%;
  margin-top: var(--headerTop);
  overflow: hidden;
  height: var(--homeBckHeight);
  background-image: v-bind("'url(' + imgurl + ')'");
  background-size: 100%;
}

.content_wrap {
  width: 680px;
  text-align: left;
  align-items: center;
  margin-left: 100px;
  margin-top: 20px;
  position: absolute;
}
</style>
