<template>
  <main class="content_container">
    <h1 class="title">单笔转账记录</h1>

    <!-- 搜索 -->
    <section class="search">
      <el-input
        v-model="contactSearch"
        placeholder="地址,Txid"
        size="large"
        style="width: 568px;"
        prefix-icon="search"
      />

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

      <!-- 查询按钮 -->
      <el-button size="large" color="#000" type="info" round @click="paramsQuery">查询</el-button>

      <!-- 重置按钮 -->
      <el-button size="large" color="#000" plain type="info" round @click="resetQuery">重置</el-button>
    </section>

    <!-- 表格 -->
    <el-table :data="recordData" style="width: 100%">
      <el-table-column label="时间">
        <template #default="scope">
          <div class="date">
            <span>{{ dayjs(scope.row.createTime ? scope.row.createTime * 1000 : '').format('YYYY-MM-DD') }}</span>
            <span class="time">{{ dayjs(scope.row.createTime ? scope.row.createTime * 1000 : '').format('HH:mm') }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="状态">
        <template #default="scope">
          {{ TRANSFER_LIST_STATUS[scope.row.status] }}
        </template>
      </el-table-column>
      <el-table-column label="从">
        <template #default="scope">
          <div class="address">
            {{ `${scope.row.fromAddress.slice(0, 4)}...${scope.row.fromAddress.slice(-4)}` }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="到">
        <template #default="scope">
          <div class="address">
            {{ `${scope.row.toAddress.slice(0, 4)}...${scope.row.toAddress.slice(-4)}` }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="链上txid">
        <template #default="scope">
          <div class="txid">
            <span>{{ `${scope.row.txid.slice(0, 4)}...${scope.row.txid.slice(-4)}` }}</span>
            <el-icon @click.stop="copyMessage(scope.row.txid)"><CopyDocument /></el-icon>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="网络">
        <template #default="scope">
          <div class="info">
            <img class="icon" :src="scope.row.chainIcon" alt="">
            {{ scope.row.chain }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="币种">
        <template #default="scope">
          <div class="info">
            <img class="icon" :src="scope.row.currencyIcon" alt="">
            {{ scope.row.currency }}
          </div>
        </template>
      </el-table-column>
      <!-- TODO - 单笔转账记录不展示燃料费 -->
      <!-- <el-table-column label="燃料费">
        <template #default="scope">
          <div class="fuel">
            <span>{{ `${scope.row.gas} ${scope.row.gasToken}` }}</span>
            <span>{{ `$0.081` }}</span>
          </div>
        </template>
      </el-table-column> -->
      <el-table-column label="手续费">
        <template #default="scope">{{ `${scope.row.fee}${scope.row.currency}` }}</template>
      </el-table-column>
      <el-table-column label="金额">
        <template #default="scope">
          <div class="amount">
            {{ `-${scope.row.amount}` }}
            <span class="smallAmount">{{ `$${Math.round(scope.row.amountUsd)}` }}</span>
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

import { TRANSFER_LIST_STATUS } from '@/constants/constants'
import { getTransferList } from '@/api/account'
import { useCommonStore } from '@/pinia/modules/common'
import { copyMessage } from '@/utils/common.js'

const commonStore = useCommonStore()

const contactSearch = ref('')
const chain = ref('')
const currency = ref('')

const recordData = ref([])
const total = ref(0)
const currentPage = ref(1)

const queryList = async (params) => {
  const { code, data = {} } = await getTransferList({ ...params, pageSize: 20 })
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
  if (chain.value) params.chain = chain.value
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
  chain.value = ''
  currency.value = ''
  contactSearch.value = ''
  queryList({ page: 1 })
}

onMounted(() => {
  if (commonStore.chainsList.length === 0) {
    commonStore.GetChainsList()
  }

  if (commonStore.currencyOptions.length === 0) {
    commonStore.QueryCurrencyOptions()
  }

  queryList({
    page: 1
  })
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

  .date {
    display: flex;
    flex-direction: column;
    font-size: 16px;
    line-height: 1.25;
    color: #000;
    font-family: -inter;

    .time {
      margin-top: 4px;
      color: grey;
      font-size: 14px;
    }
  }

  .address {
    font-family: -inter;
    color: #997300;
  }

  .txid {
    display: flex;
    align-items: center;
    font-family: -inter;
    span {
      margin-right: 8px;
      color: #997300;
    }
    :deep(.el-icon) {
      width: 1.2em !important;
      height: 1.2em !important;

      svg {
        width: 100%;
        height: 100%;
      }
    }
  }

  .info {
    display: flex;
    align-items: center;
    color: #000;

    .icon {
      width: 24px;
      height: 24px;
      margin-right: 8px;
      border-radius: 24px;
    }
  }

  .fuel {
    display: flex;
    flex-direction: column;
    line-height: 1.25;
    font-family: -inter;

    span {
      &:first-child {
        margin-bottom: .5em;
        color: rgba(0, 0, 0, .85);
        font-weight: 500;
      }

      &:last-child {
        margin-top: 4px;
        color: grey;
      }
    }
  }

  .amount {
    display: flex;
    flex-direction: column;

    color: #000;
    line-height: 1.25;
    font-size: 16px;
    font-family: -inter;

    .smallAmount {
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