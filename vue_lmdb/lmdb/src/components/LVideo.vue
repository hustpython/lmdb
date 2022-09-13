<template>
    <div class="VideoContent">

        <video
                id="lvideo"
                src="../assets/陈逗逗自弹自唱你的眼神.mp4"
                @loadstart="handleLoadStart"
                @click="handlePlay"
                @timeupdate="handleTimeProgress"
        ></video>


        <!--            下方控制台-->
        <div class="progressMask" @click="handleProgress" @mousemove="handleProgressHover">
            <span class="progress"> </span>
            <span class="bkg"> </span>
            <div class="progressBtn">
                <n-icon size="18">
                    <VehicleSubway16Regular/>
                </n-icon>
            </div>
            <div class="trangleIcon">
                <div class="trangleDown"></div>
                <div class="trangleUp"></div>
            </div>

        </div>
        <div class="videoPlayCtrl">
            <!--                播放/暂停按钮-->
            <n-icon v-show="playStatus" class="playControl" size="30" @click="handlePlay">
                <PausePresentationOutlined/>
            </n-icon>
            <n-icon v-show="!playStatus" class="playControl" size="30" @click="handlePlay">
                <LiveTvRound/>
            </n-icon>
            <!--                声音调整按钮-->

            <n-icon class="volumeControl" size="30" @click="handlePlay">
                <VolumeUpFilled/>
            </n-icon>
            <n-icon class="volumeControl" size="30">
                <VolumeOffRound/>
            </n-icon>

            <!--                显示时间-->
            <div class="timeView">1/1</div>
            <!--                全屏按钮-->

            <n-icon class="fullScreen" size="30" @click="handleFull">
                <FullScreenMaximize24Filled/>
            </n-icon>
        </div>

        <!--视频上面的显示-->
        <div class="maskSetting">
            <!--            视频暂停时显示-->
            <n-icon v-show=!playStatus class="tvOff" size="120">
                <LiveTvRound/>
            </n-icon>
            <div class="videoRightBtn">
                <div class="jietu">
                    <!--                设置封面按钮-->
                    <n-icon size="30" :depth="3">
                        <Camera/>
                    </n-icon>
                </div>
                <div class="edit">
                    <!--                设置封面按钮-->
                    <n-icon size="30" :depth="3">
                        <Camera/>
                    </n-icon>
                </div>
            </div>
        </div>
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
    import {ref} from "vue";

    let lvideo = {};
    const progressBtnLeft = ref("0%");
    const mouseLeft = ref("0px");
    const videoWidth = ref("668px");
    const videoHeight = ref("376px");

    const handleProgress = (e) => {
        progressBtnLeft.value = e.offsetX.toString() + "px";

        console.log()
    };
    const handleProgressHover = (e) => {
        mouseLeft.value = e.offsetX.toString() + "px";
    };
    const handleLoadStart = () => {
        lvideo = document.getElementById("lvideo");
    };
    const playStatus = ref(false);
    const handlePlay = () => {
        playStatus.value = !playStatus.value;
        if (lvideo.paused) {
            lvideo.play();
        } else {
            lvideo.pause();
        }
    };
    const handleTimeProgress = (e) => {
        progressBtnLeft.value = e.target.currentTime / e.target.duration * 100 + '%';
    }
    const handleFull = () => {
        let videoContent = document.querySelector('.VideoContent');
        if (!document.fullscreenElement) {
            videoContent.requestFullscreen();
            videoWidth.value = "100%";
        } else {
            document.exitFullscreen();
        }
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

            .trangleIcon {
                display: none;
            }

            &:hover {
                .trangleIcon {
                    display: none;
                    position: absolute;

                    .trangleDown {
                        width: 0;
                        height: 0;
                        position: absolute;
                        top: 4px;
                        border: 6px solid transparent;
                        border-top: none;
                        border-bottom-color: #f48fb1;
                        left: v-bind(mouseLeft);
                    }

                    .trangleUp {
                        width: 0;
                        height: 0;
                        position: absolute;
                        top: -8px;
                        border: 6px solid transparent;
                        border-bottom: none;
                        border-top-color: #f48fb1;
                        left: v-bind(mouseLeft);
                    }
                }
            }


        }

        .videoPlayCtrl {
            width: v-bind(videoWidth);
            height: 36px;
            position: absolute;
            display: flex;
            cursor: pointer;
            background-color: rgb(224, 242, 241, 0.1);

            bottom: 0px;

            .playControl {
                position: absolute;
                left: 10px;
            }

            .volumeControl {
                position: absolute;
                right: 120px;
            }

            .timeView {
                position: absolute;
                left: 100px;
            }

            .fullScreen {
                position: absolute;
                right: 60px;
            }
        }

        .maskSetting {
            width: v-bind(videoWidth);
            position: absolute;
            display: flex;
            bottom: 30px;

            .tvOff {
                position: absolute;
                right: 100px;
                bottom: 120px;
            }

            .videoRightBtn {
                height: 80px;
                border-radius: 10px;
                width: 40px;
                background-color: rgba(165, 42, 42, 0.3);
                opacity: 1;
                position: absolute;
                right: 30px;
                top: -300px;

                &:hover {
                    background-color: rgba(165, 42, 42, 0.8);
                }

                .jietu {
                    margin-top: 10px;
                    display: flex;
                    align-items: center;
                    justify-content: center;
                }

                .edit {
                    display: flex;
                    align-items: center;
                    justify-content: center;
                }
            }


        }

    }


</style>
