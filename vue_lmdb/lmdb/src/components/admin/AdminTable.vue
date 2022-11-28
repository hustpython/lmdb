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
    </div>
</template>

<script setup>
    import {ref, h, onBeforeMount} from 'vue'
    import {NTag} from "naive-ui";
    import {GetVideoTable, GetAllTags} from '@/api/videolist'

    const columns = [
        {
            type: "selection",
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
            key: "durationStr",
            sorter: (row1, row2) => row1.duration - row2.duration
        },
        {
            title: "大小",
            key: "sizeStr",
            sorter: (row1, row2) => row1.size - row2.size
        },
        {
            title: "介绍",
            key: "desc",
            width: 240,
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
            width: 150,
            sorter: (row1, row2) => row1.recent - row2.recent
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
    const allTags = ref([""]);
    const pagination = {pageSize: 30};
    onBeforeMount(() => {
        GetVideoTable().then((res) => {
            data.value = res.data;
        })
    })
</script>

<style scoped lang="scss">
    .adminTable {
        margin-left: 18px;
        margin-bottom: 20px;
        margin-top: 10px;

        ::v-deep(.n-data-table  ) {
            background: transparent;
            --n-merged-th-color: transparent;
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
            color: hsla(0, 0%, 100%, .9);
        }
    }
</style>