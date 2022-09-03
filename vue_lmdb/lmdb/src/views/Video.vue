<template>
  <Header />
  <n-layout
    class="videopage"
    v-bind:class="{
      darkThemeBck: themeData === true,
      lightThemeBck: themeData === false,
    }"
  >
    <n-spance class="videospace">
      <video
        id="localvideo"
        :src="videoUrl"
        controls="controls"
        @timeupdate="handleTimeUpdate"
        @loadeddata="handleLoadedata"
      ></video>
      <n-divider />
      <n-space style="height: 360px">
        <!-- 左侧图片 -->
        <img :src="`data:image/png;base64,` + videoData[route.query.id].Cover" />
        <!-- 右侧信息 -->
        <n-space vertical class="middleInfo" v-show="showEditForm === false">
          <div>
            <span style="font-size: 18px; float: left"
              >{{ videoData[route.query.id].Title }}
              <n-button
                v-show="videoData[route.query.id].CollStr !== undefined"
                strong
                secondary
                type="tertiary"
                round
                size="tiny"
                style="margin-left: 10px"
              >
                <template #icon>
                  <n-icon><FolderDetails /></n-icon>
                </template>
                合集
              </n-button>
            </span>
          </div>
          <n-space>
            <div v-for="(item, index) in videoData[route.query.id].TagArray">
              <n-tag :type="labelTypes[index]" round>
                {{ item }}
              </n-tag>
            </div>
          </n-space>
          <n-ellipsis expand-trigger="click" line-clamp="3" :tooltip="false">
            {{ videoData[route.query.id].Desc }}
          </n-ellipsis>
        </n-space>

        <!-- 编辑表单 -->

        <n-form
          v-show="showEditForm === true"
          class="middleInfo"
          ref="formInstRef"
          size="small"
          :model="videoFormModel"
          :rules="rules"
          label-placement="left"
        >
          <n-form-item label="标题" path="Title">
            <n-input
              v-model:value="videoFormModel.Title"
              :default-value="videoFormModel.Title"
              placeholder="请输入标题"
            />
          </n-form-item>
          <n-form-item label="标签" path="TagArray">
            <n-select
              :default-value="videoFormModel.TagArray"
              placeholder="增加或删除标签"
              v-model:value="videoFormModel.TagArray"
              multiple
              :options="videoTagOption"
            />
          </n-form-item>
          <n-form-item label="介绍" path="Desc">
            <n-input
              v-model:value="videoFormModel.Desc"
              type="textarea"
              placeholder="请输入视频介绍信息"
              :default-value="videoFormModel.Desc"
            />
          </n-form-item>

          <n-form-item label="合集" path="CollStr">
            <n-select
              placeholder="创建或选择加入已有合集"
              :default-value="videoFormModel.CollStr"
              v-model:value="videoFormModel.CollStr"
              filterable
              tag
              :options="collOptions"
            />
          </n-form-item>
          <div style="display: flex; justify-content: flex-end">
            <n-button size="small" @click="handleClose"> 关闭 </n-button>
            <n-button size="small" @click="handleRecover"> 恢复 </n-button>
            <n-button size="small" @click="handleValidateClick"> 提交 </n-button>
          </div>
        </n-form>

        <n-space vertical>
          <n-button
            @click="showEditForm = !showEditForm"
            strong
            secondary
            type="tertiary"
            round
            size="tiny"
          >
            <template #icon>
              <n-icon><Edit /></n-icon>
            </template>
            编辑
          </n-button>

          <n-button strong secondary type="tertiary" round size="tiny">
            <template #icon>
              <n-icon><Favorite /></n-icon>
            </template>
            收藏
          </n-button>
        </n-space>
      </n-space>
    </n-spance>
  </n-layout>
  <!-- <canvas id="localcanvas"></canvas>
  <button id="snapbtn" @click="handleClick">
    截取总时长{{ videoDurtion.timeStr }}
  </button> -->
</template>

<script setup>
import { useRoute } from "vue-router";
import { ref, computed } from "vue";
import Header from "@/components/Header.vue";
import { UpdateVideo } from "@/api/videolist";

import { storeToRefs } from "pinia";
import { darkTheme } from "naive-ui";
import { useDarkTheme } from "@/store/themeData";

import { useVideoData } from "@/store/videoData";

