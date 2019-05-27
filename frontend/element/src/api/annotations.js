import request from '@/utils/request'
export function getAnnonList(query) {
  return request({
    url: '/iaas/AnnotationList',
    method: 'get',
    params: query
  })
}
export function addAnnon(data) {
  return request({
    url: '/iaas/AnnotationAdd',
    method: 'post',
    data: data
  })
}

export function deleteAnnon(id) {
  return request({
    url: '/iaas/AnnotationDelete',
    method: 'post',
    data: {
        id: id
    }
  })
}
export function Tst() {
  return request({
    url: '/iaas/Tst',
    method: 'get'
  })
}

