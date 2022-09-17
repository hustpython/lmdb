<template>
    <div class="VideoContent">
        <canvas id="localVideoCanvas" style="display: none"></canvas>
        <video
                autoplay
                id="lvideo"
                :src=videoUrl
                @loadeddata="handleLoadStart"
                @click="handlePlay"
                @timeupdate="handleTimeProgress"
                @ended="handleEnd"
                @mousemove="handleMouseMove"
                @pause="handlePause"
                style="display: block"
                x5-video-player-type="h5"
                webkit-playsinline=""
                playsinline=""
                preload="metadata"
                crossorigin="anonymous"
        ></video>


        <!--            下方控制台-->
        <div class="progressMask"
             @click="handleProgress"
             v-show="showControl">
            <span class="progress"> </span>
            <span class="bkg"> </span>
            <div class="progressBtn">
                <n-icon size="16">
                    <VehicleSubway16Regular/>
                </n-icon>
            </div>
        </div>
        <div class="videoPlayCtrl"
             @mouseenter="handlePlayCtrlEnter"
             v-show="showControl">
            <n-icon
                    class="preControl" :size=controlBtnSize @click="handlePlay">
                <Previous20Regular/>
            </n-icon>
            <!--                播放/暂停按钮-->
            <n-icon
                    v-show="playStatus" class="playControl" :size=controlBtnSize @click="handlePlay">
                <PausePresentationOutlined/>
            </n-icon>
            <n-icon v-show="!playStatus" class="playControl" :size=controlBtnSize @click="handlePlay">
                <LiveTvRound/>
            </n-icon>

            <!--            下一个视频-->
            <n-icon
                    class="nextControl" :size=controlBtnSize @click="handlePlay">
                <Next20Regular/>
            </n-icon>


            <!--                显示时间-->
            <div class="timeView"
                 @click="handleTimeTextClick"
                 v-show=!showTimeEditInput
            >{{controlTimeView.playTime}}/{{controlTimeView.duration}}
            </div>


            <input class="timeEdit"
                   type="text"
                   v-model=inputTime
                   v-show=showTimeEditInput
                   @keyup.enter="handleInputBlur"
                   @blur="handleInputBlur"
                   ref="refInput">
            <n-icon class="coverSet" :size=controlBtnSize
                    @click="handleClickClap">
                <Camera/>
            </n-icon>

            <!--                声音调整按钮-->

            <n-icon
                    v-show=volumeBtnShow
                    @mouseenter="volumeControlMouseEnter"
                    @mouseleave="volumeControlMouseLeave"
                    @click="volumeControlClick"
                    class="volumeControl" :size=controlBtnSize>
                <VolumeUpFilled/>
            </n-icon>
            <n-icon
                    v-show=!volumeBtnShow
                    @mouseenter="volumeControlMouseEnter"
                    @mouseleave="volumeControlMouseLeave"
                    @click="volumeControlClick"
                    class="volumeControl" :size=controlBtnSize>
                <VolumeOffRound/>
            </n-icon>

            <!--            音量进度条显示-->
            <div
                    v-show=volumeShow
                    @click="handleVolumeProgressClick"
                    @mousedown="handleVolumeProgressDown"
                    @mouseenter="handleVolumeBoxEnter"
                    @mouseleave="handleVolumeBoxLeave"
            >
                <span class="volumeBox">
                    {{volumePercent}}
            </span>
                <div
                        class="volumeProgress"
                        ref="volumeBoxRef">
                    <div class="currentProgress">
                        <div class="volumeMoveBtn"
                        >
                        </div>
                    </div>
                </div>


            </div>
            <!--                全屏按钮-->
            <n-icon class="fullScreen" :size=controlBtnSize @click="handleFull">
                <FullScreenMaximize24Filled/>
            </n-icon>

            <div v-show="controlNotifyShow"
                 class="controlNotify">
                {{notifyMsg}}
            </div>

        </div>


        <!--            视频暂停时显示-->
        <n-icon v-show=!playStatus class="tvOff" size="90">
            <LiveTvRound/>
        </n-icon>


    </div>
