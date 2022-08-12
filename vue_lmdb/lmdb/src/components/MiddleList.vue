<template>
  <div class="cards">
    <div class="card" v-for="item in videoData">
      <img class="card-image" src="../file/1.jpg" alt="" />
      <div class="card-title">
        <a :href="'/video?id=' + item.Path" class="toggle-info btn">
          <span class="left"></span>
          <span class="right"></span>
        </a>
        <h2>
          {{ item.Title }}
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
import { ref, onBeforeMount } from "vue";

const videoData = ref([{ Path: "", Title: "" }]);
const handleClick = () => {
  GetVideList().then((res) => {
    if (res.code == 200) {
      videoData.value = res.data;
    }
  });
};
// 加载后端数据
onBeforeMount(() => {
  handleClick();
});
</script>

<style>
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
  margin: 10px auto;
  width: 100%;
  text-align: center;
}
div.card {
  background: #ffffff;
  display: inline-block;
  margin: 6px;
  perspective: 1000;
  position: relative;
  text-align: left;
  transition: all 0.3s 0s ease-in;
  z-index: 3;
  max-width: 380px;
}
div.card img {
  max-width: 380px;
}
div.card div.card-title {
  background: #ffffff;
  padding: 6px 15px 10px;
  position: relative;
  z-index: -1;
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
  font-size: 12px;
  font-weight: 700;
  letter-spacing: -0.05em;
  margin: 0;
  padding: 0;
  color: #000293;
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
</style>