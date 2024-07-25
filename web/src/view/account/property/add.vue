<!-- 添加币种 -->
<template>
  <main class="content_container">
    <header class="header">
      <el-button
        color="#000"
        plain
        icon="back"
        circle
        @click.stop="router.go(-1)"
      />
      <h1>添加币种</h1>
    </header>

    <!-- 搜索 -->
    <section class="search">
      <el-input
        v-model="currency"
        placeholder="币种"
        prefix-icon="search"
        size="large"
        style="width: 600px; margin-right: 16px;"
      />

      <el-select
        v-model="chain"
        clearable
        placeholder="网络"
        style="width: 140px"
        size="large"
      >
        <el-option
          v-for="item in commonStore.chainsList"
          :key="item.name"
          :label="item.name"
          :value="item.name"
        />
      </el-select>

      <el-button style="margin-left: 20px;" size="large" color="#000" round @click="oddsQuery">查询</el-button>

      <el-button size="large" color="#000" round plain @click="restOddsQuery">重置</el-button>
    </section>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCommonStore } from '@/pinia/modules/common'
import { getAccountCurrencyCreatableList } from '@/api/account'

const router = useRouter()
// 条件查询字段
const currency = ref('')
const chain = ref('')
const commonStore = useCommonStore()

const currencyList = ref([])

// 条件查询
const oddsQuery = () => {
  initCurrencyList({ currency: currency.value, chain: chain.value })
}

// 重置条件查询
const restOddsQuery = () => {
  currency.value = ''
  chain.value = ''
  initCurrencyList()
}

const initCurrencyList = async (params) => {
  const { code, data } = await getAccountCurrencyCreatableList(params)
  if (code === 0) {
    currencyList.value = data || []
  }
}

onMounted(() => {
  if (commonStore.chainsList.length === 0) {
    // 更新网络列表
    commonStore.GetChainsList()
  }
  // 初始化获取账户可创建币种列表
  initCurrencyList()
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

  .header {
    display: flex;
    align-items: center;

    h1 {
      font-size: 32px;
      font-weight: 700;
      margin-left: 16px;
    }
  }

  .search {
    display: flex;
    align-items: center;
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