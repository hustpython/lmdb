import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Video from '../views/Video.vue'
const routes = [
    {
        path: '/',
        name: 'home',
        component: Home
    },
    {
        path: '/video',
        name: 'video',
        component: Video
    },
]
const router = createRouter({
    history: createWebHashHistory(process.env.BASE_URL),
    routes
})
export default router
