<template>
  <BsHeader title="代码生成" description="代码生成">
    <template #actions>
      <el-button type="primary" @click="onAddClick">添加</el-button>
    </template>
  </BsHeader>
  <BsMain>
    <template #body>
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
        </el-form-item>
      </el-form>

      <div class="table-box">
        <el-table
          border
          :data="tableData"
          highlight-current-row
          size="small"
          height="100%"
          stripe
        >
          <el-table-column
            prop="id"
            label="编码"
            width="150"
            show-overflow-tooltip
          />
          <el-table-column
            prop="title"
            label="标题"
            width="150"
            show-overflow-tooltip
          />
          <el-table-column
            prop="fileName"
            label="文件名称"
            width="200"
            show-overflow-tooltip
          />
          <el-table-column
            prop="ext"
            label="文件类型"
            width="200"
            show-overflow-tooltip
          />
          <el-table-column
            prop="createdAt"
            label="创建时间"
            width="200"
            show-overflow-tooltip
          />
          <el-table-column
            prop="mark"
            label="备注"
            min-width="300"
            show-overflow-tooltip
          />
          <el-table-column
            label="操作"
            width="190"
            align="center"
            fixed="right"
          >
            <template v-slot="{ row }">
              <el-button
                type="primary"
                link
                size="small"
                @click="onEditClick(row)"
                >编辑</el-button
              >
              <el-button
                type="danger"
                text
                size="small"
                @click="onDeleteClick(row)"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
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
  </BsMain>

  <BsDialog
    :title="title"
    :width="1000"
    :visible="visible"
    @close="onClose"
    @open="onOpen"
    v-if="visible"
    :showFooter="false"
  >
    <template #body>
      <el-steps :active="active" finish-status="success" simple>
        <el-step title="编辑模型" />
        <el-step title="预览代码" />
      </el-steps>
      <div class="w-full h-[550px] mt-10">
        <EditInfo v-if="active === 0" />
        <CodeInfo v-if="active === 1" />
        <ResultInfo v-if="active === 2" />
      </div>
      <div class="flex justify-end" v-if="active < 2">
        <el-button @click="onPrevClick">上一步</el-button>
        <el-button @click="onNextClick">下一步</el-button>
      </div>
    </template>
  </BsDialog>
</template>
  
<script setup>
import { onMounted, reactive, ref } from "vue";

import EditInfo from "./components/edit_info.vue";
import CodeInfo from "./components/code_info.vue";
import ResultInfo from "./components/result_info.vue";

import {
  apiResourceList,
  apiResourceCreate,
  apiResourceModify,
  apiResourceDelete,
} from "@/apis/page/resource";

import { apiGetAllTable, apiGetTableInfo } from "@/apis/develop/index";

import { ElMessage, ElMessageBox } from "element-plus";

import { useDevelopStore } from "@/store/develop";
import { storeToRefs } from "pinia";

const developStore = useDevelopStore();

const { active, visible } = storeToRefs(developStore);

const nameTitle = "生成代码";
// 标题
const title = ref("");
// 操作类型
const operationType = ref(0);
// 数据表单
const dataFormRef = ref();

const tableRef = ref();

// 分页
const pageNum = ref(1);
const pageSize = ref(10);
const total = ref(0);

// 查询表单
const form = reactive({
  condition: "",
});

// 弹窗表单
const dataForm = reactive({
  id: "",
  title: "",
  tableName: "",
  fileName: "",
  path: "",
  ext: "",
  mark: "",
});

const onPrevClick = () => {
  if (active.value > 0) {
    active.value--;
  }
};

const onNextClick = () => {
  if (active.value < 2) {
    active.value++;
  }
};

//  表格数据
const tableData = ref([]);

const tables = ref([]);

const columnData = ref([]);

// 组件加载完成
onMounted(() => {
  getData();
});

/**
 * 判断当前行是否处于编辑状态
 * @param {*} row
 */
const isActiveStatus = (row) => {
  const $table = tableRef.value;
  if ($table) {
    return $table.isEditByRow(row);
  }
};

/**
 * 保存当前编辑行内容
 * @param {*} row
 */
const onColumnSaveClick = (row) => {
  const $table = tableRef.value;
  if ($table) {
    $table.clearEdit().then(() => {
      ElMessage.success("保存成功");
    });
  }
};

/**
 * 编辑当前行
 * @param {*} row
 */
const onColumnEditClick = (row) => {
  const $table = tableRef.value;
  console.log("$table", $table);
  if ($table) {
    $table.setEditRow(row);
  }
};

/**
 * 取消编辑
 * @param {*} row
 */
const onColumnCancelClick = (row) => {
  const $table = tableRef.value;
  if ($table) {
    $table.clearEdit().then(() => {
      // 还原行数据
      $table.revertData(row);
    });
  }
};

const getTables = async () => {
  let res = await apiGetAllTable();
  if (res.code === 200) {
    tables.value = res.data;
  }
};

const getTableColumns = async (name) => {
  let res = await apiGetTableInfo(name);
  if (res.code === 200) {
    columnData.value = res.data.fields;
  }
};

const onChangeTable = (val) => {
  getTableColumns(val);
};

// 获取数据
const getData = async () => {
  let sendData = {
    condition: form.condition,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let res = await apiResourceList(sendData);
  if (res.code === 200) {
    tableData.value = res.data;
    total.value = res.pageInfo.total;
  }
};

// 重置表单数据
const resetForm = () => {
  dataForm.id = "";
  dataForm.title = "";
  dataForm.fileName = "";
  dataForm.ext = "";
  dataForm.mark = "";
};

// 赋值表单数据
const setForm = (value) => {
  console.log("value", value);
  dataForm.id = value.id;
  dataForm.title = value.title;
  dataForm.fileName = value.fileName;
  dataForm.ext = value.ext;
  dataForm.mark = value.mark;
};

/**
 * 添加事件
 */
const onAddClick = () => {
  operationType.value = 0;
  title.value = `[添加]${nameTitle}`;
  resetForm();
  visible.value = true;
};

const onSizeChange = (value) => {
  pageSize.value = value;
  getData();
};

const onCurrentChange = (value) => {
  pageNum.value = value;
  getData();
};

const onChange = () => {
  getData();
};

const onQryClick = () => {
  getData();
};

const onDeleteClick = (value) => {
  ElMessageBox.confirm("请确认是否要删除数据？", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      let res = await apiResourceDelete(value.id);
      if (res.code !== 200) {
        ElMessage.error(res.message);
        return;
      }

      ElMessage.success("删除成功");
      getData();
    })
    .catch(() => {
      ElMessage.info("取消删除");
    });
};

const onOpen = () => {
  developStore.resetData();
  getTables();
};

const onClose = () => {
  visible.value = false;

  if (!dataFormRef.value) return;
  dataFormRef.value.resetFields();
};
</script>
  
<style lang="scss" scoped>
.table-box {
  height: calc(100% - 45px);
}
</style>
  