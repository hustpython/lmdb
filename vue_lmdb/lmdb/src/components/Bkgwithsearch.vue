<template>
  <div class="imgbkg">
    <img :src="imgurl" alt="背景图片" />
  </div>
  <video :src="videourl" controls="controls">></video>
</template>

<script>
export default {
  setup() {
    let imgurl = "";
    let videourl = "";
    return {
      imgurl: require("D://img/1 (5).jpg"),
      videourl: require("D:\\迅雷下载\\黑镜.Black.Mirror.1-4季全集【关注微信公众号：美剧叔】\\黑镜 第4季\\【关注微信公众号：美剧叔】黑镜.Black.Mirror.S04E03.mp4"),
    };
  },
};
</script>

<style>
.imgbkg {
  width: 100%;
  display: flex;
  height: 0 auto;
  position: relative;
  margin-top: var(--headerTop);
  background-color: var(--lightBule);
}
/* 图片居中显示 */
.imgbkg img {
  text-align: center;
  display: flex;
  margin: 0 auto;
  align-items: center;
  /* height: 400px; */
  /* width: 100%; */
  /* filter: hue-rotate(90) */

}

.imgbkg video {
  text-align: center;
  margin: 0 auto;
  align-items: center;
  max-height: 100%;
  max-width: 100%;
  width: 100%;
}
</style>