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
          <a href="#">{{item}}</a>
        </li>
      </ul>
    </div>
  </div>
</template>

<style>
.header {
  width: 100%;
  height: 60px;
  background-color: rgb(3, 37, 65);
  margin: 0;
  padding: 0;
  position: absolute;
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
    header_left_list:["影视","影评","剪辑"]
  }
 }
};
</script>