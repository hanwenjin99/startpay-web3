<template>
  <main class="content_container">
    <h1 class="title">提现管理</h1>

    <!-- 表格 -->
    <el-table :data="recordData" style="width: 100%">
      <el-table-column label="时间">
        <template #default="scope">
          <div class="tableItemColumn">
            {{ dayjs(scope.row.createTime).format('YYYY-MM-DD') }}
            <span>{{ dayjs(scope.row.createTime).format('HH:mm') }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="状态" prop="statusName" />
      <el-table-column label="到">
        <template #default="scope">{{ scope.row.bankAccount?.bankTitle }}</template>
      </el-table-column>
      <el-table-column label="币种" prop="currency" />
      <el-table-column label="服务费">
        <template #default="scope">
          {{ `${Number(scope.row.fee ?? 0).toFixed(2)}${scope.row.currency}` }}
        </template>
      </el-table-column>
      <el-table-column label="手续费">
        <template #default="scope">
          {{ `${scope.row.remittanceFee}${scope.row.currency}` }}
        </template>
      </el-table-column>
      <el-table-column label="扣款金额">
        <template #default="scope">
          <div class="tableItemColumn">
            {{ `-${scope.row.totalAmount}` }}
            <span>{{ `$${scope.row.totalAmount}` }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="到账金额">
        <template #default="scope">
          <div class="tableItemColumn">
            {{ scope.row.amount }}
            <span>{{ scope.row.amount }}</span>
          </div>
        </template>
      </el-table-column>
      <!-- 新增交易/审核附言 -->
      <el-table-column label="交易附言" prop="inputNote" />
      <el-table-column label="审核附言" prop="adminMemo" />
      <!-- 管理后台 - 审核 -->
      <el-table-column label="操作">
        <template #default="scope">
          <el-button link @click="reviewWithdraw(scope.row)">审核</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="footerPage">
      <el-pagination background layout="prev, pager, next" :total="total" @current-change="handleChangePage" />
    </div>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import dayjs from 'dayjs'

import { getAdminWithdrawOrderList, reviewAdminWithdrawOrder } from '@/api/account'

const recordData = ref([])
const total = ref(0)

const queryList = async (page) => {
  const { code, data = {} } = await getAdminWithdrawOrderList({ page, pageSize: 10 })
  if (code === 0) {
    recordData.value = data.content ?? []
    total.value = data.total_pages ?? 0
  }
}

// 管理后台审核提现记录
const reviewWithdraw = (item) => {
  // 填写审核备注信息
  ElMessageBox.prompt('请填写审核备注信息', '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消'
  }).then(async ({ value }) => {
    const { code } = await reviewAdminWithdrawOrder({
      id: item.id,
      chain: item.chain,
      currency: item.currency,
      status: item.status,
      memo: value // 审核备注信息
    })
    if (code === 0) {
      ElMessage.success('审核成功！')
      // 刷新列表
      queryList(1)
    }
  })
}

// 分页请求
const handleChangePage = (page) => {
  queryList(page)
}

onMounted(() => {
  queryList(1)
})

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