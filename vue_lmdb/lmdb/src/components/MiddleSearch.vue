<template>
    <div class="homebkg">
        <n-carousel autoplay class="imgbkg">
            <n-carousel-item v-for="(item,index) in rightItems">

            </n-carousel-item>
        </n-carousel>
        <div class="rightBox">
            <n-carousel
                    direction="vertical"
                    show-arrow
                    :slides-per-view="3"
                    :space-between="15"
                    :show-dots="false"
                    mousewheel
                    class="vertical"
            >
                <n-carousel-item
                        style="width: 100%"
                        class="verticalImg"
                        v-for="(item,index) in rightItems">
                </n-carousel-item>
            </n-carousel>
        </div>
    </div>


</template>

<script setup>
    import {ref} from 'vue'

    const imgurl = require("../assets/classic.jpg");
    const rightItems = ref(["小森林", "1212", "121", "121", "12"]);


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
            margin-right: 8%;
            background-color: hsla(0, 0%, 100%, .1);

            .vertical {
                float: right;
                width: 85%;

                .verticalImg {
                    background-image: v-bind("'url(' + imgurl + ')'");
                    background-size: 100%;
                }
            }
        }
    }

</style>
