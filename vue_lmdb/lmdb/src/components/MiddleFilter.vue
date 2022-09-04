<template>
  <n-space class="videoTabs" justify="space-between">
    <n-tabs type="bar" default-value="4" animated v-model:value="tabVal">
      <n-tab name="1"> 筛选 </n-tab>
      <n-tab name="2"> 同步 </n-tab>
      <n-tab name="3"> 合集 </n-tab>
      <n-tab name="4"> 隐藏 </n-tab>
    </n-tabs>
    <div class="pagenum">
      <n-pagination
        v-model:page="page"
        v-model:page-size="pageSize"
        :page-count="pageCount"
        @update:page="handleUpdatePage"
        show-size-picker
        :page-sizes="pageSizeArr"
        style="margin-right: 40px"
      />
      <n-badge
        color="var(--tmdbLightGreen)"
        style="margin-right: 20px"
        :value="videodataNum"
        processing
      >
        记录数
      </n-badge>
    </div>
  </n-space>

  <div class="videoTabs" v-show="tabVal === '1'">
    <n-space>
      <div v-for="(item, index) in allTags">
        <n-button type="tertiary" @click="handleTagClick(index)" strong>
          {{ item }}
        </n-button>
      </div>
    </n-space>
  </div>

  <n-form :model="syncVideoForm" class="videoTabs" v-show="tabVal === '2'">
    <n-form-item label="最小值(M)">
      <n-select v-model:value="syncVideoForm.MinSize" :options="syncVideoSize" />
    </n-form-item>

    <n-form-item path="category" label="类型">
      <n-select v-model:value="syncVideoForm.MovieExt" multiple :options="options" />
    </n-form-item>

    <n-form-item>
      <n-popconfirm
        @positive-click="handleSbumitClick"
        positive-text="确认!"
        negative-text="取消!"
      >
        <template #trigger>
          <n-button>提交</n-button>
        </template>
        确认从本地同步最小为 {{ syncVideoForm.MinSize }} M 的
        {{ syncVideoForm.MovieExt }}格式的视频吗? 有可能会导致网页上的数据被覆盖哦!
      </n-popconfirm>
    </n-form-item>
  </n-form>

  <!-- <pre>{{ JSON.stringify(syncVideoForm, null, 2) }}</pre> -->
</template>

<script setup>
import { ref, computed, onBeforeMount } from "vue";
import { useVideoData } from "@/store/videoData";
import { storeToRefs } from "pinia";
import { SyncVideo } from "@/api/videolist";
import { useNotification } from "naive-ui";
import { GetAllTags, GetVideList, GetMoviesByTag } from "@/api/videolist";

const notification = useNotification();

const syncVideoSize = [
  {
    label: 0,
    value: 0,
  },
  {
    label: 400,
    value: 400,
  },
  {
    label: 800,
    value: 800,
  },
  {
    label: 1024,
    value: 1024,
  },
  {
    label: 2048,
    value: 2048,
  },
];

const videoDataStore = useVideoData();
var tabVal = ref();
var syncVideoForm = ref({
  MinSize: 400,
  MovieExt: ["mp4"],
});
const options = [
  {
    label: "mp4",
    value: "mp4",
  },
  {
    label: "mkv",
    value: "mkv",
  },
  {
    label: "rmvb",
    value: "rmvb",
  },
  {
    label: "avi",
    value: "avi",
  },
  {
    label: "webm",
    value: "webm",
  },
];

let allTags = ref(["全部"]);

// 加载后端数据
onBeforeMount(() => {
  GetAllTags().then((res) => {
    if (res.code == 200) {
      allTags.value.push(...res.data);
    }
  });
});

const handleTagClick = (index) => {
  if (index === 0) {
    GetVideList().then((res) => {
      if (res.code == 200) {
        videoDataStore.setVideoData(res.data);
        notification.success({
          content: "成功查询到" + res.data.length + "个视频",
          duration: 1000,
        });
      }
    });
  } else {
    GetMoviesByTag(allTags.value[index]).then((res) => {
      if (res.code == 200) {
        videoDataStore.setVideoData(res.data);
        notification.success({
          content: "成功查询到" + res.data.length + "个视频",
          duration: 1000,
        });
      }
    });
  }
};

var { videoData } = storeToRefs(videoDataStore);
const emit = defineEmits(["videoPageChange"]);

const pageWidth = ref(document.body.clientWidth);

const page = ref(1);

// index.css 中修改时，需要同步修改
const videoCardWidth = 270;
var pageSize = ref(Math.floor(pageWidth.value / videoCardWidth) * 2);

const handleUpdatePage = (page) => {
  emit("videoPageChange", page, pageSize.value);
};

const pageSizeArr = computed(() => {
  return [
    Math.floor(pageWidth.value / videoCardWidth) * 1,
    Math.floor(pageWidth.value / videoCardWidth) * 2,
    Math.floor(pageWidth.value / videoCardWidth) * 3,
  ];
});

const pageCount = computed(() => {
  handleUpdatePage(page.value);
  return Math.ceil(videoData.value.length / pageSize.value);
});

const videodataNum = computed(() => {
  return videoData.value.length;
});

window.onresize = () => {
  pageWidth.value = document.body.clientWidth;
};

const handleSbumitClick = () => {
  SyncVideo(syncVideoForm.value).then((res) => {
    if (res.code == 200) {
      if (res.data === null) {
        notification.success({
          content: "视频已经是最新的了!",
          duration: 3000,
        });
        return;
      }
      videoDataStore.setVideoData(res.data);
      notification.success({
        content: "成功从本地磁盘同步" + res.data.length + "个视频",
        duration: 3000,
      });
    }
  });
};
</script>

<style>
.videoTabs {
  padding: 7px 17px 7px 17px;
}
.pagenum {
  display: inline-flex;
  align-items: center;
}
</style>
