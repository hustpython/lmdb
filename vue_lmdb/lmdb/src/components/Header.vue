<template>
    <n-layout-header class="nav" bordered>
        <div class="ui-logo">
            <img :src="logourl"
            />
            <a href="/">
                <n-gradient-text :size="24" type="success"> LMDB</n-gradient-text>
            </a>
        </div>
        <n-menu
                class="header-menu"
                mode="horizontal"
                :default-value="props.activeMenuItem"
                :render-label="renderMenuLabel"
                :options="options"
        />

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
            <n-button
                    @click="handleThemeClick"
                    strong
                    secondary
                    style="margin-right: 30px"
                    size="small"
            >
                {{ themeBtnText.text }}
            </n-button>
            <img style="height: 32px; width: 32px; border-radius: 16px" :src="avatarurl"/>
        </div>
        <div class="rightbar">
            <n-icon size="30">
                <Search/>
            </n-icon>
        </div>
    </n-layout-header>
</template>

<script setup>
    import {h, reactive, defineProps, onBeforeMount, computed} from "vue";

    import {storeToRefs} from "pinia";
    import {useVideoData} from "@/store/videoData";
    import {useRouter} from "vue-router";
    import {GetMoviesByColl} from "@/api/videolist";

    import {useDarkTheme} from "@/store/themeData";
    import {Search} from "@vicons/carbon";

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
    const props = defineProps(["activeMenuItem"]);
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

    const options = [
        {
            label: "首页",
            key: "首页",
            href: "/",
        },
        {
            label: "视频",
            key: "视频",
        },
        {
            label: "剪辑",
            key: "剪辑",
        },
        {
            label: "影评",
            key: "影评",
        },
    ];

    const renderMenuLabel = (option) => {
        if ("href" in option) {
            return h("a", {href: option.href}, option.label);
        }
        return option.label;
    };

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

</script>

<style lang="scss">
    .nav {
        padding: 0 var(--side-padding);
        grid-template-columns: calc(272px - var(--side-padding)) 300px 450px auto;
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


        .header-menu {
            display: flex;
            color: #8a2be2;
            font-size: 15px;
            align-items: center;
            width: 280px;
            padding: 0;
            @include phone() {
                display: none;
            }
        }

        .nav-select {
            display: flex;
            align-items: center;
            padding: 0;
            margin-left: 60px;
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
            margin-right: 0;
            cursor: pointer;
            display: flex;
            font-size: 27px;
            color: white;
            align-items: center;

            img {
                @include phone() {
                    display: none;
                }
            }
        }

        .rightbar {
            display: none;
            @include phone() {
                cursor: pointer;
                display: flex;
                color: $lightBlue;
                align-items: center;
                @include theme();
            }
        }
    }
</style>
