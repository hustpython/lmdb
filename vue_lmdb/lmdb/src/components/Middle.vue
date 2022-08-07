<template>
  <div class="imgbkg">
    <v-carousel hide-delimiters>
      <v-carousel-item
        v-for="(item, i) in items"
        :key="i"
        :src="item.src"
      ></v-carousel-item>
    </v-carousel>
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
      items: [
        {
          src: "https://cdn.vuetifyjs.com/images/carousel/squirrel.jpg",
        },
        {
          src: "https://cdn.vuetifyjs.com/images/carousel/sky.jpg",
        },
        {
          src: "https://cdn.vuetifyjs.com/images/carousel/bird.jpg",
        },
        {
          src: "https://cdn.vuetifyjs.com/images/carousel/planet.jpg",
        },
      ],
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
}
/* 图片居中显示 */
.imgbkg v-carousel-item {
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