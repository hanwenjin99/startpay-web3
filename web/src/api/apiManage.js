// 获取API管理数据
import qs from 'qs'
import service from '@/utils/request'

// 获取我的钱包列表
export const getBoundAddressList = (params) => {
  return service({
    url: `/web3/wallet_list?${qs.stringify(params)}`,
    method: 'get'
  })
}

// 我的项目列表
export const getProjectList = (params) => {
  return service({
    url: `/web3/project_list?${qs.stringify(params)}`,
    method: 'get'
  })
}

// 新建项目
export const addProject = (params) => {
  return service({
    url: '/web3/create_project',
    method: 'post',
    data: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}