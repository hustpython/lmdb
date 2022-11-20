<template>
    <n-layout-header class="nav" bordered>
        <div class="ui-logo">
            <img :src="logourl"
            />
            <a href="/">
                <n-gradient-text :size="24" type="success"> LMDB</n-gradient-text>
            </a>
        </div>

        <div class="nav-select">
            <n-select
                    class="search"
                    filterable
                    :options="selectOptions"
                    :on-update:value="handleSeleteed"
                    placeholder="搜索电影、剧集、影评..."
            />
        </div>

        <div class="nav-end">

            <n-popover trigger="hover">
                <template #trigger>
                    <div class="notification">
                        <n-icon
                                size="27px"
                        >
                            <CircleNotificationsRound/>
                        </n-icon>
                        <span>
                        消息
                       </span>
                    </div>
                </template>
                <n-space vertical>
                    <n-card title="小卡片" size="small">
                        卡片内容
                    </n-card>
                    <n-card title="中卡片" size="medium">
                        卡片内容
                    </n-card>
                    <n-card title="大卡片" size="large">
                        卡片内容
                    </n-card>
                    <n-card title="超大卡片" size="huge">
                        卡片内容
                    </n-card>
                </n-space>
            </n-popover>

            <span style="width: 0.5px;height: 51%;background-color: hsla(0, 0%, 100%, .4);margin-right: 20px">

            </span>

            <n-popover trigger="hover">
                <template #trigger>
                    <img style="height: 32px; width: 32px; border-radius: 16px" :src="avatarurl"/>
                </template>
                <n-space>
                    <n-card title="个人中心" size="small">
                        <n-space vertical>
                            <n-button @click="handleThemeClick">
                                <template #icon>
                                    <n-icon>
                                        <SkinOutlined/>
                                    </n-icon>
                                </template>
                                暗色主题
                            </n-button>
                            <n-button>
                                <template #icon>
                                    <n-icon>
                                        <UserProfile/>
                                    </n-icon>
                                </template>
                                个人资料
                            </n-button>
                            <n-button @click="handleAdminRoute">
                                <template #icon>
                                    <n-icon>
                                        <DatabasePerson24Filled/>
                                    </n-icon>
                                </template>
                                数据管理
                            </n-button>
                        </n-space>

                    </n-card>
                    <n-card title="创作中心" size="small">
                        <n-space vertical>
                            <n-button @click="handleThemeClick">
                                <template #icon>
                                    <n-icon>
                                        <Cut20Regular/>
                                    </n-icon>
                                </template>
                                剪切列表
                            </n-button>
                            <n-button>
                                <template #icon>
                                    <n-icon>
                                        <Edit/>
                                    </n-icon>
                                </template>
                                影评创作
                            </n-button>
                            <n-button>
                                <template #icon>
                                    <n-icon>
                                        <PictureOutlined/>
                                    </n-icon>
                                </template>
                                封面推荐
                            </n-button>
                        </n-space>
                    </n-card>
                </n-space>
            </n-popover>
        </div>
    </n-layout-header>
</template>

<script setup>
    import {reactive, onBeforeMount, computed} from "vue";

    import {storeToRefs} from "pinia";
    import {useVideoData} from "@/store/videoData";
    import {useRouter} from "vue-router";
    import {GetMoviesByColl} from "@/api/videolist";

    import {useDarkTheme} from "@/store/themeData";
    import {CircleNotificationsRound, CreateNewFolderOutlined} from "@vicons/material";
    import {SkinOutlined, PictureOutlined} from "@vicons/antd";
    import {Edit, UserProfile} from "@vicons/carbon";
    import {DatabasePerson24Filled, Cut20Regular} from "@vicons/fluent"

    const darkThemeStore = useDarkTheme();

    // 默认深色主题
    let themeBtnText = reactive({text: "深色"});
    onBeforeMount(() => {
        if (darkThemeStore.themeData === true) {
            themeBtnText.text = "浅色";
            setAttribute("deep");
        } else {
            setAttribute("light");
        }
    });

    const setAttribute = (theme) => {
        window.document.documentElement.setAttribute("data-theme", theme);
    };

    const logourl = require("../assets/naivelog.svg");
    const avatarurl = require("../assets/axu.png");
    const handleThemeClick = () => {
        if (themeBtnText.text === "浅色") {
            themeBtnText.text = "深色";
            setAttribute("light");
            darkThemeStore.unsetDarkTheme();
        } else {
            themeBtnText.text = "浅色";
            setAttribute("deep");
            darkThemeStore.setDarkTheme();
        }
    };
    window.addEventListener(
        "scroll",
        () => {
            if (scrollY > 63) {
                document.body.style.setProperty("--headerTopScroll", "-63px");
            } else {
                document.body.style.setProperty("--headerTopScroll", "0px");
            }
        },
        true
    );


    const selectOptions = computed(() => {
        let tmpOptions = [];
        for (var i = 0; i < videoData.value.length; i++) {
            tmpOptions.push({
                label: videoData.value[i].Title,
                value: i,
            });
        }
        return tmpOptions;
    });

    const videoDataStore = useVideoData();
    var {videoData} = storeToRefs(videoDataStore);

    const router = useRouter();
    const handleSeleteed = (value) => {
        let temMId = videoData.value[value].MId
        if (videoData.value[value].CollStr.length != 0) {
            GetMoviesByColl(videoData.value[value].CollStr).then((res) => {
                if (res.code == 200) {
                    videoDataStore.setVideoData(res.data);
                    for (let i = 0; i < res.data.length; i++) {
                        if (res.data[i].MId === temMId) {
                            router.push({name: "video", params: {id: i}});
                        }
                    }
                }
            });
        } else {
            router.push({name: "video", params: {id: value}});
        }
    };

    const handleAdminRoute = () => {
        router.push({name: "admin"});
    }

</script>

<style lang="scss">
    .nav {
        padding: 0 var(--side-padding);
        grid-template-columns: calc(272px - var(--side-padding)) auto auto;
        height: var(--headerTop);
        display: grid;
        z-index: 30;
        position: fixed;
        top: var(--headerTopScroll);
        transition: top 0.2s linear;
        @include theme();
        @include phone() {
            padding: 0 20px;
        }

        .ui-logo {
            cursor: pointer;
            display: flex;
            font-size: 27px;
            color: white;
            align-items: center;

            img {
                height: 32px;
                margin-right: 30px;
                width: 32px;
            }
        }

        .nav-select {
            display: flex;
            align-items: center;
            padding: 0;
            @include phone {
                margin-left: -60px;
                width: 250px;
            }

            .search {
                @include theme();
                width: 80%;
                border: 0;
                outline: none;
            }
        }

        .nav-end {
            cursor: pointer;
            display: flex;
            align-items: center;
            margin-left: 60%;

            .notification {
                display: flex;
                margin-right: 30px;
                color: hsla(0, 0%, 100%, .4);
                flex-direction: column;
                justify-content: space-between;
                align-items: center;

                span {
                    color: hsla(0, 0%, 100%, .8);
                    font-size: 10px;
                    font-weight: 500;
                    line-height: 20px;
                    text-align: center;
                    word-break: keep-all;
                    letter-spacing: 1px;
                }
            }

            img {
                @include phone() {
                    display: none;
                }
            }
        }
    }
</style>
