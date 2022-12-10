<template>
    <div class="VideoRec">
        <div class="RecLeft">
            <n-select
                    filterable
                    :options="selectOptions"
                    :on-update:value="handleSelected"
                    :default-value="selectDefaultValue"
                    placeholder="选择需要识别的电影/电视剧"
            />
            <div class="originCard"
                 v-bind:style="{backgroundImage:'linear-gradient(to right, rgb(3, 37, 65), transparent),url('+ getVideoCover + ')'}"
            >
                <div class="single">
                    <n-space>
                        <img
                                :src="currentData.Poster"
                        >
                        <n-space
                                class="rightDesc"
                                vertical>
                           <span style="color: black;
                           font-size: 21px;
                           font-weight: bold">
                               {{currentData.Title}}
                            </span>
                            <n-space style="color: black;font-size: 15px;">
                                {{currentData.ReleaseDate}}
                                <div v-for="(item, index) in currentData.TagArray">
                                    <n-tag size="small" :type="labelTypes(index)" round>
                                        {{ item }}
                                    </n-tag>
                                </div>
                                {{currentData.Duration}}
                            </n-space>
                            <span>
                                <n-rate size="small" allow-half readonly
                                        :default-value="getRateFromVote(currentData.Vote)"/>
                                 {{currentData.Vote}}
                            </span>

                            <span style="color: black;font-size: 18px">
                            剧情简介:
                            </span>
                            <n-ellipsis
                                    style=" max-width: 500px"
                                    expand-trigger="click"
                                    line-clamp="4"
                                    :tooltip="false"
                            >
                                {{currentData.Desc}}
                            </n-ellipsis>
                        </n-space>
                    </n-space>
                </div>

            </div>
            <n-input v-model:value="keyWord" type="text"/>
            <n-button
                    :loading="loading"
                    style="margin-top: 15px"
                    @click="handleReconClick">
                开始识别
            </n-button>


            <div class="originCard"
                 v-bind:style="{backgroundImage:'linear-gradient(to right, rgb(3, 37, 65), transparent),url('+ item.Cover + ')'}"
                 v-for="(item) in TmdbDataArray"
            >
                <div class="single">
                    <n-space>
                        <img
                                :src="item.Poster"
                        >
                        <n-space
                                class="rightDesc"
                                vertical>
                           <span style="color: black;
                           font-size: 21px;
                           font-weight: bold">
                               {{item.Title}}
                            </span>
                            <n-space style="color: black;font-size: 15px">
                                {{item.ReleaseDate}}
                                <div v-for="(item, index) in item.TagArray">
                                    <n-tag size="small" :type="labelTypes(index)" round>
                                        {{ item }}
                                    </n-tag>
                                </div>
                                {{item.Duration}}
                            </n-space>
                            <span>
                                <n-rate size="small" allow-half readonly
                                        :default-value="getRateFromVote(item.Vote)"/>
                                 {{item.Vote}}
                            </span>

                            <span style="color: black;font-size: 18px">
                            剧情简介:
                            </span>
                            <n-ellipsis
                                    style=" max-width: 500px"
                                    expand-trigger="click"
                                    line-clamp="4"
                                    :tooltip="false"
                            >
                                {{item.Desc}}
                            </n-ellipsis>
                        </n-space>
                    </n-space>
                </div>

            </div>
        </div>
    </div>
</template>

<script setup>
    import {computed, onBeforeMount, ref} from "vue";
    import {GetMovieByMId, GetTMDBByKeyWord, GetSearchMovie} from "@/api/videolist";

    const selectOptions = ref([]);
    const selectDefaultValue = ref("我们的父辈");
    const currentData = ref({});
    const setVideoData = (Id) => {
        GetMovieByMId(Id).then((res) => {
            if (res.code == 200) {
                if (res.data === null) {
                    return false;
                }
                currentData.value = res.data;
            }
        });
    }

    const getRateFromVote = (v) => {
        return v / 2;
    }
    const keyWord = ref();

    const handleSelected = (v, k) => {
        keyWord.value = k.label;
        setVideoData(v);
    };

    const labelTypes = (index) => {
        let l = ["success", "warning", "error", "info"];
        return l[index % l.length];
    }

    onBeforeMount(() => {
        GetSearchMovie().then((res) => {
            selectOptions.value = res.data
        })
    })
    const getVideoCover = computed(() => {
        if (currentData.value.Cover === undefined) {
            return "";
        }
        if (currentData.value.Cover.indexOf(";") != -1) {
            return currentData.value.Cover;
        }
        return "data:image/png;base64," + currentData.value.Cover;
    })


    const TmdbDataArray = ref([]);

    const loading = ref(false);
    const handleReconClick = () => {
        TmdbDataArray.value = [];
        loading.value = true;
        setTimeout(() => {
            loading.value = false
        }, 30000)
        GetTMDBByKeyWord(keyWord.value).then((res) => {
            if (res.code == 200) {
                if (res.data === null) {
                    return false;
                }
                TmdbDataArray.value = res.data;
                loading.value = false;
            }
        });
    }

</script>

<style scoped lang="scss">
    .VideoRec {
        margin-left: 4%;

        .RecLeft {
            width: 87%;
            height: auto;

            .originCard {
                margin-top: 30px;
                background-size: cover;
                background-repeat: no-repeat;
                min-height: 360px;

                .single {
                    padding: 45px;

                    img {
                        width: 180px;
                        border-radius: 10px;
                        align-items: center;
                    }

                    .rightDesc {
                        width: 100%;
                        background-color: hsla(0, 0%, 100%, .1);
                        border-radius: 10px;
                    }
                }
            }
        }
    }
</style>