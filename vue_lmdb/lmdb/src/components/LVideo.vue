<template>
    <n-space>
        <div class="VideoContent">
            <canvas id="localVideoCanvas" style="display: none"></canvas>
            <video

                    id="lvideo"
                    :src=videoUrl
                    autoplay
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
                        class="preControl" :size=controlBtnSize>
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
                        class="nextControl" :size=controlBtnSize>
                    <Next20Regular/>
                </n-icon>


                <!--                显示时间-->
                <div class="timeView"
                     @click="handleTimeTextClick"
                     v-show=!showTimeEditInput
                >{{controlTimeView.playTime}}/{{controlTimeView.duration}}
                </div>

                <!--            评论按钮-->
                <n-icon class="movieComment"
                        :size=controlBtnSize
                        @click="handleCommentBtn"
                >
                    <CommentEdit20Regular/>
                </n-icon>

                <div v-show="CommentContentShow"
                     class="CommentContent">
                <textarea ref="CommentContentRef"
                          v-model=CommentSubmitStr
                          @keyup.enter="handleCommentSubmit"
                          class="CommentInput"/>
                    <div class="CommentSubmit"
                    >
                        <n-checkbox
                                v-model:checked=CommentSubmitPic
                                size="large" label="提交图片"/>
                        <n-icon
                                @click="handleCommentSubmit"
                                class="CommentSend"
                                :size=controlBtnSize>
                            <Send20Regular/>
                        </n-icon>
                    </div>
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

            <!--        展开评论列表按钮-->
            <div v-show=showControl>
                <n-icon v-show=!CommentListShow
                        class="CommentListShowBtn" size="24"
                        @click="handelCommentListShow">
                    <ChevronLeft28Regular/>
                </n-icon>
                <n-icon v-show=CommentListShow
                        class="CommentListShowBtn" size="24"
                        @click="handelCommentListShow">
                    <ChevronRight28Regular/>
                </n-icon>
            </div>


            <div v-show=CommentListShow
                 class="CommentCards">
                <div class="CommentCard"
                     v-for="(item,index) in CommentListData">
                    <div class="CommentDot"
                         :style="getRandomColor(index)">
                    </div>
                    <div class="CommentTimeText"
                         :style="getRandomColor(index)"
                         @dblclick="handleCommentTimeDblclick(index)">
                        {{item.CommentFrameStr}} | {{item.CommentTime}} | {{index+1}} / {{commentListLen}}
                    </div>

                    <n-icon
                            class="CommentDel" size="21"
                            @click="handleCommentDel(index)">
                        <Delete20Regular/>
                    </n-icon>

                    <div class="CommentRight">
                        <img v-show="item.CommentImage!==''" :src="setCoverData(item.CommentImage)"/>
                        {{item.CommentStr}}
                        <div style="clear:both;"></div>
                    </div>
                </div>
            </div>
        </div>
        <div class="RightCommentList">
            <n-card class="CollList"
                    v-show="videoData[routeID.Id].CollStr !== ''"
                    title="合集列表">
                <n-grid
                        x-gap="15" y-gap="15" :cols="4"
                >
                    <n-gi v-for="(item,index) in videoData">
                        <n-tooltip placement="bottom" trigger="hover">
                            <template #trigger>
                                <n-button size="large"
                                          :color="getCollListBtnColor(index)"
                                          @click="handleCollListBtnClick(index)"
                                >
                                    {{index+1}}
                                </n-button>
                            </template>
                            <span> {{item.Title}} </span>
                        </n-tooltip>
                    </n-gi>
                </n-grid>
            </n-card>

            <n-card class="CollList" title="评论列表">
                <template #header-extra>
                    <n-switch :rail-style="railStyle">
                        <template #checked>
                            关闭
                        </template>
                        <template #unchecked>
                            展开
                        </template>
                    </n-switch>
                </template>

            </n-card>

        </div>
    </n-space>

</template>

