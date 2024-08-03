<template>
  <main class="content_container">
    <h1 class="title">充值</h1>
    <!-- 选择币种组件 -->
    <SelectCurrency :init-data="selectOneCurrency" @handle-select-callback="handleSelect" />
    <div class="inputTitle">
      充值金额
    </div>
    <div class="sumAmount">
      <span>{{ rechargeAmount }}{{ selectOneCurrency.currency }}</span>
    </div>
    <img class="smallImg" :src="reduce" alt="">
    <div class="feeContainer">
      <h2>StartPay服务费+充值手续费</h2>
      <h2>{{ serviceAmount }}{{ selectOneCurrency.currency }}+{{ selectOneCurrency.chargeFeeamount ?? 0 }}{{ selectOneCurrency.currency }}={{ totalFee }}{{ selectOneCurrency.currency }}</h2>
    </div>
    <img class="smallImg" :src="equal" alt="">
    <span class="realAmountTitle">到账金额</span>
    <div class="amountOuter">
      <!-- <span class="amountBtn" @click="writeMaxAmount">最大</span> -->
      <input
        v-model="creatForm.amount"
        :placeholder="`0 ${selectOneCurrency.currency ?? ''}`"
        class="amountInput"
        @input="e => handleChange(e.target.value)"
      >
    </div>
    <span class="payee">汇款账户</span>
    <div class="bankSelector">
      <template v-if="creatForm.bankAccountId">
        <div class="selectBA">
          <span class="icon">{{ selectBAInfo.receiverName.slice(0, 1) }}</span>
          <span class="title">{{ selectBAInfo.receiverName }}</span>
          <span class="btn" @click.stop="resetSelectBank">
            重新选择
            <img :src="reset" alt="">
          </span>
        </div>
      </template>
      <template v-else>
        <div class="bankSelectorItem" @click="selectBankAccount">
          <el-icon><HomeFilled /></el-icon>
          <span>银行账户</span>
          <el-icon><ArrowRightBold /></el-icon>
        </div>
      </template>
      <!-- TODO 隐藏签约供应商选择入口 -->
      <!-- <span class="split">/</span>
      <div class="bankSelectorItem" @click="router.push('supplier')">
        <el-icon><Checked /></el-icon>
        <span>签约供应商</span>
        <el-icon><ArrowRightBold /></el-icon>
      </div> -->
    </div>
    <span class="payee">交易附言</span>
    <input v-model="creatForm.note" class="remakeInput">

    <!-- 提交按钮 -->
    <el-button
      :disabled="!canSubmit"
      color="#000"
      plain
      class="cotinueButton"
      round
      @click="submitCreate"
    >
      继续
    </el-button>

    <!-- 充值记录 -->
    <section class="record">
      <div class="title">
        <span>充值记录</span>
        <el-button link @click="router.push('/layout/tradeRecord/rechargeRecord')">查看全部</el-button>
      </div>

      <!-- 表格 -->
      <el-table :data="recordData" style="width: 100%">
        <el-table-column label="时间">
          <template #default="scope">
            <div class="tableItemColumn">
              {{ dayjs(scope.row.createTime).format('YYYY-MM-DD') }}
              <span>{{ dayjs(scope.row.createTime).format('HH:mm') }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="状态" prop="statusName" />
        <el-table-column label="汇款银行">
          <template #default="scope">{{ scope.row.bankAccount?.bankTitle }}</template>
        </el-table-column>
        <el-table-column label="钱包" prop="currency" />
        <!-- 平台银行 -->
        <el-table-column label="平台银行">
          <template #default="scope">
            <el-popover
              placement="top"
              width="300"
              title="详细信息"
            >
              <template #reference>
                {{ scope.row.platformBank?.bankTitle }}
              </template>
              <div class="platformBank">
                <span>企业名称：{{ scope.row.platformBank?.enterpriseTitle }}</span>
                <span>银行名称：{{ scope.row.platformBank?.bankTitle }}</span>
                <span>银行代码：{{ scope.row.platformBank?.bankCode }}</span>
                <span>收款人银行账号：{{ scope.row.platformBank?.receiverNumber }}</span>
                <span>银行账户名：{{ scope.row.platformBank?.receiverName }}</span>
                <el-button plain color="#000" @click="copyPlatformBankMsg(scope.row.platformBank)">复制</el-button>
              </div>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column label="服务费">
          <template #default="scope">
            {{ `${Number(scope.row.fee ?? 0).toFixed(2)}${scope.row.currency}` }}
          </template>
        </el-table-column>
        <el-table-column label="手续费">
          <template #default="scope">
            {{ `${scope.row.remittanceFee}${scope.row.currency}` }}
          </template>
        </el-table-column>
        <el-table-column label="扣款金额">
          <template #default="scope">
            <div class="tableItemColumn">
              {{ `-${scope.row.totalAmount}` }}
              <span>{{ `$${scope.row.totalAmount}` }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="到账金额">
          <template #default="scope">
            <div class="tableItemColumn">
              {{ scope.row.amount }}
              <span>{{ scope.row.amount }}</span>
            </div>
          </template>
        </el-table-column>
        <!-- 新增交易/审核附言 -->
        <el-table-column label="交易附言" prop="inputNote" />
        <el-table-column label="审核附言" prop="adminMemo" />
        <!-- 商户端 - 审核｜撤销 -->
        <el-table-column label="操作">
          <template #default="scope">
            <el-popconfirm
              confirm-button-text="确定"
              cancel-button-text="取消"
              icon="infoFilled"
              icon-color="#626AEF"
              title="确认审核此条记录？"
              @confirm="actionRecharge(scope.row, '审核')"
            >
              <template #reference>
                <el-button link>审核</el-button>
              </template>
            </el-popconfirm>
            <el-popconfirm
              confirm-button-text="确定"
              cancel-button-text="取消"
              icon="infoFilled"
              icon-color="#626AEF"
              title="确认撤销此条记录？"
              @confirm="actionRecharge(scope.row, '撤销')"
            >
              <template #reference>
                <el-button link>撤销</el-button>
              </template>
            </el-popconfirm>
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
import { ref, onMounted, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import dayjs from 'dayjs'
import { ElMessage, ElMessageBox } from 'element-plus'

import { copyMessage } from '@/utils/common.js'
import { useCommonStore } from '@/pinia/modules/common'
import { getRechargeList, createMerchantRecharge, actionMerchantRecharge } from '@/api/account'
import SelectCurrency from '@/components/selectCurrency/index.vue'
import reduce from '@/assets/icons/reduce.svg'
import equal from '@/assets/icons/equal.svg'
import reset from '@/assets/icons/reset.svg'

const commonStore = useCommonStore()
const router = useRouter()

const creatForm = ref({
  amount: '', // 到账金额
  bankAccountId: '', // 银行账户id
  chain: '', // 网络
  currency: '', // 币种
  note: '', // 交易附言
  projectId: '' // 选择的币种id
})
// 是否可以提交
const canSubmit = computed(() => {
  return creatForm.value.amount && creatForm.value.chain && creatForm.value.currency && creatForm.value.bankAccountId
})

// 文本框输入过滤
const handleChange = (val) => {
  // 过滤非数字或者小数的值
  creatForm.value.amount = val.replace(/[^\d.]/g, "")
}

const selectBAInfo = ref({}) // 选择的银行账户信息
const selectOneCurrency = ref({})
const recordData = ref([])
const total = ref(0)

// 服务费
const serviceAmount = computed(() => (Number(creatForm.value.amount ?? 0) * Number(selectOneCurrency.value.chargeFeerate1 ?? 0)).toFixed(2))

// 总手续费 = 服务费 + 手续费
const totalFee = computed(() => (Number(serviceAmount.value ?? 0) + Number(selectOneCurrency.value.chargeFeeamount ?? 0)).toFixed(2))

// 充值金额
const rechargeAmount = ref(0)

const setRechargeAmount = (newAmount) => {
  if (newAmount !== 0) {
    rechargeAmount.value = (Number(totalFee.value ?? 0) + Number(newAmount ?? 0)).toFixed(2)
  } else if (newAmount === 0) {
    rechargeAmount.value = 0
  }
}

// 重新选择银行
const resetSelectBank = () => {
  creatForm.value.bankAccountId = ''
  selectBAInfo.value = {}
}

// 选择币种回调
const handleSelect = (selectInfo) => {
  selectOneCurrency.value = selectInfo
  // 选择的币种信息 - 提交时入参
  creatForm.value.chain = selectInfo.chain
  creatForm.value.currency = selectInfo.currency
  creatForm.value.projectId = selectInfo.id

  // 充值页面的记录查询需要带上选择的币种信息
  queryList({
    page: 1,
    currency: selectInfo.currency,
    chain: selectInfo.chain,
    id: selectInfo.id
  })
}

const queryList = async (params) => {
  const { code, data = {} } = await getRechargeList({ ...params, pageSize: 10 })
  if (code === 0) {
    recordData.value = data.content ?? []
    total.value = data.total_pages ?? 0
  }
}

// 分页请求
const handleChangePage = (page) => {
  queryList({
    page,
    currency: selectOneCurrency.value.currency,
    chain: selectOneCurrency.value.chain,
    id: selectOneCurrency.value.id
  })
}

// 复制平台银行信息
const copyPlatformBankMsg = (item) => {
  copyMessage(
    `企业名称：${item.enterpriseTitle}\n银行名称：${item.bankTitle}\n银行代码：${item.bankCode}\n收款人银行账号：${item.receiverNumber}\n银行账户名：${item.receiverName}\n`
  )
}

// 创建充值提交
const submitCreate = () => {
  if (Number(creatForm.value.amount) <= 0) {
    ElMessage.warning('充值金额需大于0！')
    return
  }

  if (Number(creatForm.value.amount) > Number(selectOneCurrency.value.maxBound)) {
    ElMessage.warning(`充值金额最多不超过${selectOneCurrency.value.maxBound}！`)
    return
  }
  // 二次确认
  ElMessageBox.confirm(
    '确认创建充值记录吗？',
    '二次确认',
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(async () => {
    const { code } = await createMerchantRecharge(creatForm.value)
    if (code === 0) {
      ElMessage.success('创建充值成功')
      // 刷新列表数据
      queryList({
        page: 1,
        currency: selectOneCurrency.value.currency,
        chain: selectOneCurrency.value.chain,
        id: selectOneCurrency.value.id
      })
    }
  })
}

// 选择银行账户
const selectBankAccount = () => {
  commonStore.ChangePageInitRecharge({
    creatForm: creatForm.value,
    selectOneCurrency: selectOneCurrency.value
  })
  router.push({ name: 'bankAccount' })
}

// 商户端 - 充值列表操作(撤销/审核)
const actionRecharge = (item, type) => {
  // 填写备注信息
  ElMessageBox.prompt(`请填写${type}备注信息`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消'
  }).then(async ({ value }) => {
    const { code } = await actionMerchantRecharge({
      id: item.id,
      chain: item.chain,
      currency: item.currency,
      status: item.status,
      memo: value // 撤销|审核备注信息
    })
    if (code === 0) {
      ElMessage.success(`${type}充值成功！`)
      // 刷新列表
      queryList({
        page: 1,
        currency: selectOneCurrency.value.currency,
        chain: selectOneCurrency.value.chain,
        id: selectOneCurrency.value.id
      })
    }
  })
}

onMounted(() => {
  // 如果是选择银行账户返回 - 填充数据
  if (commonStore.pageInitRecharge?.isSelectedBankAccount) {
    // 选择的币种信息
    selectOneCurrency.value = commonStore.pageInitRecharge.selectOneCurrency
    // 选择的银行账户信息
    selectBAInfo.value = commonStore.pageInitRecharge.bankInfo
    // 其他信息
    creatForm.value = { ...commonStore.pageInitRecharge.creatForm, bankAccountId: commonStore.pageInitRecharge.bankInfo.id }
    commonStore.ChangePageInitRecharge({
      ...commonStore.pageInitRecharge,
      isSelectedBankAccount: false
    })
  }
})

watch(creatForm, (newVal) => {
  const newAmount = Number(newVal.amount)
  setRechargeAmount(newAmount)
}, {
  deep: true
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

    .selectBA {
      display: flex;
      align-items: center;
      flex: 1 1;

      .icon {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 24px;
        height: 24px;
        background-color: #ffefcb;
        border-radius: 4px;
        margin-right: 12px;
        font-size: 14px;
        color: #faaf00;
        font-weight: 600;
      }

      .title {
        flex: 1 1;
        font-size: 16px;
        color: #000;
      }

      .btn {
        display: flex;
        align-items: center;
        color: #bb8300;
        font-size: 16px;
        font-weight: 700;
        cursor: pointer;

        img {
          width: 24px;
          height: 24px;
          margin-left: 12px;
        }
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

  .tableItemColumn {
    display: flex;
    flex-direction: column;
    font-size: 16px;
    color: #000;
    line-height: 1.25;
    font-family: -inter;

    span {
      margin-top: 4px;
      color: grey;
      font-size: 14px;
    }
  }
}

.platformBank {
  display: flex;
  flex-direction: column;
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