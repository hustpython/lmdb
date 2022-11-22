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
                    :collapsed-width="15"
                    :width="150"
                    :collapsed="collapsed"
                    show-trigger
                    @collapse="collapsed = true"
                    @expand="collapsed = false"
            >
                <n-menu
                        :collapsed="collapsed"
                        :collapsed-width="64"
                        :options="menuOptions"
                        :default-value="selectedItem"
                        :on-update:value="handleMenuUpdate"
                />
            </n-layout-sider>
            <n-layout class="theme">
                <component :is="getMenuComponent"/>
            </n-layout>
        </n-layout>
    </n-space>

    <Footer/>
</template>

<script setup>
    import Header from "@/components/Header.vue";
    import Footer from "@/components/Footer";
    import AdminTable from "@/components/admin/AdminTable.vue";
    import AdminCutList from "@/components/admin/AdminCutList.vue";
    import {ref, computed} from 'vue';
    import {useRoute, useRouter} from "vue-router";

    const route = useRoute();

    const collapsed = ref(false);

    const menuOptions = [
        {
            type: "group",
            label: "个人中心",
            children: [
                {
                    label: "个人资料",
                    key: "0",
                },
                {
                    label: "基本设置",
                    key: "1",
                },
                {
                    label: "磁盘概览",
                    key: "2",
                },
            ]
        },
        {
            type: "group",
            label: "创作中心",
            children: [
                {
                    label: "剪切列表",
                    key: "3",
                },
                {
                    label: "影评创作",
                    key: "4",
                },
                {
                    label: "封面推荐",
                    key: "5",
                }
            ]
        },
    ];

    const selectedItem = ref("0");

    if (route.params.id === undefined) {
        selectedItem.value = "0";
    } else {
        selectedItem.value = route.params.id;
    }

    const getMenuComponent = computed(() => {
        switch (selectedItem.value) {
            case "2":
                return AdminTable;
            case "3":
                return AdminCutList;
            default:
                return "";
        }
    })
    const handleMenuUpdate = (v) => {
        selectedItem.value = v;
    }

</script>

<style scoped lang="scss">
    .adminSide {
        margin-left: 5%;
        margin-right: 5%;
        height: 600px;
        margin-top: calc(var(--headerTop) + 15px);
    }

    .theme {
        @include theme();
    }
</style>