<template>
    <div class="VideoRec">
        <div class="RecLeft">
            <n-select
                    filterable
                    :options="selectOptions"
                    v-model:value="selectDefaultValue"
                    :on-update:value="handleSelected"
                    placeholder="选择需要识别的电影/电视剧"
            />
            <div class="originCard"
                 v-bind:style="{backgroundImage:'linear-gradient(to right, rgb(3, 37, 65), transparent),url('+ getVideoCover + ')'}"
            >
                <div class="single">
                    <n-space>
                        <img
                                :src="getVideoPoster"
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
                                <n-rate size="small"
                                        allow-half readonly
                                        :value="getRateFromVote(currentData.Vote)"
                                />
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
                    :disabled="disableRecon"
                    @click="handleReconClick">
                开始识别
            </n-button>
            <n-button
                    style="margin-top: 15px;margin-left: 30px"
                    :disabled="disableRecon"
                    @click="handleReconPlay">
                返回播放
            </n-button>

            <n-divider v-show="showReconSelect"/>
            <div v-show="showReconSelect">
                <n-select
                        filterable
                        style="margin-bottom: 15px"
                        :options="recOptions"
                        :render-option="renderOption"
                        v-model:value="selectRecoIndex"
                        :on-update:value="handleSynSelected"
                />

                <n-space item-style="display: flex;" align="center">
                    <n-checkbox v-model:checked="syncCheckValue.Title">
                        片名
                    </n-checkbox>
                    <n-checkbox v-model:checked="syncCheckValue.TagArray">
                        标签
                    </n-checkbox>
                    <n-checkbox v-model:checked="syncCheckValue.Vote">
                        评分
                    </n-checkbox>
                    <n-checkbox v-model:checked="syncCheckValue.Duration">
                        时长
                    </n-checkbox>
                    <n-checkbox v-model:checked="syncCheckValue.Desc">
                        描述
                    </n-checkbox>
                    <n-checkbox v-model:checked="syncCheckValue.ReleaseDate">
                        上映时间
                    </n-checkbox>
                    <n-checkbox v-model:checked="syncCheckValue.Poster">
                        海报
                    </n-checkbox>
                    <n-checkbox v-model:checked="syncCheckValue.Cover">
                        背景
                    </n-checkbox>
                    <n-button
                            :disabled="syncDisable"
                            size="small" @click="handleReconSync">
                        同步
                    </n-button>
                </n-space>
            </div>

            <n-divider v-show="showReconSelect"/>
            <div class="originCard"
                 v-bind:style="{backgroundImage:'linear-gradient(to right, rgb(3, 37, 65), transparent),url('+ item.Cover + ')'}"
                 v-for="(item,index) in TmdbDataArray"
                 v-show="selectRecoIndex===''|| selectRecoIndex===index "
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
                                <n-rate size="small"
                                        allow-half readonly
                                        :value="getRateFromVote(item.Vote)"/>
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
    import {computed, onBeforeMount, watch, ref, h} from "vue";

    import {NTooltip, NImage} from 'naive-ui'

    import {GetMovieByMId, GetTMDBByKeyWord, GetSearchMovie, UpdateVideo} from "@/api/videolist";
    import {useRouter} from "vue-router";

    const selectOptions = ref([]);
    const selectDefaultValue = ref();
    const props = defineProps(["MId"]);


    const currentData = ref({});
    const setVideoData = (Id) => {
        GetMovieByMId(Id).then((res) => {
            if (res.code == 200) {
                if (res.data === null) {
                    return false;
                }
                currentData.value = res.data;
                selectDefaultValue.value = currentData.value.Title;
            }
        });
    }

    const keyWord = ref();
    const disableRecon = ref(true);

    watch(
        () => selectDefaultValue.value,
        (newValue) => {
            if (newValue === undefined) {
                keyWord.value = "";
                disableRecon.value = true;
            } else {
                keyWord.value = newValue;
                disableRecon.value = false;
            }
        },
        {immediate: true}
    )

    if (props.MId != undefined) {
        setVideoData(props.MId);
    }
    const getRateFromVote = (v) => {
        return v / 2;
    };

    const handleSelected = (v, k) => {
        keyWord.value = k.label;
        showReconSelect.value = false;
        TmdbDataArray.value = [];
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

    const getVideoPoster = computed(() => {
        if (currentData.value.Poster === undefined
            || currentData.value.Poster === '') {
            return "";
        }
        if (currentData.value.Poster.indexOf(";") != -1) {
            return currentData.value.Poster;
        }
        return "data:image/png;base64," + currentData.value.Poster;
    })


    const TmdbDataArray = ref([]);
    const loading = ref(false);
    const showReconSelect = ref(false);
    const recOptions = ref([{
        value: ''
    }]);
    const syncDisable = ref(true);
    const handleReconClick = () => {
        TmdbDataArray.value = [];
        loading.value = true;
        recOptions.value = [{value: ''}];
        selectRecoIndex.value = '';
        syncDisable.value = true;
        setTimeout(() => {
            loading.value = false
        }, 60000)
        GetTMDBByKeyWord(keyWord.value).then((res) => {
            if (res.code == 200) {
                if (res.data === null) {
                    return false;
                }
                showReconSelect.value = true;
                TmdbDataArray.value = res.data;
                loading.value = false;
                res.data.forEach(function (v, i) {
                    recOptions.value.push({
                        label: v.Title,
                        value: i,
                        Poster: v.Poster,
                    })
                });
            }
        });
    }

    const router = useRouter();
    const handleReconPlay = () => {
        router.push({name: "video", params: {id: currentData.value.MId}})
    }

    const renderOption = ({node, option}) => h(NTooltip, null, {
        trigger: () => node,
        default: () => h(NImage, {
            src: option.Poster,
            width: "150"
        }),
    })
    const selectRecoIndex = ref();

    const syncCheckValue = ref({
        Title: false,
        TagArray: true,
        ReleaseDate: true,
        Vote: true,
        Desc: true,
        Duration: false,
        Poster: true,
        Cover: true,
    });

    const handleReconSync = () => {
        let updateRes = {MId: currentData.value.MId};
        if (syncCheckValue.value.Desc) {
            currentData.value.Desc = TmdbDataArray.value[selectRecoIndex.value].Desc;
            updateRes.Desc = currentData.value.Desc;
        }
        if (syncCheckValue.value.Title) {
            currentData.value.Title = TmdbDataArray.value[selectRecoIndex.value].Title;
            updateRes.Title = currentData.value.Title;
        }
        if (syncCheckValue.value.ReleaseDate) {
            currentData.value.ReleaseDate = TmdbDataArray.value[selectRecoIndex.value].ReleaseDate;
            updateRes.ReleaseDate = currentData.value.ReleaseDate;
        }
        if (syncCheckValue.value.Vote) {
            currentData.value.Vote = TmdbDataArray.value[selectRecoIndex.value].Vote;
            updateRes.Vote = currentData.value.Vote;
        }
        if (syncCheckValue.value.Duration) {
            currentData.value.Duration = TmdbDataArray.value[selectRecoIndex.value].Duration;
            updateRes.Duration = currentData.value.Duration;
        }
        if (syncCheckValue.value.TagArray) {
            currentData.value.TagArray = TmdbDataArray.value[selectRecoIndex.value].TagArray;
            updateRes.TagArray = currentData.value.TagArray;
        }
        if (syncCheckValue.value.Poster) {
            currentData.value.Poster = TmdbDataArray.value[selectRecoIndex.value].Poster;
            updateRes.Poster = currentData.value.Poster.slice(22);
        }
        if (syncCheckValue.value.Cover) {
            currentData.value.Cover = TmdbDataArray.value[selectRecoIndex.value].Cover;
            updateRes.Cover = currentData.value.Cover.slice(22);
        }
        UpdateVideo(updateRes);
    }

    const handleSynSelected = (v) => {
        selectRecoIndex.value = v;
        if (v !== '') {
            syncDisable.value = false;
        } else {
            syncDisable.value = true;
        }
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
                    padding-top: 45px;
                    padding-left: 30px;
                    padding-bottom: 45px;

                    img {
                        width: 180px;
                        border-radius: 10px;
                        align-items: center;
                    }

                    .rightDesc {
                        width: 100%;
                        background-color: hsla(0, 0%, 100%, .1);
                        border-radius: 10px;
                        margin-left: 10px;
                    }
                }
            }
        }
    }
</style>