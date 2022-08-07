import { createApp } from 'vue'
import 'vuetify/styles' // Global CSS has to be imported
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import { loadFonts } from './plugins/webfontloader'

loadFonts()

createApp(App).use(vuetify).use(router).mount('#app')
