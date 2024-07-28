<template>
  <main class="content_container">
    <h1 class="title">添加收款人</h1>

    <section class="addForm">
      <div class="formItem">
        <span class="formTitle">名称</span>
        <div class="inputDiv">
          <input v-model="addForm.name" class="input">
        </div>
      </div>
      <div class="formItem">
        <span class="formTitle">网络</span>
        <SelectInter @handle-select-chain="handleSelectChain" />
      </div>
      <div class="formItem">
        <span class="formTitle">地址</span>
        <div class="inputDiv">
          <input v-model="addForm.address" class="input">
        </div>
      </div>
      <div class="formItem">
        <span class="formTitle">邮箱验证码</span>
        <div class="inputDiv">
          <input v-model="addForm.emailCode" class="input">
          <span class="sendBtn">发送验证码</span>
        </div>
      </div>
      <div class="formItem">
        <span class="formTitle">谷歌验证码</span>
        <div class="inputDiv">
          <input v-model="addForm.googleCode" class="input">
        </div>
      </div>
    </section>

    <footer class="footer">
      <el-button text @click.stop="router.go(-1)">返回</el-button>
      <el-button
        :disabled="!isCanSubmit"
        size="large"
        color="#000"
        plain
        type="info"
        round
        @click="submitAddPayee"
      >
        添加
      </el-button>
    </footer>
  </main>
</template>

<script setup>
import { reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

import { addMerchantContact } from '@/api/accountInfo'
import SelectInter from '@/components/selectInter/index.vue'

const router = useRouter()

const addForm = reactive({
  name: '', // 名称
  chain: '', // 网络
  address: '', // 地址
  emailCode: '', // 邮箱验证码
  googleCode: '', // 谷歌验证码
})

const isCanSubmit = computed(() => addForm.name && addForm.chain && addForm.address && addForm.emailCode && addForm.googleCode)

const handleSelectChain = (selected) => {
  addForm.chain = selected.name ?? ''
}

// 新增收款人
const submitAddPayee = async () => {
  const { code } = await addMerchantContact(addForm)
  if (code === 0) {
    ElMessage.success('创建收款人成功')
    router.push('payee')
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

  .addForm {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 700px;
    margin-top: 8px;

    .formItem {
      display: flex;
      flex-direction: column;
      margin-top: 24px;
      width: 100%;

      .formTitle {
        font-size: 14px;
      }

      .inputDiv {
        display: flex;
        align-items: center;
        justify-content: flex-start;
        box-sizing: border-box;
        width: 100%;
        margin-top: 8px;
        padding: 16px;
        border: 1px solid #b3b3b3;
        border-radius: 8px;

        .input {
          flex: 1 1;
          height: 20px;
          padding: 0;
          font-size: 16px;
          border: none;
          box-shadow: none;
          background-color: transparent;
        }

        .sendBtn {
          cursor: pointer;
          margin-left: 12px;
          padding-right: 0;
          padding-left: 0;
          color: #bb8300;
          font-size: 16px;
        }
      }
    }
  }

  .footer {
    display: flex;
    align-items: center;
    justify-content: space-between;

    margin-top: 24px;
    width: 700px;
  }
}
</style>