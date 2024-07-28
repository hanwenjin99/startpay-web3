<template>
  <main class="content_container">
    <h1 class="title">银行账户</h1>

    <section class="top">
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
          <!-- 删除 -->
          <el-popconfirm
            confirm-button-text="确定"
            cancel-button-text="取消"
            icon="infoFilled"
            icon-color="#626AEF"
            title="确认删除此条数据？"
            @confirm="deleteBank(scope.row)"
          >
            <template #reference>
              <el-button icon="delete" type="danger" circle />
            </template>
          </el-popconfirm>
          <el-button icon="edit" type="info" circle />
          <el-button color="#000" plain round @click="comfirmSelect(scope.row)">提现</el-button>
        </template>
      </el-table-column>
    </el-table>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

import { useCommonStore } from '@/pinia/modules/common'
import { getBankAccountList, deleteBankAccount } from '@/api/accountInfo'

const commonStore = useCommonStore()
const router = useRouter()
const route = useRoute()
console.log(route)

const listData = ref([])

// 选择银行账户
const comfirmSelect = (item) => {
  commonStore.ChangePageInitPay({
    ...commonStore.pageInitPay,
    bankInfo: {
      receiverName: item.receiverName,
      id: item.id
    },
    isSelectedBankAccount: true,
  })
  router.push('/layout/account/pay')
}

const queryList = async () => {
  const { code, data } = await getBankAccountList()
  if (code === 0) {
    listData.value = data || []
  }
}

// 删除银行账户
const deleteBank = async (item) => {
  const { id } = item
  const { code } = await deleteBankAccount({ id })
  if (code === 0) {
    ElMessage.success('删除银行账户成功！')
    queryList(1) // 重新查询列表数据
  }
}

onMounted(() => {
  queryList()
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

  .top {
    margin-bottom: 40px;
    display: flex;
    align-items: center;
    justify-content: flex-end;

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