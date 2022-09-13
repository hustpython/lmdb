<template>
    <n-layout-footer class="footer" bordered> 版权所有 mxq</n-layout-footer>
</template>

<script setup>
    import {ref} from 'vue'

    const color = ref("blue");
</script>

<!--v-bind样例-->
<style scoped lang="scss">
    .footer {
        text-align: center;
        color: v-bind(color);
    }
</style>
