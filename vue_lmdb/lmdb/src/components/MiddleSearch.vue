<template>
    <div class="homebkg">
        <n-carousel autoplay class="imgbkg">
            <div v-for="(item,index) in rightItems">

            </div>
        </n-carousel>
        <div class="rightBox">
        </div>
    </div>


</template>

<script setup>
    import {ref} from 'vue'

    const imgurl = require("../assets/classic.jpg");
    const rightItems = ref(["小森林", "1212"]);

</script>

<style lang="scss">
    .homebkg {
        width: 100%;
        margin-top: calc(var(--headerTop) + 30px);
        height: $homeBckHeight;
        @include phone {
            height: 140px;
        }

        .imgbkg {
            width: 65%;
            left: 5%;
            position: relative;
            overflow: hidden;
            background-image: v-bind("'url(' + imgurl + ')'");
            background-size: 100%;
            float: left;

        }

        .rightBox {
            width: 25%;
            height: inherit;
            float: right;
            margin-right: 7%;
            background-color: hsla(0, 0%, 100%, .1);
        }
    }

</style>
