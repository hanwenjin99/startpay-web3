<template>
  <main class="content_container">
    <!-- 资产 -->
    <h1 class="title">资产</h1>
    <section class="assets">
      <h2>资产估值</h2>
      <h1>≈ ${{ Number(accountInfo.assetValuationAmount ?? 0).toFixed(3) }}</h1>
    </section>

    <!-- 添加币种 -->
    <section class="add">
      <span class="title">币种</span>
      <el-button color="#000" icon="plus" type="info" round @click="router.push('add')">添加币种</el-button>
    </section>

    <!-- 列表 -->
    <el-table :data="accountInfo.accountInfo ?? []" style="width: 100%">
      <el-table-column>
        <template #default="scope">
          <ShowCurrency
            :main-icon="scope.row.currencyIcon"
            :secondary-icon="scope.row.chainIcon"
            :main-name="scope.row.currency"
            :secondary-name="scope.row.chain"
          />
        </template>
      </el-table-column>
      <el-table-column>
        <template #default="scope">
          <section class="number">
            {{ scope.row.balance }}
            <span class="smallNumber">${{ scope.row.amountUsd.toFixed(3) }}</span>
          </section>
        </template>
      </el-table-column>
      <el-table-column align="right">
        <template #default="scope">
          <el-button color="#000" plain type="info" icon="bottom" round>收款</el-button>
          <el-button color="#000" plain type="info" icon="top" round>转账</el-button>
          <el-button color="#000" plain type="info" icon="topRight" round>提现</el-button>
        </template>
      </el-table-column>
    </el-table>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getAccountInfo } from '@/api/account'

import ShowCurrency from '@/components/showCurrency/index.vue'

const router = useRouter()

const accountInfo = ref({})

const initAccountInfo = async () => {
  const { code, data } = await getAccountInfo()
  if (code === 0) {
    accountInfo.value = data
  }
}

onMounted(() => {
  // 初始化获取资产信息
  initAccountInfo()
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
  }

  .assets {
    display: flex;
    flex-direction: column;
    
    margin-top: 32px;
    padding-bottom: 40px;

    h1 {
      margin-bottom: 0;
      margin-top: 8px;

      font-size: 48px;
      cursor: pointer;
    }

    h2 {
      font-weight: 400;
      margin: 0;
      color: grey;
      font-size: 16px;
    }
  }

  .add {
    display: flex;
    align-items: center;
    justify-content: space-between;

    margin-bottom: 40px;

    .title {
      font-size: 24px;
      font-weight: 700;
    }
  }

  .number {
    display: flex;
    flex-direction: column;

    font-size: 16px;
    line-height: 1.25;

    .smallNumber {
      font-size: 14px;
      color: grey;
    }
  }
}

:deep(.el-table thead) {
  display: none;
}
</style>