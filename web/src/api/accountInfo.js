// 获取账户信息
import service from '@/utils/request'

// 获取收款人列表
export const getMerchantContactList = (params) => {
  return service({
    url: '/backend/merchant/contact/list',
    method: 'post',
    data: JSON.stringify({
      ...params,
      pageSize: 200
    }),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}


// 获取银行账户列表
export const getBankAccountList = () => {
  return service({
    url: '/backend/merchant/bankAccount/list',
    method: 'post',
    headers: {
      'Content-Type': 'application/json'
    }
  })
}