// 获取账户数据
import qs from 'qs'
import service from '@/utils/request'

// 获取资产信息
export const getAccountInfo = () => {
  return service({
    url: '/web3/account_info',
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

// 管理后台 - 获取支付列表
export const getAdminWithdrawOrderList = (params) => {
  return service({
    url: `/web3/admin/withdraw/list?${qs.stringify(params)}`,
    method: 'get'
  })
}

// 管理后台 - 支付审核
export const reviewAdminWithdrawOrder = (params) => {
  return service({
    url: '/web3/admin/withdraw/update',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 商户端 - 获取支付列表
export const getWithdrawOrderList = (params) => {
  return service({
    url: `/web3/merchant/withdraw/list?${qs.stringify(params)}`,
    method: 'get'
  })
}

// 商户端 - 支付撤销
export const revokeMerchantWithdraw = (params) => {
  return service({
    url: '/web3/merchant/withdraw/update',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 商户端 - 支付提交
export const createMerchantWithdraw = (params) => {
  return service({
    url: '/web3/merchant/withdraw/create',
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