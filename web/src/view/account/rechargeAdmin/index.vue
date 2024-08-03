<template>
  <main class="content_container">
    <h1 class="title">充值管理</h1>

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
      <el-table-column label="汇款银行">
        <template #default="scope">{{ scope.row.bankAccount?.bankTitle }}</template>
      </el-table-column>
      <el-table-column label="钱包" prop="currency" />
      <!-- 平台银行 -->
      <el-table-column label="平台银行">
        <template #default="scope">
          <el-popover
            placement="top"
            width="300"
            title="详细信息"
          >
            <template #reference>
              {{ scope.row.platformBank?.bankTitle }}
            </template>
            <div class="platformBank">
              <span>企业名称：{{ scope.row.platformBank?.enterpriseTitle }}</span>
              <span>银行名称：{{ scope.row.platformBank?.bankTitle }}</span>
              <span>银行代码：{{ scope.row.platformBank?.bankCode }}</span>
              <span>收款人银行账号：{{ scope.row.platformBank?.receiverNumber }}</span>
              <span>银行账户名：{{ scope.row.platformBank?.receiverName }}</span>
              <el-button plain color="#000" @click="copyPlatformBankMsg(scope.row.platformBank)">复制</el-button>
            </div>
          </el-popover>
        </template>
      </el-table-column>
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
      <!-- 商户端 - 审核｜撤销 -->
      <el-table-column label="操作">
        <template #default="scope">
          <el-popconfirm
            confirm-button-text="确定"
            cancel-button-text="取消"
            icon="infoFilled"
            icon-color="#626AEF"
            title="确认审核此条记录？"
            @confirm="reviewRecharge(scope.row)"
          >
            <template #reference>
              <el-button link>审核</el-button>
            </template>
          </el-popconfirm>
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

import { copyMessage } from '@/utils/common.js'
import { getAdminRechargeList, reviewAdminRecharge } from '@/api/account'

const recordData = ref([])
const total = ref(0)

const queryList = async (page) => {
  const { code, data = {} } = await getAdminRechargeList({ page, pageSize: 10 })
  if (code === 0) {
    recordData.value = data.content ?? []
    total.value = data.total_pages ?? 0
  }
}

// 管理后台审核充值记录
const reviewRecharge = (item) => {
  // 填写审核备注信息
  ElMessageBox.prompt('请填写审核备注信息', '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消'
  }).then(async ({ value }) => {
    const { code } = await reviewAdminRecharge({
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

// 复制平台银行信息
const copyPlatformBankMsg = (item) => {
  copyMessage(
    `企业名称：${item.enterpriseTitle}\n银行名称：${item.bankTitle}\n银行代码：${item.bankCode}\n收款人银行账号：${item.receiverNumber}\n银行账户名：${item.receiverName}\n`
  )
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