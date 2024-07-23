// 获取账户信息
import qs from 'qs'
import service from '@/utils/request'

// 获取收款人列表
export const getMerchantContactList = (params) => {
  return service({
    url: `/web3/merchant/contact/list?${qs.stringify(params)}`,
    method: 'get'
  })
}

// 删除收款人信息
export const deleteMerchantContact = (params) => {
  return service({
    url: '/web3/merchant/contact/delete',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 创建收款人
export const addMerchantContact = (params) => {
  return service({
    url: '/web3/merchant/contact/create',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 获取银行账户列表
export const getBankAccountList = () => {
  return service({
    url: '/web3/merchant/bankAccount/list',
    method: 'get'
  })
}

// 删除银行账户信息
export const deleteBankAccount = (params) => {
  return service({
    url: '/web3/merchant/bankAccount/delete',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 创建银行账户
export const createBankAccount = (params) => {
  return service({
    url: '/web3/merchant/bankAccount/create',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}