import BsHeader from '@/components/bs_header/index.vue';
import BsMain from '@/components/bs_main/index.vue';
import BsDialog from '@/components/bs_dialog/index.vue';
import BsUpload from '@/components/bs_upload/index.vue';
import BsResources from '@/components/bs_resources/index.vue';

import svgIcon from "@/components/SvgIcon/index.vue";
import 'virtual:svg-icons-register';


export const customComponentsInstall = (app) => {
    app.component('SvgIcon', svgIcon)
    app.component('BsHeader', BsHeader)
    app.component('BsMain', BsMain)
    app.component('BsDialog', BsDialog)
    app.component('BsUpload', BsUpload)
    app.component('BsResources', BsResources)
}