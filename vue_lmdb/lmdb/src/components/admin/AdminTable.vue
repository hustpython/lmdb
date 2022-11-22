<template>

    <!--    <n-card-->
    <!--            title="C盘"-->
    <!--            :bordered="true"-->
    <!--            style="width: 40%;margin-left: 3%;background-color: transparent;border-width: 3px"-->
    <!--            :segmented="{-->
    <!--      content: true }"-->
    <!--    >-->
    <!--        <div ref="testChart" style="height:200px;width:300px;top: 10px"></div>-->
    <!--    </n-card>-->
    <n-card
            title="磁盘使用情况"
            :bordered="true"
            style="width: 80%;margin-left: 3%;background-color: transparent;border-width: 3px"
            :segmented="{
      content: true , footer: 'soft'}"
    >
        <template #header-extra>
            <n-space item-style="display: flex;" align="center">
                <n-checkbox size="small" label="C(磁盘)"/>
                <n-checkbox size="small" label="D(磁盘)"/>
                <n-checkbox size="small" label="E(U盘)"/>

            </n-space>
        </template>

        <n-space justify="space-between">
            剩余
            <n-progress type="circle" :percentage="20"
                        color='#61649f'
                        :rail-color="changeColor('#61649f', { alpha: 0.3 })"
            >
                <span style="text-align: center;font-size: 12px">80% 20G/150G</span>
            </n-progress>
            <n-progress type="circle" :percentage="20">
                <span style="text-align: center;font-size: 12px"> 20G/150G</span>
            </n-progress>
            <n-progress type="circle" :percentage="20">
                <span style="text-align: center;font-size: 12px"> 20G/150G</span>
            </n-progress>
        </n-space>
        <template #footer>

            <n-space vertical>
                视频使用
                <n-progress type="line" :percentage="20">
                    40G / 150G
                </n-progress>
                <n-progress type="line" :percentage="20">
                    40G / 150G
                </n-progress>
                <n-progress type="line" :percentage="20">
                    40G / 150G
                </n-progress>
            </n-space>
        </template>

    </n-card>
    <n-data-table
            :columns="columns"
            :data="data"
            :pagination="pagination"
            :bordered="false"
    />


</template>

<script setup>
    import * as echarts from 'echarts';
    import {onMounted, ref} from 'vue';
    import {changeColor} from "seemly";
    import {useThemeVars} from "naive-ui";

    const testChart = ref(null);

    const themeVars = useThemeVars();


</script>

<style scoped lang="scss">

</style>