<template>
  <div class="imgbkg">
    <img :src="imgurl" alt="搜索背景图片" />
    <div class="content_wrap">
      <div class="title">
        <h2
          style="color: white; font-size: 3em; font-weight: 700; line-height: 1"
        >
          欢迎！
        </h2>
        <h3 style="color: white; font-size: 2em; font-weight: 600; margin: 0">
          这里是你本地所有的影视，快来探索与创作！
        </h3>
      </div>

      <div class="search" style="width: 100%; margin-top: 30px">
        <form
          id="inner_search_form"
          action="/search"
          method="get"
          accept-charset="utf-8"
        >
          <label>
            <input
              dir="auto"
              id="inner_search_v4"
              name="query"
              type="text"
              tabindex="1"
              autocorrect="off"
              autofill="off"
              autocomplete="off"
              spellcheck="false"
              placeholder="搜索电影、剧集、影评..."
              value=""
              style="
                color: transparent;
                width: 100%;
                height: 46px;
                line-height: 46px;
                font-size: 1.1em;
                color: rgba(0, 0, 0, 0.5);
                border: 0;
                outline: none;
                border-radius: 30px;
                padding: 10px 20px;
                background-color: white;
              "
            />
          </label>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  setup() {
    let imgurl = "";
    let videourl = "";
    return {
      imgurl: require("D://img/ys.jpg"),
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
}
.imgbkg img {
  text-align: center;
  align-items: center;
  width: 100%;
}
.midvideo video {
  text-align: center;
  margin: 0 auto;
  position: relative;
  align-items: center;
  max-height: 100%;
  max-width: 100%;
  width: 100%;
}

.content_wrap {
  width: 900px;
  text-align: left;
  position: fixed;
  align-items: center;
  height: 0 auto;
  margin-left: 100px;
  margin-top: 6px;
}
</style>