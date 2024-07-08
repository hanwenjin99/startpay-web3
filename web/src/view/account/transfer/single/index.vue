<template>
  <main class="content_container">
    <h1 class="title">
      单笔转账
    </h1>
    <!-- 选择币种组件 -->
    <SelectCurrency />

    <section class="amountOuter">
      <div class="formOuter">
        <span class="amountBtn">最大</span>
        <input v-model="amountVal" class="input" placeholder="0 USDT">
        <span class="btnDisabled" />
      </div>
      <span>≈${{ amountVal }}</span>
    </section>
    <span class="useful">0USDT 可用</span>

    <!-- 收款人 -->
    <span class="inputTitle">收款人</span>

    <div class="payeeSelector" @click="router.push('/layout/accountInfo/payee')">
      <span>ETH地址</span>
      <el-icon><Notebook /></el-icon>
    </div>

    <!-- 手续费 -->
    <div class="sumContainer">
      <span class="name">手续费</span>
      0
    </div>

    <div class="sumContainer">
      <span class="name">总金额</span>
      0
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
    </section>
  </main>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router'
import dayjs from 'dayjs'
import SelectCurrency from '@/components/selectCurrency/index.vue'
import { copyMessage } from '@/utils/common.js'

const amountVal = ref('')
const router = useRouter()

const recordData = [{
  updateTime: "2024-05-31T10:06:54.152+00:00",
  status: 'SUCCESS',
  fromAddress: "0x11a8ba50106b0fb8db914c507cdc799fc091f04c",
  toAddress: '0xe65baff9112775663954c8d0232ce9bd9eca4d87',
  txid: '0x7028d68f82dad1f889f840139df6022990903309fcced0b2f83e3064625d67a6',
  chain: "BSC",
  chainIcon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/bnb.png",
  currency: "USDT",
  currencyIcon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/usdt.png",
  gas: 0.000154809,
  gasToken: 'BNB',
  fee: 0,
  amount: 1,
  amountUsd: 0.999105
}]
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
</style>