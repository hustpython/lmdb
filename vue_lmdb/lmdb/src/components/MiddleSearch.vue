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
        <label>
          <input
            dir="auto"
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
              width: 80%;
              height: 30px;
              line-height: 30px;
              font-size: 1em;
              color: rgba(0, 0, 0, 0.5);
              border: 0;
              outline: none;
              border-radius: 30px;
              padding: 10px 20px;
              background-color: white;
            "
          />
          <input class="search-input" type="submit" value="搜 索" />
        </label>
      </div>
    </div>
  </div>
</template>

<script setup>
const imgurl = require("../assets/backimg.jpg");
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

.search-input {
  color: rgb(255, 255, 255);
  font-size: 13px;
  font-weight: bold;
  cursor: pointer;
  width: 80px;
  margin-left: -80px;
  height: 50px;
  border: none;
  border-radius: 0 30px 30px 0;
  background-color: var(--tmdbLightGreen);
}

.search-input:hover {
  color: rgb(0, 0, 0);
}
</style>
