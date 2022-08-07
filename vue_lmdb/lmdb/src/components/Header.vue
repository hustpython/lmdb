<template>
  <!-- 头部开始，头部结束 -->
  <div class="header">
    <div class="wrap">
      <ul class="header-left">
        <li>
          <a href="/" class="logo">
            <img src="../assets/logo.svg" alt="Local Movie Database (LMDB)" />
          </a>
        </li>
        <li v-for="item in header_left_list">
          <a href="#">{{ item }}</a>
        </li>
      </ul>
    </div>
  </div>
</template>

<style>
.header {
  justify-content: center;
  align-items: center;
  height: var(--headerTop);
  background-color: var(--lightBule);
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  position: fixed;
  z-index: 30;
  transition: top 0.2s linear;
  transition-property: top;
  transition-duration: 0.2s;
  transition-timing-function: linear;
  transition-delay: 0s;
}
.header ul {
  position: relative;
  margin-left: 60px;
}
.wrap {
  width: 1226px;
  height: 0 auto;
}
.logo img {
  height: 50px;
  width: 120px;
  vertical-align: middle;
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
</style>

<script>
export default {
  setup() {
    return {
      header_left_list: ["影视", "影评", "剪辑"],
    };
  },
};
</script>