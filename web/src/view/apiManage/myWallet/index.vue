<template>
  <main class="content_container">
    <header class="header">
      <span class="title">我的钱包</span>
      <div class="btns">
        <el-button round color="#000" @click="refresh">刷新</el-button>
        <el-button round color="#000" @click.stop="dialogVisible = true">新建</el-button>
      </div>
    </header>

    <!-- 表格 -->
    <el-table :data="walletList" style="width: 100%">
      <el-table-column label="链类型" prop="chain" />
      <!-- TODO - 隐藏商户关联id -->
      <!-- <el-table-column label="商户关联ID" prop="merchantAddressId" /> -->
      <el-table-column label="钱包名称" prop="projectName" />
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

    <!-- 分页 -->
    <div class="footerPage">
      <el-pagination
        background
        layout="prev, pager, next"
        :current-page="currentPage"
        :total="total"
        @current-change="handleChangePage"
      />
    </div>

    <!-- 新建钱包弹窗 -->
    <el-dialog
      v-model="dialogVisible"
      title="新建钱包"
      width="800"
    >
      <el-form
        ref="ruleFormRef"
        label-position="top"
        :rules="formRules"
        :model="formLabelAlign"
      >
        <el-form-item label="钱包名称" prop="name">
          <el-input v-model="formLabelAlign.name" />
        </el-form-item>
        <el-form-item label="链类型" prop="assembleChain">
          <el-select
            v-model="formLabelAlign.assembleChain"
            clearable
            placeholder="请选择"
            @change="handleChangeChain"
          >
            <el-option
              v-for="item in commonStore.chainsList"
              :key="item.name"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="结算币种" prop="settleCurrency">
          <el-select v-model="formLabelAlign.settleCurrency">
            <el-option
              v-for="item in ownCurrencyList"
              :key="item.name"
              :label="item.name"
              :value="item.name"
            />
          </el-select>
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button plain color="#010101" @click="resetForm(ruleFormRef)">取消</el-button>
        <el-button color="#000" @click="submitForm(ruleFormRef)">
          确定
        </el-button>
      </template>
    </el-dialog>
  </main>
</template>

<script setup>
import { reactive, ref, onMounted } from 'vue'
import dayjs from 'dayjs'
import { ElMessage } from 'element-plus'

import { getCurrencyOptions } from '@/api/common'
import { useCommonStore } from '@/pinia/modules/common'
import { getBoundAddressList, addProject } from '@/api/apiManage'
import { copyMessage } from '@/utils/common.js'

const commonStore = useCommonStore()

const dialogVisible = ref(false)

const ruleFormRef = ref('')
const formLabelAlign = reactive({
  name: '', // 钱包名称
  assembleChain: '', // 链类型
  settleCurrency: '' // 结算币种
})
const formRules = reactive({
  name: [{ required: true, message: '请输入钱包名称', trigger: 'change', }],
  assembleChain: [{ required: true, message: '请输入链类型', trigger: 'change', }],
  settleCurrency: [{ required: true, message: '请选择结算币种', trigger: 'change', }]
})

const walletList = ref([])
const total = ref(0)
const currentPage = ref(1)

const queryList = async (page) => {
  const { code, data = {} } = await getBoundAddressList({ page, pageSize: 20 })
  if (code === 0) {
    walletList.value = data.content || []
    total.value = data.total_pages || 0
  }
}

// 翻页
const handleChangePage = (page) => {
  currentPage.value = page
  queryList(page)
}

// 取消创建提交
const resetForm = (formEl) => {
  if (!formEl) return
  formEl.resetFields()
  dialogVisible.value = false // 关闭弹窗
}

// 提交创建钱包
const submitForm = async (formEl) => {
  if (!formEl) return
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      const { code } = await addProject(formLabelAlign)
      if (code === 0) {
        ElMessage.success('创建钱包成功！')
        dialogVisible.value = false // 关闭弹窗
        queryList(1) // 重新获取列表数据
      }
    } else {
      console.log('error submit!', fields)
    }
  })
}

// 刷新
const refresh = () => {
  currentPage.value = 1
  queryList(1)
}

// 钱包的结算币种。未选择链时不展示币种列表
const ownCurrencyList = ref([])

// 选择链类型后，动态查询结算币种列表
const handleChangeChain = async (chain) => {
  // 清空结算币种的选择
  formLabelAlign.settleCurrency = ''
  const { code, data } = await getCurrencyOptions({ chain })
  if (code === 0) {
    ownCurrencyList.value = data || []
  }
}

onMounted(() => {
  if (commonStore.chainsList.length === 0) {
    // 更新链类型
    commonStore.GetChainsList()
  }
  queryList(1)
})
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