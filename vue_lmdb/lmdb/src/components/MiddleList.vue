<template>
  <canvas class="canvas_back" id="localcanvas"></canvas>
  <div v-cloak v-bind:class="{ cards: true, showing: hoverEffict.isShowing }">
    <div
      v-bind:class="{ card: true, show: index === hoverEffict.index }"
      @mouseenter="handleMouseEnter($event, index)"
      @mouseleave="handleMouseLeave($event, index)"
      v-for="(item, index) in videoData"
    >
      <img class="card-image" :src="item.Cover" alt="" />
      <img />
      <span class="card-duration">{{ item.Duration }}</span>
      <div>
        <video
          :id="'video_' + index"
          muted
          preload
          autoplay
          @loadeddata="handleLoadVideo($event, index)"
          class="card-video"
          v-show="index === hoverEffict.index"
          :src="item.TmpVideoUrl"
        ></video>
        <div class="bili-watch-later" v-show="index === hoverEffict.index">
          <svg
            class="svg-class"
            @click="handleChangeBck(index)"
            width="24"
            height="24"
            viewBox="0 0 48 48"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M5 10C5 8.89543 5.89543 8 7 8L41 8C42.1046 8 43 8.89543 43 10V38C43 39.1046 42.1046 40 41 40H7C5.89543 40 5 39.1046 5 38V10Z"
              stroke="#333"
              stroke-width="4"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M14.5 18C15.3284 18 16 17.3284 16 16.5C16 15.6716 15.3284 15 14.5 15C13.6716 15 13 15.6716 13 16.5C13 17.3284 13.6716 18 14.5 18Z"
              stroke="#333"
              stroke-width="4"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              d="M15 24L20 28L26 21L43 34V38C43 39.1046 42.1046 40 41 40H7C5.89543 40 5 39.1046 5 38V34L15 24Z"
              fill="none"
              stroke="#333"
              stroke-width="4"
              stroke-linejoin="round"
            />
          </svg>
          <span class="watch-later__tip">设为封面</span>
        </div>
      </div>

      <div class="card-title">
        <a :href="'/video?id=' + item.TmpVideoUrl" class="toggle-info btn">
          <span class="left"></span>
          <span class="right"></span>
        </a>
        <h2>
          {{ item.Title }}
          <small>INTERACTIVE HOVER EFFECT</small>
        </h2>
      </div>

      <div class="card-flap flap1">
        <div class="card-description">
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ullam,
          recusandae!
        </div>
        <div class="card-flap flap2">
          <div class="card-actions">
            <a href="" class="btn">READ MORE</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { GetVideList } from "../api/videolist";
import { reactive, onBeforeMount } from "vue";
import { timeFilter } from "../api/timefilter";

let videoData = reactive([]);

const handleChangeBck = (index) => {
  const myvideo = document.getElementById("video_" + index); // 获取视频对象
  myvideo.crossOrigin = "*";
  const mycanvas = document.getElementById("localcanvas"); // 获取 canvas 对象
  const ctx = mycanvas.getContext("2d"); // 绘制2d
  mycanvas.width = myvideo.clientWidth; // 获取视频宽度
  mycanvas.height = myvideo.clientHeight; //获取视频高度
  ctx.drawImage(myvideo, 0, 0, myvideo.clientWidth, myvideo.clientHeight);
  try {
    videoData[index].Cover = mycanvas.toDataURL("image/png"); // 导出图片
  } catch (error) {
    console.log("设置失败，稍后再试", error);
  }
};

const handleLoadVideo = (e, index) => {
  videoData[index].Duration = timeFilter(e.target.duration);
};

var hoverEffict = reactive({
  index: -1,
  isShowing: false,
});
// 加载后端数据
onBeforeMount(() => {
  GetVideList().then((res) => {
    if (res.code == 200) {
      videoData.push(...res.data);
    }
  });
});

const handleMouseEnter = (e, index) => {
  e.preventDefault();
  videoData[index].TmpVideoUrl =
    "http://localhost:9090/" + videoData[index].VideoUrl;
  // videoData[index].TmpVideoUrl = require("../file/1212.mp4");
  hoverEffict.index = index;
  hoverEffict.isShowing = true;
};

const handleMouseLeave = (e, index) => {
  videoData[index].TmpVideoUrl = "";
  e.preventDefault();
  hoverEffict.isShowing = false;
  hoverEffict.index = -1;
};
</script>

<style>
.card-duration {
  min-width: 0;
  flex: 1;
  display: -webkit-flex;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  color: #008a93;
}

.bili-watch-later {
  display: -webkit-flex;
  display: flex;
  align-items: center;
  justify-content: center;
  position: absolute;
  top: 8px;
  right: 8px;
  width: 28px;
  height: 28px;
  border-radius: 6px;
  cursor: pointer;
  z-index: 9;
  transform: translateZ(0);
}

.watch-later__tip {
  display: none;
}

