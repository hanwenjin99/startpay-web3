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
          :key="`${item.currency}_${item.chain}`"
          :value="`${item.currency}_${item.chain}`"
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
import { ref, onMounted, watch } from 'vue'
import ShowCurrency from '@/components/showCurrency/index.vue'
import { getAccountInfo } from '@/api/account'

const props = defineProps({
  initData: {
    type: Object,
    default: () => {}
  }
})

const emit = defineEmits(['handleSelectCallback'])
const radioCurrency = ref('')
const selectInfo = ref({})

watch(props, (newVal) => {
  radioCurrency.value = `${newVal?.initData?.currency}_${newVal?.initData?.chain}`
  selectInfo.value = newVal?.initData
}, {
  immediate: true,
  deep: true
})

const dialogVisible = ref(false)

const currencyData = ref([])

const handleChange = (value) => {
  selectInfo.value = currencyData.value.filter(item => `${item.currency}_${item.chain}` === value)[0]
  dialogVisible.value = false
  // 父组件回调
  emit('handleSelectCallback', selectInfo.value)
}

const initCurrencyList = async () => {
  const { code, data } = await getAccountInfo()
  if (code === 0) {
    currencyData.value = data.accountInfo ?? []
    // 初始化设置第一个选择的币种
    selectInfo.value = data.accountInfo[0] ?? {}
    // 父组件回调
    emit('handleSelectCallback', selectInfo.value)
  }
}

onMounted(() => {
  // 初始化获取币种列表
  initCurrencyList()
})
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