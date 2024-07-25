// 获取公共数据 - 网络列表/币种列表等
import service from '@/utils/request'

// 获取链类型
export const getChainsList = () => {
  return service({
    url: '/web3/chain_list',
    method: 'get'
  })
}

// 获取币种列表
export const getCurrencyOptions = () => {
  return service({
    url: '/backend/options/currency',
    method: 'get'
  })
}

// 获取记录状态列表
export const getDepositOrderStatus = () => {
  return service({
    url: '/backend/options/deposit_order_status',
    method: 'post'
  })
}