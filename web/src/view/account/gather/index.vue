<template>
  <main class="content_container">
    <h1 class="title">收款</h1>

    <!-- 选择币种组件 -->
    <span class="selectTitle">币种</span>
    <SelectCurrency @handle-select-callback="queryAddress" />

    <!-- 地址 -->
    <span class="smallTitle">地址</span>
    <section class="address">
      <div class="qrCode">
        <canvas ref="qrCanvas" class="canvas" />
      </div>
      <span class="line" />
      <div class="addressBottom">
        <span>{{ addressValue }}</span>
        <el-icon @click.stop="copyMessage(addressValue)"><CopyDocument /></el-icon>
      </div>
    </section>

    <!-- 收款记录 -->
    <section class="record">
      <div class="title">
        <span>收款记录</span>
        <el-button link @click="router.push('/layout/tradeRecord/gatherRecord')">查看全部</el-button>
      </div>

      <!-- 表格 -->
      <el-table :data="recordData" style="width: 100%">
        <el-table-column label="时间">
          <template #default="scope">
            <div class="date">
              <span>{{ dayjs(scope.row.lastUpdate).format('YYYY-MM-DD') }}</span>
              <span class="time">{{ dayjs(scope.row.lastUpdate).format('HH:mm') }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="状态">
          <template #default="scope">
            <div class="status">
              <span :class="['icon', scope.row.status === 'FINISHED' ? 'finished' : '']" />
              {{ scope.row.status === 'FINISHED' ? '已完成' : '' }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="从">
          <template #default="scope">
            <div class="fromAddress">
              {{ `${scope.row.fromAddress.slice(0, 4)}...${scope.row.fromAddress.slice(-4)}` }}
              <span class="safe">Safe</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="到">
          <template #default="scope">
            <div class="toAddress">
              {{ `${scope.row.address.slice(0, 4)}...${scope.row.address.slice(-4)}` }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="Txid">
          <template #default="scope">
            <div class="txid">
              <span>{{ `${scope.row.txHash.slice(0, 4)}...${scope.row.txHash.slice(-4)}` }}</span>
              <el-icon @click.stop="copyMessage(scope.row.txHash)"><CopyDocument /></el-icon>
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
              <img class="icon" :src="scope.row.assetIcon" alt="">
              {{ scope.row.asset }}
            </div>
          </template>
        </el-table-column>
        <el-table-column label="金额">
          <template #default="scope">
            <div class="amount">
              {{ `+${scope.row.amountUsd}` }}
              <span class="smallAmount">{{ `$${scope.row.amountUsd}` }}</span>
            </div>
          </template>
        </el-table-column>
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
import QRCode from 'qrcode'
import dayjs from 'dayjs'
import { useRouter } from 'vue-router'

import { getDepositAddress, getDepositOrderList } from '@/api/account'
import SelectCurrency from '@/components/selectCurrency/index.vue'
import { copyMessage } from '@/utils/common.js'

const router = useRouter()
const qrCanvas = ref(null)
const addressValue = ref('')

const recordData = ref([])
const total = ref(0)

// 根据选择的币种查询地址
const queryAddress = async ({currency, chain}) => {
  const { code, data } = await getDepositAddress({ asset: currency, chain })
  if (code === 0) {
    addressValue.value = data
  }
}

const queryDepositList = async (page) => {
  const { code, data } = await getDepositOrderList({ depositOrderType: 'ACCOUNT', pageSize: 10, page })
  if (code === 0) {
    recordData.value = data.content ?? []
    total.value = data.total_pages ?? 0
  }
}

// 列表翻页
const handleChangePage = (page) => {
  queryDepositList(page)
}

onMounted(() => {
  QRCode.toCanvas(qrCanvas.value, addressValue.value, error => {
    if (error) console.error(error);
  });
  // 初始化查询收款记录
  queryDepositList(1)
});
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

  .selectTitle {
    font-size: 14px;
  }

  .smallTitle {
    font-size: 14px;
    margin-top: 24px;
    margin-bottom: 8px;
  }

  .address {
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    align-items: center;

    width: 568px;
    height: 385px;
    padding-top: 16px;
    border: 1px solid #e6e6e6;
    border-radius: 8px;

    .qrCode {
      box-sizing: border-box;
      display: flex;
      align-items: center;
      justify-content: center;

      width: 280px;
      height: 280px;
      padding: 16px;

      .canvas {
        width: 248px !important;
        height: 248px !important;
      }
    }

    .line {
      width: 568px;
      height: 1px;
      margin-top: 16px;
      background-color: #e6e6e6;
    }

    .addressBottom {
      width: 70%;
      display: flex;
      flex: 1 1;
      align-items: center;
      justify-content: space-between;
      font-size: 16px;
    }
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

      .finished {
        background-color: rgb(48, 190, 55);
      }
    }

    .fromAddress {
      display: flex;
      flex-direction: column;
      color: #997300;

      .safe {
        display: flex;
        align-items: center;

        width: fit-content;
        height: 18px;
        margin-top: 8px;
        padding: 1.5px 8px;

        font-size: 12px;
        border-radius: 36px;
        color: rgb(48, 190, 55);
        background-color: rgba(48, 190, 55, 0.1);
      }
    }

    .toAddress {
      color: #997300;
    }
    .txid {
      display: flex;
      align-items: center;
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

    .amount {
      display: flex;
      flex-direction: column;

      color: #000;
      line-height: 1.5;
      font-size: 16px;

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