<template>
  <div class="middle_list">
   <button @click="handleClick">
    在这里
   </button>
  </div>
</template>

<script setup>
import {GetVideList} from "../api/videolist"
import {reactive} from 'vue'

const items = ["最近","最新","全部"]
const handleClick = ()=> {
    console.log("click")
    GetVideList().then((res) => {
       console.log(res)
	})
}
</script>

<style>
.middle_list {
  width: 100%;
  height: 500px;
  background-color: rgb(188, 150, 150);
  position: fixed;
  top: 400px;
}

</style>