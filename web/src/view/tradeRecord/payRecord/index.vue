<template>
  <main class="content_container">
    <h1 class="title">支付记录</h1>

    <!-- 搜索 -->
    <section class="search">
      <el-input
        v-model="searchString"
        placeholder="地址,Txid"
        prefix-icon="search"
        size="large"
        style="width: 600px; margin-right: 16px;"
      />

      <!-- 币种 -->
      <el-select
        v-model="selectCurrencyVal"
        clearable
        placeholder="币种"
        style="width: 100px; margin: 0 20px;"
        size="large"
      >
        <el-option
          v-for="item in currencyOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
      <!-- 重置按钮 -->
      <el-button size="large" color="#000" plain type="info" round>重置</el-button>
    </section>

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
    </el-table>
  </main>
</template>

<script setup>
import { ref, reactive } from 'vue'

const searchString = ref('')
const recordData = []

const selectCurrencyVal = ref('')
const currencyOptions = reactive([
  { label: 'ETH', value: 'ETH' },
  { label: 'BNB', value: 'BNB' },
  { label: 'USDT', value: 'USDT' },
  { label: 'USDC', value: 'USDC' },
  { label: 'MATIC', value: 'MATIC' },
  { label: 'TRX', value: 'TRX' },
  { label: 'BTC', value: 'BTC' },
  { label: 'USD', value: 'USD' },
])
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
</style>