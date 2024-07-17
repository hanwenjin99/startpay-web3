<template>
  <main class="content_container">
    <h1 class="title">新增批量支付任务</h1>

    <section class="add_container">
      <div class="add_head">
        <div class="desc">
          <span>{{ `让您的用户加入FuturesCash\n安全快捷，更省燃料费` }}</span>
          <img :src="fcLogo" alt="">
        </div>
        <span class="register_link">注册链接</span>
        <div class="link_div">
          <span>{{ link }}</span>
          <span @click.stop="copyLink">拷贝</span>
        </div>
      </div>
      <div class="add_body">
        <div class="down_load">
          任务文件
          <el-button color="#000" plain round icon="download" @click="downloadTemplate">下载模版</el-button>
        </div>
        <!-- 拖拽上传 -->
        <el-upload
          class="upload"
          drag
          :action="UPLOAD_FILE_URL"
        >
          <div class="upload_div">
            <el-icon size="32"><Document /></el-icon>
            <span class="title">拖放或点击上传</span>
            <span class="desc">csv文件，格式如下：platform · userId · chain · address · asset · amount(platform + userId 不为空 或 address 不为空)</span>
          </div>
        </el-upload>
      </div>
      <div class="add_footer">
        <el-button link @click.stop="router.go(-1)">返回</el-button>
        <el-button size="large" round plain color="#000">继续</el-button>
      </div>
    </section>
  </main>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { UPLOAD_FILE_URL } from '@/constants/constants'
import { getBatchPayoutTemplate } from '@/api/account'
import { copyMessage } from '@/utils/common.js'
import fcLogo from '@/assets/icons/fc-logo.svg'

const router = useRouter()

const link = ref('https://connect.futures.cash/?user=hanwenjin99@gmail.com')

const copyLink = () => {
  copyMessage(link.value)
}

// 下载模版
const downloadTemplate = async () => {
  const { code, data } = await getBatchPayoutTemplate()
  if (code === 0) {
    window.open(data, '_blank')
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
    margin: 0;
  }

  .add_container {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 600px;
    margin-top: 32px;

    .add_head {
      box-sizing: border-box;
      width: 100%;
      padding: 24px;
      background-color: #f70;
      border-radius: 16px;
      box-shadow: 0 4px 16px 0 rgba(0, 0, 0, 0.2);

      .desc {
        display: flex;
        justify-content: space-between;
        width: 100%;

        span {
          color: #fff;
          font-size: 20px;
          white-space: pre-wrap;
          font-weight: 700;
          font-family: -inter;
        }

        img {
          width: 90px;
          height: 20px;
        }
      }
      
      .register_link {
        display: flex;
        margin-top: 16px;
        color: #fff;
        font-size: 12px;
        font-family: -inter;
      }

      .link_div {
        display: flex;
        align-items: center;
        width: 100%;
        margin-top: 8px;

        span {
          &:first-child {
            flex: 1 1;
            margin-right: 12px;
            padding: 9px 16px;
            color: #fff;
            font-size: 14px;
            background-color: hsla(0, 0%, 100%, .2);
            border-radius: 56px;
            font-family: -inter;
          }

          &:last-child {
            padding: 9px 24px;
            color: #f70;
            font-size: 14px;
            background-color: #fff;
            border-radius: 56px;
            cursor: pointer;
            font-family: -inter;
            font-weight: 700;
          }
        }
      }
    }

    .add_body {
      box-sizing: border-box;
      display: flex;
      flex-direction: column;
      width: 100%;
      margin-top: 32px;

      .down_load {
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
      }

      .upload {
        margin-top: 10px;

        .upload_div {
          display: flex;
          flex-direction: column;
          align-items: center;

          .title {
            margin-top: 8px;
            font-size: 16px;
            line-height: 1.5;
            font-weight: 700;
          }

          .desc {
            font-family: -inter;
            margin-top: 4px;
            color: grey;
            font-size: 12px;
            line-height: 1.5;
            text-align: center;
          }
        }
      }
    }

    .add_footer {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-top: 32px;
      width: 100%;
    }
  }
}
</style>