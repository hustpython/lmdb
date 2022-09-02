import { defineStore } from 'pinia';

export const useVideoData = defineStore({
    id: "videoData",
    state: () => {
        return {
            videoData: [],
        }
    },
    actions: {
        setVideoData(val) {
            this.videoData = val;
        }
    },
    persist: {
        enabled: false,
        strategies: [{
            key: 'videoData',
            storage: localStorage,
        }]
    }
})