<template>
    <!--    选中行 受控的排序 可展开 树型数据 可切换的可编辑表格 右键菜单-->
    <!--    自定义选择项菜单-->
    <div class="adminTable">
        <n-data-table
                :columns="columns"
                :data="data"
                :bordered="false"
                :single-line="false"
                :pagination="pagination"
                style="height: 540px"
                flex-height
                :row-props="rowProps"
        />
        <n-dropdown
                style="background-color: hsla(0, 0%, 100%, .3)"
                placement="bottom-start"
                trigger="manual"
                :x="xRef"
                :y="yRef"
                size="large"
                :options="rightOptions"
                :show="showDropdownRef"
                :on-clickoutside="onClickoutside"
                @select="handleSelect"
        />
    </div>
</template>

<script setup>
    import {ref, h, onBeforeMount, nextTick} from 'vue'
    import {NTag, NIcon} from "naive-ui";
    import {GetVideoTable} from '@/api/videolist'
    import {PlayOutline, FolderDetails, Delete, Save} from "@vicons/carbon";
    import {useRouter} from "vue-router";

    const renderIcon = (icon) => {
        return () => {
            return h(NIcon, null, {
                default: () => h(icon)
            });
        };
    };
    const columns = [
        {
            type: "selection",
            options: [
                {
                    label: "批量删除",
                    key: "batchDelete",
                    onSelect: (pageData) => {
                        console.log(pageData);
                    }
                },
                {
                    label: "增加合集",
                    key: "batchAddColl",
                    onSelect: (pageData) => {
                        console.log(pageData);
                    }
                },
                {
                    label: "另存为",
                    key: "batchSave",
                    onSelect: (pageData) => {
                        console.log(pageData);
                    }
                }
            ],
        },
        {
            title: "片名(合集名)",
            key: "title",
            width: 150,
            ellipsis: {
                tooltip: true
            }
        },
        {
            title: "时长",
            width: 85,
            key: "durationStr",
            sorter: (row1, row2) => row1.duration - row2.duration
        },
        {
            title: "大小",
            width: 72,
            key: "sizeStr",
            sorter: (row1, row2) => row1.size - row2.size
        },
        {
            title: "介绍",
            key: "desc",
            width: 200,
            ellipsis: {
                tooltip: true
            }
        },
        {
            title: "标签",
            key: "tags",
            render(row) {
                const tags = row.tags.map((tagKey) => {
                    return h(
                        NTag,
                        {
                            style: {
                                marginRight: '6px'
                            },
                            type: 'info',
                            bordered: false
                        },
                        {
                            default: () => tagKey
                        }
                    )
                })
                return tags
            },
        },
        {
            title: "最近",
            key: "recentStr",
            width: 110,
            sorter: (row1, row2) => row1.recent - row2.recent
        },
        {
            title: "最新",
            key: "modifyTimeStr",
            width: 110,
            sorter: (row1, row2) => row1.modifyTime - row2.modifyTime
        },
        {
            title: '在线',
            key: 'online',
            defaultFilter: ['是', '否'],
            filterOptions: [
                {
                    label: '是',
                    value: '是'
                },
                {
                    label: '否',
                    value: '否'
                }
            ],
            filter(value, row) {
                return row.online.indexOf(value) >= 0
            }
        }
    ];
    const data = ref([]);
    const pagination = {pageSize: 30};
    onBeforeMount(() => {
        GetVideoTable().then((res) => {
            data.value = res.data;
        })
    })

    const rightOptions = [
        {
            label: "播放",
            key: "play",
            icon: renderIcon(PlayOutline)
        },
        {
            label: "打开",
            key: "open",
            icon: renderIcon(FolderDetails)
        },
        {
            label: "删除",
            key: "delete",
            icon: renderIcon(Delete)
        },
        {
            label: "另存为",
            key: "save",
            icon: renderIcon(Save)
        },

    ];
    const showDropdownRef = ref(false);
    const xRef = ref(0);
    const yRef = ref(0);
    const onClickoutside = () => {
        showDropdownRef.value = false;
    }
    let rightKeyValue = "";
    const router = useRouter();
    const handleSelect = (key) => {
        showDropdownRef.value = false;
        let v = rightKeyValue.replace('cj', '');
        switch (key) {
            case "play":
                router.push({name: "video", params: {id: v}});
                break;
            default:
                console.log(v);
        }
    }
    const rowProps = (row) => {
        return {
            onContextmenu: (e) => {
                rightKeyValue = row.key;
                e.preventDefault();
                showDropdownRef.value = false;
                nextTick().then(() => {
                    showDropdownRef.value = true;
                    xRef.value = e.clientX;
                    yRef.value = e.clientY;
                });
            }
        };
    }
</script>

<style scoped lang="scss">
    .adminTable {
        margin-left: 18px;
        margin-bottom: 20px;
        margin-top: 10px;

        ::v-deep(.n-data-table  ) {
            background: transparent;
            --n-merged-th-color: hsla(0, 0%, 100%, .1);
            --n-merged-td-color: transparent;
            --n-merged-th-color-hover: transparent;
            --n-merged-td-color-hover: transparent;
            --n-merged-border-color: hsla(0, 0%, 100%, .4);

        }

        ::v-deep(.n-checkbox ) {
            --n-merged-color-table: transparent;
        }

        ::v-deep(.n-data-table-th) {
            --n-th-font-weight: 900;
        }
    }
</style>