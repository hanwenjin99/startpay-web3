<template>
  <main class="content_container">
    <h1 class="title">支付</h1>
    <!-- 选择币种组件 -->
    <SelectCurrency @handle-select-callback="handleSelect" />
    <div class="inputTitle">
      提现金额
      <span>0{{ selectOneCurrency.currency }} 可用</span>
    </div>
    <div class="sumAmount">
      <span>0{{ selectOneCurrency.currency }}</span>
    </div>
    <img class="smallImg" :src="reduce" alt="">
    <div class="feeContainer">
      <h2>Satogate服务费+汇款手续费</h2>
      <h2>0{{ selectOneCurrency.currency }}+{{ selectOneCurrency.remittanceFeeAmount ?? 0 }}{{ selectOneCurrency.currency }}=35{{ selectOneCurrency.currency }}</h2>
    </div>
    <img class="smallImg" :src="equal" alt="">
    <span class="realAmountTitle">到账金额</span>
    <div class="amountOuter">
      <span class="amountBtn">最大</span>
      <input :placeholder="`0 ${selectOneCurrency.currency ?? ''}`" class="amountInput">
    </div>
    <span class="payee">到</span>
    <div class="bankSelector">
      <div class="bankSelectorItem" @click="router.push('/layout/accountInfo/bankAccount')">
        <el-icon><HomeFilled /></el-icon>
        <span>银行账户</span>
        <el-icon><ArrowRightBold /></el-icon>
      </div>
      <span class="split">/</span>
      <div class="bankSelectorItem" @click="router.push('supplier')">
        <el-icon><Checked /></el-icon>
        <span>签约供应商</span>
        <el-icon><ArrowRightBold /></el-icon>
      </div>
    </div>
    <span class="payee">交易附言</span>
    <input class="remakeInput">

    <!-- 提交按钮 -->
    <el-button color="#000" plain class="cotinueButton" round>继续</el-button>

    <!-- 支付记录 -->
    <section class="record">
      <div class="title">
        <span>支付记录</span>
        <el-button link @click="router.push('/layout/tradeRecord/payRecord')">查看全部</el-button>
      </div>

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

      <!-- 分页 -->
      <div class="footerPage">
        <el-pagination background layout="prev, pager, next" :total="total" @current-change="handleChangePage" />
      </div>
    </section>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

import { getWithdrawOrderList } from '@/api/account'
import SelectCurrency from '@/components/selectCurrency/index.vue'
import reduce from '@/assets/icons/reduce.svg'
import equal from '@/assets/icons/equal.svg'

const router = useRouter()

const selectOneCurrency = ref({})
const recordData = ref([])
const total = ref(0)

// 选择币种回调
const handleSelect = (selectInfo) => {
  selectOneCurrency.value = selectInfo
}

// 分页请求
const handleChangePage = (page) => {
  queryList(page)
}

const queryList = async (page) => {
  const { code, data = {} } = await getWithdrawOrderList({ page, pageSize: 10 })
  if (code === 0) {
    recordData.value = data.content ?? []
    total.value = data.total_pages ?? 0
  }
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
  }

  .inputTitle {
    margin-top: 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 568px;

    span {
      color: grey;
    }
  }

  .sumAmount {
    box-sizing: border-box;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 568px;
    height: 102px;
    margin-top: 8px;
    padding: 0 16px;
    background-color: #fafafa;
    border: 1px solid #e6e6e6;
    border-radius: 8px;

    span {
      margin: 0 16px;
      font-size: 48px;
      line-height: 112%;
      letter-spacing: -1.92px;
      text-align: center;
      font-weight: 700;
      font-family: -inter;
    }
  }

  .smallImg {
    width: 16px;
    height: 16px;
    margin-top: 4px;
  }

  .feeContainer {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 568px;
    margin-top: 4px;

    h2 {
      font-weight: 400;
      margin: 0;
      font-size: 14px;
      &:first-child {
        color: grey;
      }

      &:last-child {
        color: #000;
      }
    }
  }

  .realAmountTitle {
    color: #000;
    font-size: 14px;
  }

  .amountOuter {
    box-sizing: border-box;
    display: flex;
    align-items: center;
    width: 568px;
    height: 102px;
    margin-top: 8px;
    padding: 24px 16px;
    border: 1px solid #b3b3b3;
    border-radius: 8px;

    .amountBtn {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 36px;
      height: 36px;
      border: 1px solid #e6e6e6;
      border-radius: 36px;
      cursor: pointer;
      font-size: 12px;
      text-align: center;
    }

    .amountInput {
      width: 100%;
      flex: 1 1;
      padding: 0 8px;
      font-weight: 700;
      font-size: 48px;
      line-height: 112%;
      letter-spacing: -1.92px;
      text-align: center;
      border: none;
      box-shadow: none;
      background-color: transparent;
    }
  }

  .payee {
    margin-top: 24px;
  }

  .bankSelector {
    box-sizing: border-box;
    display: flex;
    align-items: center;
    width: 568px;
    height: 56px;
    margin-top: 8px;
    padding: 16px;
    border: 1px solid #b3b3b3;
    border-radius: 8px;

    .split {
      margin: 0 20px;
    }

    .bankSelectorItem {
      display: flex;
      flex: 1 1;
      flex-direction: row;
      align-items: center;
      cursor: pointer;

      span {
        flex: 1 1;
        margin: 0 12px;
        font-size: 16px;
      }
    }
  }

  .remakeInput {
    box-sizing: border-box;
    margin-top: 8px;
    font-size: 16px;
    font-family: Inter;
    font-style: normal;
    line-height: 125%;
    width: 568px;
    height: 56px;
    border: 1px solid #b3b3b3;
    border-radius: 8px;
    padding: 4px 11px;
    &:focus {
      box-shadow: 0 0 0 2px rgba(0, 0, 0, 0.2);
    }
  }

  .cotinueButton {
    margin-top: 32px;
    width: 568px;
    height: 56px;
    border-radius: 48px;
  }

  .record {
    display: flex;
    flex-direction: column;

    .title {
      margin: 56px 0 32px 0;
      display: flex;
      align-items: center;
      justify-content: space-between;

      span {
        font-size: 24px;
        font-weight: 700;
      }
    }
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