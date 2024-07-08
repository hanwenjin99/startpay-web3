<template>
  <main class="content_container">
    <h1 class="title">签约供应商</h1>

    <!-- 条件搜索 -->
    <section class="search">
      <el-input
        v-model="searchString"
        placeholder="企业名、服务"
        prefix-icon="search"
        size="large"
        style="width: 600px; margin-right: 16px;"
      />

      <!-- 分类 -->
      <el-select
        v-model="selectClassify"
        clearable
        placeholder="分类"
        style="width: 100px; margin-right: 20px;"
        size="large"
      >
        <el-option
          v-for="item in classifyOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>

      <!-- 重置按钮 -->
      <el-button size="large" color="#000" plain type="info" round>重置</el-button>
    </section>

    <!-- 供应商列表 -->
    <section class="list">
      <div v-for="item in list" :key="item.catId" class="item">
        <img class="icon" :src="item.icon" alt="">
        <span class="name">{{ item.catName }}</span>
        <h1>{{ `${item.companyName} · ${item.serviceName}` }}</h1>
        <h2>{{ item.description }}</h2>
        <img class="rightIcon" :src="supplierHover" alt="">
      </div>
    </section>
  </main>
</template>

<script setup>
import { ref } from 'vue'
import supplierHover from '@/assets/icons/supplier_hover.png'

const searchString = ref('')
const selectClassify = ref('')
const classifyOptions = [
  { value: '653e213d2b92e913265063b8', label: 'RTC技术' },
  { value: '6539111af1889a133e57dc78', label: '云服务' },
]

const list = [{
  catId: '653e213d2b92e913265063b8',
  catName: 'RTC技术',
  companyName: '声网',
  serviceName: '全球实时互动API平台 - 覆盖全球200+国家/地区',
  description: '在应用内构建语音通话、视频通话、互动直播、实时消息等多种实时互动场景。',
  icon: 'https://futurespay.s3.ap-southeast-1.amazonaws.com/token/token_icon_4934d629b3d1976c1b7e265f67362c2d.png'
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

  .search {
    display: flex;
    align-items: center;
    
    margin-bottom: 40px;
  }

  .list {
    width: 100%;
    display: flex;
    flex-wrap: wrap;

    .item {
      position: relative;
      display: flex;
      flex-direction: column;
      box-sizing: border-box;
      width: 250px;
      margin: 0 24px  24px 0;
      padding: 24px;
      background: #fff;
      border: 1px solid #e6e6e6;
      border-radius: 16px;
      cursor: pointer;

      &:hover {
        box-shadow: 0 0 12px rgba(0, 0, 0, 0.12);
        .rightIcon {
          opacity: 1;
        }
      }

      .rightIcon {
        position: absolute;
        top: 24px;
        right: 24px;
        width: 40px;
        height: 40px;
        opacity: 0;
      }

      .icon {
        width: 64px;
        height: 64px;
        border-radius: 64px;
      }

      .name {
        width: fit-content;
        margin-top: 24px;
        padding: 5px 12px;
        font-size: 14px;
        border: 1px solid #e6e6e6;
        border-radius: 36px;
      }

      h1 {
        margin: 16px 0 0;
        font-size: 16px;
      }

      h2 {
        margin: 8px 0 0;
        color: grey;
        font-size: 14px;
        line-height: 1.75;
        font-weight: 400;
      }
    }
  }
}

:deep(.el-input--large .el-input__wrapper),
:deep(.el-select--large .el-select__wrapper) {
  border-radius: 20px;
}

:deep(.el-input__wrapper.is-focus),
:deep(.el-select__wrapper.is-focused) {
  box-shadow: 0 0 0 2px rgba(0,0,0,0.2);
}
</style>