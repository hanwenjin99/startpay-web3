// 获取交易记录
import service from '@/utils/request'


// 查询取消记录
export const getRejectList = (params) => {
  return service({
    url: '/backend/reject_audit_record/list',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}