<template>
    <div class="VideoContent">

        <video
                id="lvideo"
                src="../assets/陈逗逗自弹自唱你的眼神.mp4"
                @loadeddata="handleLoadStart"
                @click="handlePlay"
                @timeupdate="handleTimeProgress"
                @ended="handleEnd"
                @mousemove="handleMouseMove"
        ></video>


        <!--            下方控制台-->
        <div class="progressMask"
             @click="handleProgress"
             v-show="showControl">
            <span class="progress"> </span>
            <span class="bkg"> </span>
            <div class="progressBtn">
                <n-icon size="18">
                    <VehicleSubway16Regular/>
                </n-icon>
            </div>
        </div>
        <div class="videoPlayCtrl"
             @mouseenter="handlePlayCtrlEnter"
             v-show="showControl">
            <!--                播放/暂停按钮-->
            <n-icon
                    v-show="playStatus" class="playControl" size="30" @click="handlePlay">
                <PausePresentationOutlined/>
            </n-icon>
            <n-icon v-show="!playStatus" class="playControl" size="30" @click="handlePlay">
                <LiveTvRound/>
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
            <n-icon class="coverSet" size="30" :depth="2">
                <Camera/>
            </n-icon>

            <!--                声音调整按钮-->

            <n-icon
                    v-show=volumeBtnShow
                    @mouseenter="volumeControlMouseEnter"
                    @mouseleave="volumeControlMouseLeave"
                    @click="volumeControlClick"
                    class="volumeControl" size="30">
                <VolumeUpFilled/>
            </n-icon>
            <n-icon
                    v-show=!volumeBtnShow
                    @mouseenter="volumeControlMouseEnter"
                    @mouseleave="volumeControlMouseLeave"
                    @click="volumeControlClick"
                    class="volumeControl" size="30">
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
            <n-icon class="fullScreen" size="30" @click="handleFull">
                <FullScreenMaximize24Filled/>
            </n-icon>
        </div>


        <!--            视频暂停时显示-->
        <n-icon v-show=!playStatus class="tvOff" size="120">
            <LiveTvRound/>
        </n-icon>


    </div>
</template>

<script setup>
    import {VehicleSubway16Regular, FullScreenMaximize24Filled} from "@vicons/fluent";
    import {
        VolumeUpFilled,
        PausePresentationOutlined,
        VolumeOffRound,
        LiveTvRound,
    } from "@vicons/material";
    import {Camera} from "@vicons/carbon";
    import {timeFilter, timeStrToSec} from "@/api/timefilter";
    import {ref, reactive, nextTick} from "vue";

    let lvideo = {};
    const progressBtnLeft = ref("0%");
    const mouseLeft = ref("0px");
    const videoWidth = ref("668px");
    const videoHeight = ref("376px");
    const showControl = ref(true);
    const showTimeEditInput = ref(false);
    let videoControlTimer;


    document.body.onkeypress = function (e) {
        if (e.keyCode == 32) {
            e.preventDefault();
            handlePlay();
        }
    }

    let controlTimeView = reactive(
        {
            playTime: timeFilter(0),
            duration: timeFilter(0),
        })

    const handleProgress = (e) => {
        progressBtnLeft.value = e.offsetX.toString() + "px";
        lvideo.currentTime = e.offsetX / e.target.clientWidth * lvideo.duration;
    };

    const handleLoadStart = () => {
        lvideo = document.getElementById("lvideo");
        controlTimeView.duration = timeFilter(lvideo.duration);
    };

    const playStatus = ref(false);
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
        }
    };

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

</script>

<style scoped lang="scss">
    .VideoContent {
        width: v-bind(videoWidth);
        position: relative;
        display: flex;
        height: v-bind(videoHeight);

        video {
            position: absolute;
            width: 100%;
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
            bottom: 36px;

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
                height: 3px;
                background: $lightPink;
                opacity: 0.1;
            }

            .progress {
                width: v-bind(videoWidth);
                position: relative;
                height: 3px;
                width: v-bind(progressBtnLeft);
                background: #f48fb1;
            }
        }

        .videoPlayCtrl {
            width: v-bind(videoWidth);
            height: 36px;
            position: absolute;
            display: flex;
            bottom: 0px;

            .playControl {
                position: absolute;
                left: 10px;
                cursor: pointer;
            }

            .timeView {
                position: absolute;
                left: 100px;
                font-size: 15px;
                top: 5px;
                cursor: text;
                user-select: none;
                @include theme()
            }

            .timeEdit {
                position: absolute;
                left: 100px;
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
                left: 250px;
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
                right: 60px;
                cursor: pointer;
            }
        }


        .tvOff {
            display: flex;
            position: absolute;
            right: 75px;
            bottom: 180px;
        }

    }
</style>
