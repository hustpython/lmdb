<template>
    <canvas id="localcanvas" style="display: none"></canvas>
    <div
            class="cards"
            v-bind:class="{
      showing: hoverEffict.isShowing,
    }"
    >
        <MiddleFilter @videoPageChange="handlePageChange"/>
        <div
                class="card"
                v-bind:class="{
        show: index === hoverEffict.index,
      }"
                @mouseenter="handleMouseEnter(index)"
                @mouseleave="handleMouseLeave(index)"
                v-for="(item, index) in videoData.slice(begin, end)"
        >
            <div class="card-image videoCardSize">
                <img
                        @click="setTitleHref(index)"
                        class="videoCardSize"
                        :src="item.Cover.length === 0 ? defaultCover : setCoverData(item.Cover)"
                />

                <span class="card-duration">{{ item.Duration }}</span>
            </div>

            <div class="video-mask videoCardSize" v-show="index === hoverEffict.videoIndex">
                <video
                        :id="'video_' + index"
                        crossorigin="anonymous"
                        class="videoCardSize"
                        muted
                        preload
                        autoplay
                        @loadeddata="handleLoadVideo($event, index)"
                        @mousemove="handleProgress($event)"
                        @click="handleVideoClick"
                        :src="item.TmpVideoUrl"
                ></video>
                <div class="set-cover">
                    <img
                            class="set-cover-svg"
                            @click="handleChangeBck(index)"
                            src="../assets/yl.svg"
                    />
                    <span class="set-cover-tip">设为封面</span>
                </div>
                <div class="video-duration-progress">
                    <span :style="'width:' + hoverEffict.progress + '%'"></span>
                </div>
            </div>

            <div @click="setTitleHref(index)" class="card-title">
                <div>
                    <n-ellipsis class="ellipsistitle">
                        {{ item.Title }}
                    </n-ellipsis>
                </div>
            </div>

            <div
                    class="card-flap flap1"
                    v-bind:class="{
          showing: hoverEffict.isShowing,
        }"
            >
                <n-space class="description">
                    <div v-for="(item, index) in item.TagArray">
                        <n-tag size="small" :type="labelTypes[index]" round>
                            {{ item }}
                        </n-tag>
                    </div>
                </n-space>
                <n-ellipsis class="description" :line-clamp="3" style="text-align: center">
                    {{ item.Desc }}
                </n-ellipsis>
            </div>
        </div>
    </div>
</template>

<script setup>
    import {GetMoviesByColl, GetVideList, UpdateVideo} from "@/api/videolist";
    import {reactive, ref, onBeforeMount} from "vue";
    import {timeFilter} from "@/api/timefilter";
    import {useVideoData} from "@/store/videoData";
    import {storeToRefs} from "pinia";
    import {useNotification} from "naive-ui";
    import {useRouter} from "vue-router";

    import MiddleFilter from "@/components/MiddleFilter.vue";

    const router = useRouter();

    const defaultCover = require("../assets/videocardbck.png");
    const videoDataStore = useVideoData();
    const notification = useNotification();

    const labelTypes = ["success", "warning", "error", "info"];

    var begin = ref(0);
    var end = ref(0);

    const absoluteIndex = (index) => {
        return begin.value + index;
    };

    const setCoverData = (cover) => {
        if (cover.indexOf(";") != -1) {
            return cover;
        }
        return "data:image/png;base64," + cover;
    };

    var {videoData} = storeToRefs(videoDataStore);

    const handlePageChange = (page, pageCount) => {
        begin.value = (page - 1) * pageCount;
        end.value = page * pageCount;
    };

    handlePageChange(1);
    const handleChangeBck = (index) => {
        const myvideo = document.getElementById("video_" + index); // 获取视频对象
        const mycanvas = document.getElementById("localcanvas"); // 获取 canvas 对象
        const ctx = mycanvas.getContext("2d"); // 绘制2d
        mycanvas.width = myvideo.clientWidth; // 获取视频宽度
        mycanvas.height = myvideo.clientHeight; //获取视频高度
        ctx.drawImage(myvideo, 0, 0, mycanvas.width, mycanvas.height);
        index = absoluteIndex(index);
        try {
            videoData.value[index].Cover = mycanvas.toDataURL("image/png"); // 导出图片
            let tmpCover = {
                MId: videoData.value[index].MId,
                Cover: videoData.value[index].Cover.slice(22),
            };
            UpdateVideo(tmpCover);
            notification.success({
                content: videoData.value[index].Title + " : 设置背景成功",
                duration: 3000,
            });
        } catch (error) {
            console.log("设置失败，稍后再试", error);
            videoData.value[index].TmpVideoUrl = "";
            notification.error({
                content: videoData.value[index].Title + " : 设置背景失败",
                duration: 2000,
            });
            setTimeout(function () {
                videoData.value[index].TmpVideoUrl =
                    config.SERVER_API + videoData.value[index].VideoUrl;
                hoverEffict.progress = 0;
            }, 1);
        }
    };

    const handleLoadVideo = (e, index) => {
        index = absoluteIndex(index);
        if (videoData.value[index].Duration.length !== 0) {
            return;
        }
        videoData.value[index].Duration = timeFilter(e.target.duration);
        let temDuration = {
            MId: videoData.value[index].MId,
            Duration: videoData.value[index].Duration,
        };
        UpdateVideo(temDuration);
    };

    var hoverEffict = reactive({
        index: -1,
        videoIndex: -1,
        isShowing: false,
        progress: 0,
    });
    // 加载后端数据
    onBeforeMount(() => {
        GetVideList().then((res) => {
            if (res.code == 200) {
                videoDataStore.setVideoData(res.data);
            }
        });
    });

    let videoTimer;

    const handleMouseEnter = (index) => {
        hoverEffict.index = index;
        hoverEffict.isShowing = true;
        videoTimer = setTimeout(() => {
            hoverEffict.videoIndex = index;
            index = absoluteIndex(index);
            videoData.value[index].TmpVideoUrl =
                config.SERVER_API + videoData.value[index].VideoUrl;
        }, 1500);
    };

    const handleMouseLeave = (index) => {
        clearTimeout(videoTimer);
        hoverEffict.isShowing = false;
        hoverEffict.index = -1;
        hoverEffict.videoIndex = -1;
        hoverEffict.progress = 0;

        index = absoluteIndex(index);
        videoData.value[index].TmpVideoUrl = "";
    };

    const handleVideoClick = (e) => {
        if (e.target.paused) {
            e.target.play();
        } else {
            e.target.pause();
        }
    };

    const handleProgress = (e) => {
        if (e.target.paused) {
            return;
        }
        hoverEffict.progress = (e.offsetX / e.target.clientWidth) * 100;
        e.target.currentTime = (e.target.duration * hoverEffict.progress) / 100;
    };

    const setTitleHref = (id) => {
        let temMId = videoData.value[absoluteIndex(id)].MId
        if (videoData.value[absoluteIndex(id)].CollStr.length != 0) {
            GetMoviesByColl(videoData.value[absoluteIndex(id)].CollStr).then((res) => {
                if (res.code == 200) {
                    videoDataStore.setVideoData(res.data);
                    for (let i = 0; i < res.data.length; i++) {
                        if (res.data[i].MId === temMId) {
                            router.push({name: "video", query: {id: i}});
                        }
                    }
                }
            });
        } else {
            router.push({name: "video", query: {id: absoluteIndex(id)}});
        }
    }

