<template>
  <main class="content_container">
    <h1 class="title">充值记录</h1>

    <!-- 搜索 -->
    <section class="search">
      <el-input
        v-model="contactSearch"
        placeholder="地址,Txid"
        size="large"
        style="width: 568px;"
        prefix-icon="search"
      />
      <!-- 币种 -->
      <el-select
        v-model="currency"
        clearable
        placeholder="币种"
        style="width: 100px; margin: 0 20px;"
        size="large"
      >
        <el-option
          v-for="item in commonStore.currencyOptions"
          :key="item.name"
          :label="item.name"
          :value="item.name"
        />
      </el-select>

      <!-- 查询按钮 -->
      <el-button size="large" color="#000" type="info" round @click="paramsQuery">查询</el-button>

      <!-- 重置按钮 -->
      <el-button size="large" color="#000" plain type="info" round @click="resetQuery">重置</el-button>
    </section>

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
    </el-table>

    <!-- 分页 -->
    <div class="footerPage">
      <el-pagination
        background
        layout="prev, pager, next"
        :current-page="currentPage"
        :total="total"
        @current-change="handleChangePage"
      />
    </div>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import dayjs from 'dayjs'

import { copyMessage } from '@/utils/common.js'
import { useCommonStore } from '@/pinia/modules/common'
import { getRechargeList } from '@/api/account'

const commonStore = useCommonStore()

const currency = ref('')
const contactSearch = ref('')

const recordData = ref([])
const total = ref(0)
const currentPage = ref(1)

// 翻页
const handleChangePage = (page) => {
  currentPage.value = page
  const params = { page }
  if (currency.value) params.currency = currency.value
  if (contactSearch.value) {
    params.contactSearch = contactSearch.value
    params.txid = contactSearch.value
  }
  queryList(params)
}

// 条件查询
const paramsQuery = () => {
  currentPage.value = 1
  const params = { page: 1 }
  if (currency.value) params.currency = currency.value
  if (contactSearch.value) {
    params.contactSearch = contactSearch.value
    params.txid = contactSearch.value
  }
  queryList(params)
}

// 重置条件查询
const resetQuery = () => {
  currentPage.value = 1
  currency.value = ''
  contactSearch.value = ''
  queryList({ page: 1 })
}

const queryList = async (params) => {
  const { code, data = {} } = await getRechargeList({ ...params, pageSize: 20 })
  if (code === 0) {
    recordData.value = data.content || []
    total.value = data.total_pages || 0
  }
}

// 复制平台银行信息
const copyPlatformBankMsg = (item) => {
  copyMessage(
    `企业名称：${item.enterpriseTitle}\n银行名称：${item.bankTitle}\n银行代码：${item.bankCode}\n收款人银行账号：${item.receiverNumber}\n银行账户名：${item.receiverName}\n`
  )
}

onMounted(() => {
  if (commonStore.currencyOptions.length === 0) {
    commonStore.QueryCurrencyOptions()
  }
  queryList({ page: 1 })
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
    margin: 0 0 30px 0;
  }

  .search {
    display: flex;
    align-items: center;
    
    margin-bottom: 40px;
  }

  .tableItemColumn {
    display: flex;
    flex-direction: column;
    font-size: 16px;
    color: #000;
    line-height: 1.25;
    font-family: -inter;

    span {
      margin-top: 4px;
      color: grey;
      font-size: 14px;
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