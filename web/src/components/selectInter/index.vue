<template>
  <main class="selectInter">
    <section class="showSelect" @click.stop="dialogVisible = true">
      <div class="showInfo">
        <img :src="selectInfo.chainIcon" alt="">
        {{ selectInfo.chain }}
      </div>
      <el-icon><ArrowRightBold /></el-icon>
    </section>

    <!-- 选择币种弹窗 -->
    <el-dialog v-model="dialogVisible" title="选择网络" width="800">
      <el-radio-group v-model="radioInter" class="groupStyle" @change="handleChange">
        <el-radio
          v-for="item in interData"
          :key="item.chain"
          :value="item.chain"
          size="large"
        >
          <!-- 选项展示 -->
          <div class="showSelectItem">
            <img :src="item.chainIcon" alt="">
            {{ item.chain }}
          </div>
        </el-radio>
      </el-radio-group>
    </el-dialog>
  </main>
</template>

<script setup>
import { ref } from 'vue'

const interData = [{
  chain: "ETH",
  chainIcon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/eth.png"
}]

const radioInter = ref('')
const selectInfo = ref(interData[0])

const dialogVisible = ref(false)

const handleChange = (key) => {
  selectInfo.value = interData.filter(item => item.chain === key)[0]
  dialogVisible.value = false
}

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