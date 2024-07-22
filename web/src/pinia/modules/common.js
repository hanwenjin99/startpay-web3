// 公共数据存储 - 网络列表/币种列表等
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getChainsInfo, getCurrencyOptions, getDepositOrderStatus, getChainsList } from '@/api/common'

export const useCommonStore = defineStore('common', () => {
  const chainsList = ref(["ETH", "BSC", "TRON", "POLYGON"]) // 链类型
  const chainsInfoList = ref([]) // 网络类型
  const currencyOptions = ref([]) // 币种类型
  const depositOrderStatus = ref([])


  const setChainsInfoList = (val) => {
    chainsInfoList.value = val
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

  /* 获取网络列表信息*/
  const GetChainsInfo = async () => {
    const { code, data } = await getChainsInfo()
    if (code === 0) {
      setChainsInfoList(data || [])
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
    GetChainsList,
    chainsInfoList,
    QueryCurrencyOptions,
    QueryDepositOrderStatus,
    chainsList,
    currencyOptions,
    depositOrderStatus
  }
})
