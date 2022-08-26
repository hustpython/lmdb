<template>
  <Header />
  <video
    id="localvideo"
    autoplay="true"
    :src="videourl"
    controls="controls"
    @timeupdate="handleTimeUpdate"
    @loadeddata="handleLoadedata"
  ></video>
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
const route = useRoute();
const videoDurtion = reactive({
  timeStr: "",
});
const videourl = route.query.id;
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
video {
  position: relative;
  height: 500px;
  max-height: 100%;
  max-width: 100%;
  width: 900px;
  display: block;
  margin: auto;
  margin-top: 60px;
  border: 0;
}
</style>
