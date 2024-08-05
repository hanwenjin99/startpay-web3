<template>
  <main class="content_container">
    <h1 class="title">批量转账记录</h1>
    
    <section class="search">
      <div class="left">
        <el-input
          v-model="search"
          placeholder="ID、任务"
          prefix-icon="search"
          size="large"
          style="width: 600px; margin-right: 16px;"
        />
        <!-- 状态 -->
        <el-select
          v-model="status"
          clearable
          placeholder="状态"
          style="width: 100px"
          size="large"
        >
          <el-option
            v-for="item in TRANSFER_STATUS_OPTIONS"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>

        <!-- 查询按钮 -->
        <el-button style="margin-left: 20px;" size="large" color="#000" round @click="paramsQuery">查询</el-button>
        <!-- 重置按钮 -->
        <el-button size="large" color="#000" round plain @click="resetQuery">重置</el-button>
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

import { TRANSFER_STATUS_OPTIONS } from '@/constants/constants'
import { getBatchPayoutList } from '@/api/account'

const search = ref('')
const status = ref('')

const recordData = ref([])
const currentPage = ref(1)
const total = ref(0)

// 获取批量转账记录
const queryBatchPayoutList = async (params) => {
  const { code, data = {} } = await getBatchPayoutList(params)
  if (code === 0) {
    recordData.value = data.content ?? []
    total.value = data.total_pages ?? 0
  }
}

// 条件查询列表
const paramsQuery = () => {
  currentPage.value = 1 // 第一页开始
  const params = { page: 1 }
  if (search.value) params.search = search.value
  if (status.value) params.status = status.value
  queryBatchPayoutList(params)
}

// 重置条件查询
const resetQuery = () => {
  currentPage.value = 1
  search.value = ''
  status.value = ''
  queryBatchPayoutList({ page: 1 })
}

// 翻页查询
const handleChangePage = (page) => {
  const params = { page }
  if (search.value) params.search = search.value
  if (status.value) params.status = status.value
  queryBatchPayoutList(params)
}

onMounted(() => {
  queryBatchPayoutList({ page: 1 })
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