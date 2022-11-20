<template>
    <Header/>
    <n-space
            class="adminSide"
            vertical>
        <n-layout

                has-sider>
            <n-layout-sider
                    class="theme"
                    bordered
                    collapse-mode="width"
                    :collapsed-width="45"
                    :width="150"
                    show-trigger
                    @collapse="collapsed = true"
                    @expand="collapsed = false"
            >
                <n-menu
                        :options="menuOptions"
                        :collapsed="collapsed"
                        :collapsed-width="64"
                        key-field="key"
                        label-field="label"
                        children-field="children"
                />
            </n-layout-sider>
            <n-layout class="theme">
                <span>内容</span>
            </n-layout>
        </n-layout>
    </n-space>
</template>

<script setup>
    import Header from "@/components/Header.vue";
    import {ref} from 'vue';

    const collapsed = ref(true);
    const menuOptions = [
        {
            type: "group",
            label: "个人中心",
            key: "personal",
            children: [
                {
                    label: "个人资料",
                    key: "0-1",
                },
                {
                    label: "数据管理",
                    key: "0-2",
                }
            ]
        },
        {
            type: "group",
            label: "创作中心",
            key: "create",
            children: [
                {
                    label: "剪切列表",
                    key: "1-1",
                },
                {
                    label: "影评创作",
                    key: "1-2",
                },
                {
                    label: "封面推荐",
                    key: "1-3",
                }
            ]
        },
    ];
</script>

<style scoped lang="scss">
    .adminSide {
        margin-left: 5%;
        margin-right: 5%;
        margin-top: calc(var(--headerTop) + 15px);
    }

    .theme {
        @include theme();
    }
</style>