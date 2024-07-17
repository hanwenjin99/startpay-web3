<template>
  <main class="content_container">
    <h1 class="title">签约供应商</h1>

    <!-- 条件搜索 -->
    <section class="search">
      <el-input
        v-model="search"
        placeholder="企业名、服务"
        prefix-icon="search"
        size="large"
        style="width: 600px; margin-right: 16px;"
      />

      <!-- 分类 -->
      <el-select
        v-model="supplierCatId"
        clearable
        placeholder="分类"
        style="width: 100px; margin-right: 20px;"
        size="large"
      >
        <el-option
          v-for="item in classifyOptions"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <!-- 查询 -->
      <el-button size="large" color="#000" type="info" round @click="paramsQuery">查询</el-button>

      <!-- 重置按钮 -->
      <el-button size="large" color="#000" plain type="info" round @click="resetQuery">重置</el-button>
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
import { ref, onMounted } from 'vue'

import { getSupplierCatList, getSignedSupplierList } from '@/api/account'
import supplierHover from '@/assets/supplier_hover.png'

const search = ref('')
const supplierCatId = ref('')
const classifyOptions = ref([])

const list = ref([])

const initSupplierCatList = async () => {
  const { code, data } = await getSupplierCatList()
  if (code === 0) {
    classifyOptions.value = data || []
  }
}

const querySignedSupplierList = async (params) => {
  const { code, data = {} } = await getSignedSupplierList(params)
  if (code === 0) {
    list.value = data.content ?? []
  }
}

// 条件查询
const paramsQuery = () => {
  const params = {}
  if (search.value) params.search = search.value
  if (supplierCatId.value) params.supplierCatId = supplierCatId.value
  querySignedSupplierList(params)
}

// 重置条件查询
const resetQuery = () => {
  search.value = ''
  supplierCatId.value = ''
  querySignedSupplierList()
}

onMounted(() => {
  initSupplierCatList()
  querySignedSupplierList()
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