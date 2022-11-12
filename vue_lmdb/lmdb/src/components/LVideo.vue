<template>
    <n-space justify="space-between">
        <n-space vertical>
            <div class="VideoContent"
                 ref="VideoContentRef">
                <canvas id="localVideoCanvas" style="display: none"></canvas>
                <video
                        id="lvideo"
                        ref="videoCon"
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
                <!--                 剪切控制界面-->
                <div class="cutMask"
                     v-show="cutMaskShow">

                    <div class="upMask"
                         @click="handleCutMaskClick"
                    >
                        <span class="cutDuration">

                        </span>
                    </div>
                    <div class="lowMask">
                        <div class="leftCut"
                             @mousedown="handleNeedleMove(0)"
                             ref="leftNeedleRef">
                            <n-icon size="24"
                            >
                                <WaterSharp/>
                            </n-icon>
                        </div>
                        <div class="rightCut"
                             @mousedown="handleNeedleMove(1)">
                            <n-icon size="24">
                                <WaterSharp/>
                            </n-icon>
                        </div>

                        <n-space justify="space-around">
                            <n-button type="tertiary" size="tiny"
                                      @click="handleTimeModify">
                                {{cutTimeModLabel}}
                            </n-button>
                            <div v-show="getLeftCutTimeShow"
                            >
                                {{getLeftCutTime}}
                            </div>

                            <input v-show="!getLeftCutTimeShow"
                                   class="cutTimeInput"
                                   type="text"
                                   v-model=cutTimeInputModel.left
                            >

                            <div
                                    v-show="getRightCutTimeShow"
                            >
                                {{getRightCutTime}}
                            </div>

                            <input
                                    v-show="!getRightCutTimeShow"
                                    class="cutTimeInput"
                                    type="text"
                                    v-model=cutTimeInputModel.right>

                            <div class="currentTimeShow">
                                截取时长 {{cutPlayTime}} / {{cutDuration}}
                            </div>
                            <n-button type="tertiary" size="tiny"
                                      @click="handleCutPlay">
                                {{cutPlayLabel}}
                            </n-button>
                            <n-button type="tertiary" size="tiny"
                                      @click="handleCutSubmit"
                            >
                                确定
                            </n-button>
                            <n-button type="tertiary" size="tiny"
                                      @click="handleCutCancle"
                            >
                                取消
                            </n-button>
                        </n-space>
                    </div>


                </div>
                <!--            下方控制台-->
                <div class="progressMask"
                     @click="handleProgress"
                     v-show="showControl && !cutMaskShow"
                >
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
                     v-show="showControl && !cutMaskShow">
                    <n-icon
                            v-show="currentData.CollStr !== ''"
                            @click="handlePlayPre"
                            @mouseenter="handlePlayPreEnter"
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
                            v-show="currentData.CollStr !== ''"
                            @click="handlePlayNext"
                            @mouseenter="handlePlayNextEnter"
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
                    <!--                     视频剪切-->
                    <n-icon
                            class="cutControl" :size=controlBtnSize
                            @click="handleCutBtn">

                        <Cut20Regular/>
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
                         v-for="(item,index) in CommentListData[CollListSelected]">
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
            <div class="LeftDownForm">
                <n-space style="height: 360px;margin-top:30px">
                    <!-- 左侧图片 -->
                    <img class="videoCardSize" :src="setCoverData(currentData.Cover)"/>

                    <div v-show="!showEditForm">
                        <n-space vertical style="width: 320px">
                            <div>
            <span style="font-size: 18px; float: left"
            >{{ currentData.Title }}
              <n-button
                      v-show="currentData.CollStr !== ''"
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
                                <div v-for="(item, index) in currentData.TagArray">
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
                                {{ currentData.Desc }}
                            </n-ellipsis>
                        </n-space>

                    </div>

                    <n-form
                            v-show="showEditForm"
                            class="middleInfo"
                            ref="formInstRef"
                            style="width: 320px"
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

            </div>
        </n-space>
        <div class="RightCommentList">
            <n-card class="CollList"
                    v-show="currentData.CollStr !== ''"
                    title="合集列表">
                <template #header-extra>
                    <n-switch
                            v-model:value="CollListSwitchList"
                            :rail-style="railStyle"
                    >
                        <template #checked>
                            连播
                        </template>
                        <template #unchecked>
                            单播
                        </template>
                    </n-switch>
                </template>
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
                    <n-switch v-model:value="CommTableSwitch"
                              :rail-style="railStyle">
                        <template #checked>
                            展开
                        </template>
                        <template #unchecked>
                            关闭
                        </template>
                    </n-switch>
                </template>

                <table class="CommentTable" v-show="CommTableSwitch">
                    <tbody v-for="(item,index) in CommentListData">
                    <tr v-for="d in item">
                        <td width="25%">{{d.CommentTime}}</td>
                        <td width="60%">{{d.CommentStr}}</td>
                        <td width="15%">
                            <n-tooltip trigger="hover">
                                <template #trigger>
                                    <n-button
                                            @click="handleCommentTableClick(index,d.CommentFrame)"
                                    >跳转
                                    </n-button>
                                </template>
                                跳转至{{d.Title}}{{d.CommentFrameStr}}
                            </n-tooltip>
                        </td>
                    </tr>

                    </tbody>

                </table>

            </n-card>

            <n-card class="CollList" title="精彩片段">


            </n-card>

        </div>
    </n-space>

