<template>
  <div class="header-box">
    <div class="grape-title">
      <img
        src="@/assets/images/icon.webp"
        style="width: 24px; height: 24px"
      /><span style="margin-left: 10px">一个小工具</span>
    </div>
    <div class="header-actions">
      <el-avatar :size="25">{{ info.name ? info.name[0] : "A" }}</el-avatar>
      <el-dropdown class="user-name" @command="dropdownClick">
        <span class="el-dropdown-link" style="margin-left: 10px">
          {{ info.name }}
          <el-icon>
            <ArrowDown />
          </el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="修改密码">
              <svg-icon
                iconName="changepwd"
                color="#D9D9D9"
                :size="15"
                style="margin-right: 10px"
              />
              修改密码
            </el-dropdown-item>
            <el-dropdown-item command="退出登录">
              <svg-icon
                iconName="logout"
                color="#D9D9D9"
                :size="15"
                style="margin-right: 10px"
              />
              退出登录
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import { ArrowDown, Switch, BottomRight } from "@element-plus/icons-vue";
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/user";
import { storeToRefs } from "pinia";

const userStore = useUserStore();
let { info } = storeToRefs(userStore);

const route = useRouter();

const dropdownClick = (value) => {
  switch (value) {
    case "修改密码": {
      break;
    }
    case "退出登录": {
      route.push("/login");
      localStorage.clear();
      sessionStorage.clear();
      break;
    }
  }
};
</script>

<style lang="scss" scoped>
.header-box {
  height: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .grape-title {
    color: #fff;
    font-size: 20px;
    font-weight: 600;
    display: inline-flex;
    align-items: center;
  }

  .header-actions {
    margin-right: 10px;
    display: inline-flex;
    align-items: center;
  }

  .el-dropdown-link:focus {
    outline: none;
  }

  .user-name {
    display: inline-block;
    font-size: 15px;
    font-weight: bold;
    vertical-align: middle;
    color: white;
  }
}
</style>
