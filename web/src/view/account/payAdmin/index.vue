<template>
  <main class="content_container">
    <h1 class="title">支付管理</h1>

    <!-- 表格 -->
    <el-table :data="recordData" style="width: 100%">
      <el-table-column label="时间"></el-table-column>
      <el-table-column label="状态"></el-table-column>
      <el-table-column label="到"></el-table-column>
      <el-table-column label="币种"></el-table-column>
      <el-table-column label="服务费"></el-table-column>
      <el-table-column label="手续费"></el-table-column>
      <el-table-column label="扣款金额"></el-table-column>
      <el-table-column label="到账金额"></el-table-column>
      <el-table-column label="操作"></el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="footerPage">
      <el-pagination background layout="prev, pager, next" :total="total" @current-change="handleChangePage" />
    </div>
  </main>
</template>

<script setup>
import { ref } from 'vue'

const recordData = ref([])
const total = ref(0)

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
    margin: 0;
    margin-bottom: 20px;
  }
}

.footerPage {
  width: 100%;
  display: flex;
  justify-content: center;
}

:deep(.el-pagination.is-background .el-pager li),
:deep(.el-pagination.is-background .btn-prev),
:deep(.el-pagination.is-background .btn-next) {
  border-radius: 24px;
}

:deep(.el-pagination.is-background .el-pager li.is-active) {
  background-color: #000;
}

:deep(.el-pagination.is-background .el-pager li):hover {
  color: #000;
}
:deep(.el-pagination.is-background .el-pager li.is-active):hover {
  color: #fff;
}
</style>