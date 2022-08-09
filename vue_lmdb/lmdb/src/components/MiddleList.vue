<template>
  <div class="middle_list">
   <button @click="handleClick">
    在这里
   </button>
   <ul>
      <li v-for="item in videoData">
          {{ item.Path}}
        </li>
    </ul>
  </div>
</template>

<script setup>
import {GetVideList} from "../api/videolist"
import {ref} from 'vue'

const videoData = ref([{Path:""}]);
const handleClick = ()=> {
    GetVideList().then((res) => {
      if (res.code == 200) {
         videoData.value = res.data;
      }
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