<template>
  <el-form inline>
    <el-form-item label="检索">
      <el-input
        v-model="form.condition"
        placeholder="请输入检索内容"
        clearable
        @change="onChange"
      />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onQryClick">查询</el-button>
      <el-button type="primary" @click="onSelectClick">选择</el-button>
      <el-button type="primary" @click="onCancelClick">取消</el-button>
    </el-form-item>
  </el-form>

  <div class="table-box">
    <el-scrollbar :height="250">
      <div class="table-items">
        <div
          v-for="(item, index) in tableData"
          :key="index"
          style="margin: 2px; padding: 2px"
          :style="{ border: imgVars[index].color }"
          @click="onImgClick(index)"
        >
          <el-image
            style="width: 100px; height: 100px; margin: 2px"
            :src="getPath(item.fileName, item.ext)"
            fit="cover"
          />
          <div>{{ item.title }}</div>
        </div>
      </div>
    </el-scrollbar>
  </div>
  <div style="margin-top: 10px">
    <el-pagination
      small
      background
      :current-page="pageNum"
      :page-size="pageSize"
      :page-sizes="[5, 10, 20]"
      :total="total"
      layout="total, sizes, prev, pager, next"
      @size-change="onSizeChange"
      @current-change="onCurrentChange"
    />
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from "vue";
import { apiResourceList } from "@/apis/page/resource";
import { getImgPath } from "@/utils/utils";

const emits = defineEmits(["onSelect", "onCancel"]);
// 分页
const pageNum = ref(1);
const pageSize = ref(10);
const total = ref(0);

// 查询表单
const form = reactive({
  condition: "",
});

const imgVars = ref({});

//  表格数据
const tableData = ref([]);

onMounted(() => {
  getData();
});

const onQryClick = () => {
  getData();
};

const onChange = () => {
  getData();
};

const getPath = (name, ext) => {
  console.log("getPath", name, ext);
  return getImgPath(name, ext);
};

const onSelectClick = () => {
  for (let key in imgVars.value) {
    let item = imgVars.value[key];

    if (item.click) {
      emits("onSelect", item.name);
      return;
    }
  }
};

const onCancelClick = () => {
  emits("onCancel")
}


const onImgClick = (index) => {
  for (let key in imgVars.value) {
    let item = imgVars.value[key];

    console.log("item", item);
    item.color = "1px solid #D9D9D9";
    item.click = false;
  }

  imgVars.value[index].click = true;
  imgVars.value[index].color = "1px solid #397AFD";
};

const getData = async () => {
  let sendData = {
    condition: form.condition,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let res = await apiResourceList(sendData);
  if (res.code === 200) {
    const data = [...res.data];

    for (let index = 0; index < data.length; index++) {
      imgVars.value[index] = {
        color: "1px solid #D9D9D9",
        click: false,
        name: `${data[index].fileName}${data[index].ext}`,
      };
    }

    tableData.value = res.data;
    total.value = res.pageInfo.total;
  }
};

const onSizeChange = (value) => {
  pageSize.value = value;
  getData();
};

const onCurrentChange = (value) => {
  pageNum.value = value;
  getData();
};
</script>

<style lang="scss" scoped>
.table-box {
  height: 300px;

  .table-items {
    display: flex;
    flex-wrap: wrap;
    align-content: flex-start;
  }
}
</style>