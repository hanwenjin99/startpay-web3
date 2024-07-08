<template>
  <main class="selectCurrency">
    <section class="showSelect" @click.stop="dialogVisible = true">
      <ShowCurrency
        :main-icon="selectInfo?.currencyIcon"
        :secondary-icon="selectInfo?.chainIcon"
        :main-name="selectInfo?.currency"
        :secondary-name="selectInfo?.chain"
      />
      <el-icon><ArrowRightBold /></el-icon>
    </section>

    <!-- 选择币种弹窗 -->
    <el-dialog v-model="dialogVisible" title="选择币种" width="800">
      <el-radio-group v-model="radioCurrency" class="groupStyle" @change="handleChange">
        <el-radio
          v-for="item in currencyData"
          :key="item.id"
          :value="item.id"
          size="large"
        >
          <!-- 选项展示 -->
          <ShowCurrency
            :main-icon="item.currencyIcon"
            :secondary-icon="item.chainIcon"
            :main-name="item.currency"
            :secondary-name="item.chain"
          />
        </el-radio>
      </el-radio-group>
    </el-dialog>
  </main>
</template>

<script setup>
import { ref } from 'vue'
import ShowCurrency from '@/components/showCurrency/index.vue'

const radioCurrency = ref('')
const selectInfo = ref(null)

const dialogVisible = ref(false)

const currencyData = [{
  id: '1',
  amountUsd: 0.08240702647012606,
  balance: 0.000145191,
  chain: "BSC",
  chainIcon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/bnb.png",
  currency: "BNB",
  currencyIcon: "https://pifutures.oss-cn-shanghai.aliyuncs.com/cash/bnb.png"
}]

const handleChange = (id) => {
  selectInfo.value = currencyData.filter(item => item.id === id)[0]
  dialogVisible.value = false
}

</script>

<style lang="scss" scoped>
.selectCurrency {
  display: flex;
  flex-direction: column;

  .showSelect {
    box-sizing: border-box;
    display: flex;
    flex-shrink: 0;
    align-items: center;
    justify-content: space-between;

    width: 568px;
    height: 56px;
    margin-top: 8px;
    padding: 0 16px;
    border: 1px solid #b3b3b3;
    border-radius: 8px;
    cursor: pointer;
  }
}

.groupStyle {
  display: flex !important;
  flex-direction: column !important;
  align-items: flex-start;
}
</style>