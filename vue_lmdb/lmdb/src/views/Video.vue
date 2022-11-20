<template>
    <Header/>
    <n-layout class="videopage">
        <div class="videospace">
            <LVideo :Id=routeID></LVideo>
        </div>
    </n-layout>
    <Footer/>
</template>

<script setup>
    import {useRoute} from "vue-router";
    import Header from "@/components/Header.vue";
    import LVideo from "@/components/LVideo.vue";
    import Footer from "@/components/Footer";

    const route = useRoute();
    const routeID = route.params.id;

</script>
<style lang="scss">
    .videopage {
        margin-top: var(--headerTop);
        height: 100%;
        @include theme();
    }

    .videospace {
        float: left;
        margin-top: 30px;
        margin-left: 60px;
        margin-right: 60px;
        @include phone() {
            margin-left: 0px;
            margin-top: 0px;
        }
    }
</style>
