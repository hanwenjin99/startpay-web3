<template>
  <main class="content_container">
    <header class="header">
      <span class="title">我的项目</span>
      <div class="btns">
        <el-button round color="#000">刷新</el-button>
        <el-button round color="#000" @click.stop="dialogVisible = true">新建</el-button>
      </div>
    </header>

    <!-- 表格 -->
    <el-table :data="projectList" style="width: 100%">
      <el-table-column label="项目名称"></el-table-column>
      <el-table-column label="App Key"></el-table-column>
      <el-table-column label="App Secret"></el-table-column>
      <el-table-column label="创建时间"></el-table-column>
      <el-table-column label="结算币种"></el-table-column>
      <el-table-column label="归集地址"></el-table-column>
      <el-table-column label="Payment Page 返回URL"></el-table-column>
      <el-table-column label="回调地址"></el-table-column>
      <el-table-column label="操作"></el-table-column>
    </el-table>

    <!-- 新建项目弹窗 -->
    <el-dialog
      v-model="dialogVisible"
      title="新建项目"
      width="800"
    >
      <el-form
        label-position="top"
        :model="formLabelAlign"
      >
        <el-form-item label="项目名称">
          <el-input />
        </el-form-item>
        <el-form-item label="结算币种">
          <el-input />
        </el-form-item>
        <el-form-item label="Payment Page 返回URL">
          <el-input />
        </el-form-item>
        <el-form-item label="回调地址URL">
          <el-input>
            <template #prepend>https://</template>
          </el-input>
        </el-form-item>
        <el-form-item label="邮箱验证码">
          <div class="rowFlex">
            <el-input style="width: 660px;" />
            <el-button color="#000" plain>获取验证码</el-button>
          </div>
        </el-form-item>
        <el-form-item label="谷歌验证码">
          <el-input />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button plain color="#010101" @click="dialogVisible = false">取消</el-button>
        <el-button color="#000" @click="dialogVisible = false">确定</el-button>
      </template>
    </el-dialog>
  </main>
</template>

<script setup>
import { ref, reactive } from 'vue'

const projectList = ref([])
const dialogVisible = ref(false)

const formLabelAlign = reactive({
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

:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 2px rgba(0, 0, 0, 0.2);
}
</style>