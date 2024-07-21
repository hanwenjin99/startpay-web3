// 获取API管理数据
import service from '@/utils/request'

// 获取我的钱包列表
export const getBoundAddressList = (params) => {
  return service({
    url: '/web3/wallet_list',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}