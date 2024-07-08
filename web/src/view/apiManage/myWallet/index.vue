<template>
  <main class="content_container">
    <header class="header">
      <span class="title">我的钱包</span>
      <div class="btns">
        <el-button round color="#000">刷新</el-button>
        <el-button round color="#000" @click.stop="dialogVisible = true">新建</el-button>
      </div>
    </header>

    <!-- 表格 -->
    <el-table :data="walletList" style="width: 100%">
      <el-table-column label="链类型" prop="chain" />
      <el-table-column label="商户关联ID" prop="merchantAddressId" />
      <el-table-column label="关联项目" prop="projectName" />
      <el-table-column label="地址">
        <template #default="scope">
          <div class="address">
            <span>{{ `${scope.row.address.slice(0, 25)}...` }}</span>
            <el-icon @click="copyMessage(scope.row.address)"><CopyDocument /></el-icon>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="创建时间">
        <template #default="scope">
          {{ dayjs(scope.row.createTime).format('YYYY-MM-DD HH:mm:ss') }}
        </template>
      </el-table-column>
    </el-table>

    <!-- 新建项目弹窗 -->
    <el-dialog
      v-model="dialogVisible"
      title="新建钱包"
      width="800"
    >
      <el-form
        label-position="top"
        :model="formLabelAlign"
      >
        <el-form-item label="链类型">
          <el-select />
        </el-form-item>
        <el-form-item label="商户关联ID">
          <el-input />
        </el-form-item>
        <el-form-item label="关联项目">
          <el-select />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button plain color="#010101" @click="dialogVisible = false">取消</el-button>
        <el-button color="#000" @click="dialogVisible = false">
          确定
        </el-button>
      </template>
    </el-dialog>
  </main>
</template>

<script setup>
import { reactive, ref } from 'vue'
import dayjs from 'dayjs'

import { copyMessage } from '@/utils/common.js'

const dialogVisible = ref(false)

const formLabelAlign = reactive({})

const walletList = ref([
  {
    "id": "66638b3850429b342b0f0139",
    "chain": "BTC",
    "address": "bc1qmf7w88zl8ctsryvammyjar97jjk0lg57u8qxhv",
    "merchantAddressId": null,
    "projectId": null,
    "projectName": null,
    "createTime": "2024-06-07T22:35:36.968+00:00"
  }
])
</script>

<style lang="scss" scoped>
.content_container {
  box-sizing: border-box;
  display: flex;
  flex-direction: column;

  padding: 10px;

  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
    margin-bottom: 20px;

    .title {
      color: rgba(0, 0, 0, .85);
      font-weight: 500;
      font-size: 16px;
    }

    .btns {
      display: flex;
      align-items: center;
    }
  }

  .address {
    display: flex;
    align-items: center;
    font-family: -initer;
  }
}

:deep(.el-input__wrapper.is-focus),
:deep(.el-select__wrapper.is-focused) {
  box-shadow: 0 0 0 2px rgba(0, 0, 0, 0.2);
}
</style>