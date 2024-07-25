<template>
  <main class="content_container">
    <h1 class="title">取消记录</h1>

    <!-- 搜索 -->
    <section class="search">
      <!-- 网络 -->
      <el-select
        v-model="chain"
        clearable
        placeholder="网络"
        style="width: 100px"
        size="large"
      >
        <el-option
          v-for="item in commonStore.chainsList"
          :key="item.name"
          :label="item.name"
          :value="item.name"
        />
      </el-select>

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

      <!-- 类型 -->
      <el-select
        v-model="modelType"
        clearable
        placeholder="类型"
        style="width: 100px; margin: 0 20px;"
        size="large"
      >
        <el-option
          v-for="item in typeOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>

      <!-- 查询按钮 -->
      <el-button size="large" color="#000" type="info" round @click="paramsQuery">查询</el-button>

      <!-- 重置按钮 -->
      <el-button size="large" color="#000" plain type="info" round @click="resetQuery">重置</el-button>
    </section>

    <!-- 表格 -->
    <el-table :data="recordData" style="width: 100%">
      <el-table-column label="时间"></el-table-column>
      <el-table-column label="类型"></el-table-column>
      <el-table-column label="状态"></el-table-column>
      <el-table-column label="到"></el-table-column>
      <el-table-column label="币种"></el-table-column>
      <el-table-column label="手续费"></el-table-column>
      <el-table-column label="金额"></el-table-column>
      <el-table-column label="到账金额"></el-table-column>
      <el-table-column label="备注"></el-table-column>
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
import { ref, reactive, onMounted } from 'vue'
import { useCommonStore } from '@/pinia/modules/common'
import { getRejectList } from '@/api/tradeRecord'

const commonStore = useCommonStore()

const chain = ref('')
const currency = ref('')

const recordData = ref([])
const total = ref(0)
const currentPage = ref(1)

const modelType = ref('')
const typeOptions = reactive([
{ label: '支付', value: 'WITHDRAW_ORDER' },
{ label: '批量支付', value: 'BATCH_PAYOUT' },
])

const queryList = async (params) => {
  const { code, data = {} } = await getRejectList({ ...params, pageSize: 20 })
  if (code === 0) {
    recordData.value = data.content || []
    total.value = data.total_pages || 0
  }
}

// 翻页
const handleChangePage = (page) => {
  currentPage.value = page
  const params = { page }
  if (chain.value) params.chain = chain.value
  if (currency.value) params.currency = currency.value
  if (modelType.value) params.modelType = modelType.value
  queryList(params)
}

// 条件查询
const paramsQuery = () => {
  currentPage.value = 1
  const params = { page: 1 }
  if (chain.value) params.chain = chain.value
  if (currency.value) params.currency = currency.value
  if (modelType.value) params.modelType = modelType.value
  queryList(params)
}

// 重置条件查询
const resetQuery = () => {
  currentPage.value = 1
  chain.value = ''
  currency.value = ''
  modelType.value = ''
  queryList({ page: 1 })
}

onMounted(() => {
  if (commonStore.chainsList.length === 0) {
    commonStore.GetChainsList()
  }

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