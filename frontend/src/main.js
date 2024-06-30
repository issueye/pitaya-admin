import { createApp } from 'vue'
import 'normalize.css/normalize.css'

import '@/assets/css/tailwind.css'
// import 'element-plus/dist/index.css'

import router from './router/index';
import { install } from './install'
import pinia from '@/store/index';

import './assets/css/style.css'
import './assets/css/vars.scss'


import App from './App.vue'




let app = createApp(App)
app.use(router)
app.use(pinia)
app.use(install)


app.mount('#app')
