<template>
  <canvas id="localcanvas" style="display: none"></canvas>
  <div v-cloak v-bind:class="{ cards: true, showing: hoverEffict.isShowing }">
    <div
      v-bind:class="{ card: true, show: index === hoverEffict.index }"
      @mouseenter="handleMouseEnter($event, index)"
      @mouseleave="handleMouseLeave($event, index)"
      v-for="(item, index) in videoData"
    >
      <div class="card-image">
        <img :src="item.Cover" alt="" />
        <span class="card-duration">{{ item.Duration }}</span>
      </div>

      <div class="video-mask" v-show="index === hoverEffict.index">
        <video
          :id="'video_' + index"
          muted
          preload
          autoplay
          @loadeddata="handleLoadVideo($event, index)"
          @mousemove="handleProgress($event)"
          @click="handleVideoClick"
          :src="item.TmpVideoUrl"
        ></video>
        <div class="set-cover">
          <img
            class="set-cover-svg"
            @click="handleChangeBck(index)"
            src="../assets/yl.svg"
          />
          <span class="set-cover-tip">设为封面</span>
        </div>
        <div class="video-duration-progress">
          <span :style="'width:' + hoverEffict.progress + '%'"></span>
        </div>
      </div>

      <div class="card-title">
        <a :href="'/video?id=' + item.TmpVideoUrl" class="toggle-info">
          <img src="../assets/openmovie.svg" />
        </a>
        <small>{{ item.Title }}</small>
      </div>

      <div class="card-flap flap1">
        <div class="card-description">还是挺好看的哦</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { GetVideList } from "@/api/videolist";
import { reactive, onMounted } from "vue";
import { timeFilter } from "@/api/timefilter";
import { useVideoData } from "@/store/videoData";
import { storeToRefs } from "pinia";

const videoDataStore = useVideoData();
var { videoData } = storeToRefs(videoDataStore);

const handleChangeBck = (index) => {
  const myvideo = document.getElementById("video_" + index); // 获取视频对象
  myvideo.crossOrigin = "*";
  const mycanvas = document.getElementById("localcanvas"); // 获取 canvas 对象
  const ctx = mycanvas.getContext("2d"); // 绘制2d
  mycanvas.width = myvideo.clientWidth; // 获取视频宽度
  mycanvas.height = myvideo.clientHeight; //获取视频高度
  ctx.drawImage(myvideo, 0, 0, mycanvas.width, mycanvas.height);
  try {
    videoData.value[index].Cover = mycanvas.toDataURL("image/png"); // 导出图片
  } catch (error) {
    console.log("设置失败，稍后再试", error);
    videoData.value[index].TmpVideoUrl = "";
    setTimeout(function () {
      videoData.value[index].TmpVideoUrl =
        "http://localhost:9090/" + videoData.value[index].VideoUrl;
      hoverEffict.progress = 0;
    }, 1);
  }
};

const handleLoadVideo = (e, index) => {
  videoData.value[index].Duration = timeFilter(e.target.duration);
};

var hoverEffict = reactive({
  index: -1,
  isShowing: false,
  progress: 0,
});
// 加载后端数据
onMounted(() => {
  if (videoData.value.length >= 1) {
    return;
  }
  GetVideList().then((res) => {
    if (res.code == 200) {
      videoDataStore.setVideoData(res.data);
    }
  });
});

const handleMouseEnter = (e, index) => {
  e.preventDefault();
  videoData.value[index].TmpVideoUrl =
    "http://localhost:9090/" + videoData.value[index].VideoUrl;
  // videoData[index].TmpVideoUrl = require("../file/1212.mp4");
  hoverEffict.index = index;
  hoverEffict.isShowing = true;
};

const handleMouseLeave = (e, index) => {
  videoData.value[index].TmpVideoUrl = "";
  e.preventDefault();
  hoverEffict.isShowing = false;
  hoverEffict.index = -1;
  hoverEffict.progress = 0;
};

const handleVideoClick = (e) => {
  if (e.target.paused) {
    e.target.play();
  } else {
    e.target.pause();
  }
};

const handleProgress = (e) => {
  if (e.target.paused) {
    return;
  }
  hoverEffict.progress = (e.offsetX / e.target.clientWidth) * 100;
  e.target.currentTime = (e.target.duration * hoverEffict.progress) / 100;
};
</script>

<style>
div.cards {
  margin: 10px auto;
  max-width: 100%;
  text-align: center;
  height: 0 auto;
}
div.card {
  background: #ffffff;
  display: inline-block;
  margin: 10px;
  max-width: 380px;
  perspective: 1000;
  position: relative;
  text-align: left;
  transition: all 0.3s 0s ease-in;
  z-index: 1;
}

/* 封面相关 */

.card-image {
  height: 214px;
  width: 380px;
  cursor: pointer;
}

.card-image img {
  height: 214px;
  width: 380px;
}

.card-image .card-duration {
  position: absolute;
  bottom: 42px;
  right: 10px;
  height: 20px;
  line-height: 20px;
  transition: opacity 0.3s;
  z-index: 5;
  font-size: 13px;
  background-color: rgba(43, 104, 165, 0.4);
  color: #fff;
  border-radius: 2px;
  z-index: inherit;
}

/* 视频相关 */
.video-mask {
  height: 214px;
  width: 380px;
  top: 0;
  position: absolute;
  background-color: #000;
}

.video-mask video {
  margin: 0;
}

video {
  height: 214px;
  width: 380px;
}

.set-cover {
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

.set-cover-tip {
  display: none;
}

.set-cover-svg:hover + .set-cover-tip {
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

.set-cover-svg {
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

.video-duration-progress {
  position: absolute;
  bottom: 0px;
  width: 100%;
  height: 10px;
  border-color: rgb(0, 0, 0);
  border-style: solid;
  border-width: 4px 8px;
  background: #444;
  box-sizing: border-box;
}

.video-duration-progress span {
  display: block;
  background: #fff;
  height: 2px;
  transition: width 0.12s;
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
  top: 4px;
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

div.card div.card-title small {
  display: block;
  font-size: 15px;
  font-weight: 400;
  letter-spacing: -0.025em;
  color: cadetblue;
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

div.card.show div.card-title a.toggle-info span {
  top: 15px;
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