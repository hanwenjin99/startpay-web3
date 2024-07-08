<template>
  <main class="content_container">
    <h1 class="title">收款记录</h1>

    <!-- 搜索 -->
    <section class="search">
      <el-input
        v-model="searchString"
        placeholder="地址,Txid"
        prefix-icon="search"
        size="large"
        style="width: 600px; margin-right: 16px;"
      />
      <!-- 网络 -->
      <el-select
        v-model="selectInteVal"
        clearable
        placeholder="网络"
        style="width: 100px"
        size="large"
      >
        <el-option
          v-for="item in inteOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>

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
      `<el-table-column label="时间">
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
            <span :class="['icon', scope.row.status === 'FINISHED' ? 'finished' : '']"></span>
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
      <el-table-column label="链上txid">
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
      <el-table-column label="商户订单id"></el-table-column>
    </el-table>`
  </main>
</template>

<script setup>
import { ref, reactive } from 'vue'
import dayjs from 'dayjs'

const searchString = ref('')
const selectInteVal = ref('')
const inteOptions = reactive([
  { label: 'ETH', value: 'ETH' },
  { label: 'BSC', value: 'BSC' },
  { label: 'TRON', value: 'TRON' },
  { label: 'POLYGON', value: 'POLYGON' },
  { label: 'BTC', value: 'BTC' },
])

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

const recordData = ref([
  {
    lastUpdate: "2024-05-31T09:58:10.710+00:00",
    status: 'FINISHED',
    fromAddress: "0xa180fe01b906a1be37be6c534a3300785b20d947",
    address: "0x11a8ba50106b0fb8db914c507cdc799fc091f04c",
    txHash: "0xca5391f8e1b61d6f2b93823ce7d994f796e19d2843f6584140e7130583bfc767",
    chain: "BSC",
    chainIcon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/bnb.png",
    asset: "USDT",
    assetIcon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
    amountUsd: 2766.67
  }
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

:deep(.el-input--large .el-input__wrapper),
:deep(.el-select--large .el-select__wrapper) {
  border-radius: 20px;
}

:deep(.el-input__wrapper.is-focus),
:deep(.el-select__wrapper.is-focused) {
  box-shadow: 0 0 0 2px rgba(0,0,0,0.2);
}
</style>