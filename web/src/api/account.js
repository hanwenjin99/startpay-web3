// 获取账户数据
import qs from 'qs'
import service from '@/utils/request'

// 获取资产信息
export const getAccountInfo = () => {
  return service({
    url: '/backend/account/info',
    method: 'get'
  })
}

// 获取账户可创建币种列表
export const getAccountCurrencyCreatableList = (params) => {
  return service({
    url: '/backend/options/account_currency_creatable/list',
    method: 'post',
    data: JSON.stringify({
      ...params,
      page: 0,
      pageSize: 200
    }),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 根据币种查询地址
export const getDepositAddress = (params) => {
  return service({
    url: '/backend/deposit_address/get',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 查询收款记录
export const getDepositOrderList = (params) => {
  return service({
    url: '/backend/deposit_order/list',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 单笔转账 - 转化接口
export const postBackendQuote = (chain) => {
  return service({
    url: '/backend/quote',
    method: 'post',
    data: JSON.stringify({ chain }),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 查询单笔转账记录
export const getTransferList = (params) => {
  return service({
    url: '/backend/transfer/list ',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 查询批量转账记录
export const getBatchPayoutList = (params) => {
  return service({
    url: `/backend/batchPayout/list?${qs.stringify({
      ...params,
      pageSize: 20
    })}`,
    method: 'get',
  })
}

// 获取模版地址
export const getBatchPayoutTemplate = () => {
  return service({
    url: '/backend/batchPayout/template',
    method: 'get',
  })
}

// 获取支付列表
export const getWithdrawOrderList = (params) => {
  return service({
    url: '/backend/withdraw_order/list',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 查询供应商分类列表
export const getSupplierCatList = () => {
  return service({
    url: '/backend/supplier_cat/list',
    method: 'get'
  })
}

// 查询签约供应商列表
export const getSignedSupplierList = (params) => {
  return service({
    url: '/backend/signedSupplier/list',
    method: 'post',
    data: JSON.stringify({
      ...params,
      page: 0,
      pageSize: 200
    }),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}