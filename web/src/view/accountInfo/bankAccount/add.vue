<template>
  <main class="content_container">
    <h1 class="title">添加银行账户</h1>
    <!-- 银行/地区 -->
    银行国家/地区
    <el-select
      v-model="addForm.region"
      class="selectStyle"
      style="width: 568px;"
      size="large"
    >
      <el-option label="US 美国" value="US" />
      <el-option label="HK 香港" value="HK" />
      <el-option label="SG 新加坡" value="SG" />
    </el-select>
    <!-- 收款方式 -->
    <span class="smallTitle">收款方式</span>
    <div class="payContainer">
      <img :src="swiftBank" alt="">
      <div class="desc">
        <span>Swift 国际银行 货币为USD</span>
        <span>每笔收取35.00USD费用 · 0-3个工作日</span>
      </div>
      <span class="point" />
    </div>

    <!-- 企业名称 -->
    <span class="smallTitle">企业名称</span>
    <el-input v-model="addForm.enterpriseTitle" style="width: 568px;" />

    <!-- 银行信息 -->
    <div class="rowFlex">
      <div class="columnFlex">
        <span class="smallTitle">银行名称</span>
        <el-input v-model="addForm.bankTitle" style="width: 272px;" />
      </div>

      <div class="columnFlex">
        <span class="smallTitle">银行代码(SWIFT / BIC)</span>
        <el-input v-model="addForm.bankCode" style="width: 272px;" />
      </div>
    </div>

    <!-- 路由号 -->
    <span class="smallTitle">Fedwire汇款路由号</span>
    <el-input v-model="addForm.fedWire" style="width: 568px;" />

    <!-- 账号信息 -->
    <div class="rowFlex">
      <div class="columnFlex">
        <span class="smallTitle">收款人银行账号</span>
        <el-input v-model="addForm.receiverNumber" style="width: 272px;" />
      </div>

      <div class="columnFlex">
        <span class="smallTitle">银行账户名</span>
        <el-input v-model="addForm.receiverName" style="width: 272px;" />
      </div>
    </div>

    <!-- 地址 -->
    <span class="smallTitle">收款人地址</span>
    <el-input v-model="addForm.receiverAddress" style="width: 568px;" />

    <footer class="footer">
      <el-button text @click.stop="router.go(-1)">返回</el-button>
      <el-button
        size="large"
        color="#000"
        plain
        type="info"
        round
        @click="submitAddForm"
      >
        确认
      </el-button>
    </footer>
  </main>
</template>

<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

import { createBankAccount } from '@/api/accountInfo'
import swiftBank from '@/assets/swift-bank.png'

const router = useRouter()
const addForm = reactive({
  region: 'US', // 银行国家/地区
  enterpriseTitle: '', // 企业名称
  bankTitle: '', // 银行名称
  bankCode: '', // 银行代码
  receiverNumber: '', // 收款人银行账户
  receiverName: '', // 银行账户名
  receiverAddress: '', // 收款人地址
  fedWire: '', // Fedwire汇款路由号
  remittanceType: "SWIFT" // 汇款方式
})

// 提交创建银行账户
const submitAddForm = async () => {
  if (!/^[A-Za-z0-9]{3,33}$/.test(addForm.receiverNumber)) {
    ElMessage.warning('收款人银行账号必须为3-33 位英文，数字')
    return
  }
  const { code } = await createBankAccount(addForm)
  if (code === 0) {
    ElMessage.success('创建银行账户成功！')
    router.push('bankAccount')
  }
}

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

  .smallTitle {
    margin-top: 24px;
  }

  .payContainer {
    box-sizing: border-box;
    display: flex;
    align-items: center;
    width: 568px;
    height: 56px;
    margin-top: 8px;
    padding-left: 16px;
    border: 2px solid #000;
    border-radius: 8px;

    img {
      width: 36px;
      height: 36px;
      margin-right: 12px;
    }

    .desc {
      flex: 1 1;
      display: flex;
      flex-direction: column;

      span {
        font-family: -inter;
        line-height: 1.25;
        &:first-child {
          font-size: 16px;
          font-weight: 700;
        }
        &:last-child {
          color: grey;
          font-size: 12px;
        }
      }
    }

    .point {
      box-sizing: border-box;
      width: 22px;
      height: 22px;
      margin-right: 24px;
      border: 6px solid #000;
      border-radius: 22px;
    }
  }

  .columnFlex {
    display: flex;
    flex-direction: column;
  }

  .rowFlex {
    width: 568px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .footer {
    display: flex;
    align-items: center;
    justify-content: space-between;

    margin-top: 24px;
    width: 568px;
  }

  :deep(.el-select--large .el-select__wrapper) {
    height: 56px;
    padding-top: 0 !important;
    padding-bottom: 0 !important;
    border-radius: 8px;
    margin-top: 8px;
  }

  :deep(.el-select__wrapper.is-focused) {
    box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.2);
  }

  :deep(.el-input__wrapper) {
    margin-top: 8px;
    height: 56px;
    border-radius: 8px;
  }

  :deep(.el-input__wrapper.is-focus) {
    box-shadow: 0 0 0 2px rgba(0, 0, 0, 0.2);
  }
}
</style>