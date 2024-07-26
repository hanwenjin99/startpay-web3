// 获取总览数据
import service from '@/utils/request'

// 获取估值统计数据
export const getDashboard = () => {
  return service({
    url: '/web3/dashboard',
    method: 'get'
  })
}

// 获取公告列表
export const getAnnouncementList = () => {
  return service({
    url: '/web3/announcement',
    method: 'get',
    data: JSON.stringify({
      page: 1,
      pageSize: 5
    }),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}