</script>

<style lang="scss">
    div.cards {
        margin: 15px;
        height: 100%;
        width: 0 auto;
        right: 10px;
        left: 10px;
        @include phone() {
            display: none;
        }
    }

    div.cards::after {
        content: "";
        width: 100%;
        height: 150px;
        display: block;
    }

    .card {
        display: inline-block;
        margin: 10px;
        position: relative;
        text-align: left;
        transition: all 0.3s 0s ease-in;
        max-width: var(--videoCardWidth);
        z-index: 1;
        border-radius: 10px 10px 0 0;
        box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);

        .card-title {
            padding: 6px 15px 10px;
            position: relative;
            z-index: 0;

            .ellipsistitle {
                max-width: calc($videoCardWidth * 0.9);
            }
        }
    }

    /* 封面相关 */

    .card-image {
        cursor: pointer;
        border-radius: inherit;

        img {
            border-radius: inherit;
        }

        .card-duration {
            position: absolute;
            bottom: 42px;
            right: 10px;
            height: 20px;
            line-height: 20px;
            transition: opacity 0.3s;
            z-index: 5;
            font-size: 13px;
            background-color: rgba(43, 104, 165, 0.4);
            color: #fff;
            border-radius: 2px;
            z-index: inherit;
        }
    }

    /* 视频相关 */
    .video-mask {
        top: 0;
        position: absolute;
        background-color: #000;
        border-radius: inherit;

        video {
            margin: 0;
            border-radius: inherit;
        }
    }

    .set-cover {
        display: -webkit-flex;
        display: flex;
        align-items: center;
        justify-content: center;
        position: absolute;
        top: 8px;
        right: 8px;
        width: 28px;
        height: 28px;
        border-radius: 6px;
        cursor: pointer;
        z-index: 9;
        transform: translateZ(0);
    }

    .set-cover-tip {
        display: none;
    }

    .set-cover-svg:hover + .set-cover-tip {
        pointer-events: none;
        user-select: none;
        position: absolute;
        bottom: -6px;
        right: -10px;
        transform: translateY(100%);
        font-size: 12px;
        color: #fff;
        border-radius: 4px;
        line-height: 18px;
        padding: 4px 8px;
        background-color: rgba(0, 0, 0, 0.8);
        white-space: nowrap;
        display: -webkit-flex;
    }

    .set-cover-svg {
        display: -webkit-flex;
        display: flex;
        align-items: center;
        justify-content: center;
        position: absolute;
        top: 8px;
        right: 8px;
        width: 28px;
        height: 28px;
        border-radius: 6px;
        cursor: pointer;
        transform: translateZ(0);
    }

    .video-duration-progress {
        position: absolute;
        bottom: 0px;
        width: 100%;
        height: 10px;
        border-color: rgb(0, 0, 0);
        border-style: solid;
        border-width: 4px 8px;
        background: #444;
        box-sizing: border-box;
    }

    .video-duration-progress span {
        display: block;
        background: #fff;
        height: 2px;
        transition: width 0.12s;
    }

    div.card div.card-flap {
        position: absolute;
        width: 100%;
        transform-origin: top;
        transform: rotateX(-90deg);
    }

    div.card div.flap1 {
        transition: all 0.3s 0.3s ease-out;
        z-index: -1;
        @include theme();
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
        z-index: 10;
    }

    div.card.show div.card-title a.toggle-info span {
        top: 15px;
    }

    div.card.show div.card-flap {
        transform: rotateX(0deg);
    }

    div.card.show div.flap1 {
        transition: all 0.3s 0s ease-out;
    }

    .description {
        padding: 6px 15px 10px;
        position: relative;
        font-size: 14px;
    }
</style>
