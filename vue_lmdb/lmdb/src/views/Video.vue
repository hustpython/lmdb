<template>
    <Header/>
    <n-layout class="videopage">
        <n-spance class="videospace">
            <div>
                <LVideo :Id=routeID></LVideo>
            </div>

            <n-space style="height: 360px;margin-top:30px">
                <!-- 左侧图片 -->
                <img class="videoCardSize" :src="setCoverData(videoData[routeID].Cover)"/>
                <!-- 右侧信息 -->
                <n-space vertical class="middleInfo" v-show="showEditForm === false">
                    <div>
            <span style="font-size: 18px; float: left"
            >{{ videoData[routeID].Title }}
              <n-button
                      v-show="videoData[routeID].CollStr !== ''"
                      strong
                      secondary
                      type="tertiary"
                      round
                      size="tiny"
                      style="margin-left: 10px"
              >
                <template #icon>
                  <n-icon><FolderDetails/></n-icon>
                </template>
                合集
              </n-button>
            </span>
                    </div>
                    <n-space>
                        <div v-for="(item, index) in videoData[routeID].TagArray">
                            <n-tag :type="labelTypes[index]" round>
                                {{ item }}
                            </n-tag>
                        </div>
                    </n-space>
                    <n-ellipsis
                            style="max-width: 320px"
                            expand-trigger="click"
                            line-clamp="3"
                            :tooltip="false"
                    >
                        {{ videoData[routeID].Desc }}
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
                        <n-button size="small" @click="handleClose"> 关闭</n-button>
                        <n-button size="small" @click="handleRecover"> 恢复</n-button>
                        <n-button size="small" @click="handleValidateClick"> 提交</n-button>
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
                            <n-icon>
                                <Edit/>
                            </n-icon>
                        </template>
                        编辑
                    </n-button>

                    <n-button strong secondary type="tertiary" round size="tiny">
                        <template #icon>
                            <n-icon>
                                <Favorite/>
                            </n-icon>
                        </template>
                        收藏
                    </n-button>
                </n-space>
            </n-space>
        </n-spance>
    </n-layout>
</template>

<script setup>
    import {useRoute} from "vue-router";
    import {ref, computed, onBeforeMount} from "vue";
    import Header from "@/components/Header.vue";
    import LVideo from "@/components/LVideo.vue";
    import {UpdateVideo, GetAllColl, DeleteMovieColl} from "@/api/videolist";
    import {storeToRefs} from "pinia";
    import {useVideoData} from "@/store/videoData";
    import {Edit, FolderDetails, Favorite, Camera} from "@vicons/carbon";
    import {getUTCTime, timeFilter} from "@/api/timefilter";

    var showEditForm = ref(false);

    const videoDataStore = useVideoData();
    var {videoData} = storeToRefs(videoDataStore);

    const route = useRoute();

    const routeID = route.query.id;

    const formInstRef = ref(null);
    const labelTypes = ["success", "warning", "error", "info"];

    const setCoverData = (cover) => {
        if (cover.indexOf(";") != -1) {
            return cover;
        }
        return "data:image/png;base64," + cover;
    };

    var videoFormModel = ref({
        Title: videoData.value[routeID].Title,
        TagArray: videoData.value[routeID].TagArray,
        Desc: videoData.value[routeID].Desc,
        CollStr: videoData.value[routeID].CollStr,
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

    let collOptions = ref([{label: "", value: ""}]);

    onBeforeMount(() => {
        GetAllColl().then((res) => {
            if (res.code == 200) {
                for (var i = 0; i < res.data.length; i++) {
                    collOptions.value.push({
                        label: res.data[i],
                        value: res.data[i],
                    });
                }
            }
        });
    });

    const rules = {
        Title: {
            required: true,
            min: 1,
            max: 15,
            message: "标题最短长度为 1,最大长度为15",
        },
        TagArray: {
            required: true,
            type: "array",
            min: 1,
            max: 4,
            message: "最新需要1个标签,最多可以填写4个标签",
        },
        Desc: {
            required: false,
            max: 300,
            message: "介绍内容不能超过300",
        },
        CollStr: {
            required: false,
            max: 15,
            message: "集合名称最大长度为15",
        },
    };

    const handleRecover = () => {
        videoFormModel.value.Title = videoData.value[routeID].Title;
        videoFormModel.value.TagArray = [];
        videoFormModel.value.TagArray.push(...videoData.value[routeID].TagArray);
        videoFormModel.value.Desc = videoData.value[routeID].Desc;
        videoFormModel.value.CollStr = videoData.value[routeID].CollStr;
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
            videoFormModel.value.MId = videoData.value[routeID].MId;
            UpdateVideo(videoFormModel.value);
            videoData.value[routeID].Title = videoFormModel.value.Title;
            videoData.value[routeID].Desc = videoFormModel.value.Desc;
            videoData.value[routeID].TagArray = videoFormModel.value.TagArray;
            videoData.value[routeID].CollStr = videoFormModel.value.CollStr;
            if (videoFormModel.value.CollStr === "") {
                DeleteMovieColl(videoFormModel.value.MId);
            }
            showEditForm.value = false;
        }

        setTimeout(procForm, 100);
    };
</script>
<style lang="scss">
    .videopage {
        margin-top: var(--headerTop);
        height: 100%;
        @include theme();
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
        @include phone() {
            width: 300px;
            height: 200px;
        }
    }

    .middleInfo {
        width: 320px;
    }
</style>
