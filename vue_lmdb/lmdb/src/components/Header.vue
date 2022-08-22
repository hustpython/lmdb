<template>
  <!-- 头部开始，头部结束 -->
  <div class="header">
    <div class="wrap">
      <ul class="header-left">
        <li>
          <a href="/" class="logo">
            <img :src="logourl" alt="Local Movie Database (LMDB)" />
          </a>
        </li>
        <li v-for="item in header_left_list">
          <a href="#">{{ item }}</a>
        </li>
      </ul>
      <div class="header-right">
        <img :src="avatarurl" alt="mxq" />
      </div>
    </div>
  </div>
</template>

<style>
.header {
  justify-content: center;
  align-items: center;
  height: var(--headerTop);
  background-color: var(--lightBule);
  top: var(--headerTopScroll);
  left: 0;
  width: 100%;
  display: flex;
  position: fixed;
  z-index: 30;
  transition: top 0.2s linear;
}
.header ul {
  position: relative;
}
.wrap {
  width: 1000px;
  height: 0 auto;
}
.logo img {
  height: 50px;
  width: 120px;
  vertical-align: middle;
}

.header-right {
  float: right;
}

.header-right > img {
  height: 40px;
  width: 40px;
  border-radius: 40px;
  margin-top: 25px;
}

.header-left {
  float: left;
}
.header li {
  float: left;
  margin-right: 30px;
}

.header a {
  line-height: 60px;
  color: rgb(210, 221, 209);
  font-size: 21px;
  font-weight: 556px;
}

.header a:hover {
  color: var(--tmdbLightGreen);
}
</style>

<script>
export default {
  mounted() {
    window.addEventListener(
      "scroll",
      () => {
        if (scrollY > 64) {
          document.body.style.setProperty("--headerTopScroll", "-64px");
        } else {
          document.body.style.setProperty("--headerTopScroll", "0px");
        }
      },
      true
    );
  },
  setup() {
    return {
      header_left_list: ["影视", "影评", "剪辑"],
      logourl: require("../assets/logo.svg"),
      avatarurl: require("../assets/axu.png"),
    };
  },
};
</script>
