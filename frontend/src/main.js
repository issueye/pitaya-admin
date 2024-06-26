import { createApp } from 'vue'
import 'normalize.css/normalize.css'

import ElementPlus from 'element-plus'

import 'vxe-table/lib/style.css'
import '@/assets/css/tailwind.css'
import 'element-plus/dist/index.css'
import './assets/css/style.css'
import './assets/css/vars.scss'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import pinia from '@/store/index';

import svgIcon from "@/components/SvgIcon/index.vue";
import 'virtual:svg-icons-register';

import BsHeader from '@/components/bs_header/index.vue';
import BsMain from '@/components/bs_main/index.vue';
import BsDialog from '@/components/bs_dialog/index.vue';
import BsUpload from '@/components/bs_upload/index.vue';
import BsResources from '@/components/bs_resources/index.vue';
import { iconifyInstall } from '@/iconify/index'
import { directive } from 'vue3-menus';

import router from './router/index';
import App from './App.vue'



// 安装图标库
iconifyInstall()

let app = createApp(App)
app.use(router)
app.use(pinia)
app.use(ElementPlus)
app.component('SvgIcon', svgIcon)
app.component('BsHeader', BsHeader)
app.component('BsMain', BsMain)
app.component('BsDialog', BsDialog)
app.component('BsUpload', BsUpload)
app.component('BsResources', BsResources)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.directive('menus', directive); // 只注册指令

app.mount('#app')
