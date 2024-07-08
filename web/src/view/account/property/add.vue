<!-- 添加币种 -->
<template>
  <main class="content_container">
    <header class="header">
      <el-button
        color="#000"
        plain
        icon="back"
        circle
        @click.stop="router.go(-1)"
      />
      <h1>添加币种</h1>
    </header>

    <!-- 搜索 -->
    <section class="search">
      <el-input
        v-model="searchString"
        placeholder="币种"
        prefix-icon="search"
        size="large"
        style="width: 600px; margin-right: 16px;"
      />

      <el-select
        v-model="selectVal"
        clearable
        placeholder="网络"
        style="width: 140px"
        size="large"
      >
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </section>
  </main>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const searchString = ref('')
const selectVal = ref('')

const options = reactive([
  { label: 'ETH', value: 'ETH' },
  { label: 'BSC', value: 'BSC' },
  { label: 'TRON', value: 'TRON' },
  { label: 'POLYGON', value: 'POLYGON' },
  { label: 'BTC', value: 'BTC' },
])
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

  .header {
    display: flex;
    align-items: center;

    h1 {
      font-size: 32px;
      font-weight: 700;
      margin-left: 16px;
    }
  }

  .search {
    display: flex;
    align-items: center;
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