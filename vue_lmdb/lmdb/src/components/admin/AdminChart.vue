<template>
    <div class="adminChart">
        <div class="diskStyle">
            <n-card v-for="(item,index) in diskInfos"
                    :title="item.DriveTypeStr"
                    :bordered="true"
                    size="small"
                    class="diskCard"
                    :segmented="{
      content: true , footer: 'soft'}"
            >
                <template #header-extra>
                    <n-button text style="font-size: 24px"
                              @click="handleOpenDir(index)">
                        <n-icon size="24"
                        >
                            <FolderOpen20Regular/>
                        </n-icon>
                    </n-button>


                </template>
                <n-progress type="circle" :percentage="getRestPercent(index)"
                            :color='getDiskUsageColor(index)'
                            style="margin-left: 10%;width: 80%;"
                            :rail-color="changeColor(getDiskUsageColor(index), { alpha: 0.3 })"
                >
                <span style="text-align: center;font-size: 12px">剩余 {{getRestPercent(index)}}%<br>
                    {{item.FreeSpaceStr}}/{{item.SizeStr}}
                </span>
                </n-progress>

                <template #footer>
                <span style="text-align: center;font-size: 12px">
                    视频占用 {{getVideoUsePercent(index)}}%
                </span>
                    <n-progress type="line" :percentage="getVideoUsePercent(index)"
                                :color='getDiskUsageColor(index)'
                                :rail-color="changeColor(getDiskUsageColor(index), { alpha: 0.3 })">
                    <span style="text-align: center;font-size: 12px">
                        {{item.VideoSizeStr}}/{{item.SizeStr}}
                    </span>
                    </n-progress>

                </template>

            </n-card>
        </div>
        <div class="diskPie">
            <div ref="videoCountPie" class="videoCount"></div>
            <div ref="videoSizePie" class="videoCount"></div>
        </div>
    </div>

</template>

<script setup>
    import {onBeforeMount, ref} from 'vue';
    import {changeColor} from "seemly";
    import {FolderOpen20Regular} from '@vicons/fluent';
    import * as echarts from 'echarts';
    import {bytesToSize} from "@/api/timefilter"
    import {GetDiskInfo, OpenCutVideoByPath} from "@/api/videolist";

    const videoCountPie = ref(null);
    const videoSizePie = ref(null);

    function getDiskUsageColor(i) {
        let diskUsageColor = ["#2177b8", "#d276a3", "#41ae3c", "#fecc11", "#f2481b"];
        return diskUsageColor[i % diskUsageColor.length];
    }

    let diskInfos = ref([]);
    let videoDisData = {
        counts: [],
        countTotal: 0,
        sizes: [],
        sizeTotal: 0,
    };

    onBeforeMount(() => {
        GetDiskInfo().then((res) => {
            if (res.code === 200) {
                diskInfos.value = res.data;
                let totalCount = 0;
                let totalSize = 0;
                diskInfos.value.forEach(e => {
                    videoDisData.counts.push({
                        value: e.VideoCount,
                        name: e.DriveTypeStr + ":" + e.VideoCount,
                    })
                    totalCount += e.VideoCount;
                    videoDisData.sizes.push({
                        value: e.VideoSize,
                        name: e.DriveTypeStr + ":" + e.VideoSizeStr,
                    })
                    totalSize += e.VideoSize;
                })
                videoDisData.sizeTotal = totalSize;
                videoDisData.countTotal = totalCount;
                setVideoPie();
                setVideoSize();
            }
        })
    })

    const setVideoSize = () => {
        const videoSizePieInst = echarts.init(videoSizePie.value);
        videoSizePieInst.setOption({
            title: {
                text: '视频大小分布',
                subtext: "总:" + bytesToSize(videoDisData.sizeTotal),
                left: 'center',
                textStyle: {
                    color: "#fff"
                },
                subtextStyle: {
                    color: "#fff"
                }
            },
            tooltip: {
                trigger: 'item'
            },
            legend: {
                bottom: 'bottom',
                textStyle: {
                    color: "#fff"
                }
            },
            series: [
                {
                    name: '大小',
                    type: 'pie',
                    radius: '50%',
                    data: videoDisData.sizes,
                    emphasis: {
                        itemStyle: {
                            shadowBlur: 10,
                            shadowOffsetX: 0,
                            shadowColor: 'rgba(0, 0, 0, 0.5)'
                        }
                    }
                }
            ]
        })
    }

    const setVideoPie = () => {
        const videoCountPieInst = echarts.init(videoCountPie.value);
        videoCountPieInst.setOption({
            title: {
                text: '视频数量分布',
                subtext: "总:" + videoDisData.countTotal,
                left: 'center',
                textStyle: {
                    color: "#fff"
                },
                subtextStyle: {
                    color: "#fff"
                }
            },
            tooltip: {
                trigger: 'item'
            },
            legend: {
                bottom: 'bottom',
                textStyle: {
                    color: "#fff"
                }
            },
            series: [
                {
                    name: '数量',
                    type: 'pie',
                    radius: '50%',
                    data: videoDisData.counts,
                    emphasis: {
                        itemStyle: {
                            shadowBlur: 10,
                            shadowOffsetX: 0,
                            shadowColor: 'rgba(0, 0, 0, 0.5)'
                        }
                    }
                }
            ]
        })
    };

    const getRestPercent = (i) => {
        let s = diskInfos.value[i].FreeSpace / diskInfos.value[i].Size * 100;
        return s.toFixed(1);
    }
    const getVideoUsePercent = (i) => {
        let s = diskInfos.value[i].VideoSize / diskInfos.value[i].Size * 100;
        return s.toFixed(1);
    }
    const handleOpenDir = (i) => {
        OpenCutVideoByPath(diskInfos.value[i].DeviceID + ":");
    }


</script>

<style scoped lang="scss">
    .adminChart {
        width: 100%;
        height: 100%;
        display: block;

        .diskStyle {
            display: flex;
            flex-flow: row wrap;
            height: auto;

            .diskCard {
                width: 30%;
                margin-left: 2.5%;
                margin-top: 10px;
                background-color: transparent;
                border-width: 1.8px;
            }
        }

        .diskPie {
            width: 100%;
            display: flex;
            margin-top: 3%;

            .videoCount {
                height: 300px;
                width: 350px;
                margin-left: 5%;
            }
        }
    }
</style>