// 获取公共数据 - 网络列表/币种列表等
import qs from 'qs'
import service from '@/utils/request'

// 获取链类型
export const getChainsList = () => {
  return service({
    url: '/web3/chain_list',
    method: 'get'
  })
}

// 获取币种列表
export const getCurrencyOptions = (params) => {
  return service({
    url: `/web3/token_list?${qs.stringify(params)}`,
    method: 'get'
  })
}

// 获取 收款记录 状态列表
export const getDepositOrderStatus = () => {
  return service({
    url: '/backend/options/deposit_order_status',
    method: 'post'
  })
}