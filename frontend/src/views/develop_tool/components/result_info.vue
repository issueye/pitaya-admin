<template>
  <div>
    <el-result
      icon="success"
      title="代码保存成功"
      sub-title="请返回主页面下载代码"
    >
      <template #extra>
        <el-button type="primary" @click="onBackClick">返回</el-button>
        <el-button type="primary" @click="onCloseClick">关闭</el-button>
      </template>
    </el-result>
  </div>
</template>

<script setup>
import { useDevelopStore } from "@/store/develop";
import { storeToRefs } from "pinia";

const developStore = useDevelopStore();

const { active, visible } = storeToRefs(developStore);

const onBackClick = () => {
  // 返回上一步
  active.value--;
};

const onCloseClick = () => {
  visible.value = false;
};
</script>