<template>
  <BsHeader title="用户组管理" description="用户组管理">
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
          size="small"
          highlight-current-row
          stripe
        >
          <el-table-column
            prop="id"
            label="编码"
            show-overflow-tooltip
            width="150"
          />
          <el-table-column prop="name" label="姓名" width="150" />
          <el-table-column
            prop="mark"
            label="备注"
            show-overflow-tooltip
            min-width="300"
          />
          <el-table-column
            prop="state"
            label="状态"
            width="70"
            align="center"
            fixed="right"
          >
            <template v-slot="{ row }">
              <el-tag
                effect="plain"
                :type="row.state === 1 ? 'success' : 'danger'"
              >
                {{ row.state === 1 ? "启用" : "停用" }}
              </el-tag>
            </template>
          </el-table-column>
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
                @click="onEditStateClick(row)"
                >{{ row.state === 0 ? "启用" : "停用" }}</el-button
              >
              <el-button
                type="primary"
                link
                size="small"
                @click="onEditClick(row)"
                >编辑</el-button
              >
              <el-button
                type="danger"
                link
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
    :width="500"
    :visible="visible"
    @close="onClose"
    @save="onSave"
  >
    <template #body>
      <el-form
        label-width="auto"
        :model="dataForm"
        :rules="rules"
        ref="dataFormRef"
      >
        <el-form-item label="名称" prop="name">
          <el-input
            v-model="dataForm.name"
            placeholder="请输入账户"
            :disabled="dataForm.sys === 1 && operationType === 1"
            clearable
          />
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model="dataForm.mark"
            placeholder="请输入备注"
            :disabled="dataForm.sys === 1 && operationType === 1"
            type="textarea"
            :row="2"
            clearable
          />
        </el-form-item>
      </el-form>
    </template>
  </BsDialog>
</template>
  
  <script setup>
import { onMounted, reactive, ref } from "vue";

import {
  apiUserGroupList,
  apiUserGroupCreate,
  apiUserGroupModify,
  apiUserGroupModifyState,
  apiUserGroupDelete,
} from "@/apis/sys/user";
import { ElMessage, ElMessageBox } from "element-plus";

const nameTitle = "用户组信息";
// 标题
const title = ref("用户组信息");
// 显示弹窗
const visible = ref(false);
// 操作类型
const operationType = ref(0);
// 数据表单
const dataFormRef = ref();
// 表单验证规则
const rules = reactive({
  name: [{ required: true, message: "请输入组名称", trigger: "blur" }],
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
  name: "",
  mark: "",
  sys: 0,
});

//  表格数据
const tableData = ref([]);

// 组件加载完成
onMounted(() => {
  getData();
});

// 获取数据
const getData = async () => {
  let sendData = {
    condition: form.condition,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let res = await apiUserGroupList(sendData);
  if (res.code === 200) {
    tableData.value = res.data;
    total.value = res.pageInfo.total;
  }
};

// 重置表单数据
const resetForm = () => {
  dataForm.id = "";
  dataForm.name = "";
  dataForm.mark = "";
  dataForm.sys = 0;
};

// 赋值表单数据
const setForm = (value) => {
  console.log("value", value);
  dataForm.id = value.id;
  dataForm.name = value.name;
  dataForm.mark = value.mark;
  dataForm.sys = value.sys;
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

const onCurrentChange = () => {
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

const onEditStateClick = async (value) => {
  const res = await apiUserGroupModifyState(value.id);
  if (res.code !== 200) {
    ElMessage.error(res.message);
    return;
  }

  ElMessage.success(res.message);
  getData();
};

const onDeleteClick = (value) => {
  ElMessageBox.confirm("请确认是否要删除数据？", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      let res = await apiUserGroupDelete(value.id);
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
          const res = await apiUserGroupCreate(dataForm);
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
          const res = await apiUserGroupModify(dataForm);
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
  
  <style lang="scss">
.table-box {
  height: calc(100% - 45px);
}
</style>
  