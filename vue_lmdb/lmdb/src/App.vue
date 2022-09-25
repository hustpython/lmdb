<template>
    <n-config-provider :locale="zhCN" :date-locale="dateZhCN" :theme="theme">
        <n-layout class="main">
            <n-notification-provider>
                <router-view/>
            </n-notification-provider>
        </n-layout>
    </n-config-provider>
</template>

<script setup>
    import {computed} from "vue";
    import {storeToRefs} from "pinia";
    import {darkTheme, zhCN, dateZhCN} from "naive-ui";
    import {useDarkTheme} from "@/store/themeData";

    const darkThemeStore = useDarkTheme();
    var {themeData} = storeToRefs(darkThemeStore);

    const theme = computed(() => {
        if (themeData.value === true) {
            return darkTheme;
        } else {
            return null;
        }
    });
</script>

<style lang="scss">
    .main {
        @include theme();
    }
</style>
