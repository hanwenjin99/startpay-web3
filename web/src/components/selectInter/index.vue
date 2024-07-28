<template>
  <main class="selectInter">
    <section class="showSelect" @click.stop="dialogVisible = true">
      <div class="showInfo">
        <img :src="selectInfo?.icon" alt="">
        {{ selectInfo?.name }}
      </div>
      <el-icon><ArrowRightBold /></el-icon>
    </section>

    <!-- 选择币种弹窗 -->
    <el-dialog v-model="dialogVisible" title="选择网络" width="800">
      <el-radio-group v-model="radioInter" class="groupStyle" @change="handleChange">
        <el-radio
          v-for="item in commonStore.chainsList"
          :key="item.name"
          :value="item.name"
          size="large"
        >
          <!-- 选项展示 -->
          <div class="showSelectItem">
            <img :src="item.icon" alt="">
            {{ item.name }}
          </div>
        </el-radio>
      </el-radio-group>
    </el-dialog>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue'

import { useCommonStore } from '@/pinia/modules/common'

const commonStore = useCommonStore()

const emits = defineEmits('handleSelectChain')

const radioInter = ref(commonStore.chainsList[0]?.name)
const selectInfo = ref(commonStore.chainsList[0])

const dialogVisible = ref(false)

const handleChange = (key) => {
  selectInfo.value = commonStore.chainsList.filter(item => item.chain === key)[0]
  emits('handleSelectChain', selectInfo.value)
  dialogVisible.value = false
}

onMounted(async () => {
  let initList = []
  if (commonStore.chainsList.length === 0) {
    // 更新网络列表
    const chainList = await commonStore.GetChainsList()
    if (Array.isArray(chainList) && chainList.length > 0) {
      initList = [...chainList]
    }
  } else {
    initList = [...commonStore.chainsList]
  }
  // 更新默认选择的网络信息
  radioInter.value = initList[0].name
  selectInfo.value = initList[0]
  // 返回给父组件
  emits('handleSelectChain', initList[0])
})
</script>

<style lang="scss" scoped>
.selectInter {
  display: flex;
  flex-direction: column;

  .showSelect {
    box-sizing: border-box;
    display: flex;
    flex-shrink: 0;
    align-items: center;
    justify-content: space-between;

    width: 700px;
    height: 56px;
    margin-top: 8px;
    padding: 0 16px;
    border: 1px solid #b3b3b3;
    border-radius: 8px;
    cursor: pointer;

    .showInfo {
      display: flex;
      align-items: center;
      img {
        width: 24px;
        height: 24px;
        margin-right: 12px;
        border-radius: 24px;
      }
    }
  }
}

.groupStyle {
  display: flex !important;
  flex-direction: column !important;
  align-items: flex-start;
}

.showSelectItem {
  display: flex;
  align-items: center;
  font-family: -inter;
  font-size: 16px;
  font-weight: 700;
  img {
    width: 36px;
    height: 36px;
    border-radius: 36px;
    margin-right: 12px;
  }
}
</style>