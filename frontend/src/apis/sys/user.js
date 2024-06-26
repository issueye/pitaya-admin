import request from '../axios';

/**
 * 登录接口
 * @param {*} data 
 * @returns 
 */
export const apiLogin = (data) => {
    return request({
        url: `/api/v1/login`,
        method: 'post',
        data: data,
    })
}

/**
 * 获取用户列表
 * @param {*} params 
 * @returns 
 */
export const apiUserList = (params) => {
    return request({
        url: 'api/v1/user',
        method: 'get',
        params: params,
    })
}

/**
 * 创建用户
 * @param {*} data 
 * @returns 
 */
export const apiUserCreate = (data) => {
    return request({
        url: 'api/v1/user',
        method: 'post',
        data: data,
    })
}

/**
 * 编辑用户
 * @param {*} data 
 * @returns 
 */
export const apiUserModify = (data) => {
    return request({
        url: `api/v1/user/${data.id}`,
        method: 'put',
        data: data,
    })
}

/**
 * 修改用户状态
 * @param {*} data 
 * @returns 
 */
export const apiUserModifyState = (id) => {
    return request({
        url: `api/v1/user/state/${id}`,
        method: 'put',
    })
}


/**
 * 删除用户
 * @data {*} data 
 * @returns 
 */
export const apiUserDelete = (id) => {
    return request({
        url: `api/v1/user/${id}`,
        method: 'delete',
    })
}

/**
 * 获取用户组列表
 * @param {*} params 
 * @returns 
 */
export const apiUserGroupList = (params) => {
    return request({
        url: 'api/v1/userGroup',
        method: 'get',
        params: params,
    })
}


/**
 * 添加用户组
 * @data {*} data 
 * @returns 
 */
export const apiUserGroupCreate = (data) => {
    return request({
        url: 'api/v1/userGroup',
        method: 'post',
        data: data,
    })
}

/**
 * 修改用户组
 * @data {*} data 
 * @returns 
 */
export const apiUserGroupModify = (data) => {
    return request({
        url: `api/v1/userGroup/${data.id}`,
        method: 'put',
        data: data,
    })
}


/**
 * 修改用户组状态
 * @param {*} data 
 * @returns 
 */
export const apiUserGroupModifyState = (id) => {
    return request({
        url: `api/v1/userGroup/state/${id}`,
        method: 'put',
    })
}


/**
 * 修改用户组
 * @data {*} data 
 * @returns 
 */
export const apiUserGroupDelete = (id) => {
    return request({
        url: `api/v1/userGroup/${id}`,
        method: 'delete',
    })
}
