import { iconifyInstall } from './iconify'
import { vxeTableInstall } from './vxe-table'
import { customComponentsInstall } from './custom_components'
import { elementInstall } from './element'


// 安装
export const install = (app) => {
    // element 
    elementInstall(app)

    // 图标库
    iconifyInstall()

    // vxe-table
    vxeTableInstall(app)

    // custom_components
    customComponentsInstall(app)
}