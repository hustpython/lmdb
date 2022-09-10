<template>
  <n-layout-header class="nav" bordered>
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
    <div class="rightbar">
      <n-icon size="30"><Search /></n-icon>
    </div>
  </n-layout-header>
</template>

<script setup>
import { h, reactive, onBeforeMount } from "vue";
import { useDarkTheme } from "@/store/themeData";
import { Search } from "@vicons/carbon";

const darkThemeStore = useDarkTheme();

// 默认深色主题
let themeBtnText = reactive({ text: "深色" });
onBeforeMount(() => {
  if (darkThemeStore.themeData === true) {
    themeBtnText.text = "浅色";
    setAttribute("deep");
  } else {
    setAttribute("light");
  }
});

const setAttribute = (theme) => {
  window.document.documentElement.setAttribute("data-theme", theme);
};

const logourl = require("../assets/naivelog.svg");
const avatarurl = require("../assets/axu.png");
const handleThemeClick = () => {
  if (themeBtnText.text === "浅色") {
    themeBtnText.text = "深色";
    setAttribute("light");
    darkThemeStore.unsetDarkTheme();
  } else {
    themeBtnText.text = "浅色";
    setAttribute("deep");
    darkThemeStore.setDarkTheme();
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

<style lang="scss">
.nav {
  padding: 0 var(--side-padding);
  grid-template-columns: calc(272px - var(--side-padding)) 1fr auto;
  height: var(--headerTop);
  display: grid;
  z-index: 30;
  position: fixed;
  top: var(--headerTopScroll);
  transition: top 0.2s linear;
  @include theme();
  @include phone() {
    padding: 0 20px;
  }
  .ui-logo {
    cursor: pointer;
    display: flex;
    font-size: 27px;
    color: white;
    img {
      height: 32px;
      margin-right: 30px;
      width: 32px;
    }
  }
  .ui-logo {
    align-items: center;
  }
  .header-menu {
    display: flex;
    color: #8a2be2;
    font-size: 15px;
    align-items: center;
    width: 400px;
    @include phone() {
      display: none;
    }
  }
  .nav-end {
    cursor: pointer;
    display: flex;
    font-size: 27px;
    color: white;
    align-items: center;
    img {
      @include phone() {
        display: none;
      }
    }
  }
  .rightbar {
    display: none;
    @include phone() {
      cursor: pointer;
      display: flex;
      color: $lightBlue;
      align-items: center;
      @include theme();
    }
  }
}
</style>
