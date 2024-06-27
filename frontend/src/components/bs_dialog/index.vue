<template>
  <el-dialog
    draggable
    v-model="props.visible"
    :label="props.title"
    :width="props.width"
    :before-close="onClose"
    :close-on-click-modal="false"
    @open="onOpen"
  >
    <div class="dialog-body-box">
      <slot name="body"> </slot>
    </div>
    <div class="dialog-footer-box" v-if="props.showFooter">
      <slot name="footer">
        <el-button type="primary" @click="submitForm">确定</el-button>
        <el-button @click="onClose">关闭</el-button>
      </slot>
    </div>
  </el-dialog>
</template>

<script setup>
const props = defineProps({
  title: {
    type: String,
    default: "",
  },
  visible: {
    type: Boolean,
    default: false,
  },
  showFooter: {
    type: Boolean,
    default: true,
  },
  width: 500,
});

const emits = defineEmits(["close", "open", "save"]);

const onClose = () => {
  emits("close");
};

const onOpen = () => {
  emits("open");
};

const submitForm = () => {
  emits("save");
};
</script>

<style lang="scss" scoped>
.dialog-footer-box {
  padding: 20px;
  display: flex;
  justify-content: flex-end;
}

.dialog-body-box {
  padding: 20px;
}
</style>