</template>

<script setup>
    import {
        VehicleSubway16Regular, FullScreenMaximize24Filled, Next20Regular,
        Previous20Regular
    } from "@vicons/fluent";
    import {
        VolumeUpFilled,
        PausePresentationOutlined,
        VolumeOffRound,
        LiveTvRound,
    } from "@vicons/material";
    import {Camera} from "@vicons/carbon";
    import {getUTCTime, timeFilter, timeStrToSec} from "@/api/timefilter";
    import {ref, reactive, nextTick, defineProps, onBeforeUnmount} from "vue";
    import {UpdateVideo} from "@/api/videolist";

    import {storeToRefs} from "pinia";
    import {useVideoData} from "@/store/videoData";

    const videoDataStore = useVideoData();
    var {videoData} = storeToRefs(videoDataStore);

    const routeID = defineProps(["Id"]);
    const videoUrl = config.SERVER_API + videoData.value[routeID.Id].VideoUrl;

    let lvideo = {};

    const controlBtnSize = "24px";

    const progressBtnLeft = ref("0");
    const mouseLeft = ref("0px");
    const videoWidth = ref("668px");
    const videoHeight = ref("376px");
    const showControl = ref(true);
    const showTimeEditInput = ref(false);
    const controlNotifyShow = ref(false);
    let videoControlTimer;

    let controlNotifyShowTimer;

    document.body.onkeydown = function (e) {
        var e = event || window.event || arguments.callee.caller.arguments[0];
        // 空格
        if (e.keyCode == 32) {
            e.preventDefault();
            handlePlay();
        }
        // 左键
        if (e.keyCode == 39) {
            lvideo.currentTime += 6;
        }
        // 右键
        if (e.keyCode == 37) {
            lvideo.currentTime -= 6;
        }
        // 插入键,截图
        if (e.keyCode == 45) {
            handleClickClap()
        }
    }
    const notifyMsg = ref();
    let controlTimeView = reactive(
        {
            playTime: timeFilter(0),
            duration: timeFilter(0),
        })

    const handleProgress = (e) => {
        progressBtnLeft.value = e.offsetX.toString() + "px";
        lvideo.currentTime = e.offsetX / e.target.clientWidth * lvideo.duration;
    };
    const delayClearNotifyMsg = (z) => {
        notifyMsg.value = z;
        clearTimeout(controlNotifyShowTimer);
        controlNotifyShow.value = true;
        setTimeout(function () {
            controlNotifyShow.value = false;
        }, 3000)
    }
    const handleLoadStart = (e) => {
        lvideo = document.getElementById("lvideo");
        controlTimeView.duration = timeFilter(lvideo.duration);
        if (videoData.value[routeID.Id].LastWatch > 0) {
            e.target.currentTime = videoData.value[routeID.Id].LastWatch;
            delayClearNotifyMsg("为您定位至:" + timeFilter(videoData.value[routeID.Id].LastWatch));

        }
        if (videoData.value[routeID.Id].Duration.length !== 0) {
            return;
        }
        let updateTime = {
            MId: videoData.value[routeID.Id].MId,
            Duration: timeFilter(e.target.duration),
        };
        UpdateVideo(updateTime);
    };

    const playStatus = ref(true);
    const handlePlay = () => {
        playStatus.value = !playStatus.value;
        if (lvideo.paused) {
            lvideo.play();
            showControl.value = false;
        } else {
            showControl.value = true;
            lvideo.pause();
        }
    };
    const handleEnd = () => {
        playStatus.value = false;
    }
    const handleTimeProgress = (e) => {
        progressBtnLeft.value = e.target.currentTime / e.target.duration * 100 + '%';
        controlTimeView.playTime = timeFilter(e.target.currentTime)
    }
    const handleFull = () => {
        if (!document.fullscreenElement) {
            document.querySelector('.VideoContent').requestFullscreen();
            videoWidth.value = "100%";
        } else {
            document.exitFullscreen();
            videoWidth.value = "668px";
        }
    };

    document.addEventListener('fullscreenchange', () => {
        if (!document.fullscreenElement) {
            videoWidth.value = "668px";
        }
    })

    const handleMouseMove = () => {
        clearTimeout(videoControlTimer);
        showControl.value = true;
        if (lvideo.paused) {
            return;
        }
        videoControlTimer = setTimeout(function () {
            showControl.value = false;
        }, 3000);
    }
    const handlePlayCtrlEnter = () => {
        clearTimeout(videoControlTimer);
        showControl.value = true;
    }
    const inputTime = ref();
    const refInput = ref();
    const handleTimeTextClick = () => {
        showTimeEditInput.value = true;
        inputTime.value = timeFilter(lvideo.currentTime);
        nextTick(() => {
            refInput.value.focus()
        })
    }
    const handleInputBlur = () => {
        let inputTimeSec = timeStrToSec(inputTime.value);
        if (inputTimeSec > lvideo.duration) {
            lvideo.currentTime = lvideo.duration;
            inputTime.value = timeFilter(lvideo.currentTime);
        } else {
            lvideo.currentTime = inputTimeSec;
        }
        showTimeEditInput.value = false;
    }


    const volumeHeight = ref(
        "70px"
    )

    const volumePercent = ref(
        100
    )
    const volumeBtnShow = ref(true);
    const handleVolumeProgressDown = (e) => {
        document.onmousemove = (e) => {
            handleVolumeProgressClick(e);
        };
        document.onmouseup = function () {
            document.onmousemove = document.onmouseup = null;
        };
        return false;
    }

    const defaultHeight = 60;
    const volumeBoxRef = ref();
    const handleVolumeProgressClick = (e) => {
        let tmpHeight = volumeBoxRef.value.getBoundingClientRect().y + defaultHeight - e.y;

        if (tmpHeight < 0) {
            tmpHeight = 0;
        }
        if (tmpHeight > defaultHeight) {
            tmpHeight = defaultHeight;
        }
        volumePercent.value = Math.floor(tmpHeight / defaultHeight * 100);
        volumeHeight.value = tmpHeight + "px";
        lvideo.volume = volumePercent.value / 100;
        if (volumePercent.value === 0) {
            volumeBtnShow.value = false;
        } else {
            volumeBtnShow.value = true;
        }
    }


    const volumeShow = ref(false);
    let volumeShowTimer;
    const volumeControlMouseEnter = () => {
        volumeShow.value = true;
    }

    const volumeControlMouseLeave = () => {
        document.onmousemove = document.onmouseup = null;
        clearTimeout(volumeShowTimer);
        volumeShowTimer = setTimeout(function () {
            volumeShow.value = false;
        }, 1000);
    }
    const handleVolumeBoxEnter = () => {
        clearTimeout(volumeShowTimer);
        volumeShow.value = true;
    }
    const handleVolumeBoxLeave = () => {
        volumeShow.value = false;
    }
    const volumeControlClick = () => {
        volumeBtnShow.value = !volumeBtnShow.value;
        if (volumeBtnShow.value === false) {
            volumeHeight.value = "0px";
            volumePercent.value = 0;
            lvideo.volume = 0;
        }
    }

    const handleClickClap = () => {
        const mycanvas = document.getElementById("localVideoCanvas"); // 获取 canvas 对象
        const ctx = mycanvas.getContext("2d"); // 绘制2d
        mycanvas.width = lvideo.clientWidth; // 获取视频宽度
        mycanvas.height = lvideo.clientHeight; //获取视频高度
        ctx.drawImage(lvideo, 0, 0, mycanvas.width, mycanvas.height);
        try {
            videoData.value[routeID.Id].Cover = mycanvas.toDataURL("image/png"); // 导出图片
            let tmpCover = {
                MId: videoData.value[routeID.Id].MId,
                Cover: videoData.value[routeID.Id].Cover.slice(22),
            };
            UpdateVideo(tmpCover);
            delayClearNotifyMsg("为您截取一张图片");
        } catch (error) {
            console.log("设置失败，稍后再试", error);
        }
    };

    onBeforeUnmount(() => {
        handlePause();
    });

    const handlePause = () => {
        let updateTime = {
            MId: videoData.value[routeID.Id].MId,
            RecentWatch: getUTCTime(),
            LastWatch: Math.floor(lvideo.currentTime),
        };
        UpdateVideo(updateTime);
    };


