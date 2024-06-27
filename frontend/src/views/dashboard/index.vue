<template>
  <el-row :gutter="20">
    <el-col :span="8">
      <el-card>
        <template #header>
          <div
            style="
              display: inline-flex;
              justify-content: space-between;
              width: 100%;
            "
          >
            <div>端口号</div>
            <el-icon><Tools /></el-icon>
          </div>
        </template>
        <div style="display: inline-flex">
          <el-statistic label="总数" :value="2" class="tic-item" />
          <el-statistic label="启用数" :value="2" class="tic-item" />
        </div>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card>
        <template #header>
          <div
            style="
              display: inline-flex;
              justify-content: space-between;
              width: 100%;
            "
          >
            <div>端口号</div>
            <el-icon><Tools /></el-icon>
          </div>
        </template>
        <div style="display: inline-flex">
          <el-statistic label="总数" :value="2" class="tic-item" />
          <el-statistic label="启用数" :value="2" class="tic-item" />
        </div>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card>
        <template #header>
          <div
            style="
              display: inline-flex;
              justify-content: space-between;
              width: 100%;
            "
          >
            <div>系统</div>
            <el-icon><Tools /></el-icon>
          </div>
        </template>
        <div style="display: inline-flex">
          <el-statistic label="总数" :value="2" class="tic-item" />
          <el-statistic label="启用数" :value="2" class="tic-item" />
        </div>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card>
        <template #header>
          <div
            style="
              display: inline-flex;
              justify-content: space-between;
              width: 100%;
            "
          >
            <div>规则</div>
            <el-icon><Tools /></el-icon>
          </div>
        </template>
        <div style="display: inline-flex">
          <el-statistic label="总数" :value="2" class="tic-item" />
          <el-statistic label="启用数" :value="2" class="tic-item" />
        </div>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card>
        <template #header>
          <div
            style="
              display: inline-flex;
              justify-content: space-between;
              width: 100%;
            "
          >
            <div>转发地址</div>
            <el-icon><Tools /></el-icon>
          </div>
        </template>
        <div style="display: inline-flex">
          <el-statistic label="总数" :value="2" class="tic-item" />
          <el-statistic label="启用数" :value="2" class="tic-item" />
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
</script>

<style lang="scss" scoped>
.el-card {
  --el-card-padding: 12px;
}

.tic-item {
  margin-right: 50px;
}
</style>
