<template>
    <div class="imgbkg">
        <div class="content_wrap">
            <div class="title">
                <h2 style="color: white; font-size: 2.2em; font-weight: 600; line-height: 2">
                    欢迎！
                </h2>
                <h3 style="color: white; font-size: 1.8em; font-weight: 500; margin: 0">
                    这里是你本地所有的影视，快来探索与创作！
                </h3>
            </div>

            <div class="select">
                <n-select
                        class="search"
                        filterable
                        :options="options"
                        :on-update:value="handleSeleteed"
                        placeholder="搜索电影、剧集、影评..."
                />
            </div>
        </div>
    </div>
</template>

<script setup>
    import {storeToRefs} from "pinia";
    import {useVideoData} from "@/store/videoData";
    import {useRouter} from "vue-router";
    import {computed} from "vue";
    import {GetMoviesByColl} from "@/api/videolist";

    const imgurl = require("../assets/backimg.jpg");

    const videoDataStore = useVideoData();
    var {videoData} = storeToRefs(videoDataStore);

    const options = computed(() => {
        let tmpOptions = [];
        for (var i = 0; i < videoData.value.length; i++) {
            tmpOptions.push({
                label: videoData.value[i].Title,
                value: i,
            });
        }
        return tmpOptions;
    });

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
    .imgbkg {
        width: 100%;
        margin-top: var(--headerTop);
        overflow: hidden;
        height: $homeBckHeight;
        background-image: v-bind("'url(' + imgurl + ')'");
        background-size: 100%;
        @include phone {
            height: 32px;
            background-image: none;
        }
    }

    .select {
        width: 100%;
        margin-top: 30px;
        @include phone {
            margin-top: 0;
        }

        .search {
            @include theme();
            width: 80%;
            height: 32px;
            line-height: 30px;
            font-size: 1em;
            border: 0;
            outline: none;
            border-radius: 30px;
            padding: 10px 20px;
            @include phone() {
                display: none;
            }
        }
    }

    .content_wrap {
        width: 680px;
        text-align: left;
        align-items: center;
        margin-left: 100px;
        margin-top: 20px;
        position: absolute;
        @include phone {
            margin: 0;
        }

        .title {
            @include phone {
                display: none;
            }
        }
    }
</style>
