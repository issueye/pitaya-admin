<template>
  <div class="w-full h-full">
    <el-container class="h-full">
      <el-header>
        <Header />
      </el-header>
      <div class="flex flex-row w-full h-full">
        <Menu class="shadow-xl z-50" :style="{ width: menuWidth }" />
        <div
          class="flex flex-col bg-[#F0F2F5]"
          :style="{ width: `calc(100% - ${menuWidth})` }"
        >
          <NavBar />
          <div class="m-[10px] grow">
            <RouterView />
          </div>
        </div>
      </div>
    </el-container>
  </div>
</template>

<script setup>
import { storeToRefs } from "pinia";
import Header from "./components/header.vue";
import Menu from "./components/menu.vue";
import NavBar from "./components/navbar.vue";

import { useMenuStore } from "@/store/menu";

const menuStore = useMenuStore();

const { menuWidth } = storeToRefs(menuStore);
console.log("menuWidth", menuWidth);
</script>
