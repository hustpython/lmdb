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
                    :on-update:value="handleSelected"
                    placeholder="搜索电影、剧集、影评..."
            />
        </div>

        <div class="nav-end">
            <n-popover trigger="hover"
                       style="background-color: hsla(0, 0%, 100%, .3)">
                <template #trigger>

                    <n-badge class="notification" value="1" dot>
                        <n-icon
                                size="32px"
                        >
                            <CircleNotificationsRound/>
                        </n-icon>
                    </n-badge>

                </template>
                <n-space vertical

                >
                    <n-card
                            style="background-color: hsla(0, 0%, 100%, .3)"
                            title="小卡片" size="small">
                        卡片内容
                    </n-card>

                </n-space>
            </n-popover>

            <span class="headerDiv"
            >
            </span>
            <n-popover
                    style="background-color: hsla(0, 0%, 100%, .3)"
                    trigger="hover">
                <template #trigger>
                    <img style="height: 32px; width: 32px; border-radius: 16px" :src="avatarurl"/>
                </template>
                <n-space>
                    <n-card v-for:="(item,index) in headerAdminMenu"
                            style="background-color: hsla(0, 0%, 100%, .3)"
                            :title=item.label
                            size="small"
                            hoverable>
                        <n-space vertical>
                            <n-button v-for="(item1,index) in item.children"
                                      @click="handleMenuClick(item1.key)">
                                {{item1.label}}
                            </n-button>
                        </n-space>
                    </n-card>
                    <n-card
                            style="background-color: hsla(0, 0%, 100%, .3)"
                            title=其他
                            size="small"
                            hoverable>
                        <n-space vertical>
                            <n-button
                                    @click="handleThemeClick">
                                切换主题
                            </n-button>
                            <n-button>
                                退出登录
                            </n-button>
                        </n-space>
                    </n-card>
                </n-space>
            </n-popover>
        </div>
    </n-layout-header>
</template>

<script setup>
    import {reactive, onBeforeMount, ref} from "vue";
    import {CircleNotificationsRound} from "@vicons/material";

    import {useRouter} from "vue-router";
    import {GetSearchMovie} from "@/api/videolist";

    import {useDarkTheme} from "@/store/themeData";

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


    const selectOptions = ref([]);

    onBeforeMount(() => {
        GetSearchMovie().then((res) => {
            selectOptions.value = res.data
        })
    })
    const router = useRouter();
    const handleSelected = (value) => {
        router.push({name: "video", params: {id: value}});
    };

    const headerAdminMenu = menuConfig;

    const handleMenuClick = (key) => {
        router.push({name: "admin", params: {id: key}});
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
                width: 90%;
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
                flex-direction: column;
                justify-content: space-between;
                align-items: center;
                @include themeHeaderNotify();
            }

            img {
                @include phone() {
                    display: none;
                }
            }

            .headerDiv {
                width: 0.5px;
                height: 51%;
                background-color: hsla(0, 0%, 100%, .4);
                margin-right: 20px;

            }
        }
    }
</style>
