<template>
  <n-layout-header
    class="nav"
    bordered
    v-bind:class="{
      darkThemeBck: themeData === true,
      lightThemeBck: themeData === false,
    }"
  >
    <div class="ui-logo">
      <img :src="logourl" />
      <n-gradient-text :size="24" type="success"> LMDB </n-gradient-text>
    </div>
    <n-menu
      class="header-menu"
      mode="horizontal"
      default-value="首页"
      :render-label="renderMenuLabel"
      :options="options"
    />
    <div class="nav-end">
      <n-button
        @click="handleThemeClick"
        strong
        secondary
        style="margin-right: 30px"
        size="small"
      >
        {{ themeBtnText.text }}
      </n-button>
      <img style="height: 32px; width: 32px; border-radius: 16px" :src="avatarurl" />
    </div>
  </n-layout-header>
</template>

<script setup>
import { h, reactive } from "vue";
import { useDarkTheme } from "@/store/themeData";
import { storeToRefs } from "pinia";
import { useNotification } from "naive-ui";
const notification = useNotification();

const darkThemeStore = useDarkTheme();
var { themeData } = storeToRefs(darkThemeStore);

let themeBtnText = reactive({ text: "浅色" });

const logourl = require("../assets/naivelog.svg");
const avatarurl = require("../assets/axu.png");
const handleThemeClick = () => {
  if (themeBtnText.text === "浅色") {
    themeBtnText.text = "深色";
    darkThemeStore.unsetDarkTheme();
    notification.success({
      content: "主题调整为浅色",
      duration: 1000,
    });
  } else {
    themeBtnText.text = "浅色";
    darkThemeStore.setDarkTheme();
    notification.success({
      content: "主题调整为深色",
      duration: 1000,
    });
  }
};
window.addEventListener(
  "scroll",
  () => {
    if (scrollY > 63) {
      document.body.style.setProperty("--headerTopScroll", "-63px");
    } else {
      document.body.style.setProperty("--headerTopScroll", "0px");
    }
  },
  true
);

const options = [
  {
    label: "首页",
    key: "首页",
    href: "/",
  },
  {
    label: "视频",
    key: "视频",
  },
  {
    label: "剪辑",
    key: "剪辑",
  },
  {
    label: "影评",
    key: "影评",
  },
];

const renderMenuLabel = (option) => {
  if ("href" in option) {
    return h("a", { href: option.href }, option.label);
  }
  return option.label;
};
</script>

<style>
.nav {
  padding: 0 var(--side-padding);
  grid-template-columns: calc(272px - var(--side-padding)) 1fr auto;
  height: var(--headerTop);
  display: grid;
  z-index: 30;
  position: fixed;
  top: var(--headerTopScroll);
  transition: top 0.2s linear;
}
/* 以后可以考虑使用这种变量来改变主题 */
/* background-color: v-bind(color); */
.ui-logo {
  cursor: pointer;
  display: flex;
  font-size: 27px;
  color: white;
}

.nav,
.ui-logo {
  align-items: center;
}

.ui-logo > img {
  height: 32px;
  margin-right: 30px;
  width: 32px;
}

.nav-end {
  cursor: pointer;
  display: flex;
  font-size: 27px;
  color: white;
}

.nav,
.nav-end {
  align-items: center;
}

.header-menu {
  display: flex;
  color: #8a2be2;
  font-size: 15px;
  align-items: center;
  width: 400px;
}
</style>
