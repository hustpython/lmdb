<template>
    <n-layout-footer class="footer" bordered>
        Copyright © 2022-{{getCurrentYear}} LMDB Made by MXQ
    </n-layout-footer>
</template>

<script setup>
    import {computed} from "vue";

    const getCurrentYear = computed(() => {
        return new Date().getFullYear()
    });

</script>

<!--v-bind样例-->
<style scoped lang="scss">
    .footer {
        text-align: center;
        line-height: 40px;
        @include themeHeaderNotify();
    }
</style>
