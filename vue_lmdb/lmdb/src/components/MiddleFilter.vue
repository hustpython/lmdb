<template>
    <div class="filter">
        <n-space class="videoTabs" justify="space-between">
            <n-tabs
                    type="bar"
                    default-value="4"
                    animated
                    v-model:value="tabVal"
                    :on-update:value="handleUpdateValue"
            >
                <n-tab name="1"> 筛选</n-tab>
                <n-tab name="2"> 同步</n-tab>
                <n-tab name="3">
                    <n-dropdown
                            :options="collMenuOptions"
                            @select="handleCollMenuSelect">
                        <n-button>合集</n-button>
                    </n-dropdown>
                </n-tab>
                <n-tab name="5"> 最近</n-tab>
                <n-tab name="4"> 隐藏</n-tab>
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
                    <n-button
                            size="small"
                            bordered="false"
                            color="#42A5F5"
                            :ghost="tagIndex !== index"
                            @click="handleTagClick(index)"
                            round
                    >
                        {{ item }}
                    </n-button>
                </div>
            </n-space>
        </div>

        <n-form :model="syncVideoForm" class="videoTabs" v-show="tabVal === '2'">
            <n-form-item label="最小值(M)">
                <n-select v-model:value="syncVideoForm.MinSize" :options="syncVideoSize"/>
            </n-form-item>

            <n-form-item path="category" label="类型">
                <n-select v-model:value="syncVideoForm.MovieExt" multiple :options="options"/>
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

        <div class="videoTabs" v-show="tabVal === '3'">
            <n-collapse
                    accordion="true"
                    v-show="collMenuSelectKey===1"
                    arrow-placement="right"
                    :expanded-names="expandedCollNames"
                    @item-header-click="handleCollHeaderClick"
            >
                <n-collapse-item
                        :title="item.Title" :name="index" v-for="(item, index) in collEditForm">
                    <n-form
                            label-placement="left"
                            label-width="auto"
                    >
                        <n-form-item label="选择视频" path="transferValue">
                            <n-transfer :options="item.TranOptions"
                                        v-model:value="item.SelectedValues"/>
                        </n-form-item>
                        <n-form-item label="合集名称" path="collName">
                            <n-input
                                    v-model:value="item.DefaultCollName"
                                    maxlength="15"
                                    show-count clearable
                            />
                        </n-form-item>
                        <n-form-item>
                            <n-space>
                                <n-button attr-type="button" @click="handleAddCollFormSubmit(index)">
                                    提交
                                </n-button>
                                <n-button attr-type="button" @click="handleAddCollFormCancel">
                                    取消
                                </n-button>
                            </n-space>
                        </n-form-item>
                    </n-form>
                </n-collapse-item>
            </n-collapse>

            <n-list v-show="collMenuSelectKey===0" class="coll" v-for="(coll, index) in allColl" hoverable clickable
                    bordered>
                <n-list-item @click="handleCollClick(index)">
                    <n-thing
                            :title="coll"
                            content-style="margin-top: 10px; font-size:12px!important"
                    >
                    </n-thing>
                </n-list-item>
            </n-list>
        </div>
    </div>
</template>

<script setup>
    import {ref, computed, onBeforeMount} from "vue";
    import {useVideoData} from "@/store/videoData";
    import {storeToRefs} from "pinia";
    import {SyncVideo, GetMoviesByRecent} from "@/api/videolist";
    import {useNotification} from "naive-ui";

    import {
        GetAllTags,
        GetVideList,
        GetMoviesByTag,
        GetAllColl,
        GetMoviesByColl,
        GetMovieAndColl,
        BatchAddColl,
    } from "@/api/videolist";

    import {useRouter} from "vue-router";

    const router = useRouter();

    const notification = useNotification();

    const syncVideoSize = [
        {
            label: 0,
            value: 0,
        },
        {
            label: 200,
            value: 200,
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
    var tagIndex = ref(0);

    var syncVideoForm = ref({
        MinSize: 200,
        MovieExt: ["mp4", "MP4"],
    });
    const options = [
        {
            label: "mp4",
            value: "mp4",
        },
        {
            label: "MP4",
            value: "MP4",
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
    let allColl = ref([]);
    // 加载后端数据
    onBeforeMount(() => {
        GetAllTags().then((res) => {
            if (res.code == 200) {
                allTags.value.push(...res.data);
            }
        });
        GetAllColl().then((res) => {
            if (res.code == 200) {
                allColl.value = res.data;
            }
        });
    });

    const handleUpdateValue = (value) => {
        tabVal.value = value;
        if (value === "5") {
            GetMoviesByRecent().then((res) => {
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

    const handleTagClick = (index) => {
        tagIndex.value = index;
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

    const handleCollClick = (index) => {
        GetMoviesByColl(allColl.value[index]).then((res) => {
            if (res.code == 200) {
                videoDataStore.setVideoData(res.data);
                notification.success({
                    content: "合集:" + allColl.value[index] + "共有" + res.data.length + "个视频",
                    duration: 1000,
                });
            }
        });
    };

    var {videoData} = storeToRefs(videoDataStore);
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
                } else {
                    videoDataStore.addVideoData(res.data);
                    notification.success({
                        content: "成功从本地磁盘同步" + res.data.length + "个视频",
                        duration: 3000,
                    });
                }
                router.go(0);
            }
        });
    };

    const collMenuOptions = [{
        label: "查询",
        key: 0,
    }, {
        label: "编辑",
        key: 1
    }];
    let collMenuSelectKey = ref(0);


    const collEditForm = ref();
    const expandedCollNames = ref([]);

    const handleAddCollFormSubmit = (index) => {
        if (collEditForm.value[index].SelectedValues.length === 0 ||
            collEditForm.value[index].DefaultCollName.length == 0 ||
            collEditForm.value[index].DefaultCollName.length > 15) {
            notification.error({
                content: "输入数据无效",
                duration: 1000,
            });
            return;
        } else {
            let temData = {
                CollName: collEditForm.value[index].DefaultCollName,
                MIds: collEditForm.value[index].SelectedValues,
            }
            BatchAddColl(temData).then((res) => {
                if (res.code == 200) {
                    notification.success({
                        content: collEditForm.value[index].SelectedValues.length +
                            "个视频超过加入合集" + collEditForm.value[index].DefaultCollName,
                        duration: 3000,
                    });
                } else {
                    notification.error({
                        content: "添加合集失败",
                        duration: 1000,
                    });
                }
            });
        }
        expandedCollNames.value = [];
    }

    const handleLoadMovieAndColl = () => {
        GetMovieAndColl().then((res) => {
            if (res.code == 200) {
                collEditForm.value = res.data;
            }
        });
    }

    const handleCollMenuSelect = (key) => {
        collMenuSelectKey.value = key;
        if (key === 1) {
            handleLoadMovieAndColl();
        } else {
            GetAllColl().then((res) => {
                if (res.code == 200) {
                    allColl.value = res.data;
                }
            });
        }
    }

    const handleAddCollFormCancel = () => {
        expandedCollNames.value = [];
    }
    const handleCollHeaderClick = (name) => {
        if (name.expanded) {
            expandedCollNames.value = [name.name];
        } else {
            expandedCollNames.value = [];
        }
    }

</script>

<style lang="scss">
    .filter {
        @include phone {
            display: none;
        }

        .videoTabs {
            padding: 7px 17px 7px 17px;
            max-height: 300px;
            overflow: auto;
            @include phone {
                display: none;
            }
        }

        .pagenum {
            display: inline-flex;
            align-items: center;
        }

        .coll {
            @include theme();
        }
    }
</style>
