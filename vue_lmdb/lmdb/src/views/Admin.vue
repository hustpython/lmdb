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
                <component :is="getMenuComponent">
                </component>
            </n-layout>
        </n-layout>
    </n-space>
    <Footer/>
</template>

<script setup>
    import Header from "@/components/Header.vue";
    import Footer from "@/components/Footer";
    import AdminChart from "@/components/admin/AdminChart.vue";
    import AdminTable from "@/components/admin/AdminTable.vue"
    import AdminCutList from "@/components/admin/AdminCutList.vue";
    import {ref, computed} from 'vue';
    import {useRoute} from "vue-router";

    const route = useRoute();

    const collapsed = ref(false);

    const menuOptions = menuConfig;

    const selectedItem = ref("0");

    if (route.params.id === undefined) {
        selectedItem.value = "0";
    } else {
        selectedItem.value = route.params.id;
    }


    const getMenuComponent = computed(() => {
        switch (selectedItem.value) {
            case "磁盘概览":
                return AdminChart;
            case "数据表格":
                return AdminTable;
            default:
                return AdminCutList;
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
        height: 100%;
        margin-top: calc(var(--headerTop) + 15px);

        .theme {
            @include theme();
        }
    }

</style>