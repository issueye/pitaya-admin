import { createRouter, createWebHistory } from 'vue-router'

import Login from '@/views/login.vue';
import Home from '@/views/home/index.vue';
import UserManagement from '@/views/system/user_management/index.vue';
import AuthGroupManagement from '@/views/system/auth_group_management/index.vue';
import Dashboard from '@/views/dashboard/index.vue';

import ResourceManagement from '@/views/page_mana/resource_management/index.vue';

import DevelopTool from '@/views/develop_tool/index.vue';

import { createWebHashHistory } from 'vue-router';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        redirect: '/dashboard',
        children: [
            {
                path: 'dashboard',
                name: 'Dashboard',
                component: Dashboard,
                meta: {
                    title: '首页',
                }
            },
            {
                path: 'page',
                name: 'Page',
                children: [

                    {
                        path: 'resource_management',
                        name: 'ResourceManagement',
                        component: ResourceManagement,
                    }
                ],
            },
            {
                path: 'develop',
                name: 'Develop',
                children: [
                    {
                        path: 'develop_tool',
                        name: 'DevelopTool',
                        component: DevelopTool,
                    }
                ]
            },
            {
                path: 'system',
                name: 'System',
                children: [
                    {
                        path: 'user_management',
                        name: 'UserManagement',
                        component: UserManagement,
                    },
                    {
                        path: 'auth_group_management',
                        name: 'AuthGroupManagement',
                        component: AuthGroupManagement,
                        meta: {
                            title: '用户组管理',
                        },
                    },
                    {
                        path: 'menu_management',
                        name: 'MenuManagement',
                        component: UserManagement,
                    }
                ]
            }
        ]
    },
    {
        path: '/login',
        name: 'login',
        component: Login,
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

router.beforeEach((to, from, next) => {
    if (to.path === '/login') next()

    let token = localStorage.getItem('token')
    if (!token) {
        next('/login')
    } else {
        next()
    }

    if (!from.name) {
        from.name = 'Home';
        next('/')
    }
})

export default router