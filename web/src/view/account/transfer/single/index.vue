<template>
  <main class="content_container">
    <h1 class="title">
      单笔转账
    </h1>
    <!-- 选择币种组件 -->
    <SelectCurrency @handle-select-callback="handleSelect" />

    <section class="amountOuter">
      <div class="formOuter">
        <span class="amountBtn" @click="handleMaxBtn">最大</span>
        <input
          v-model="amountInputVal"
          class="input"
          :placeholder="`0${selectOneCurrency.currency ?? ''}`"
          @input="e => handleChange(e.target.value)"
        >
        <span class="btnDisabled" />
      </div>
      <span>≈${{ amountEqualVal }}</span>
    </section>
    <span class="useful">{{ selectOneCurrency.balance ?? 0 }}{{ selectOneCurrency.currency ?? '' }} 可用</span>

    <!-- 收款人 -->
    <span class="inputTitle">收款人</span>

    <div class="payeeSelector" @click="router.push('/layout/accountInfo/payee')">
      <span>ETH地址</span>
      <el-icon><Notebook /></el-icon>
    </div>

    <!-- 手续费 -->
    <div class="sumContainer">
      <span class="name">手续费</span>
      {{ feeAmount }}
    </div>

    <div class="sumContainer">
      <span class="name">总金额</span>
      {{ sumAmount }}
    </div>

    <!-- 按钮 -->
    <el-button class="continueBtn" round color="#000" plain>继续</el-button>

    <!-- 记录 -->
    <section class="record">
      <div class="title">
        <span>转账记录</span>
        <el-button link @click="router.push('/layout/tradeRecord/gatherRecord')">查看全部</el-button>
      </div>

      <!-- 表格 -->
      <el-table :data="recordData" style="width: 100%">
        <el-table-column label="时间">
          <template #default="scope">
            <div class="date">
              <span>{{ dayjs(scope.row.updateTime).format('YYYY-MM-DD') }}</span>
              <span class="time">{{ dayjs(scope.row.updateTime).format('HH:mm') }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="状态">
          <template #default="scope">
            <div class="status">
              <span :class="['icon', scope.row.status === 'SUCCESS' ? 'success' : '']" />
              {{ scope.row.status === 'SUCCESS' ? '已完成' : '' }}
            </div>
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
        <el-table-column label="Txid">
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
        <el-table-column label="燃料费">
          <template #default="scope">
            <div class="fuel">
              <span>{{ `${scope.row.gas} ${scope.row.gasToken}` }}</span>
              <span>{{ `$0.081` }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="手续费">
          <template #default="scope">{{ `${scope.row.fee}${scope.row.currency}` }}</template>
        </el-table-column>
        <el-table-column label="金额">
          <template #default="scope">
            <div class="amount">
              {{ `+${scope.row.amount}` }}
              <span class="smallAmount">{{ `$${scope.row.amountUsd.toFixed(3)}` }}</span>
            </div>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="footerPage">
        <el-pagination background layout="prev, pager, next" :total="recordTotal" @current-change="handleChangePage" />
      </div>
    </section>
  </main>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router'
import dayjs from 'dayjs'

import { postBackendQuote, getTransferList } from '@/api/account'
import SelectCurrency from '@/components/selectCurrency/index.vue'
import { copyMessage } from '@/utils/common.js'

const router = useRouter()
const selectOneCurrency = ref({})
// 填入的值
const amountInputVal = ref(0)
// 手续费
const feeAmount = ref(0)

// 等价的值
const amountEqualVal = computed(() => (
  ((selectOneCurrency.value.usdPrice ?? 0) * (amountInputVal.value ?? 0)).toFixed(2)
))

// 总金额
const sumAmount = computed(() => (
  Number(amountInputVal.value ?? 0) + Number(feeAmount.value ?? 0)
))

// 文本框输入
const handleChange = (val) => {
  // 过滤非数字或者小数的值
  amountInputVal.value = val.replace(/[^\d.]/g, "")
  // 调接口
  postBackendQuote(selectOneCurrency.value.chain)
}

const recordData = []
const recordTotal = 0

// 点击最大按钮
const handleMaxBtn = () => {
  // 文本框填入最大值
  amountInputVal.value = selectOneCurrency.value.balance
  // 调接口
  postBackendQuote(selectOneCurrency.value.chain)
}

// 选择币种回调
const handleSelect = (selectInfo) => {
  selectOneCurrency.value = selectInfo
}

// 翻页
const handleChangePage = (page) => {
  queryTransferList(page)
}

const queryTransferList = async (page) => {
  const { code, data = {} } = await getTransferList({ page, pageSize: 10 })
  if (code === 0) {
    recordData.value = data.content ?? []
    recordTotal.value = data.total_pages ?? 0
  }
}

onMounted(() => {
  queryTransferList(1)
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

  .amountOuter {
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 568px;
    height: 134px;
    margin-top: 24px;
    padding: 32px 16px 24px;
    border: 1px solid #b3b3b3;
    border-radius: 8px;

    .formOuter {
      display: flex;
      align-items: center;

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
      }

      .input {
        flex: 1 1;
        padding: 0 8px;
        font-weight: 700;
        font-size: 48px;
        line-height: 112%;
        letter-spacing: -1.92px;
        text-align: center;
        width: 100%;
        border: none;
        box-shadow: none;
        background-color: transparent;
        font-family: Inter;
      }

      .btnDisabled {
        width: 36px;
        height: 36px;
      }
    }

    span {
      color: grey;
      font-size: 14px;
      text-align: center;
    }
  }

  .useful {
    margin-top: 8px;
    color: grey;
    font-size: 14px;
  }

  .inputTitle {
    margin-top: 24px;
    font-size: 14px;
  }

  .payeeSelector {
    box-sizing: border-box;

    display: flex;
    align-items: center;
    justify-content: space-between;

    width: 568px;
    height: 56px;
    margin-top: 8px;
    padding: 16px;
    border: 1px solid #b3b3b3;
    border-radius: 8px;

    cursor: pointer;

    span {
      color: grey;
      font-size: 16px;
    }
  }

  .sumContainer {
    margin-top: 28px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 568px;
    .name {
      color: grey;
      font-size: 14px;
    }
  }

  .continueBtn {
    margin-top: 32px;
    width: 568px;
    height: 56px;
    border-radius: 48px
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

    .status {
      display: flex;
      align-items: center;
      color: #000;

      .icon {
        width: 8px;
        height: 8px;
        margin-right: 8px;
        border-radius: 8px;
      }

      .success {
        background-color: rgb(48, 190, 55);
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