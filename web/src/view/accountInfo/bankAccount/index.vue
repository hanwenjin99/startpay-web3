<template>
  <main class="content_container">
    <h1 class="title">银行账户</h1>

    <section class="top">
      <section class="search">
        <el-input
          v-model="searchString"
          placeholder="名称、地址"
          prefix-icon="search"
          size="large"
          style="width: 600px; margin-right: 16px;"
        />

        <!-- 重置按钮 -->
        <el-button size="large" color="#000" plain type="info" round>重置</el-button>
      </section>
      <el-button color="#000" icon="plus" type="info" round @click="router.push('addBankAccount')">添加银行账户</el-button>
    </section>

    <!-- 银行账户列表 -->
    <el-table :data="listData" style="width: 100%">
      <el-table-column>
        <template #default="scope">
          <div class="bankInfo">
            <span class="icon">{{ scope.row.bankTitle.slice(0, 1) }}</span>
            <div class="bankMsg">
              <span>{{ scope.row.receiverName }}</span>
              <span>{{ scope.row.receiverNumber }}</span>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column align="right">
        <template #default="scope">
          <el-button icon="delete" type="danger" circle></el-button>
          <el-button icon="edit" type="info" circle></el-button>
          <el-button color="#000" plain round @click="router.push('/layout/account/pay')">支付</el-button>
        </template>
      </el-table-column>
    </el-table>
  </main>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const searchString = ref('')

const listData = [{
  bankTitle: "BISON BANK S.A.",
  receiverName: "Start Pay Technology Limited",
  receiverNumber: "PT50006301087042000188015"
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

  .top {
    margin-bottom: 40px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .search {
      display: flex;
      align-items: center;
    }
  }

  .bankInfo {
    display: flex;
    align-content: center;

    .icon {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 36px;
      height: 36px;
      background-color: #ffefcb;
      border-radius: 36px;
      color: #faaf00;
      font-weight: 600;
      font-size: 20px;
      line-height: 125%;
      letter-spacing: -.6px;
    }

    .bankMsg {
      margin-left: 12px;
      display: flex;
      flex-direction: column;

      span {
        font-size: 16px;
        font-family: -inter;
        font-weight: 700;
        line-height: 1.25;
        &:last-child {
          margin-top: 4px;
          color: grey;
          font-size: 14px;
          font-weight: 400;
        }
      }
    }
  }
}

:deep(.el-input--large .el-input__wrapper) {
  border-radius: 20px;
}

:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 2px rgba(0,0,0,0.2);
}

:deep(.el-table thead) {
  display: none;
}
</style>