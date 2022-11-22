<template>
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
</template>

<script setup>
    import {onBeforeMount, ref} from 'vue';
    import {changeColor} from "seemly";
    import {FolderOpen20Regular} from '@vicons/fluent';
    import {GetDiskInfo, OpenCutVideoByPath} from "@/api/videolist";

    function getDiskUsageColor(i) {
        let diskUsageColor = ["#2177b8", "#d276a3", "#41ae3c", "#fecc11", "#f2481b"];
        return diskUsageColor[i % diskUsageColor.length];
    }

    let diskInfos = ref([]);
    onBeforeMount(() => {
        GetDiskInfo().then((res) => {
            if (res.code === 200) {
                diskInfos.value = res.data;
            }
        })
    })
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
    .diskStyle {
        display: flex;
        flex-flow: row wrap;
        height: 100%;

        .diskCard {
            width: 30%;
            margin-left: 2.5%;
            margin-top: 10px;
            background-color: transparent;
            border-width: 1.8px;
        }
    }
</style>