</script>

<style scoped lang="scss">
    .VideoContent {
        width: v-bind(videoWidth);
        position: relative;
        display: flex;
        height: v-bind(videoHeight);

        video {
            position: absolute;
            width: v-bind(videoWidth);
            height: 100%;
            align-items: center;
            bottom: 0;
        }

        .progressMask {
            width: v-bind(videoWidth);
            cursor: pointer;
            position: absolute;
            display: flex;
            height: 15px;
            border-radius: 3px;
            bottom: 30px;

            &:hover {
                .progressBtn {
                    display: block;
                    position: absolute;
                    left: v-bind(progressBtnLeft);
                    top: -8px;
                }
            }

            .progressBtn {
                display: none;
            }

            .bkg {
                width: v-bind(videoWidth);
                position: absolute;
                height: 2px;
                background: $lightPink;
                opacity: 0.5;
            }

            .progress {
                width: v-bind(videoWidth);
                position: relative;
                height: 2px;
                width: v-bind(progressBtnLeft);
                background: #D81B60;
            }
        }

        .videoPlayCtrl {
            width: v-bind(videoWidth);
            height: 36px;
            position: absolute;
            display: flex;
            bottom: 0px;
            background-color: rgb(21, 21, 21, 0.1);

            .playControl {
                position: absolute;
                left: 50px;
                cursor: pointer;
            }

            .nextControl {
                position: absolute;
                left: 90px;
                cursor: pointer;
            }

            .preControl {
                position: absolute;
                left: 10px;
                cursor: pointer;
            }

            .timeView {
                position: absolute;
                left: 150px;
                font-size: 15px;
                top: 5px;
                cursor: text;
                user-select: none;
            }

            .timeEdit {
                position: absolute;
                left: 150px;
                font-size: 15px;
                top: 5px;
                border: none;
                width: 70px;
                outline: medium;
                cursor: text;
                background: rgb(189, 189, 189, 0.3);
                user-select: none;
            }

            .coverSet {
                position: absolute;
                left: 300px;
                font-size: 15px;
                top: 5px;
                cursor: pointer;
            }


            .volumeControl {
                position: absolute;
                right: 120px;
                cursor: pointer;
            }


            $volumeViewBottom: 55px;
            $volumeViewHeight: 100px;


            .volumeBox {
                position: absolute;
                height: $volumeViewHeight;
                width: 30px;
                right: 120px;
                font-size: 12px;
                text-align: center;
                bottom: $volumeViewBottom;
                background: rgb(19, 19, 19, 0.6);
                user-select: none;
            }

            .volumeProgress {
                position: absolute;
                cursor: pointer;
                height: $volumeViewHeight - 40;
                width: 3px;
                right: 133px;
                bottom: $volumeViewBottom + 20;
                background: white;

                .currentProgress {
                    position: absolute;
                    max-height: $volumeViewHeight - 40;
                    height: v-bind(volumeHeight);
                    width: 3px;
                    background: lightblue;
                    bottom: 0;

                    .volumeMoveBtn {
                        position: absolute;
                        height: 12px;
                        width: 12px;
                        left: -5px;
                        border-radius: 6px;
                        background: lightblue;
                        bottom: v-bind(volumeHeight)-6px;

                        &:hover {
                            box-shadow: 0 0 2px 2px lightblue;
                        }
                    }
                }
            }


            .fullScreen {
                position: absolute;
                right: 40px;
                cursor: pointer;
            }

            .controlNotify {
                position: absolute;
                font-size: 15px;
                left: 20px;
                bottom: $volumeViewBottom;
                background-color: rgb(21, 21, 21, 0.6);
                line-height: 40px;
                text-align: center;
                width: 180px;
                border-radius: 15px;
                user-select: none;
            }
        }


        .tvOff {
            display: flex;
            position: absolute;
            right: 10px;
            bottom: 75px;
        }

    }
</style>
