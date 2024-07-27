// 公共数据存储 - 网络列表/币种列表等
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getCurrencyOptions, getDepositOrderStatus, getChainsList } from '@/api/common'

export const useCommonStore = defineStore('common', () => {
  const chainsList = ref([]) // 链类型
  
  const currencyOptions = ref([]) // 币种类型
  const depositOrderStatus = ref([])

  // 页面初始化数据
  const pageInitPay = ref({}) //提现页面
  const singleTransfer = ref({}) // 单笔转账页面

  const ChangePageInitPay = (newVal) => {
    pageInitPay.value = { ...newVal }
  }

  const ChangeSingleTransfer = (newVal) => {
    singleTransfer.value = { ...newVal }
  }

  const setChainsList = (val) => {
    chainsList.value = val
  }

  const setCurrencyOptions = (val) => {
    currencyOptions.value = val
  }

  const setDepositOrderStatus = (val) => {
    depositOrderStatus.value = val
  }

  // 获取链类型
  const GetChainsList = async () => {
    const { code, data } = await getChainsList()
    if (code === 0) {
      setChainsList(data || [])
    }
  }
  
  // 获取币种列表
  const QueryCurrencyOptions = async (params) => {
    const { code, data } = await getCurrencyOptions(params)
    if (code === 0) {
      setCurrencyOptions(data || [])
    }
  }

  // 获取记录状态列表
  const QueryDepositOrderStatus = async () => {
    const { code, data } = await getDepositOrderStatus()
    if (code === 0) {
      setDepositOrderStatus(data || [])
    }
  }

  return {
    GetChainsList,
    QueryCurrencyOptions,
    QueryDepositOrderStatus,
    chainsList,
    currencyOptions,
    depositOrderStatus,
    ChangePageInitPay,
    pageInitPay,
    singleTransfer,
    ChangeSingleTransfer
  }
})