.svg-class:hover + .watch-later__tip {
  pointer-events: none;
  user-select: none;
  position: absolute;
  bottom: -6px;
  right: -10px;
  transform: translateY(100%);
  font-size: 12px;
  color: #fff;
  border-radius: 4px;
  line-height: 18px;
  padding: 4px 8px;
  background-color: rgba(0, 0, 0, 0.8);
  white-space: nowrap;
  display: -webkit-flex;
}

.svg-class {
  display: -webkit-flex;
  display: flex;
  align-items: center;
  justify-content: center;
  position: absolute;
  top: 8px;
  right: 8px;
  width: 28px;
  height: 28px;
  border-radius: 6px;
  cursor: pointer;
  transform: translateZ(0);
}
.canvas_back {
  height: 300px;
  width: 300px;
  position: absolute;
  z-index: -100px;
  display: none;
}
.card-image {
  height: 300px;
  width: 300px;
  cursor: pointer;
}
.card-video {
  height: 300px;
  width: 300px;
  top: -60px;
  bottom: 0px;
  position: absolute;
  background-color: black;
}
a.btn {
  background: #0096a0;
  border-radius: 4px;
  box-shadow: 0 2px 0px 0 rgba(0, 0, 0, 0.25);
  color: #ffffff;
  display: inline-block;
  padding: 6px 30px 8px;
  position: relative;
  text-decoration: none;
  transition: all 0.1s 0s ease-out;
}
.no-touch a.btn:hover {
  background: #00a2ad;
  box-shadow: 0px 8px 2px 0 rgba(0, 0, 0, 0.075);
  transform: translateY(-2px);
  transition: all 0.25s 0s ease-out;
}
.no-touch a.btn:active,
a.btn:active {
  background: #008a93;
  box-shadow: 0 1px 0px 0 rgba(255, 255, 255, 0.25);
  transform: translate3d(0, 1px, 0);
  transition: all 0.025s 0s ease-out;
}
div.cards {
  margin: 80px auto;
  max-width: 100%;
  text-align: center;
  height: 0 auto;
}
div.card {
  background: #ffffff;
  display: inline-block;
  margin: 20px;
  max-width: 300px;
  perspective: 1000;
  position: relative;
  text-align: left;
  transition: all 0.3s 0s ease-in;
  z-index: 1;
}

div.card div.card-title {
  background: #ffffff;
  padding: 6px 15px 10px;
  position: relative;
  z-index: 0;
}
div.card div.card-title a.toggle-info {
  border-radius: 32px;
  height: 32px;
  padding: 0;
  position: absolute;
  right: 15px;
  top: 10px;
  width: 32px;
}
div.card div.card-title a.toggle-info span {
  background: #ffffff;
  display: block;
  height: 2px;
  position: absolute;
  top: 16px;
  transition: all 0.15s 0s ease-out;
  width: 12px;
}
div.card div.card-title a.toggle-info span.left {
  right: 14px;
  transform: rotate(45deg);
}
div.card div.card-title a.toggle-info span.right {
  left: 14px;
  transform: rotate(-45deg);
}
div.card div.card-title h2 {
  font-size: 24px;
  font-weight: 700;
  letter-spacing: -0.05em;
  margin: 0;
  padding: 0;
}
div.card div.card-title h2 small {
  display: block;
  font-size: 18px;
  font-weight: 600;
  letter-spacing: -0.025em;
}
div.card div.card-description {
  padding: 0 15px 10px;
  position: relative;
  font-size: 14px;
}
div.card div.card-actions {
  box-shadow: 0 2px 0px 0 rgba(0, 0, 0, 0.075);
  padding: 10px 15px 20px;
  text-align: center;
}
div.card div.card-flap {
  background: #d9d9d9;
  position: absolute;
  width: 100%;
  transform-origin: top;
  transform: rotateX(-90deg);
}
div.card div.flap1 {
  transition: all 0.3s 0.3s ease-out;
  z-index: -1;
}
div.card div.flap2 {
  transition: all 0.3s 0s ease-out;
  z-index: -2;
}
div.cards.showing div.card {
  cursor: pointer;
  opacity: 0.6;
  transform: scale(0.88);
}
.no-touch div.cards.showing div.card:hover {
  opacity: 0.94;
  transform: scale(0.92);
}
div.card.show {
  opacity: 1 !important;
  transform: scale(1) !important;
  z-index: 10;
}
div.card.show div.card-title a.toggle-info {
  background: #ff6666 !important;
}
div.card.show div.card-title a.toggle-info span {
  top: 15px;
}
div.card.show div.card-title a.toggle-info span.left {
  right: 10px;
}
div.card.show div.card-title a.toggle-info span.right {
  left: 10px;
}
div.card.show div.card-flap {
  background: #ffffff;
  transform: rotateX(0deg);
}
div.card.show div.flap1 {
  transition: all 0.3s 0s ease-out;
}
div.card.show div.flap2 {
  transition: all 0.3s 0.2s ease-out;
}

[v-cloak] {
  display: none;
}
</style>