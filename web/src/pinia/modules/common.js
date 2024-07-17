// 公共数据存储 - 网络列表/币种列表等
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getChainsInfo, getCurrencyOptions, getDepositOrderStatus } from '@/api/common'

export const useCommonStore = defineStore('common', () => {
  const chainsList = ref([])
  const currencyOptions = ref([])
  const depositOrderStatus = ref([])


  const setChainsList = (val) => {
    chainsList.value = val
  }

  const setCurrencyOptions = (val) => {
    currencyOptions.value = val
  }

  const setDepositOrderStatus = (val) => {
    depositOrderStatus.value = val
  }

  /* 获取网络列表信息*/
  const GetChainsInfo = async () => {
    const { code, data } = await getChainsInfo()
    if (code === 0) {
      setChainsList(data || [])
    }
  }
  
  // 获取币种列表
  const QueryCurrencyOptions = async () => {
    const { code, data } = await getCurrencyOptions()
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
    GetChainsInfo,
    QueryCurrencyOptions,
    QueryDepositOrderStatus,
    chainsList,
    currencyOptions,
    depositOrderStatus
  }
})