</template>

<script setup>
    import {
        VehicleSubway16Regular, FullScreenMaximize24Filled, Next20Regular,
        Previous20Regular, CommentEdit20Regular, Send20Regular, ChevronLeft28Regular,
        ChevronRight28Regular, Delete20Regular, Cut20Regular
    } from "@vicons/fluent";
    import {Edit, FolderDetails, Favorite, Camera} from "@vicons/carbon";
    import {
        VolumeUpFilled,
        PausePresentationOutlined,
        VolumeOffRound,
        LiveTvRound,
    } from "@vicons/material";
    import {WaterSharp} from "@vicons/ionicons5"
    import {useRouter} from "vue-router";

    const router = useRouter();

    import {getUTCTime, timeFilter, timeStrToSec, getCurrentTime, percent2Point} from "@/api/timefilter";
    import {ref, reactive, nextTick, computed, onBeforeUnmount, onBeforeMount} from "vue";
    import {
        UpdateVideo,
        AddComment,
        GetComment,
        GetCommentByColl,
        DeleteComment,
        GetAllColl,
        DeleteMovieColl,
        CutVideoByMId
    } from "@/api/videolist";

    import {storeToRefs} from "pinia";
    import {useVideoData} from "@/store/videoData";

    const videoDataStore = useVideoData();
    var {videoData} = storeToRefs(videoDataStore);

    const routeID = defineProps(["Id"]);
    const currentData = ref(videoData.value[routeID.Id]);

    const videoUrl = computed(() => {
        return config.SERVER_API + currentData.value.VideoUrl;
    });

    const getVideoWidth = () => {
        if (document.body.clientWidth > 1000) {
            return "668px"
        } else {
            return document.body.clientWidth + "px";
        }
    }
    const getVideoHeight = () => {
        if (document.body.clientWidth > 1000) {
            return "376px"
        } else {
            return document.body.clientWidth * 0.56 + "px";
        }
    }
    let isFull = false;

    window.onresize = () => {
        if (document.body.clientWidth > 1000) {
            videoCon.value.controls = false;
            isMobile = false;
        } else {
            videoCon.value.controls = true;
            isMobile = true;
        }
        if (isFull) {
            videoWidth.value = "100%";
        } else {
            videoWidth.value = getVideoWidth();
            videoHeight.value = getVideoHeight();
        }
    };

    let lvideo = {};

    const controlBtnSize = "24px";

    const progressBtnLeft = ref("0");
    const mouseLeft = ref("0px");
    const videoWidth = ref(getVideoWidth());
    const videoHeight = ref(getVideoHeight());
    const showControl = ref(true);
    const showTimeEditInput = ref(false);
    const controlNotifyShow = ref(false);
    // 剪切模式
    let VideoContentRef = ref();

    const cutWaterColorLow = "rgb(39, 117, 182, 0.6)";
    const cutWaterColor = "rgb(39, 117, 182)";

    const leftNeedleColor = ref(cutWaterColor);
    const rightNeedleColor = ref(cutWaterColorLow);
    const leftNeedlePosi = ref(0);
    const rightNeedlePosi = ref("100%");

    const isLeftNeedleActive = ref(true);
    const cutDuration = ref(0);
    const cutPlayTime = ref("00:00:00");
    const cutMaskShow = ref(false);

    const cutPlayLabel = ref("开始预览");

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
        let bfb = e.offsetX / e.target.clientWidth;
        progressBtnLeft.value = bfb * 100 + '%';
        lvideo.currentTime = bfb * lvideo.duration;
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
    const videoCon = ref();


    const handleLoadStart = (e) => {
        let RecentWatch = {
            MId: currentData.value.MId,
            RecentWatch: getUTCTime(),
        };
        UpdateVideo(RecentWatch);

        if (isMobile) {
            videoCon.value.controls = true;
        }
        playStatus.value = true;

        if (CommentListData.value.length == 0) {
            if (currentData.value.CollStr == "") {
                GetComment(currentData.value.MId).then((res) => {
                    CommentListData.value = [res.data];
                })
            } else {
                GetCommentByColl(currentData.value.CollStr).then((res) => {
                    CommentListData.value = res.data;
                })
            }
        }

        lvideo = document.getElementById("lvideo");
        controlTimeView.duration = timeFilter(lvideo.duration);
        if (currentData.value.LastWatch > 0) {
            e.target.currentTime = currentData.value.LastWatch;
            delayClearNotifyMsg("为您定位至:" + timeFilter(currentData.value.LastWatch));

        }
        if (currentData.value.Duration.length !== 0) {
            return;
        }
        let updateTime = {
            MId: currentData.value.MId,
            Duration: timeFilter(e.target.duration),
        };
        UpdateVideo(updateTime);
    };

    const playStatus = ref(true);
    const handlePlay = () => {
        if (cutMaskShow.value) {
            return;
        }
        playStatus.value = !playStatus.value;
        if (isMobile) {
            return
        }
        if (lvideo.paused) {
            lvideo.play();
            showControl.value = false;
        } else {
            showControl.value = true;
            lvideo.pause();
        }
    };

    const CollListSwitchList = ref(true);
    const CommTableSwitch = ref(false);
    const handleEnd = () => {
        playStatus.value = false;
        if (!CollListSwitchList.value) {
            return;
        }
        handlePlayNextEnter();
        setTimeout(function () {
            handlePlayNext();
        }, 3000)
    }


    const handleTimeProgress = (e) => {
        if (isMobile) {
            return;
        }
        progressBtnLeft.value = e.target.currentTime / e.target.duration * 100 + '%';
        controlTimeView.playTime = timeFilter(e.target.currentTime);
        if (cutMaskShow.value && !lvideo.paused) {
            cutPlayTime.value = timeFilter(e.target.currentTime -
                percent2Point(leftNeedlePosi.value) * e.target.duration);
            if (e.target.currentTime / e.target.duration > percent2Point(rightNeedlePosi.value)) {
                lvideo.pause();
                cutPlayLabel.value = "开始预览";
                playStatus.value = false;
            }
        }
    }


    const handleFull = () => {
        if (!document.fullscreenElement) {
            isFull = true;
            document.querySelector('.VideoContent').requestFullscreen();
            videoWidth.value = "100%";
        } else {
            isFull = false;
            document.exitFullscreen();
            videoWidth.value = getVideoWidth();
        }
    };

    document.addEventListener('fullscreenchange', () => {
        if (!document.fullscreenElement) {
            isFull = false;
            videoWidth.value = getVideoWidth();
        }
    })

    const handleMouseMove = () => {
        if (isMobile) {
            return
        }
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
        videoData.value[CollListSelected.value].Cover = getVideoCoverDataURL();
        let tmpCover = {
            MId: currentData.value.MId,
            Cover: videoData.value[CollListSelected.value].Cover.slice(22),
        };
        UpdateVideo(tmpCover);
        delayClearNotifyMsg("为您截取一张图片");
    };

    onBeforeUnmount(() => {
        handlePause();
    });

    const handlePause = () => {
        let updateTime = {
            MId: currentData.value.MId,
            RecentWatch: getUTCTime(),
            LastWatch: Math.floor(lvideo.currentTime),
        };
        videoData.value[CollListSelected.value].LastWatch = Math.floor(lvideo.currentTime);
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
        if (CommentListData.value[CollListSelected.value].length > 15) {
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
            MId: currentData.value.MId,
        };
        try {
            AddComment(addData).then((res) => {
                if (res.code === 200) {
                    CommentListData.value[CollListSelected.value].push(addData);
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
        lvideo.currentTime = CommentListData.value[CollListSelected.value][index].CommentFrame;
    }

    const handleCommentDel = (index) => {
        DeleteComment({
            MId: currentData.value.MId,
            CommentFrame: CommentListData.value[CollListSelected.value][index].CommentFrame,
        });
        CommentListData.value[CollListSelected.value].splice(index, 1);
    }

    document.body.onkeydown = function (e) {
        if (cutMaskShow.value) {
            return;
        }
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
        return CommentListData.value[CollListSelected.value].length;
    });

    const CollListSelected = ref(routeID.Id);

    if (currentData.value.CollStr == "") {
        CollListSelected.value = 0;
    }

    const getCollListBtnColor = (index) => {
        if (index == CollListSelected.value) {
            return "#26C6DA";
        }
    }
    const handleCollListBtnClick = (index) => {
        CollListSelected.value = index;
        currentData.value = videoData.value[index];
        updateVideoFormModel();
    }

    const handlePlayPre = () => {
        let temIndex = parseInt(CollListSelected.value) - 1;
        if (temIndex < 0) {
            return
        }
        CollListSelected.value = temIndex;
        currentData.value = videoData.value[temIndex];
        updateVideoFormModel();
    }

    const handlePlayNext = () => {
        let temIndex = parseInt(CollListSelected.value) + 1;
        if (temIndex >= videoData.value.length) {
            return
        }
        CollListSelected.value = temIndex;
        currentData.value = videoData.value[temIndex];
        updateVideoFormModel();
    }

    const handlePlayPreEnter = () => {
        let temIndex = parseInt(CollListSelected.value) - 1;
        if (temIndex >= videoData.value.length) {
            return
        }
        delayClearNotifyMsg("上一个:" + videoData.value[temIndex].Title);
    }

    const handlePlayNextEnter = () => {
        if (currentData.value.CollStr === '') {
            return;
        }
        let temIndex = parseInt(CollListSelected.value) + 1;
        if (temIndex >= videoData.value.length) {
            return
        }
        delayClearNotifyMsg("下一个:" + videoData.value[temIndex].Title);
    }


    var showEditForm = ref(false);
    const formInstRef = ref(null);
    const labelTypes = ["success", "warning", "error", "info"];
    var videoFormModel = ref({
        Title: currentData.value.Title,
        TagArray: currentData.value.TagArray,
        Desc: currentData.value.Desc,
        CollStr: currentData.value.CollStr,
    });

    const updateVideoFormModel = () => {
        videoFormModel.value.Title = currentData.value.Title
        videoFormModel.value.TagArray = currentData.value.TagArray
        videoFormModel.value.Desc = currentData.value.Desc
        videoFormModel.value.CollStr = currentData.value.CollStr
    }
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

    let collOptions = ref([{label: "", value: ""}]);
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
    let isMobile = false;
    onBeforeMount(() => {
        if (document.body.clientWidth <= 1000) {
            isMobile = true;
        }
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


    const handleRecover = () => {
        videoFormModel.value.Title = currentData.value.Title;
        videoFormModel.value.TagArray = [];
        videoFormModel.value.TagArray.push(...currentData.value.TagArray);
        videoFormModel.value.Desc = currentData.value.Desc;
        videoFormModel.value.CollStr = currentData.value.CollStr;
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
            videoFormModel.value.MId = currentData.value.MId;
            UpdateVideo(videoFormModel.value);
            videoData.value[CollListSelected.value].Title = videoFormModel.value.Title;
            videoData.value[CollListSelected.value].Desc = videoFormModel.value.Desc;
            videoData.value[CollListSelected.value].TagArray = videoFormModel.value.TagArray;
            videoData.value[CollListSelected.value].CollStr = videoFormModel.value.CollStr;
            if (videoFormModel.value.CollStr === "") {
                DeleteMovieColl(videoFormModel.value.MId);
            }
            showEditForm.value = false;
        }

        setTimeout(procForm, 100);
    };

    const handleCommentTableClick = (index, t) => {
        if (CollListSelected.value != index) {
            CollListSelected.value = index
            currentData.value = videoData.value[index];
            updateVideoFormModel();
        }
        setTimeout(function () {
            lvideo.currentTime = t;
        }, 1000)
    }

    // 全部转成百分比

    const handleCutPlay = () => {
        cutPlayTime.value = "00:00:00";
        if (lvideo.paused) {
            lvideo.currentTime = percent2Point(leftNeedlePosi.value) * lvideo.duration;
            lvideo.play();
            cutPlayLabel.value = "预览暂停"
            playStatus.value = true;
        } else {
            lvideo.pause();
            cutPlayLabel.value = "开始预览";
            playStatus.value = false;
        }
    }

    const handleCutBtn = () => {
        cutPlayTime.value = "00:00:00";
        showControl.value = false;
        cutMaskShow.value = true;
        playStatus.value = false;
        if (leftNeedlePosi.value === 0) {
            leftNeedlePosi.value = progressBtnLeft.value;
        }
        cutDuration.value = timeFilter((percent2Point(rightNeedlePosi.value) - percent2Point(leftNeedlePosi.value)) * lvideo.duration);
        getLeftCutTime.value = timeFilter(percent2Point(leftNeedlePosi.value) * lvideo.duration);
        getRightCutTime.value = timeFilter(percent2Point(rightNeedlePosi.value) * lvideo.duration);
        if (lvideo.play()) {
            lvideo.pause();
        }
    }

    const handleCutCancle = () => {
        cutMaskShow.value = false;
    }

    const getLeftCutTime = ref("0");
    const getRightCutTime = ref("0");

    const getLeftCutTimeShow = ref(true);
    const getRightCutTimeShow = ref(true);

    const handleCutMaskClick = (e) => {
        cutPlayTime.value = "00:00:00";
        if (lvideo.play()) {
            lvideo.pause();
            cutPlayLabel.value = "开始预览";
        }
        let point = (e.clientX - VideoContentRef.value.offsetLeft) /
            VideoContentRef.value.clientWidth;

        if (isLeftNeedleActive.value) {
            if (percent2Point(rightNeedlePosi.value) - point < 0) {
                point = percent2Point(rightNeedlePosi.value);
            }
            leftNeedlePosi.value = point * 100 + '%';
        } else {
            if (point - percent2Point(leftNeedlePosi.value) < 0) {
                point = percent2Point(leftNeedlePosi.value);
            }
            rightNeedlePosi.value = point * 100 + '%';
        }
        lvideo.currentTime = lvideo.duration * point;
        cutDuration.value = timeFilter((percent2Point(rightNeedlePosi.value) - percent2Point(leftNeedlePosi.value)) * lvideo.duration);
        getLeftCutTime.value = timeFilter(percent2Point(leftNeedlePosi.value) * lvideo.duration);
        getRightCutTime.value = timeFilter(percent2Point(rightNeedlePosi.value) * lvideo.duration);
    }

    const handleNeedleMove = (key, e) => {
        cutPlayTime.value = "00:00:00";
        if (lvideo.play()) {
            lvideo.pause();
            cutPlayLabel.value = "开始预览";
        }
        let point = 0;
        document.onmousemove = (e) => {
            point = (e.clientX - VideoContentRef.value.offsetLeft) /
                VideoContentRef.value.clientWidth;
            if (point < 0) {
                point = 0;
            } else if (point > 1) {
                point = 1;
            }
            switch (key) {
                case 0:
                    leftNeedleColor.value = cutWaterColor;
                    rightNeedleColor.value = cutWaterColorLow;
                    isLeftNeedleActive.value = true;
                    if (percent2Point(rightNeedlePosi.value) - point < 0) {
                        return
                    }
                    leftNeedlePosi.value = point * 100 + '%';
                    break;
                case 1:
                    rightNeedleColor.value = cutWaterColor;
                    leftNeedleColor.value = cutWaterColorLow;
                    isLeftNeedleActive.value = false;
                    if (point - percent2Point(leftNeedlePosi.value) < 0) {
                        return
                    }
                    rightNeedlePosi.value = point * 100 + '%';
                    break;
            }
            lvideo.currentTime = lvideo.duration * point;
            cutDuration.value = timeFilter((percent2Point(rightNeedlePosi.value) - percent2Point(leftNeedlePosi.value)) * lvideo.duration);
            getLeftCutTime.value = timeFilter(percent2Point(leftNeedlePosi.value) * lvideo.duration);
            getRightCutTime.value = timeFilter(percent2Point(rightNeedlePosi.value) * lvideo.duration);
        };
        document.onmouseup = function () {
            document.onmousemove = document.onmouseup = null;
        };

        return false;
    }

    const cutTimeModLabel = ref("手动调整");
    const cutTimeInputModel = ref(
        {
            left: "",
            right: ""
        }
    )
    const handleTimeModify = () => {
        if (getLeftCutTimeShow.value) {
            cutTimeModLabel.value = "确认调整"
            cutTimeInputModel.value.left = getLeftCutTime.value;
            cutTimeInputModel.value.right = getRightCutTime.value;
        } else {
            cutTimeModLabel.value = "手动调整"
            let t1 = timeStrToSec(cutTimeInputModel.value.left);
            let t2 = timeStrToSec(cutTimeInputModel.value.right);
            if (t2 === getRightCutTime.value && t1 === getLeftCutTime.value) {
                return;
            }
            if (0 < t1 < lvideo.duration && 0 < t2 < lvideo.duration && t1 < t2) {
                let point = t1 / lvideo.duration;
                lvideo.currentTime = lvideo.duration * point;
                leftNeedlePosi.value = point * 100 + '%';
                let point2 = t2 / lvideo.duration;
                rightNeedlePosi.value = point2 * 100 + '%';
                cutDuration.value = timeFilter(Math.ceil((percent2Point(rightNeedlePosi.value) - percent2Point(leftNeedlePosi.value)) * lvideo.duration));
                getLeftCutTime.value = timeFilter(percent2Point(leftNeedlePosi.value) * lvideo.duration);
                getRightCutTime.value = timeFilter(percent2Point(rightNeedlePosi.value) * lvideo.duration);
            }
        }
        getLeftCutTimeShow.value = !getLeftCutTimeShow.value;
        getRightCutTimeShow.value = !getRightCutTimeShow.value;

    }

    const handleCutSubmit = () => {
        let cutDurationSec = timeStrToSec(cutDuration.value);
        if (cutDurationSec === 0) {
            return;
        }
        let data = {
            start: timeStrToSec(getLeftCutTime.value),
            duration: cutDurationSec,
            mid: currentData.value.MId
        }
        CutVideoByMId(data);
        cutMaskShow.value = false;
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

        $MaskHeight: 70px;
        $NeedleMaskHeight: 10px;
        $LowMaskHeight: 60px;
        $UpMaskHeight: 4px;
        $LeftRightNeedleColor: #63bbd0;

        .cutMask {
            width: v-bind(videoWidth);
            position: absolute;
            display: flex;
            height: $MaskHeight;
            bottom: 0;
            $cutWaterColorLow: rgb(39, 117, 182, 0.6);
            $cutWaterColor: rgb(39, 117, 182);
            $waterIconOffset: 12px;

            .upMask {
                height: $UpMaskHeight;
                width: inherit;
                position: absolute;
                display: flex;
                background: rgb(17, 101, 154, 0.3);
                bottom: $LowMaskHeight;
                cursor: pointer;

                .cutDuration {
                    display: block;
                    position: absolute;
                    left: v-bind(leftNeedlePosi);
                    width: calc(v-bind(rightNeedlePosi) - v-bind(leftNeedlePosi));
                    height: inherit;
                    background: $cutWaterColor;
                }
            }


            .lowMask {
                height: $LowMaskHeight;
                position: absolute;
                display: flex;
                width: inherit;
                background: #1c2938;
                bottom: 0;

                .cutTimeInput {
                    border: none;
                    width: 70px;
                    outline: medium;
                    cursor: text;
                    background: rgb(189, 189, 189, 0.3);
                    user-select: none;
                }

                .leftCut {
                    left: calc(v-bind(leftNeedlePosi) - $waterIconOffset);
                    display: block;
                    position: absolute;
                    cursor: pointer;

                    .n-icon {
                        display: block;
                        color: v-bind(leftNeedleColor);
                        transition: color .5s;

                        &:hover {
                            color: $cutWaterColor;
                        }
                    }
                }

                .rightCut {
                    left: calc(v-bind(rightNeedlePosi) - $waterIconOffset);
                    display: block;
                    position: absolute;
                    cursor: pointer;

                    .n-icon {
                        display: block;
                        color: v-bind(rightNeedleColor);
                        transform: rotateY(180deg);
                        transition: color .5s;

                        &:hover {
                            color: $cutWaterColor;
                        }
                    }
                }

                .n-space {
                    display: block;
                    position: absolute;
                    top: 30px;
                }
            }

        }

        .progressMask {
            @include phone() {
                display: none;
            }
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
                position: relative;
                height: 2px;
                width: v-bind(progressBtnLeft);
                background: $BtnHoverColor;
            }
        }

        $videoPlayCtrlBottom: 4px;


        .videoPlayCtrl {
            @include phone() {
                display: none;
            }
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

            .cutControl {
                position: absolute;
                right: 190px;
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
            @include phone() {
                display: none;
            }
            display: flex;
            position: absolute;
            right: 10px;
            bottom: 30%;
        }

        .CommentListShowBtn {
            @include phone() {
                display: none;
            }
            display: flex;
            position: absolute;
            right: 5px;
            bottom: 60%;
            align-items: center;
            height: 80px;
            background-color: rgb(21, 21, 21, .4);
            z-index: 1;

            &:hover {
                color: $BtnHoverColor;
            }
        }

        .CommentCards {
            @include phone() {
                display: none;
            }
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

    .LeftDownForm {
        @include phone() {
            display: none;
        }
    }

    .RightCommentList {
        margin-left: 30px;
        width: 400px;
        max-height: 1000px;
        background-color: rgb(21, 21, 21, 0.3);
        @include phone() {
            width: 350px;
            margin-left: 10px;
        }

        .collapse {
            position: relative;
            font-size: 18px;
            text-align-all: center;
            align-items: center;
            margin-top: 10px;
        }

        .CollList {
            max-height: 600px;
            background-color: rgb(21, 21, 21, 0.3);
            overflow: auto;
            @include theme();
        }
    }

    .CommentTable, td, th {
        @include theme();
        border: 1px solid $lightPink;
        border-collapse: collapse;
    }
</style>
