<template>
  <div class="flex items-center mb-4">
    <span class="mr-2 text-gray-700">数据表</span>
    <div class="grow">
      <el-select
        placeholder="请选择表"
        v-model="selectTable"
        @change="onChangeTable"
      >
        <el-option
          v-for="(table, index) in tables"
          :value="table"
          :label="table"
          :key="index"
        />
      </el-select>
    </div>
  </div>

  <div>
    <vxe-table
      border
      stripe
      auto-resize
      height="475px"
      ref="tableRef"
      :data="columns"
      size="mini"
      empty-text="暂无数据"
      :column-config="{ resizable: true }"
      :row-config="{ isCurrent: true, isHover: true }"
      :edit-config="{
        trigger: 'manual',
        mode: 'row',
        autoClear: false,
        showStatus: true,
      }"
    >
      <vxe-column field="cid" title="序号" :width="80" />
      <vxe-column field="name" title="字段名" :min-width="100" show-overflow />
      <vxe-column field="type" title="类型" :width="150" show-overflow />
      <vxe-column
        field="is_primary_key"
        title="主键"
        :width="50"
        align="center"
        show-overflow
      >
        <template v-slot="{ row }">
          <el-checkbox :checked="row.is_primary_key" :disabled="true" />
        </template>
      </vxe-column>
      <vxe-column
        field="comment"
        :min-width="100"
        title="说明"
        :edit-render="{}"
        show-overflow
      >
        <template #edit="{ row }">
          <el-input v-model="row.comment" />
        </template>
      </vxe-column>
      <vxe-column
        field="capacity"
        :min-width="200"
        title="功能"
        :edit-render="{}"
        show-overflow
      >
        <template #default="{ row }">
          <el-tag
            v-for="(item, index) in row.capacity"
            :key="index"
            class="mr-2"
            size="small"
          >
            {{ getName(item) }}
          </el-tag>
        </template>
        <template #edit="{ row }">
          <el-select v-model="row.capacity" multiple>
            <el-option value="isCondition" label="条件" />
            <el-option value="isAsc" label="正排序" />
            <el-option value="isDesc" label="倒排序" />
          </el-select>
        </template>
      </vxe-column>
      <vxe-column title="操作" :width="100" fixed="right">
        <template #default="{ row }">
          <div v-if="isActiveStatus(row)">
            <el-button type="primary" link @click="onColumnSaveClick(row)"
              >保存</el-button
            >
            <el-button link @click="onColumnCancelClick(row)">取消</el-button>
          </div>
          <div v-else>
            <el-button type="primary" link @click="onColumnEditClick(row)"
              >编辑</el-button
            >
          </div>
        </template>
      </vxe-column>
    </vxe-table>
  </div>
</template>


<script setup>
import { onMounted, ref } from "vue";
import { useDevelopStore } from "@/store/develop";
import { storeToRefs } from "pinia";
import { ElMessage } from "element-plus";

const developStore = useDevelopStore();

const { tables, columns, selectTable } = storeToRefs(developStore);

const tableRef = ref();

onMounted(() => {
  getTables();
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
      // 处理数据
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
  developStore.getTables();
};

const onChangeTable = (val) => {
  selectTable.value = val;
  developStore.getColumns();
};

const getName = (capacity) => {
  switch (capacity) {
    case "isCondition":
      return "条件";
    case "isAsc":
      return "正排序";
    case "isDesc":
      return "倒排序";
    default:
      return "";
  }
};
</script>