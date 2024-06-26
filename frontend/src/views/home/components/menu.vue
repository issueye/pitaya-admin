<template>
  <div class="menu-box">
    <el-menu
      router
      :collapse-transition="false"
      :unique-opened="true"
      :collapse="isCollapse"
      :default-active="activeMenu"
      :default-openeds="defaultOpeneds"
      @select="onSelect"
    >
      <el-menu-item index="/dashboard" style="border-bottom: 1px solid #d9d9d9">
        <el-icon>
          <Icon icon="mdi-light:home" />
        </el-icon>
        <template #title>
          <span>首页</span>
        </template>
      </el-menu-item>
      <el-sub-menu index="/page" style="border-bottom: 1px solid #d9d9d9">
        <template #title>
          <el-icon>
            <Icon icon="dashicons:text-page" />
          </el-icon>
          <span>页面管理</span>
        </template>
        <el-menu-item index="/page/resource_management">资源管理</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="/system">
        <template #title>
          <el-icon>
            <Setting />
          </el-icon>
          <span>系统设置</span>
        </template>

        <el-menu-item index="/system/user_management">用户管理</el-menu-item>
        <el-menu-item index="/system/auth_group_management"
          >用户组管理</el-menu-item
        >
      </el-sub-menu>
    </el-menu>

    <div class="collapse-box">
      <!-- <img src="@/assets/collapse.png" alt="" @click="collapseClick" /> -->
      <svg-icon
        class="cvg-collapse-box"
        iconName="collapse"
        size="20"
        @click="collapseClick"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useMenuStore } from "@/store/menu";
import { storeToRefs } from "pinia";
import { Icon } from "@iconify/vue";

const menus = [
  {
    index: "/dashboard",
    title: "首页",
    icon: "mdi-light:home",
  },
  {
    index: "/page",
    title: "页面管理",
    icon: "dashicons:text-page",
    children: [
      {
        index: "/page/page_management",
        title: "页面管理",
        icon: "iconoir:page",
      },
      {
        index: "/page/target_management",
        title: "服务地址管理",
        icon: "mdi:server",
      },
      {
        index: "/page/resource_management",
        title: "资源管理",
        icon: "carbon:software-resource-cluster",
      },
      {
        index: "/page/http_message_query",
        title: "转发记录查询",
        icon: "carbon:software-resource-cluster",
      },
    ],
  },
  {
    index: "/system",
    title: "系统设置",
    icon: "grommet-icons:system",
    children: [
      { index: "/system/user_management", title: "用户管理", icon: "mdi:user" },
      {
        index: "/system/auth_group_management",
        title: "用户组管理",
        icon: "mdi:user-group",
      },
    ],
  },
];

const menuStore = useMenuStore();
const { isCollapse, collapseLeft, collapseRotate, navBar, activeMenu } =
  storeToRefs(menuStore);

const defaultOpeneds = ref([]);

const collapseClick = () => {
  menuStore.changeCollapse();
};

const findMenu = (index) => {
  let rtnData = null;
  menus.find((e) => {
    if (e.children) {
      if (e.children.length > 0) {
        const child = e.children;
        const menu = child.find((e) => e.index == index);
        if (menu) {
          rtnData = menu;
        }
      } else {
        if (e.index == index) {
          rtnData = e;
        }
      }
    } else {
      if (e.index == index) {
        rtnData = e;
      }
    }
  });

  return rtnData;
};

const onSelect = (value) => {
  const menu = findMenu(value);
  if (menu) {
    menuStore.setNavBar(menu);
    menuStore.setActiveMenu(menu);
  }
};
</script>

<style lang="scss" scoped>
.menu-box {
  .collapse-box {
    position: absolute;
    left: v-bind(collapseLeft);
    top: 50%;

    .cvg-collapse-box {
      width: 70px;
      height: 30px;
      transform: v-bind(collapseRotate);
    }
  }
}
</style>
