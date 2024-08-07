<template>
  <BsHeader title="用户管理" description="用户管理">
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
          <el-table-column prop="id" label="编码" width="150" show-overflow />
          <el-table-column prop="account" label="账户" width="130" />
          <el-table-column prop="name" label="姓名" width="150" />
          <el-table-column prop="groupName" label="用户组" width="150" />
          <el-table-column prop="mark" label="备注" />
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
                >{{ row.state === 1 ? "停用" : "启用" }}</el-button
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
    @open="onOpen"
  >
    <template #body>
      <el-form
        label-width="auto"
        :model="dataForm"
        :rules="rules"
        ref="dataFormRef"
      >
        <el-form-item label="账户" prop="account">
          <el-input
            v-model="dataForm.account"
            placeholder="请输入账户"
            clearable
          />
        </el-form-item>
        <el-form-item label="姓名" prop="name">
          <el-input
            v-model="dataForm.name"
            placeholder="请输入账户"
            clearable
          />
        </el-form-item>
        <el-form-item label="用户组" prop="groupId">
          <el-select
            v-model="dataForm.groupId"
            value-key=""
            placeholder=""
            clearable
            filterable
          >
            <el-option
              v-for="item in groupTableData"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
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
  apiUserList,
  apiUserCreate,
  apiUserModify,
  apiUserDelete,
  apiUserModifyState,
  apiUserGroupList,
} from "@/apis/sys/user";
import { ElMessage, ElMessageBox } from "element-plus";

const nameTitle = "用户信息";
// 标题
const title = ref("用户信息");
// 显示弹窗
const visible = ref(false);
// 操作类型
const operationType = ref(0);
// 数据
const dataFormRef = ref();

const rules = reactive({
  name: [{ required: true, message: "请输入姓名", trigger: "blur" }],
  account: [{ required: true, message: "请输入账户", trigger: "blur" }],
  groupId: [{ required: true, message: "请选择用户组", trigger: "blur" }],
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
  account: "",
  name: "",
  groupId: "",
  sys: 0,
  mark: "",
});

onMounted(() => {});

const tableData = ref([]);
const groupTableData = ref([]);

const getData = async () => {
  let sendData = {
    pageSize: pageSize.value,
    pageNum: pageNum.value,
    condition: form.condition,
  };
  let res = await apiUserList(sendData);
  if (res.code === 200) {
    tableData.value = res.data;
    total.value = res.pageInfo.total;
  }
};

// 重置表单数据
const resetForm = () => {
  dataForm.id = "";
  dataForm.account = "";
  dataForm.name = "";
  dataForm.groupId = "";
  dataForm.sys = 0;
  dataForm.mark = "";
};

// 赋值表单数据
const setForm = (value) => {
  dataForm.id = value.id;
  dataForm.account = value.account;
  dataForm.name = value.name;
  dataForm.groupId = value.groupId;
  dataForm.sys = value.sys;
  dataForm.mark = value.mark;
};

onMounted(() => {
  getData();
});

// 事件

/**
 * 添加事件
 */
const onAddClick = () => {
  operationType.value = 0;
  title.value = `[添加]${nameTitle}`;
  resetForm();
  visible.value = true;
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
  const res = await apiUserModifyState(value.id);
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
      let res = await apiUserDelete(value.id);
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
          const res = await apiUserCreate(dataForm);
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
          const res = await apiUserModify(dataForm);
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

// 弹窗打开时
const onOpen = async () => {
  const res = await apiUserGroupList();
  if (res.code === 200) {
    groupTableData.value = res.data;
  }
};

const onSizeChange = (value) => {
  pageSize.value = value;
  getData();
};

const onCurrentChange = () => {
  getData();
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
