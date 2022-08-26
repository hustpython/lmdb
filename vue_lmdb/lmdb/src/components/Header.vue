<template>
  <n-layout>
    <n-layout-header class="nav">
      <div class="ui-logo">
        <img :src="logourl" />
        <n-gradient-text :size="24" type="success"> LMDB </n-gradient-text>
      </div>
      <n-menu
        class="header-menu"
        mode="horizontal"
        default-value="首页"
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
  </n-layout>
</template>

<script setup>
import { reactive } from "vue";
import { useDarkTheme } from "@/store/themeData";

const darkThemeStore = useDarkTheme();

let themeBtnText = reactive({ text: "浅色" });

const logourl = require("../assets/naivelog.svg");
const avatarurl = require("../assets/axu.png");
const handleThemeClick = () => {
  if (themeBtnText.text === "浅色") {
    themeBtnText.text = "深色";
    darkThemeStore.unsetDarkTheme();
  } else {
    themeBtnText.text = "浅色";
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
</script>

<style>
.nav {
  padding: 0 var(--side-padding);
  grid-template-columns: calc(272px - var(--side-padding)) 1fr auto;
  background-color: var(--lightBule);
  height: var(--headerTop);
  display: grid;
  z-index: 30;
  position: fixed;
  top: var(--headerTopScroll);
  transition: top 0.2s linear;
}

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
