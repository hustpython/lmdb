import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store';
import naive from 'naive-ui'
import '@/config/config.js'
import '@/config/adminMenu.js'

createApp(App).use(router).use(store).use(naive).mount('#app')
