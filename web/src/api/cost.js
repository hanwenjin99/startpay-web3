// 获取费用数据
import qs from 'qs'
import service from '@/utils/request'

// 获取账单和趋势数据
export const getBillSummary = (month) => {
  return service({
    url: `/web3/bill/summary?month=${month}`,
    method: 'get'
  })
}

// 查询费用明细
export const getBillList = (params) => {
  return service({
    url: `/web3/bill/list?${qs.stringify({
      ...params,
      pageSize: 20
    })}`,
    method: 'get'
  })
}

// 导出明细文件
export const exportBillList = (month) => {
  return service({
    url: `/web3/list/export?month=${month}`,
    method: 'get'
  })
}