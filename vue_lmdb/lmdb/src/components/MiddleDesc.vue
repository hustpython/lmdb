<template>
  <n-space class="description">
    <div v-for="(item, index) in props.moreInfo.Labels">
      <n-tag :type="labelTypes[index]" round>
        {{ item }}
      </n-tag>
    </div>
  </n-space>

  <n-space class="description" vertical :size="10">
    <n-descriptions bordered>
      <n-descriptions-item label="导演/主演">
        {{ props.moreInfo.Actors }}
      </n-descriptions-item>
      <n-descriptions-item label="国家/时间">
        {{ props.moreInfo.National }}
      </n-descriptions-item>
    </n-descriptions>
  </n-space>
  <n-ellipsis class="description" :line-clamp="2" style="text-align: center">
    {{ props.moreInfo.Desc }}
  </n-ellipsis>
</template>

<script setup>
const props = defineProps(["moreInfo"]);
const labelTypes = ["success", "warning", "error", "info"];
</script>
<style>
.description {
  padding: 0 15px 10px;
  position: relative;
  font-size: 14px;
}
</style>