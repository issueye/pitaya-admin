<template>
  <div class="flex flex-row w-full bg-white shadow">
    <div
      v-for="(item, index) in navBar"
      :key="index"
      class="flex flex-row h-8 items-center first:ml-1"
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
    </div>
  </div>
</template>

<script setup>
import { useMenuStore } from "@/store/menu";
import { storeToRefs } from "pinia";
import { useRouter } from "vue-router";
import { Icon } from "@iconify/vue";

const router = useRouter();
const menuStore = useMenuStore();
const { navBar, activeMenu, dashboard } = storeToRefs(menuStore);

const rightMenuData = [
  { label: "关闭当前", value: "now" },
  { label: "关闭其他", value: "other" },
  { label: "关闭全部", value: "all" },
];

const onMenuItemClick = (item) => {
  menuStore.setActiveMenu(item);
};

const onMenuRemoveClick = async (item, type) => {
  const menu = menuStore.removeBar(item, type);
  console.log("menu", menu);

  if (type == 0) {
    menuStore.setActiveMenu(menu);
  }
};
</script>