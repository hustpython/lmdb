<template>
  <Header />
  <n-layout
    class="videopage"
    v-bind:class="{
      darkThemeBck: themeData === true,
      lightThemeBck: themeData === false,
    }"
  >
    <n-spance class="videospace">
      <div style="font-size: 18px">{{ title }}</div>
      <video
        id="localvideo"
        autoplay="true"
        :src="videourl"
        controls="controls"
        @timeupdate="handleTimeUpdate"
        @loadeddata="handleLoadedata"
      ></video>
    </n-spance>
  </n-layout>
  <!-- <canvas id="localcanvas"></canvas>
  <button id="snapbtn" @click="handleClick">
    截取总时长{{ videoDurtion.timeStr }}
  </button> -->
</template>

<script setup>
import { useRoute } from "vue-router";
import { reactive } from "vue";
import { timeFilter } from "../api/timefilter";
import Header from "@/components/Header.vue";

import { storeToRefs } from "pinia";
import { darkTheme } from "naive-ui";
import { useDarkTheme } from "@/store/themeData";

import { useVideoData } from "@/store/videoData";

const videoDataStore = useVideoData();
var { videoData } = storeToRefs(videoDataStore);

const darkThemeStore = useDarkTheme();
var { themeData } = storeToRefs(darkThemeStore);

const route = useRoute();
const videoDurtion = reactive({
  timeStr: "",
});
const title = videoData.value[route.query.id].Title;
const videourl = "http://localhost:9090/" + videoData.value[route.query.id].VideoUrl;
const handleTimeUpdate = (ev) => {};

const handleLoadedata = (ev) => {
  videoDurtion.timeStr = timeFilter(ev.target.duration);
};

const handleClick = () => {
  updateCanvas();
};

const updateCanvas = () => {
  //console.log(ev,"timeupdate")
  const myvideo = document.getElementById("localvideo"); // 获取视频对象
  const mycanvas = document.getElementById("localcanvas"); // 获取 canvas 对象
  //console.log(ev.target.duration) // 总时长

  const ctx = mycanvas.getContext("2d"); // 绘制2d
  myvideo.crossOrigin = "anonymous"; // 解决跨域问题，也就是提示污染资源无法转换视频
  // myvideo.currentTime = 100; // 第一帧

  mycanvas.width = myvideo.clientWidth; // 获取视频宽度
  mycanvas.height = myvideo.clientHeight; //获取视频高度
  // 利用canvas对象方法绘图
  ctx.drawImage(myvideo, 0, 0, myvideo.clientWidth, myvideo.clientHeight);
  // mycanvas.toDataURL("image/png");
};
</script>
<style>
.videopage {
  margin-top: var(--headerTop);
}
.videospace {
  float: left;
  margin-top: 30px;
  margin-left: 60px;
}
video {
  position: relative;
  height: 376px;
  width: 668x;
  display: block;
  margin: 0;
  padding: 0;
  margin-top: 15px;
  border: 0;
}
</style>
