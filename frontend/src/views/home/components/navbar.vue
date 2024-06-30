<template>
  <div class="flex flex-row items-center w-full bg-white shadow">
    <div
      v-for="(item, index) in navBar"
      :key="index"
      class="flex flex-row h-8 items-center first:ml-1"
    >
      <el-dropdown
        :teleported="false"
        trigger="contextmenu"
        @command="(commond) => onMenuCommand(commond, item)"
        ref="itemCtxMenuRef"
        class="h-full items-center"
      >
        <div
          v-if="item.index == activeMenu"
          class="flex items-center justify-between font-normal text-[#3395FF] h-4/5 w-[100px] pl-2 pr-2 ml-[5px] border border-solid border-[#d9d9d9] rounded-sm hover:border-[#E0EFFF] text-xs"
          @click="onMenuItemClick(item)"
        >
          <span>{{ item.title }} </span>
          <Icon
            icon="carbon:close-outline"
            @click.stop="onMenuRemoveClick(item, 0)"
            v-if="item.index !== dashboard.index"
          />
        </div>
        <div
          v-else
          class="flex items-center justify-between font-normal text-[#595959] h-4/5 w-[100px] pl-2 pr-2 ml-[5px] border border-solid border-[#d9d9d9] rounded-sm hover:border-[#E0EFFF] text-xs group"
          @click="onMenuItemClick(item)"
        >
          {{ item.title }}
          <Icon
            icon="carbon:close-outline"
            @click.stop="onMenuRemoveClick(item, 1)"
            v-if="item.index !== dashboard.index"
            class="hidden group-hover:block group-hover: text-[#3395FF]"
          />
        </div>
        <template #dropdown>
          <el-dropdown-menu @mouseleave.native="onMouseleave(index)">
            <el-dropdown-item command="all">关闭全部</el-dropdown-item>
            <el-dropdown-item command="current">关闭当前</el-dropdown-item>
            <el-dropdown-item command="other">关闭其他</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useMenuStore } from "@/store/menu";
import { storeToRefs } from "pinia";
import { useRouter } from "vue-router";
import { Icon } from "@iconify/vue";

const router = useRouter();
const menuStore = useMenuStore();
const { navBar, activeMenu, dashboard } = storeToRefs(menuStore);

const onMenuItemClick = (item) => {
  menuStore.setActiveMenu(item);
};

const itemCtxMenuRef = ref([]);

const onMouseleave = (index) => {
  let ctxRef = itemCtxMenuRef.value[index];
  console.log("onMouseleave", ctxRef);

  itemCtxMenuRef.value[index].handleClose();
};

const onMenuCommand = (commond, item) => {
  switch (commond) {
    case "all":
      menuStore.removeAllBar();
      break;
    case "current":
      menuStore.removeCurrentBar(item);
      break;
    case "other":
      menuStore.removeOtherBar(item);
      break;
    default:
      break;
  }
};

const onMenuRemoveClick = (item, type) => {};

const onMenuClick = (item) => {
  router.push({
    path: item.path,
  });
};
</script>