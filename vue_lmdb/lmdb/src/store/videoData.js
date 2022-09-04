import { defineStore } from 'pinia';

export const useVideoData = defineStore({
    id: "videoData",
    state: () => {
        return {
            videoData: [],
        }
    },
    actions: {
        addVideoData(val) {
            this.videoData.push(...val);
        },
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