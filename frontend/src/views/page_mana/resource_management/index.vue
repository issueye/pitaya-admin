<template>
  <BsHeader title="资源管理" description="资源管理">
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
        <vxe-table
          round
          border
          :data="tableData"
          size="mini"
          height="100%"
          stripe
          empty-text="没有数据"
          auto-resize
          :row-config="{ isCurrent: true, isHover: true }"
        >
          <vxe-column field="id" title="编码" width="150" show-overflow />
          <vxe-column field="title" title="标题" width="150" show-overflow />
          <vxe-column
            field="fileName"
            title="文件名称"
            width="200"
            show-overflow
          />
          <vxe-column field="ext" title="文件类型" width="200" show-overflow />
          <vxe-column
            field="createdAt"
            title="创建时间"
            width="200"
            show-overflow
          />
          <vxe-column field="mark" title="备注" min-width="300" show-overflow />
          <vxe-column title="操作" width="190" align="center" fixed="right">
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
          </vxe-column>
        </vxe-table>
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
    :width="600"
    :visible="visible"
    @close="onClose"
    @save="onSave"
    v-if="visible"
  >
    <template #body>
      <el-form
        label-width="auto"
        :model="dataForm"
        :rules="rules"
        ref="dataFormRef"
      >
        <el-row :gutter="20">
          <el-col>
            <el-form-item label="标题" prop="title">
              <el-input
                v-model="dataForm.title"
                placeholder="请输入标题"
                clearable
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="16">
            <el-form-item label="名称" prop="fileName">
              <el-input
                v-model="dataForm.fileName"
                placeholder="请输入名称"
                clearable
                :disabled="true"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="类型" prop="ext">
              <el-input
                v-model="dataForm.ext"
                placeholder="请输入文件类型"
                clearable
                :disabled="true"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="备注">
          <el-input
            v-model="dataForm.mark"
            placeholder="请输入备注"
            type="textarea"
            :row="2"
            clearable
          />
        </el-form-item>
        <el-form-item label="资源">
          <BsUpload
            :name="getFileName()"
            @upload="onUpload"
            @unUpload="unUpload"
          />
        </el-form-item>
      </el-form>
    </template>
  </BsDialog>
</template>

<script setup>
import { onMounted, reactive, ref } from "vue";

import {
  apiResourceList,
  apiResourceCreate,
  apiResourceModify,
  apiResourceDelete,
} from "@/apis/page/resource";
import { ElMessage, ElMessageBox } from "element-plus";

const nameTitle = "资源信息";
// 标题
const title = ref("");
// 显示弹窗
const visible = ref(false);
// 操作类型
const operationType = ref(0);
// 数据表单
const dataFormRef = ref();
// 表单验证规则
const rules = reactive({
  title: [
    {
      required: true,
      message: "请输入标题",
      trigger: "blur",
    },
  ],
  fileName: [
    {
      required: true,
      message: "请输入文件名称",
      trigger: "blur",
    },
  ],
  ext: [
    {
      required: true,
      message: "请输入文件类型",
      trigger: "blur",
    },
  ],
});
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
  fileName: "",
  path: "",
  ext: "",
  mark: "",
});

//  表格数据
const tableData = ref([]);

// 组件加载完成
onMounted(() => {
  getData();
});

const getFileName = () => {
  return dataForm.fileName + dataForm.ext;
};

const onUpload = async (data) => {
  console.log("data", data);

  dataForm.fileName = data.name;
  dataForm.ext = data.ext;
};

const unUpload = () => {
  dataForm.fileName = "";
  dataForm.ext = "";
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

const onEditClick = (value) => {
  operationType.value = 1;
  title.value = `[编辑]${nameTitle}`;
  setForm(value);
  visible.value = true;
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

const onSave = () => {
  if (!dataFormRef.value) return;
  dataFormRef.value.validate(async (valid) => {
    if (valid) {
      switch (operationType.value) {
        case 0: {
          const res = await apiResourceCreate(dataForm);
          if (res.code !== 200) {
            ElMessage.error(res.message);
            return;
          }

          ElMessage.success(res.message);
          visible.value = false;
          getData();
          break;
        }

        case 1: {
          const res = await apiResourceModify(dataForm);
          if (res.code !== 200) {
            ElMessage.error(res.message);
            return;
          }

          ElMessage.success(res.message);
          visible.value = false;
          getData();
          break;
        }
      }
    } else {
      console.log("表单验证失败");
    }
  });
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
