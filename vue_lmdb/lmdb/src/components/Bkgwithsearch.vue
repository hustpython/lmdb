<template>
  <div class="imgbkg">
    <img :src="imgurl" 
    alt="背景图片" />
  </div>
</template>

<script>
import { onMounted} from 'vue'
import axios from 'axios'
export default {
  setup() {
    let imgurl = "";
     onMounted(() => {
      console.log('Component is mounted!')
      axios.get("D://%E8%BF%85%E9%9B%B7%E4%B8%8B%E8%BD%BD//%E9%BB%91%E9%95%9C.Black.Mirror.1-4%E5%AD%A3%E5%85%A8%E9%9B%86%E3%80%90%E5%85%B3%E6%B3%A8%E5%BE%AE%E4%BF%A1%E5%85%AC%E4%BC%97%E5%8F%B7%EF%BC%9A%E7%BE%8E%E5%89%A7%E5%8F%94%E3%80%91/%E9%BB%91%E9%95%9C%20%E7%AC%AC4%E5%AD%A3/%E3%80%90%E5%85%B3%E6%B3%A8%E5%BE%AE%E4%BF%A1%E5%85%AC%E4%BC%97%E5%8F%B7%EF%BC%9A%E7%BE%8E%E5%89%A7%E5%8F%94%E3%80%91%E9%BB%91%E9%95%9C.Black.Mirror.S04E01.mp4",{
   responseType:'blob'}).then(res=>{
        console.log(res.data)
      })
    })
    return {
      imgurl:"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg.jj20.com%2Fup%2Fallimg%2F4k%2Fs%2F02%2F2109242332225H9-0-lp.jpg&refer=http%3A%2F%2Fimg.jj20.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1662407583&t=4a445dc15f5e215ab3170c4baa925bc5",
    };
  },
};
</script>

<style>
.imgbkg {
  width: 100%;
  display: flex;
  height: 0 auto;
  background-color: rgb(0, 0, 0);
}
/* 图片居中显示 */
.imgbkg img {
  text-align: center;
  margin: 0 auto;
  align-items: center;
  /* width: 100%; */
}
</style>