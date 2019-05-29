import request from '@/utils/request'
export function getApiList(query) {
  return request({
    url: '/manager/APIList',
    method: 'get',
    params: query
  })
}
export function addApi(data) {
  return request({
    url: '/manager/APIAdd',
    method: 'post',
    data: data
  })
}

export function deleteApi(id) {
  return request({
    url: '/manager/ApiDelete',
    method: 'post',
    data: {
        id: id
    }
  })
}

export function addApiKey(data) {
    return request({
      url: '/manager/APIKeyAdd',
      method: 'post',
      data: data
    })
}

export function deleteApiKey(id) {
  return request({
    url: '/manager/APIKeyDelete',
    method: 'post',
    data: {
      id: id
    }
  })
}