<script setup>
    import {
        VehicleSubway16Regular, FullScreenMaximize24Filled, Next20Regular,
        Previous20Regular, CommentEdit20Regular, Send20Regular, ChevronLeft28Regular,
        ChevronRight28Regular, Delete20Regular
    } from "@vicons/fluent";
    import {
        VolumeUpFilled,
        PausePresentationOutlined,
        VolumeOffRound,
        LiveTvRound,
    } from "@vicons/material";
    import {Camera} from "@vicons/carbon";
    import {useRouter} from "vue-router";

    const router = useRouter();

    import {getUTCTime, timeFilter, timeStrToSec, getCurrentTime} from "@/api/timefilter";
    import {ref, reactive, nextTick, computed, onBeforeUnmount} from "vue";
    import {UpdateVideo, AddComment, GetComment, DeleteComment} from "@/api/videolist";

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

    const setCoverData = (cover) => {
        if (cover.indexOf(";") != -1) {
            return cover;
        }
        return "data:image/png;base64," + cover;
    };

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

    const CommentListData = ref(
        []
    )

    const handleLoadStart = (e) => {
        GetComment(videoData.value[routeID.Id].MId).then((res) => {
            CommentListData.value = res.data;
        })
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

    const getVideoCoverDataURL = () => {
        const mycanvas = document.getElementById("localVideoCanvas"); // 获取 canvas 对象
        const ctx = mycanvas.getContext("2d"); // 绘制2d
        mycanvas.width = lvideo.clientWidth; // 获取视频宽度
        mycanvas.height = lvideo.clientHeight; //获取视频高度
        ctx.drawImage(lvideo, 0, 0, mycanvas.width, mycanvas.height);
        return mycanvas.toDataURL("image/png"); // 导出图片
    }

    const handleClickClap = () => {
        videoData.value[routeID.Id].Cover = getVideoCoverDataURL();
        let tmpCover = {
            MId: videoData.value[routeID.Id].MId,
            Cover: videoData.value[routeID.Id].Cover.slice(22),
        };
        UpdateVideo(tmpCover);
        delayClearNotifyMsg("为您截取一张图片");
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

    const CommentContentShow = ref(false);
    const CommentContentRef = ref();
    const handleCommentBtn = () => {
        if (!lvideo.paused) {
            playStatus.value = false;
            showControl.value = true;
            lvideo.pause();
        }
        showControl.value = true;
        CommentContentShow.value = !CommentContentShow.value;
        nextTick(() => {
            CommentContentRef.value.focus()
        })
    }

    const CommentListShow = ref(false);

    const handelCommentListShow = () => {
        CommentListShow.value = !CommentListShow.value;
    }

    const CommentSubmitPic = ref(
        true
    );
    const CommentSubmitStr = ref("");

    const CommentSubmitStrMaxLen = 300;
    const handleCommentSubmit = () => {
        if (CommentSubmitStr.value.length <= 0 || CommentSubmitStr.value.length > CommentSubmitStrMaxLen) {
            return
        }
        if (CommentListData.value.length > 15) {
            delayClearNotifyMsg("当前评论数量已经超过15");
            return;
        }
        let tempImage = "";
        if (CommentSubmitPic.value) {
            tempImage = getVideoCoverDataURL();
        }

        let addData = {
            CommentImage: tempImage.slice(22),
            CommentStr: CommentSubmitStr.value,
            CommentFrameStr: timeFilter(lvideo.currentTime),
            CommentFrame: lvideo.currentTime,
            CommentTime: getCurrentTime(),
            MId: videoData.value[routeID.Id].MId,
        };
        try {
            AddComment(addData).then((res) => {
                if (res.code === 200) {
                    CommentListData.value.push(addData);
                }
            })
        } catch (e) {

        }


        CommentSubmitStr.value = "";
        CommentSubmitPic.value = true;
        CommentListShow.value = true;
        CommentContentShow.value = false;
    }
    const randomColor = ["#EF9A9A", "#F48FB1", "#CE93D8", "#9575CD", "#3949AB",
        "#1E88E5", "#4FC3F7", "#4DD0E1", "#26A69A", "#66BB6A", "#9CCC65", "#DCE775",
        "#FDD835", "#FFA000", "#FB8C00", "#F4511E", "#8D6E63", "#78909C", "#BDBDBD"];
    const getRandomColor = (index) => {
        return "background:" + randomColor[index % randomColor.length]
    }
    const handleCommentTimeDblclick = (index) => {
        lvideo.currentTime = CommentListData.value[index].CommentFrame;
    }

    const handleCommentDel = (index) => {
        DeleteComment({
            MId: videoData.value[routeID.Id].MId,
            CommentFrame: CommentListData.value[index].CommentFrame,
        });
        CommentListData.value.splice(index, 1);
    }

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
        // 插入键，截图
        if (e.keyCode == 45) {
            handleClickClap()
        }
        // Ctrl键，添加评论
        if (e.keyCode == 17) {
            handleCommentBtn()
        }
        // Alt + A键 打开/关闭评论列表
        if (e.altKey && e.keyCode == 65) {
            handelCommentListShow()
        }
    }

    const commentListLen = computed(() => {
        return CommentListData.value.length;
    });

    const CollListSelected = ref(routeID.Id);
    const getCollListBtnColor = (index) => {
        if (index == CollListSelected.value) {
            return "#26C6DA";
        }
    }
    const handleCollListBtnClick = (index) => {
        CollListSelected.value = index;
        router.push({name: "#", query: {id: index}});
    }

</script>

<style scoped lang="scss">
    $BtnHoverColor: #FF80AB;
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
                background: $BtnHoverColor;
            }
        }

        $videoPlayCtrlBottom: 4px;


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
                top: $videoPlayCtrlBottom;

                &:hover {
                    color: $BtnHoverColor;
                }
            }

            .nextControl {
                position: absolute;
                left: 90px;
                cursor: pointer;
                top: $videoPlayCtrlBottom;

                &:hover {
                    color: $BtnHoverColor;
                }
            }

            .preControl {
                position: absolute;
                left: 10px;
                cursor: pointer;
                top: $videoPlayCtrlBottom;

                &:hover {
                    color: $BtnHoverColor;
                }
            }

            .timeView {
                position: absolute;
                left: 150px;
                font-size: 15px;
                top: $videoPlayCtrlBottom;
                cursor: text;
                user-select: none;

                &:hover {
                    color: $BtnHoverColor;
                }
            }

            .timeEdit {
                position: absolute;
                left: 150px;
                font-size: 15px;
                top: $videoPlayCtrlBottom;
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
                top: $videoPlayCtrlBottom;
                cursor: pointer;

                &:hover {
                    color: $BtnHoverColor;
                }
            }


            .movieComment {
                position: absolute;
                left: 380px;
                top: $videoPlayCtrlBottom;
                cursor: pointer;

                &:hover {
                    color: $BtnHoverColor;
                }

            }

            $commentBkgColor: rgb(255, 128, 171, .6);
            $commentLeft: 100px;
            $commentBottom: -30px;
            $commentBorder: 15px;
            $commentWidth: 240px;
            $commentHeight: 150px;

            .CommentContent {
                width: $commentWidth;
                height: $commentHeight;
                border: none;
                border-radius: $commentBorder;
                position: absolute;
                left: 280px;
                top: -180px;
                background: $commentBkgColor;
                z-index: 1;

                .CommentInput {
                    background: transparent;
                    border: none;
                    border-radius: $commentBorder;
                    width: $commentWidth - 4;
                    height: $commentHeight - 40;
                    position: relative;
                    top: 0;
                    left: 0;
                    resize: none;
                    outline: none;
                    font-size: 18px;
                    color: #ffffff;
                    font-weight: bold;
                }

                .CommentSubmit {
                    position: absolute;
                    bottom: 6px;
                    left: 30px;

                    .CommentSend {
                        bottom: -6px;
                        cursor: pointer;
                        margin-left: 10px;

                        &:hover {
                            color: $BtnHoverColor;
                        }
                    }
                }

                &:before {
                    content: '';
                    width: 0;
                    height: 0;
                    border: $commentBorder solid;
                    position: absolute;
                    bottom: $commentBottom;
                    left: $commentLeft;
                    border-color: $commentBkgColor transparent transparent;
                }

                &:after {
                    content: '';
                    width: 0;
                    height: 0;
                    border: $commentBorder solid;
                    position: absolute;
                    bottom: $commentBottom;
                    left: $commentLeft;
                    border-color: $commentBkgColor transparent transparent;
                }
            }

            .volumeControl {
                position: absolute;
                right: 120px;
                cursor: pointer;
                top: $videoPlayCtrlBottom;

                &:hover {
                    color: $BtnHoverColor;
                }
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
                right: 10px;
                cursor: pointer;
                top: $videoPlayCtrlBottom;

                &:hover {
                    color: $BtnHoverColor;
                }
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

        .CommentListShowBtn {
            display: flex;
            position: absolute;
            right: 5px;
            bottom: 210px;
            align-items: center;
            height: 80px;
            background-color: rgb(21, 21, 21, .4);
            z-index: 1;

            &:hover {
                color: $BtnHoverColor;
            }
        }

        .CommentCards {
            top: 5px;
            position: absolute;
            width: 360px;
            max-height: 85%;
            right: 0;
            flex-direction: column;
            z-index: 0;
            overflow-y: scroll;

            .CommentCard {
                display: flex;
                overflow: hidden;
                padding-left: 30px;
                background: repeating-linear-gradient(#0000 0 calc(1.2rem - 1px), #66afe1 0 1.2rem) right bottom /100% 100%, linear-gradient(#ff0000 0 0) 20px 0/2px 100% rgb(255, 255, 255, 1);
                background-repeat: no-repeat;
                line-height: 1.2rem;
                -webkit-mask: radial-gradient(circle 0.8rem at 2px 50%, #0000 98%, #000) 0 0/100% 2.4rem;
                flex-direction: column;

                $CommentDotWeight: 14px;

                .CommentDot {
                    position: absolute;
                    left: $CommentDotWeight;
                    width: $CommentDotWeight;
                    height: $CommentDotWeight;
                    border-radius: calc($CommentDotWeight / 2);
                }

                .CommentTimeText {
                    position: relative;
                    left: 0px;
                    width: 100%;
                    height: 20px;
                    display: flex;
                    user-select: none;

                    &:hover {
                        box-shadow: 0 0 2px 2px lightblue;
                    }
                }

                .CommentDel {
                    position: absolute;
                    right: 20px;
                    display: flex;
                    color: floralwhite;

                    &:hover {
                        color: blueviolet;
                    }
                }

                .CommentRight {
                    width: 100%;
                    word-break: break-all;
                    color: black;
                    font-family: 华文行楷;
                    border-radius: 10px;

                    img {
                        float: left;
                        margin: 6px 6px 6px 0px;
                        width: 150px;
                    }
                }
            }
        }

    }

    .RightCommentList {
        margin-left: 30px;
        width: 400px;
        height: 40px;
        background-color: rgb(21, 21, 21, 0.3);

        .collapse {
            position: relative;
            font-size: 18px;
            text-align-all: center;
            align-items: center;
            margin-top: 10px;
        }

        .CollList {
            height: auto;
            background-color: rgb(21, 21, 21, 0.3);
            height: 240px;
            overflow-y: scroll;
            @include theme();
        }
    }
</style>
