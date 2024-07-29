export const UPLOAD_FILE_URL = '///backend/batchPayout/upload/csv'

export const TRANSFER_STATUS_OPTIONS = [
  { label: '已创建', value: 'CREATED' },
  { label: '已完成', value: 'FINISHED' },
  { label: '进行中', value: 'PROCESSING' },
  { label: '错误', value: 'FAILED' },
  { label: '待审核', value: 'UNDER_REVIEW' },
]

// 收款状态信息
export const GATHER_LIST_STATUS = {
  ERROR: '错误',
  FINISHED: '完成',
  INSUFFICIENT_AMOUNT: '充值金额不足（归集）',
  PROCESSING: '处理中'
}

// 转账状态信息
export const TRANSFER_LIST_STATUS = {
  UNDER_REVIEW: '审核中',
  CREATED: '已创建',
  GAS_CHARGING: '燃料费充值中',
  PROCESSING: '处理中',
  TO_BE_CONFIRMED: '待确认',
  SUCCESS: '成功',
  ERROR: '错误',
  CANCELED: '已取消'
}