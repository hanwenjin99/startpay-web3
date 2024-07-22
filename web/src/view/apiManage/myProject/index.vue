<template>
  <main class="content_container">
    <header class="header">
      <span class="title">我的项目</span>
      <div class="btns">
        <el-button round color="#000" @click="refresh">刷新</el-button>
        <el-button round color="#000" @click.stop="dialogVisible = true">新建</el-button>
      </div>
    </header>

    <!-- 表格 -->
    <el-table :data="projectList" style="width: 100%">
      <el-table-column label="项目名称" prop="name" />
      <el-table-column label="App Key">
        <template #default="scope">
          <div class="copy_item">
            <span>{{ scope.row.appKey }}</span>
            <el-icon @click="copyMessage(scope.row.appKey)"><CopyDocument /></el-icon>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="App Secret">
        <template #default="scope">
          <div class="copy_item">
            <span>{{ scope.row.appSecret }}</span>
            <el-icon @click="copyMessage(scope.row.appSecret)"><CopyDocument /></el-icon>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="创建时间">
        <template #default="scope">
          {{ dayjs(scope.row.createTime).format('YYYY-MM-DD HH:mm:ss') }}
        </template>
      </el-table-column>
      <el-table-column label="结算币种">
        <template #default="scope">
          {{ scope.row.assembleChain }} - {{ scope.row.settleCurrency }}
        </template>
      </el-table-column>
      <el-table-column label="归集地址">
        <template #default="scope">
          <div class="copy_item">
            <span>{{ scope.row.assembleAddress }}</span>
            <el-icon @click="copyMessage(scope.row.assembleAddress)"><CopyDocument /></el-icon>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Payment Page 返回URL" prop="paymentPageUrl" />
      <el-table-column label="回调地址" prop="callbackUrl" />
      <el-table-column label="操作">
        <el-button link>编辑</el-button>
        <el-button link>删除</el-button>
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

    <!-- 新建项目弹窗 -->
    <el-dialog
      v-model="dialogVisible"
      title="新建项目"
      width="800"
    >
      <el-form
        ref="ruleFormRef"
        label-position="top"
        :rules="formRules"
        :model="formLabelAlign"
      >
        <el-form-item label="项目名称" prop="name">
          <el-input v-model="formLabelAlign.name" />
        </el-form-item>
        <el-form-item label="结算币种" prop="settleCurrency">
          <el-select v-model="formLabelAlign.settleCurrency">
            <el-option value="USDT">USDT</el-option>
            <el-option value="USDC">USDC</el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="链类型" prop="assembleChain">
          <el-select
            v-model="formLabelAlign.assembleChain"
            clearable
            placeholder="请选择"
          >
            <el-option
              v-for="item in commonStore.chainsList"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="钱包地址" prop="assemble_address">
          <el-input v-model="formLabelAlign.assemble_address" />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button plain color="#010101" @click="resetForm(ruleFormRef)">取消</el-button>
        <el-button color="#000" @click="submitForm(ruleFormRef)">确定</el-button>
      </template>
    </el-dialog>
  </main>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import dayjs from 'dayjs'
import { ElMessage } from 'element-plus'

import { useCommonStore } from '@/pinia/modules/common'
import { copyMessage } from '@/utils/common.js'
import { getProjectList, addProject } from '@/api/apiManage'

const commonStore = useCommonStore()

const projectList = ref([])
const total = ref(0)
const currentPage = ref(1)

const dialogVisible = ref(false)
const ruleFormRef = ref('')

const formLabelAlign = reactive({
  name: '', // 项目名称
  settleCurrency: '', // 结算币种
  assembleChain: '', // 链类型
  assemble_address: '', // 钱包地址
})

const formRules = reactive({
  name: [{ required: true, message: '请输入项目名称', trigger: 'change', }],
  settleCurrency: [{ required: true, message: '请选择结算币种', trigger: 'change', }],
  assembleChain: [{ required: true, message: '请输入链类型', trigger: 'change', }]
})

const queryList = async (page) => {
  const { code, data = {} } = await getProjectList({ page, pageSize: 20 })
  if (code === 0) {
    projectList.value = data.content || []
    total.value = data.total_pages || 0
  }
}

// 提交创建项目
const submitForm = async (formEl) => {
  if (!formEl) return
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      const { code } = await addProject(formLabelAlign)
      if (code === 0) {
        ElMessage.success('创建项目成功！')
        dialogVisible.value = false // 关闭弹窗
        queryList(1) // 重新获取列表数据
      }
    } else {
      console.log('error submit!', fields)
    }
  })
}

// 取消创建提交
const resetForm = (formEl) => {
  if (!formEl) return
  formEl.resetFields()
  dialogVisible.value = false // 关闭弹窗
}

// 翻页
const handleChangePage = (page) => {
  currentPage.value = page
  queryList(page)
}

// 刷新
const refresh = () => {
  currentPage.value = 1
  queryList(1)
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
  
  .rowFlex {
    display: flex;
    align-items: center;
  }
}

.copy_item {
  display: flex;
  align-items: center;
  font-family: -initer;
}

:deep(.el-input__wrapper.is-focus) {
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