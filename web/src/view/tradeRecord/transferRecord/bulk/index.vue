<template>
  <main class="content_container">
    <h1 class="title">批量转账记录</h1>
    
    <section class="search">
      <div class="left">
        <el-input
          v-model="searchString"
          placeholder="ID、任务"
          prefix-icon="search"
          size="large"
          style="width: 600px; margin-right: 16px;"
        />
        <!-- 状态 -->
        <el-select
          v-model="selectStatus"
          clearable
          placeholder="状态"
          style="width: 100px"
          size="large"
        >
          <el-option
            v-for="item in statusOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </div>
    </section>

    <el-table :data="recordData" style="width: 100%">
      <el-table-column label="时间"></el-table-column>
      <el-table-column label="状态"></el-table-column>
      <el-table-column label="任务"></el-table-column>
      <el-table-column label="ID"></el-table-column>
      <el-table-column label="支付人数"></el-table-column>
      <el-table-column label="燃料费"></el-table-column>
      <el-table-column label="手续费"></el-table-column>
      <el-table-column label="支付总金额"></el-table-column>
      <el-table-column label="操作"></el-table-column>
    </el-table>
  </main>
</template>

<script setup>
import { ref, reactive } from 'vue'

const searchString = ref('')
const selectStatus = ref('')

const statusOptions = reactive([
  { label: '已创建', value: '1' },
  { label: '已完成', value: '2' },
  { label: '进行中', value: '3' },
  { label: '错误', value: '4' },
  { label: '待审核', value: '5' },
])

const recordData = []
</script>

<style lang="scss" scoped>
.content_container {
  box-sizing: border-box;
  display: flex;
  flex-direction: column;

  padding: 10px;
  color: #000;

  /* 抗锯齿渲染文字 */
  -webkit-font-smoothing: antialiased;

  .title {
    font-size: 30px;
    margin: 0 0 30px 0;
  }

  .search {
    display: flex;
    align-items: center;
    justify-content: space-between;
    
    margin-bottom: 40px;

    .left {
      display: flex;
      align-items: center;
    }
  }
}

:deep(.el-input--large .el-input__wrapper),
:deep(.el-select--large .el-select__wrapper) {
  border-radius: 20px;
}

:deep(.el-input__wrapper.is-focus),
:deep(.el-select__wrapper.is-focused) {
  box-shadow: 0 0 0 2px rgba(0,0,0,0.2);
}
</style>