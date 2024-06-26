import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useRouter } from 'vue-router';


const collapse = '100px'
const expand = '60px'
const selfDash = {
    index: '/dashboard',
    title: '首页',
}

export const useMenuStore = defineStore(
    'menu',
    () => {
        const router = useRouter()

        const isCollapse = ref(false)
        const menuWidth = ref('200px')
        const collapseLeft = ref(collapse)
        const collapseRotate = ref('none')

        const dashboard = ref({ ...selfDash })

        const changeCollapse = () => {
            isCollapse.value = !isCollapse.value;

            if (isCollapse.value) {
                menuWidth.value = '64px'
                collapseLeft.value = expand
                collapseRotate.value = 'rotate(180deg)'
            } else {
                menuWidth.value = '200px'
                collapseLeft.value = collapse
                collapseRotate.value = 'none'
            }
        }

        const activeMenu = ref(selfDash.index)

        const navBar = ref([{ ...selfDash }])

        /**
         * 添加标签
         * @param {*} data 
         */
        const setNavBar = (data) => {
            const tab = navBar.value.find(e => e.index == data.index)
            if (!tab) {
                navBar.value.push(data)
            }
        }

        /**
         * 移除标签
         * @param {*} data 
         */
        const removeBar = (data) => {
            let menu = {}
            const tab = navBar.value.find(e => e.index == data.index)
            if (tab) {
                // 查找所在的位置
                const indexNum = navBar.value.findIndex(e => e.index == data.index)
                if (indexNum == 0) {
                    menu = { ...selfDash }
                } else {
                    // 获取上一个菜单
                    const lastMenu = navBar.value[indexNum - 1];
                    menu = lastMenu;
                }

                navBar.value = navBar.value.filter(e => e.index !== data.index)
            } else {
                menu = { ...selfDash }
            }
            // console.log('removeBar:menu', menu);
            return menu
        }

        const setActiveMenu = (data) => {
            activeMenu.value = data.index;
            console.log('activeMenu', activeMenu.value);
            router.push(activeMenu.value)
        }

        return {
            // 属性
            isCollapse,
            menuWidth,
            collapseLeft,
            collapseRotate,

            dashboard,
            activeMenu,
            navBar,

            // 方法
            changeCollapse,
            setNavBar,
            removeBar,
            setActiveMenu,
        }
    },
    {
        persist: {
            key: 'menu_store',
            storage: localStorage,
        },
    })