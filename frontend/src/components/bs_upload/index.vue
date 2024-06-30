<template>
  <el-upload
    v-model:file-list="fileList"
    :http-request="onUpload"
    action=""
    list-type="picture-card"
    :on-preview="onPictureCardPreview"
    :on-remove="onRemove"
  >
    <el-icon>
      <Plus />
    </el-icon>
  </el-upload>

  <el-image-viewer
    @close="onViewerClose"
    v-if="showViewer"
    :url-list="previewList"
  />
</template>

<script setup>
import { onMounted, reactive, ref, toRefs } from "vue";
import { apiResourceUpload, apiResourceUnUpload } from "@/apis/page/resource";
import { getImgPath } from "@/utils/utils";

const props = defineProps({
  name: "",
});
const emits = defineEmits(["upload", "unUpload"]);

const previewList = ref([]);
const showViewer = ref(false);

const fileList = ref([]);
const dataForm = reactive({
  name: "",
  ext: "",
});

const displayAdd = ref("none");

onMounted(() => {
  fileList.value = [];
  previewList.value = [];
  dataForm.name = "";
  dataForm.ext = "";
  showViewer.value = false;

  // console.log("props.name", props.name);

  if (props.name) {
    const info = props.name.split(".");
    dataForm.name = info[0];
    dataForm.ext = `.${info[1]}`;

    fileList.value.push({ url: getImgPath(dataForm.name, dataForm.ext) });
  }

  changeDisplayAdd();
});

const changeDisplayAdd = () => {
  displayAdd.value = fileList.value.length == 0 ? "inline-flex" : "none";
};

const onUpload = async (fileObject) => {
  let fd = new FormData(); // 新建一个FormData()对象，这就相当于你新建了一个表单
  fd.append("file", fileObject.file);
  fd.append("type", "");

  let { data } = await apiResourceUpload(fd);

  dataForm.name = data.name;
  dataForm.ext = data.ext;

  emits("upload", dataForm);

  changeDisplayAdd();
};

const onRemove = async () => {
  await apiResourceUnUpload(`${dataForm.name}${dataForm.ext}`);
  fileList.value = [];

  emits("unUpload");

  changeDisplayAdd();
};

const onPictureCardPreview = () => {
  showViewer.value = true;
  previewList.value = [];

  previewList.value.push(getImgPath(dataForm.name, dataForm.ext));
};

const onViewerClose = () => {
  showViewer.value = false;
};
</script>

<style lang="scss" scoped>
::v-deep(.el-upload--picture-card) {
  display: v-bind(displayAdd);
}
</style>