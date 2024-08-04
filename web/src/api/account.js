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
    url: `/web3/token_list?${qs.stringify({
      ...params
      /*,
      page: 0,
      pageSize: 200*/
    })}`,
    method: 'get'
  })
}

// 查询收款记录
export const getDepositOrderList = (params) => {
  return service({
    url: `/web3/deposit_order?${qs.stringify(params)}`,
    method: 'get'
  })
}

// 单笔转账 - 转化接口
export const postBackendQuote = (chain) => {
  return service({
    url: `/web3/quote?chain=${chain}`,
    method: 'get'
  })
}

// 创建单笔转账
export const createSingleTransfer = (params) => {
  return service({
    url: '/web3/create_transfer',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 查询单笔转账记录
export const getTransferList = (params) => {
  return service({
    url: `/web3/transfer_order?${qs.stringify(params)}`,
    method: 'get'
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

// 商户端获取充值列表
export const getRechargeList = (params) => {
  return service({
    url: `/web3/merchant/charge/list?${qs.stringify(params)}`,
    method: 'get'
  })
}

// 商户端 - 充值提交
export const createMerchantRecharge = (params) => {
  return service({
    url: '/web3/merchant/charge/create',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 商户端 - 充值列表操作
export const actionMerchantRecharge = (params) => {
  return service({
    url: '/web3/merchant/charge/update',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}

// 管理后台 - 获取充值列表
export const getAdminRechargeList = (params) => {
  return service({
    url: `/web3/admin/charge/list?${qs.stringify(params)}`,
    method: 'get'
  })
}

// 管理后台 - 充值审核
export const reviewAdminRecharge = (params) => {
  return service({
    url: '/web3/admin/charge/update',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}
