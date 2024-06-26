import request from '../axios';

/**
 * 获取资源列表
 * @param {*} params 
 * @returns 
 */
export const apiResourceList = (params) => {
    return request({
        url: 'api/v1/resource',
        method: 'get',
        params: params,
    })
}

/**
 * 创建资源
 * @param {*} data 
 * @returns 
 */
export const apiResourceCreate = (data) => {
    return request({
        url: 'api/v1/resource',
        method: 'post',
        data: data,
    })
}

/**
 * 上传资源
 * @param {*} data 
 * @returns 
 */
export const apiResourceUpload = (data) => {
    return request({
        url: 'api/v1/resource/upload',
        method: 'post',
        data: data,
    })
}

/**
 * 移除资源
 * @param {*} data 
 * @returns 
 */
export const apiResourceUnUpload = (id) => {
    return request({
        url: `api/v1/resource/upload/${id}`,
        method: 'delete',
    })
}

/**
 * 编辑资源
 * @param {*} data 
 * @returns 
 */
export const apiResourceModify = (data) => {
    return request({
        url: `api/v1/resource`,
        method: 'put',
        data: data,
    })
}


/**
 * 删除资源
 * @data {*} data 
 * @returns 
 */
export const apiResourceDelete = (id) => {
    return request({
        url: `api/v1/resource/${id}`,
        method: 'delete',
    })
}