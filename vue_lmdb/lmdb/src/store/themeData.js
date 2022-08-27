import { defineStore } from 'pinia';

export const useDarkTheme = defineStore({
    id: "themeData",
    state: () => {
        return {
            themeData: true,
        }
    },
    actions: {
        unsetDarkTheme() {
            this.themeData = false;
        },
        setDarkTheme() {
            this.themeData = true;
        }
    },
    persist: {
        enabled: true,
        strategies: [{
            key: 'themeData',
            storage: localStorage,
        }]
    }
})