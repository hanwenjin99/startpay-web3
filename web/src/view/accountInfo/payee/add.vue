<template>
  <main class="content_container">
    <h1 class="title">添加收款人</h1>

    <section class="addForm">
      <div class="formItem">
        <span class="formTitle">名称</span>
        <div class="inputDiv">
          <input class="input">
        </div>
      </div>
      <div class="formItem">
        <span class="formTitle">网络</span>
        <SelectInter />
      </div>
      <div class="formItem">
        <span class="formTitle">地址</span>
        <div class="inputDiv">
          <input class="input">
        </div>
      </div>
      <div class="formItem">
        <span class="formTitle">邮箱验证码</span>
        <div class="inputDiv">
          <input class="input">
          <span class="sendBtn">发送验证码</span>
        </div>
      </div>
      <div class="formItem">
        <span class="formTitle">谷歌验证码</span>
        <div class="inputDiv">
          <input class="input">
        </div>
      </div>
    </section>

    <footer class="footer">
      <el-button text @click.stop="router.go(-1)">返回</el-button>
      <el-button size="large" color="#000" plain type="info" round>添加</el-button>
    </footer>
  </main>
</template>

<script setup>
import { useRouter } from 'vue-router'

import SelectInter from '@/components/selectInter/index.vue'

const router = useRouter()
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