import { Edit, FolderDetails, Favorite } from "@vicons/carbon";

var showEditForm = ref(false);

const videoDataStore = useVideoData();
var { videoData } = storeToRefs(videoDataStore);

const darkThemeStore = useDarkTheme();
var { themeData } = storeToRefs(darkThemeStore);

const route = useRoute();

const videoUrl = config.SERVER_API + videoData.value[route.query.id].VideoUrl;

const formInstRef = ref(null);
const labelTypes = ["success", "warning", "error", "info"];

var videoFormModel = ref({
  Title: videoData.value[route.query.id].Title,
  TagArray: videoData.value[route.query.id].TagArray,
  Desc: videoData.value[route.query.id].Desc,
  CollStr: videoData.value[route.query.id].CollStr,
});

const videoTagOption = computed(() => {
  let tmpOptions = [];
  for (var i = 0; i < config.AllVideoTags.length; i++) {
    tmpOptions.push({
      label: config.AllVideoTags[i],
      value: config.AllVideoTags[i],
    });
  }
  return tmpOptions;
});

const collOptions = [
  {
    label: "",
    value: "",
  },
];

const rules = {
  Title: {
    required: true,
    min: 1,
    max: 15,
    message: "标题最短长度为 1,最大长度为15",
  },
  TagArray: {
    required: false,
    type: "array",
    max: 4,
    message: "最多可以填写4个标签",
  },
  Desc: {
    required: false,
    max: 100,
    message: "介绍内容不能超过100",
  },
  CollStr: {
    required: false,
    max: 15,
    message: "集合名称最大长度为15",
  },
};

const handleRecover = () => {
  videoFormModel.value.Title = videoData.value[route.query.id].Title;
  videoFormModel.value.TagArray = [];
  videoFormModel.value.TagArray.push(...videoData.value[route.query.id].TagArray);
  videoFormModel.value.Desc = videoData.value[route.query.id].Desc;
  videoFormModel.value.CollStr = videoData.value[route.query.id].CollStr;
  formInstRef.value?.restoreValidation();
};

const handleClose = () => {
  showEditForm.value = false;
};

const handleValidateClick = () => {
  let getError = false;
  formInstRef.value?.validate((errors) => {
    if (errors) {
      getError = true;
      console.error(errors);
    }
  });
  function procForm() {
    if (getError === true) {
      return;
    } 
    videoFormModel.value.MId = videoData.value[route.query.id].MId;
    UpdateVideo(videoFormModel.value);
    videoData.value[route.query.id].Title = videoFormModel.value.Title;
    videoData.value[route.query.id].Desc = videoFormModel.value.Desc;
    videoData.value[route.query.id].TagArray = videoFormModel.value.TagArray;
    videoData.value[route.query.id].CollStr = videoFormModel.value.CollStr;
    showEditForm.value = false;
  }
  setTimeout(procForm, 100);
};

const handleTimeUpdate = (ev) => {};

const handleClick = () => {
  updateCanvas();
};

const updateCanvas = () => {
  //console.log(ev,"timeupdate")
  const myvideo = document.getElementById("localvideo"); // 获取视频对象
  const mycanvas = document.getElementById("localcanvas"); // 获取 canvas 对象
  //console.log(ev.target.duration) // 总时长

  const ctx = mycanvas.getContext("2d"); // 绘制2d
  myvideo.crossOrigin = "anonymous"; // 解决跨域问题，也就是提示污染资源无法转换视频
  // myvideo.currentTime = 100; // 第一帧

  mycanvas.width = myvideo.clientWidth; // 获取视频宽度
  mycanvas.height = myvideo.clientHeight; //获取视频高度
  // 利用canvas对象方法绘图
  ctx.drawImage(myvideo, 0, 0, myvideo.clientWidth, myvideo.clientHeight);
  // mycanvas.toDataURL("image/png");
};
</script>
<style>
.videopage {
  margin-top: var(--headerTop);
  height: 100%;
}

.videospace {
  float: left;
  margin-top: 30px;
  margin-left: 60px;
}
video {
  position: relative;
  height: 376px;
  width: 668px;
  display: block;
  margin: 0;
  padding: 0;
  margin-top: 15px;
  border: 0;
}
.middleInfo {
  width: 320px;
}
</style>
